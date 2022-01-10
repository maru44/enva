package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type ProjectController struct {
	in  domain.IProjectInteractor
	oIn domain.IOrgInteractor
}

func NewProjectController(sql database.ISqlHandler) *ProjectController {
	return &ProjectController{
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

func (con *ProjectController) ListAllView(w http.ResponseWriter, r *http.Request) {
	ps, err := con.in.ListAll(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": ps})
	return
}

// filtered by user @sql
func (con *ProjectController) ListByUserView(w http.ResponseWriter, r *http.Request) {
	ps, err := con.in.ListByUser(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": ps})
	return
}

// filtered by user @sql
func (con *ProjectController) ListByOrgView(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get(QueryParamsID)
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
	return
}

// filtered by user @sql
func (con *ProjectController) SlugListByUserView(w http.ResponseWriter, r *http.Request) {
	slugs, err := con.in.SlugListByUser(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": slugs})
	return
}

// filtered by user @sql
func (con *ProjectController) ProjectDetailView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	slug := r.URL.Query().Get(QueryParamsSlug)
	if slug == "" {
		response(w, r, perr.New("No slug was given", perr.BadRequest), nil)
		return
	}

	// with ?orgId=
	orgID := r.URL.Query().Get(QueryParamsOrgID)
	if orgID != "" {
		p, err := con.in.GetBySlugAndOrgID(ctx, slug, domain.OrgID(orgID))
		if err != nil {
			response(w, r, perr.Wrap(err, perr.NotFound), nil)
			return
		}

		response(w, r, nil, map[string]interface{}{"data": p})
		return
	}

	// with ?orgSlug=
	orgSlug := r.URL.Query().Get(QueryParamsOrgSlug)
	if orgSlug != "" {
		p, err := con.in.GetBySlugAndOrgSlug(ctx, slug, orgSlug)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.NotFound), nil)
			return
		}

		response(w, r, nil, map[string]interface{}{"data": p})
		return
	}

	p, err := con.in.GetBySlug(ctx, slug)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": p})
	return
}

// @TODO if orgID not equal null >>
// user need to be a member of that org
func (con *ProjectController) CreateView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input domain.ProjectInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	if input.OrgID != nil {
		ut, err := con.oIn.MemberGetCurrentUserType(ctx, *input.OrgID)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.NotFound), nil)
			return
		}
		if err := ut.IsAdmin(); err != nil {
			response(w, r, perr.Wrap(err, perr.Forbidden), nil)
			return
		}
	}

	id, err := con.in.Create(ctx, input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}

func (con *ProjectController) DeleteView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := r.URL.Query().Get(QueryParamsProjectID)

	p, err := con.in.GetByID(ctx, domain.ProjectID(projectID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	// validate user access
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
	}
	if p.OwnerType == "user" {
		if p.OwnerUser.ID != user.ID {
			response(w, r, perr.New("user is not owner of this project", perr.Forbidden), nil)
			return
		}
	} else {
		ut, err := con.oIn.MemberGetCurrentUserType(ctx, p.OwnerOrg.ID)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.NotFound), nil)
			return
		}
		if err := ut.IsAdmin(); err != nil {
			response(w, r, perr.Wrap(err, perr.Forbidden), nil)
			return
		}
	}

	affected, err := con.in.Delete(ctx, domain.ProjectID(projectID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": affected})
	return
}
