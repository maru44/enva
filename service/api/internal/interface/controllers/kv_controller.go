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

type KvController struct {
	in  domain.IKvInteractor
	pIn domain.IProjectInteractor
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
	}
}

func (con *KvController) ListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := r.URL.Query().Get(QueryParamsProjectID)

	if err := con.userAccessToProject(ctx, domain.ProjectID(projectID)); err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	kvs, err := con.in.ListValid(ctx, domain.ProjectID(projectID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
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
	json.NewDecoder(r.Body).Decode(&input)

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
	user := domain.UserFromCtx(ctx)

	// find parent project
	p, err := con.pIn.GetByID(ctx, projectID)
	if err != nil {
		return perr.Wrap(err, perr.NotFound, "Project is not found")
	}

	// validate user can access to project
	if err := p.ValidateUserGet(user); err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}

	return nil
}
