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

type (
	CliKvController struct {
		in  domain.IKvInteractor
		pIn domain.IProjectInteractor
		cIn domain.ICliKvInteractor
		oIn domain.IOrgInteractor
	}
)

func NewCliKvController(sql database.ISqlHandler) *CliKvController {
	return &CliKvController{
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
		cIn: usecase.NewCliKvInteractor(
			&database.CliKvRepository{
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

func (con *CliKvController) ListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	projectID, err := con.userGuestAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	kvs, err := con.in.ListValid(ctx, *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	sort.Slice(kvs, func(i, j int) bool { return kvs[i].Key.String() > kvs[j].Key.String() })
	response(w, r, nil, map[string]interface{}{"data": kvs})
	return
}

func (con *CliKvController) DetailView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := r.URL.Query().Get(QueryParamsKvKey)
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	projectID, err := con.userGuestAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	kv, err := con.in.DetailValid(ctx, domain.KvKey(key), *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": kv})
	return
}

func (con *CliKvController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)
	var input domain.KvInputWithProjectID

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}
	input.ProjectID = *projectID

	id, err := con.in.Create(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}

func (con *CliKvController) BulkInsertView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)
	var inputs []domain.KvInput

	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	// if there are any key value sets
	// it is not validated
	kvs, err := con.in.ListValid(ctx, *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	if kvs != nil {
		response(w, r, perr.New(perr.BadRequest.Error(), perr.BadRequest, "Key Value sets already exists"), nil)
		return
	}

	if err := con.cIn.BulkInsert(ctx, *projectID, inputs); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "success"})
	return
}

func (con *CliKvController) UpdateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	var input domain.KvInputWithProjectID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}
	input.ProjectID = *projectID

	id, err := con.in.Update(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}

func (con *CliKvController) DeleteView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	key := r.URL.Query().Get(QueryParamsKvKey)
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	affected, err := con.in.DeleteByKey(ctx, domain.KvKey(key), *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": affected})
	return
}

func (con *CliKvController) userAccessToProject(ctx context.Context, projectSlug, orgSlug string) (*domain.ProjectID, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	// find parent project
	var p *domain.Project
	if orgSlug == "" {
		p, err = con.pIn.GetBySlug(ctx, projectSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.NotFound, "Project is not found")
		}
	} else {
		p, err = con.pIn.GetBySlugAndOrgSlug(ctx, projectSlug, orgSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.NotFound, "Project is not found")
		}
	}

	// validate user can access to project
	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			return nil, perr.New("user is guest", perr.Forbidden)
		}
	} else {
		ut, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			return nil, perr.Wrap(err, perr.Forbidden)
		}
		if *ut == domain.UserTypeGuest {
			return nil, perr.New("user is guest", perr.Forbidden, "you are guest user")
		}
	}

	return &p.ID, nil
}

func (con *CliKvController) userGuestAccessToProject(ctx context.Context, projectSlug, orgSlug string) (*domain.ProjectID, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	// find parent project
	var p *domain.Project
	if orgSlug == "" {
		p, err = con.pIn.GetBySlug(ctx, projectSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.NotFound, "Project is not found")
		}
	} else {
		p, err = con.pIn.GetBySlugAndOrgSlug(ctx, projectSlug, orgSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.NotFound, "Project is not found")
		}
	}

	// validate user can access to project
	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			return nil, perr.New("user is not owner of this project", perr.Forbidden)
		}
	} else {
		_, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			return nil, perr.Wrap(err, perr.Forbidden)
		}
	}

	return &p.ID, nil
}
