package controllers

import (
	"context"
	"net/http"
	"strings"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type CliUserController struct {
	in domain.IUserInteractor
}

func NewCliUserController(sql database.ISqlHandler, pass password.IPassword) *CliUserController {
	return &CliUserController{
		in: usecase.NewUserInteractor(
			&database.UserRepository{
				ISqlHandler: sql,
				IPassword:   pass,
			},
		),
	}
}

/********************************
    Middle ware
********************************/

func (con *CliUserController) BaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(
			r.Context(),
			domain.CtxAccessKey,
			domain.CtxAccess{
				Method: r.Method,
				URL:    r.URL.Path,
			},
		)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (con *CliUserController) LoginRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iPass := r.Header.Get("Authorization")
		iPassArr := strings.SplitN(iPass, domain.CLI_HEADER_SEP, 2)
		if len(iPassArr) != 2 {
			response(w, r, perr.New(perr.ErrForbidden.Error(), perr.ErrForbidden), nil)
			return
		}

		input := &domain.UserCliValidationInput{
			EmailOrUsername: iPassArr[0],
			CliPassword:     iPassArr[1],
		}

		user, err := con.in.GetUserCli(r.Context(), input)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
			return
		}

		r = setUserToContext(r, *user)
		next.ServeHTTP(w, r)
	})
}
