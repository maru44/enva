package controllers

import (
	"context"
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/myjwt"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/config"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type (
	BaseController struct {
		ji domain.JwtIntectactor
	}
)

const ConBase = ConName("BASE")

func (con *BaseController) Name() ConName {
	return ConBase
}

func NewBaseController(jp myjwt.JwtParserAbstract) *BaseController {
	return NewBaseControllerFromUsecase(
		usecase.NewJwtInteractor(
			&myjwt.JwtRepository{
				JwtParserAbstract: jp,
			},
		),
	)
}

func NewBaseControllerFromUsecase(in domain.JwtIntectactor) *BaseController {
	return &BaseController{ji: in}
}

/********************************
    End points
********************************/

func (con *BaseController) NotFoundView(w http.ResponseWriter, r *http.Request) {
	response(w, r, perr.New("", perr.ErrNotFound), nil)
}

func (con *BaseController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response(w, r, nil, map[string]interface{}{"data": "OK"})
}

/********************************
    Middleware
********************************/

func (con *BaseController) BaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keySet, err := con.ji.FetchJwk(r.Context(), config.COGNITO_KEYS_URL)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.ErrInternalServerErrorWithUrgency), nil)
			return
		}

		if err := con.corsMiddleware(w, r); err != nil {
			response(w, r, perr.Wrap(err, perr.ErrCorsError), nil)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			domain.CtxAccessKey,
			domain.CtxAccess{
				Method: r.Method,
				URL:    r.URL.Path,
			},
		)
		ctx = context.WithValue(
			ctx,
			domain.CtxCognitoKeySetKey,
			keySet,
		)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (con *BaseController) corsMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Origin") != config.FRONT_URL && r.Header.Get("Origin") != "" {
		return perr.New("cors error", perr.ErrCorsError)
	}
	w.Header().Set("Access-Control-Allow-Origin", config.FRONT_URL)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie, Id-Token, Refresh-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Max-Age", "3600")
	return nil
}

func (con *BaseController) PostOnlyMiddleware(next http.Handler) http.Handler {
	return con.methodsOnlyMiddleware(next, []string{http.MethodPost})
}

func (con *BaseController) PutOnlyMiddleware(next http.Handler) http.Handler {
	return con.methodsOnlyMiddleware(next, []string{http.MethodPut})
}

func (con *BaseController) DeleteOnlyMiddleware(next http.Handler) http.Handler {
	return con.methodsOnlyMiddleware(next, []string{http.MethodDelete})
}

func (con *BaseController) GetOnlyMiddleware(next http.Handler) http.Handler {
	return con.methodsOnlyMiddleware(next, []string{http.MethodGet})
}

func (con *BaseController) methodsOnlyMiddleware(next http.Handler, methods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			response(w, r, nil, nil)
			return
		}
		for _, m := range methods {
			if r.Method == m {
				next.ServeHTTP(w, r)
				return
			}
		}
		response(w, r, perr.New("", perr.ErrMethodNotAllowed), nil)
	})
}

func (con *BaseController) AnyMethodMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			response(w, r, nil, nil)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (con *BaseController) GiveUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(domain.JwtCookieKeyIdToken)
		if err != nil {
			next.ServeHTTP(w, r)
		} else {
			user, err := con.ji.GetUserByJwt(r.Context(), cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// set user to context
			r = setUserToContext(r, *user)

			next.ServeHTTP(w, r)
		}
	})
}

func (con *BaseController) LoginRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(domain.JwtCookieKeyIdToken)
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				response(w, r, perr.Wrap(err, perr.ErrExpired), nil)
			default:
				response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
			}
			return
		}

		user, err := con.ji.GetUserByJwt(r.Context(), cookie.Value)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.ErrForbidden), nil)
			return
		}

		r = setUserToContext(r, *user)
		next.ServeHTTP(w, r)
	})
}

func setUserToContext(r *http.Request, u domain.User) *http.Request {
	ctx := context.WithValue(
		r.Context(),
		domain.CtxUserKey,
		u,
	)
	return r.WithContext(ctx)
}
