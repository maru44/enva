package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgMemberController struct {
	in domain.IOrgMemberInteractor
}

func NewOrgMemberController(sql database.ISqlHandler) *OrgMemberController {
	return &OrgMemberController{
		in: usecase.NewOrgMemberInteractor(
			&database.OrgMemberRepository{
				ISqlHandler: sql,
			},
		),
	}
}

func (con *OrgMemberController) CreateView(w http.ResponseWriter, r *http.Request) {
	var input domain.OrgMemberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	if err := con.in.Create(r.Context(), input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "OK"})
	return
}

// @TODO deny (inv con)
