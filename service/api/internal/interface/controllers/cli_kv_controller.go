package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type CliKvController struct {
	in  domain.IKvInteractor
	pIn domain.IProjectInteractor
}

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
	}
}

func (con *CliKvController) ListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)

	projectID, err := con.userAccessToProject(ctx, projectSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	kvs, err := con.in.ListValid(ctx, *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": kvs})
	return
}

func (con *CliKvController) DetailView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := r.URL.Query().Get(QueryParamsKvKey)
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)

	projectID, err := con.userAccessToProject(ctx, projectSlug)
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
	var input domain.KvInputWithProjectID

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	projectID, err := con.userAccessToProject(ctx, projectSlug)
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

func (con *CliKvController) UpdateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)
	var input domain.KvInputWithProjectID
	json.NewDecoder(r.Body).Decode(&input)

	projectID, err := con.userAccessToProject(ctx, projectSlug)
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

	kvId := r.URL.Query().Get(QueryParamsKvID)
	projectSlug := r.URL.Query().Get(QueryParamsProjectSlug)

	projectID, err := con.userAccessToProject(ctx, projectSlug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	affected, err := con.in.Delete(ctx, domain.KvID(kvId), *projectID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": affected})
	return
}

func (con *CliKvController) userAccessToProject(ctx context.Context, projectSlug string) (*domain.ProjectID, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	// find parent project
	p, err := con.pIn.GetBySlug(ctx, projectSlug)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound, "Project is not found")
	}

	// validate user can access to project
	if err := p.ValidateUserGet(user); err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	return &p.ID, nil
}
