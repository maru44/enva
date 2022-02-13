package controllers

import (
	"context"
	"encoding/json"
	"fmt"
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
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}

	kvs, err := con.in.ListValid(ctx, *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrNotFound), nil)
		return
	}
	sort.Slice(kvs, func(i, j int) bool { return kvs[i].Key.String() > kvs[j].Key.String() })
	response(w, r, nil, map[string]interface{}{"data": kvs})
}

func (con *CliKvController) DetailView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := r.URL.Query().Get(QueryParamsKvKey)
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	projectID, err := con.userGuestAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}

	kv, err := con.in.DetailValid(ctx, domain.KvKey(key), *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrNotFound), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": kv})
}

func (con *CliKvController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)
	var input domain.KvInputWithProjectID

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}
	input.ProjectID = *projectID

	id, err := con.in.Create(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
}

func (con *CliKvController) BulkInsertView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)
	fmt.Println(projectSlug)
	fmt.Println(orgSlug)
	var inputs []domain.KvInput

	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}

	// if there are any key value sets
	// it is not validated
	kvs, err := con.in.ListValid(ctx, *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrNotFound), nil)
		return
	}
	if kvs != nil {
		response(w, r, perr.New(perr.ErrBadRequest.Error(), perr.ErrBadRequest, "Key Value sets already exists"), nil)
		return
	}

	if err := con.cIn.BulkInsert(ctx, *projectID, inputs); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "success"})
}

func (con *CliKvController) UpdateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	var input domain.KvInputWithProjectID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}
	input.ProjectID = *projectID

	id, err := con.in.Update(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
	}

	response(w, r, nil, map[string]interface{}{"data": id})
}

func (con *CliKvController) DeleteView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	key := r.URL.Query().Get(QueryParamsKvKey)
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)

	projectID, err := con.userAccessToProject(ctx, projectSlug, orgSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}

	affected, err := con.in.DeleteByKey(ctx, domain.KvKey(key), *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": affected})
}

func (con *CliKvController) userAccessToProject(ctx context.Context, projectSlug, orgSlug string) (*domain.ProjectID, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	// find parent project
	var p *domain.Project
	if orgSlug == "" {
		p, err = con.pIn.GetBySlug(ctx, projectSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrNotFound, "Project is not found")
		}
	} else {
		p, err = con.pIn.GetBySlugAndOrgSlug(ctx, projectSlug, orgSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrNotFound, "Project is not found")
		}
	}

	// validate user can access to project
	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			return nil, perr.New("user is guest", perr.ErrForbidden)
		}
	} else {
		ut, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrForbidden)
		}
		if *ut == domain.UserTypeGuest {
			return nil, perr.New("user is guest", perr.ErrForbidden, "you are guest user")
		}
	}

	return &p.ID, nil
}

func (con *CliKvController) userGuestAccessToProject(ctx context.Context, projectSlug, orgSlug string) (*domain.ProjectID, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	// find parent project
	var p *domain.Project
	if orgSlug == "" {
		p, err = con.pIn.GetBySlug(ctx, projectSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrNotFound, "Project is not found")
		}
	} else {
		p, err = con.pIn.GetBySlugAndOrgSlug(ctx, projectSlug, orgSlug)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrNotFound, "Project is not found")
		}
	}

	// validate user can access to project
	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			return nil, perr.New("user is not owner of this project", perr.ErrForbidden)
		}
	} else {
		_, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrForbidden)
		}
	}

	return &p.ID, nil
}
