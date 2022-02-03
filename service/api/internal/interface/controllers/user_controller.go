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
		response(w, r, perr.New("no user in context", perr.BadRequest), nil)
		return
	}

	user, err := con.in.GetByID(ctx, ctxUser.ID)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": user})
}

func (con *UserController) ExistsCliPasswordView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cU, _ := domain.UserFromCtx(ctx)

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
}

func (con *UserController) UpdateCliPasswordView(w http.ResponseWriter, r *http.Request) {
	pass, err := con.in.UpdateCliPassword(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": pass})
}

func (con *UserController) CreateView(w http.ResponseWriter, r *http.Request) {
	id, err := con.in.CreateOrDoNothing(r.Context())
	if err != nil {
		// destroy cookie
		destroyCookie(w, domain.JwtCookieKeyAccessToken)
		destroyCookie(w, domain.JwtCookieKeyIdToken)
		destroyCookie(w, domain.JwtCookieKeyRefreshToken)

		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil, map[string]interface{}{"data": id})
}

func (con *UserController) UpdateToInvalidView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.Forbidden), nil)
		return
	}

	input := domain.UserUpdateIsValidInput{
		ID:      user.ID,
		IsValid: false,
	}
	if err := con.in.UpdateValid(ctx, input); err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": user.ID})
}
