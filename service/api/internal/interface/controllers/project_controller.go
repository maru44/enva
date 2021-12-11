package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/ichigo/service/api/internal/interface/database"
	"github.com/maru44/ichigo/service/api/internal/usecase"
	"github.com/maru44/ichigo/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type ProjectController struct {
	in domain.IProjectInteractor
}

func NewProjectController(sql database.ISqlHandler) *ProjectController {
	return &ProjectController{
		in: usecase.NewProjectInteractor(
			&database.ProjectReposotory{
				ISqlHandler: sql,
			},
		),
	}
}

func (con *ProjectController) ListByUserView(w http.ResponseWriter, r *http.Request) {
	ps, err := con.in.ListByUser(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": ps})
}

func (con *ProjectController) ListByProjectView(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get(QueryParamsOrgID)
	if orgID == "" {
		response(w, r, perr.New(ErrorNoOrgIdParams.Error(), perr.BadRequest), nil)
		return
	}

	ps, err := con.in.ListByOrg(r.Context(), domain.OrgID(orgID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": ps})
}

func (con *ProjectController) CreateView(w http.ResponseWriter, r *http.Request) {
	var input domain.ProjectInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	id, err := con.in.Create(r.Context(), input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}
