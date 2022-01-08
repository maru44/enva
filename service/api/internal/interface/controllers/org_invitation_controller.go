package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/mysmtp"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgInvitationController struct {
	in  domain.IOrgInvitationInteractor
	mIn domain.IOrgMemberInteractor
	uIn domain.IUserInteractor
}

func NewOrgInvitationController(sql database.ISqlHandler, smtp mysmtp.ISmtpHandler) *OrgInvitationController {
	return &OrgInvitationController{
		in: usecase.NewOrgInvitaionInteractor(
			&database.OrgInvitationRepository{
				ISqlHandler:  sql,
				ISmtpHandler: smtp,
			},
		),
		mIn: usecase.NewOrgMemberInteractor(
			&database.OrgMemberRepository{
				ISqlHandler: sql,
			},
		),
		uIn: usecase.NewUserInteractor(
			&database.UserRepository{
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

func (con *OrgInvitationController) DetailView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(QueryParamsID)

	ctx := r.Context()
	inv, err := con.in.Detail(ctx, domain.OrgInvitationID(id))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": inv})
	return
}

// @TODO send mail
func (con *OrgInvitationController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input domain.OrgInvitationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	// validate by current user type
	userType, err := con.mIn.GetCurrentUserType(ctx, input.OrgID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}
	if !userType.IsAdmin() {
		response(w, r, perr.New("current user does not admin or owner of this org", perr.Forbidden), nil)
		return
	}

	// add input invited userID
	invitedUser, err := con.uIn.GetByEmail(ctx, input.Eamil)
	if err == nil {
		input.User = invitedUser
	}

	if err := con.in.Create(ctx, input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "OK"})
	return
}
