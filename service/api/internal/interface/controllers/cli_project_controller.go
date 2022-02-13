package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type (
	CliProjectController struct {
		in  domain.IProjectInteractor
		oIn domain.IOrgInteractor
	}
)

func NewCliProjectController(sql database.ISqlHandler) *CliProjectController {
	return &CliProjectController{
		in: usecase.NewProjectInteractor(
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

func (con *CliProjectController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var cliInput domain.CliProjectInput
	if err := json.NewDecoder(r.Body).Decode(&cliInput); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}
	if err := cliInput.Validate(); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	input := cliInput.ToProjectInputWithoutOrg()
	if cliInput.OrgSlug != nil {
		org, ut, err := con.oIn.DetailBySlug(ctx, *cliInput.OrgSlug)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.ErrNotFound), nil)
			return
		}

		if err := ut.IsAdmin(); err != nil {
			response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
			return
		}
		input.OrgID = &org.ID
	}

	id, err := con.in.Create(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
}
