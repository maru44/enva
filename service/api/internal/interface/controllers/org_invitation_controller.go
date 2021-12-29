package controllers

import (
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgInvitationController struct {
	in domain.IOrgInvitationInteractor
}

func NewOrgInvitationController(sql database.ISqlHandler) *OrgInvitationController {
	return &OrgInvitationController{
		in: usecase.NewOrgInvitaionInteractor(
			&database.OrgInvitationRepository{
				ISqlHandler: sql,
			},
		),
	}
}

func (con *OrgInvitationController) ListFromOrgView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get(QueryParamsID)

	invs, err := con.in.ListFromOrg(ctx, domain.OrgID(id))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": invs})
	return
}

func (con *OrgInvitationController) ListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	invs, err := con.in.List(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": invs})
	return
}

func (con *OrgInvitationController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	var input domain.OrgInvitationInput
	if err := con.in.Create(ctx, input, cu.ID); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "OK"})
	return
}
