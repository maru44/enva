package controllers

import (
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type UserController struct {
	in domain.IUserInteractor
}

func NewUserController(sql database.ISqlHandler, pass password.IPassword) *UserController {
	return &UserController{
		in: usecase.NewUserInteractor(
			&database.UserRepository{
				ISqlHandler: sql,
				IPassword:   pass,
			},
		),
	}
}

func (con *UserController) ExistsCliPasswordView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cU := ctx.Value(domain.CtxUserKey).(domain.User)

	user, err := con.in.GetByID(ctx, cU.ID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}

	if user.CliPassword != nil {
		response(w, r, nil, map[string]interface{}{"data": "exists"})
		return
	}

	response(w, r, perr.New("Cli password not found", perr.BadRequest, "Cli password not found"), nil)
	return
}

func (con *UserController) UpdateCliPasswordView(w http.ResponseWriter, r *http.Request) {
	pass, err := con.in.UpdateCliPassword(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": pass})
	return
}

func (con *UserController) CreateView(w http.ResponseWriter, r *http.Request) {
	id, err := con.in.Create(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}
