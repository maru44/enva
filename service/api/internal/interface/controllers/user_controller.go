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

func (con *UserController) GetUserView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctxUser, _ := domain.UserFromCtx(ctx)
	if ctxUser == nil {
		response(w, r, perr.New("no user in context", perr.ErrBadRequest), nil)
		return
	}

	user, err := con.in.GetByID(ctx, ctxUser.ID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": user})
}

func (con *UserController) ExistsCliPasswordView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cU, _ := domain.UserFromCtx(ctx)

	user, err := con.in.GetByID(ctx, cU.ID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrNotFound), nil)
		return
	}

	if user.CliPassword != nil {
		response(w, r, nil, map[string]interface{}{"data": "exists"})
		return
	}

	response(w, r, perr.New("Cli password not found", perr.ErrBadRequest, "Cli password not found"), nil)
}

func (con *UserController) UpdateCliPasswordView(w http.ResponseWriter, r *http.Request) {
	pass, err := con.in.UpdateCliPassword(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": pass})
}

func (con *UserController) UpsertView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := con.in.UpsertIfNotInvalid(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": user})
}

func (con *UserController) UpdateToInvalidView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
		return
	}

	input := domain.UserUpdateIsValidInput{
		ID:      user.ID,
		IsValid: false,
	}
	if err := con.in.UpdateValid(ctx, input); err != nil {
		response(w, r, perr.Wrap(err, perr.ErrBadRequest), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": nil})
}
