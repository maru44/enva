package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/mysmtp"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgController struct {
	in  domain.IOrgInteractor
	uIn domain.IUserInteractor
}

func NewOrgController(sql database.ISqlHandler, smtp mysmtp.ISmtpHandler) *OrgController {
	return &OrgController{
		in: usecase.NewOrgInteractor(
			&database.OrgRepository{
				ISqlHandler:  sql,
				ISmtpHandler: smtp,
			},
		),
		uIn: usecase.NewUserInteractor(
			&database.UserRepository{
				ISqlHandler: sql,
			},
		),
	}
}

func (con *OrgController) ListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orgs, err := con.in.List(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": orgs})
	return
}

func (con *OrgController) ListOwnerAdminView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orgs, err := con.in.ListOwnerAdmin(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": orgs})
	return
}

func (con *OrgController) DetailBySlugView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := r.URL.Query().Get(QueryParamsSlug)

	org, cuUserType, err := con.in.DetailBySlug(ctx, slug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{
		"data": map[string]interface{}{
			"org":               org,
			"current_user_type": *cuUserType,
		}})
	return
}

func (con *OrgController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input domain.OrgInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	slug, err := con.in.Create(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": slug})
	return
}

/* inv */

func (con *OrgController) InvitationListFromOrgView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get(QueryParamsID)

	invs, err := con.in.InvitationListFromOrg(ctx, domain.OrgID(id))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": invs})
	return
}

// func (con *OrgController) InvitationListView(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	invs, err := con.in.InvitationList(ctx)
// 	if err != nil {
// 		response(w, r, perr.Wrap(err, perr.NotFound), nil)
// 		return
// 	}

// 	response(w, r, nil, map[string]interface{}{"data": invs})
// 	return
// }

func (con *OrgController) InvitationDetailView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(QueryParamsID)

	ctx := r.Context()
	inv, err := con.in.InvitationDetail(ctx, domain.OrgInvitationID(id))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	if inv.Status != domain.OrgInvitationStatusNew {
		err := fmt.Errorf("this invitation is %s", string(inv.Status))
		response(w, r, perr.Wrap(err, perr.BadRequest, err.Error()), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": inv})
	return
}

func (con *OrgController) InviteView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input domain.OrgInvitationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	// validate by current user type
	userType, err := con.in.MemberGetCurrentUserType(ctx, input.OrgID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}
	if err := userType.IsAdmin(); err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	// add input invited userID
	invitedUser, err := con.uIn.GetByEmail(ctx, input.Eamil)
	if err == nil {
		input.User = invitedUser
	}

	if err := con.in.Invite(ctx, input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "OK"})
	return
}

func (con *OrgController) DenyView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get(QueryParamsID)

	if err := con.in.InvitationDeny(ctx, domain.OrgInvitationID(id)); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "OK"})
	return
}

/* member */

func (con *OrgController) MemberCreateView(w http.ResponseWriter, r *http.Request) {
	var input domain.OrgMemberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	if err := con.in.MemberCreate(r.Context(), input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "OK"})
	return
}

func (con *OrgController) MemberListView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get(QueryParamsID)

	members, err := con.in.MemberList(ctx, domain.OrgID(id))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": members})
	return
}
