package controllers

import (
	"context"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/enva/service/api/internal/interface/myjwt"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type (
	BaseController struct {
		ji domain.JwtIntectactor
	}
)

func NewBaseController(jp myjwt.JwtParserAbstract) *BaseController {
	return &BaseController{
		ji: usecase.NewJwtInteractor(
			&myjwt.JwtRepository{
				JwtParserAbstract: jp,
			},
		),
	}
}

/********************************
    Get jwks
********************************/

func (con *BaseController) GetKeySet() (jwk.Set, error) {
	return jwk.Fetch(context.Background(), config.COGNITO_KEYS_URL)
}

/********************************
    End points
********************************/

func (con *BaseController) NotFoundView(w http.ResponseWriter, r *http.Request) {
	response(w, r, perr.New("", perr.NotFound), nil)
}

func (con *BaseController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response(w, r, nil, map[string]interface{}{"data": "OK"})
}

/********************************
    Middleware
********************************/

func (con *BaseController) BaseMiddleware(keySet jwk.Set, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		con.corsMiddleware(w, r)
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

func (con *BaseController) corsMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", config.FRONT_URL)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie, Id-Token, Refresh-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Max-Age", "3600")
}

func (con *BaseController) PostOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
		} else if r.Method == http.MethodOptions {
			response(w, r, nil, nil)
			return
		} else {
			response(w, r, perr.New("", perr.MethodNotAllowed), nil)
			return
		}
	})
}

func (con *BaseController) PutOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			next.ServeHTTP(w, r)
		} else if r.Method == http.MethodOptions {
			response(w, r, nil, nil)
			return
		} else {
			response(w, r, perr.New("", perr.MethodNotAllowed), nil)
			return
		}
	})
}

func (con *BaseController) DeleteOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			next.ServeHTTP(w, r)
		} else if r.Method == http.MethodOptions {
			response(w, r, nil, nil)
			return
		} else {
			response(w, r, perr.New("", perr.MethodNotAllowed), nil)
			return
		}
	})
}

func (con *BaseController) GetOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
		} else if r.Method == http.MethodOptions {
			response(w, r, nil, nil)
			return
		} else {
			response(w, r, perr.New("", perr.MethodNotAllowed), nil)
			return
		}
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
				response(w, r, perr.Wrap(err, perr.Expired), nil)
			default:
				response(w, r, perr.Wrap(err, perr.Forbidden), nil)
			}
			return
		}

		user, err := con.ji.GetUserByJwt(r.Context(), cookie.Value)
		if err != nil {
			response(w, r, perr.Wrap(err, perr.Forbidden), nil)
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
