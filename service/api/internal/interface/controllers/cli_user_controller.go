package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type CliUserController struct {
	in domain.ICliUserInteractor
}

func NewCliUserController(sql database.ISqlHandler, pass password.IPassword) *CliUserController {
	return &CliUserController{
		in: usecase.NewCliUserInteractor(
			&database.CliUserRepository{
				ISqlHandler: sql,
				IPassword:   pass,
			},
		),
	}
}

func (con *CliUserController) CreateView(w http.ResponseWriter, r *http.Request) {
	pass, err := con.in.Create(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": pass})
	return
}

func (con *CliUserController) UpdateView(w http.ResponseWriter, r *http.Request) {
	pass, err := con.in.Update(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": pass})
	return
}

func (con *CliUserController) ValidateView(w http.ResponseWriter, r *http.Request) {
	input := &domain.CliUserValidateInput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	if err := con.in.Validate(r.Context(), input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "valid user"})
	return
}

func (con *CliUserController) ExistsView(w http.ResponseWriter, r *http.Request) {
	if err := con.in.Exists(r.Context()); err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": "exists"})
	return
}

/********************************
    Middle ware
********************************/

func (con *CliUserController) LoginRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input := &domain.CliUserValidateInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			response(w, r, perr.Wrap(err, perr.BadRequest), nil)
			return
		}

		user, err := con.in.GetUser(r.Context(), input)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.Forbidden), nil)
			return
		}

		r = setUserToContext(r, *user)
		next.ServeHTTP(w, r)
	})
}
