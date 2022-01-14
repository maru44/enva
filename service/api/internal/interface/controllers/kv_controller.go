package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type KvController struct {
	in  domain.IKvInteractor
	pIn domain.IProjectInteractor
	oIn domain.IOrgInteractor
}

func NewKvController(sql database.ISqlHandler) *KvController {
	return &KvController{
		in: usecase.NewKvInteractor(
			&database.KvRepository{
				ISqlHandler: sql,
			},
		),
		pIn: usecase.NewProjectInteractor(
			&database.ProjectReposotory{
				ISqlHandler: sql,
			},
		),
		oIn: usecase.NewOrgInteractor(
			&database.OrgRepository{
				ISqlHandler: sql,
			},
		),
	}
}

func (con *KvController) ListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := r.URL.Query().Get(QueryParamsProjectID)

	if err := con.userGuestAccessToProject(ctx, domain.ProjectID(projectID)); err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	kvs, err := con.in.ListValid(ctx, domain.ProjectID(projectID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	sort.Slice(kvs, func(i, j int) bool { return kvs[i].Key.String() > kvs[j].Key.String() })
	response(w, r, nil, map[string]interface{}{"data": kvs})
	return
}

func (con *KvController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input domain.KvInputWithProjectID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	if err := con.userAccessToProject(ctx, input.ProjectID); err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	id, err := con.in.Create(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}

func (con *KvController) UpdateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input domain.KvInputWithProjectID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	if err := con.userAccessToProject(ctx, input.ProjectID); err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	id, err := con.in.Update(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}

func (con *KvController) DeleteView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectID := r.URL.Query().Get(QueryParamsProjectID)
	kvId := r.URL.Query().Get(QueryParamsKvID)

	if err := con.userAccessToProject(ctx, domain.ProjectID(projectID)); err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	affected, err := con.in.Delete(ctx, domain.KvID(kvId), domain.ProjectID(projectID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": affected})
	return
}

func (con *KvController) userAccessToProject(ctx context.Context, projectID domain.ProjectID) error {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}

	// find parent project
	p, err := con.pIn.GetByID(ctx, projectID)
	if err != nil {
		return perr.Wrap(err, perr.NotFound, "Project is not found")
	}

	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			return perr.New("user is guest", perr.Forbidden)
		}
	} else {
		ut, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			return perr.Wrap(err, perr.Forbidden)
		}
		if err := ut.IsUser(); err != nil {
			return perr.Wrap(err, perr.Forbidden)
		}
	}

	return nil
}

func (con *KvController) userGuestAccessToProject(ctx context.Context, projectID domain.ProjectID) error {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}

	// find parent project
	p, err := con.pIn.GetByID(ctx, projectID)
	if err != nil {
		return perr.Wrap(err, perr.NotFound, "Project is not found")
	}

	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			return perr.New("user is not owner of this project", perr.Forbidden)
		}
	} else {
		_, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			return perr.Wrap(err, perr.Forbidden)
		}
	}

	return nil
}
