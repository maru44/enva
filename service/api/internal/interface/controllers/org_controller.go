package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgController struct {
	in domain.IOrgInteractor
}

func NewOrgController(sql database.ISqlHandler) *OrgController {
	return &OrgController{
		in: usecase.NewOrgInteractor(
			&database.OrgRepository{
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
