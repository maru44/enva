package controllers

import "net/http"

type (
	ConName string

	Controller interface {
		Name() ConName
	}

	// Controllers struct {
	// 	Base    *BaseController
	// 	Kv      *KvController
	// 	Project *ProjectController
	// 	Org     *OrgController
	// 	User    *UserController
	// 	Health  *HealthController
	// 	// for cli
	// 	CliKv      *CliKvController
	// 	CliProject *CliProjectController
	// 	CliUser    *CliUserController
	// }

	Server struct {
		Mux         *http.ServeMux
		Routes      []Route
		Controllers []Controller
	}

	Route struct {
		Path               string
		Method             string
		Func               func(http.ResponseWriter, *http.Request)
		InitialMiddlewares []func(http.Handler) http.Handler
	}

	Middleware string

	MiddlewareMap map[Middleware]func(http.Handler) http.Handler
)

const (
	BaseMiddleware     = Middleware("base")
	LoginMiddleware    = Middleware("login")
	UserMiddleware     = Middleware("user")
	CliBaseMiddleware  = Middleware("cli")
	CliLoginMiddleware = Middleware("cliLogin")

	ConKv      = ConName("KV")
	ConProject = ConName("PROJECT")

	AnyMethod    = "ANY"
	UpsertMethod = "UPSERT"
)

// func NewServer()

func GenMiddlewares(base *BaseController, cliU *CliUserController) MiddlewareMap {
	if base == nil {
		return MiddlewareMap{
			CliBaseMiddleware:  cliU.BaseMiddleware,
			CliLoginMiddleware: cliU.LoginRequiredMiddleware,
		}
	}
	if cliU == nil {
		return MiddlewareMap{
			BaseMiddleware:  base.BaseMiddleware,
			LoginMiddleware: base.LoginRequiredMiddleware,
			UserMiddleware:  base.GiveUserMiddleware,
		}
	}

	return MiddlewareMap{
		BaseMiddleware:  base.BaseMiddleware,
		LoginMiddleware: base.LoginRequiredMiddleware,
		UserMiddleware:  base.GiveUserMiddleware,
		// for cli
		CliBaseMiddleware:  cliU.BaseMiddleware,
		CliLoginMiddleware: cliU.LoginRequiredMiddleware,
	}
}

// func (s *Server) Sv(routes []Route, preMiddlewares []Middleware, middlewares ...Middleware) {
// 	if len(middlewares) == 0 {{
// 		for _, r := range routes {
// 			s.Mux.Handle(r.Path, )
// 		}
// 	}}
// }

func S(path string, method string, fun func(http.ResponseWriter, *http.Request), ms []Middleware, middlewareMap MiddlewareMap) Route {
	initialMiddlewares := make([]func(http.Handler) http.Handler, len(ms))
	for i, m := range ms {
		initialMiddlewares[i] = middlewareMap[m]
	}

	return Route{
		Path:   path,
		Method: method,
		Func:   fun,
		// InitialMiddlewares: initialMiddlewares,
	}
}
