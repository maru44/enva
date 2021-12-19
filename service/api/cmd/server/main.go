package main

import (
	"log"
	"net/http"

	"github.com/maru44/enva/service/api/internal/infra"
	"github.com/maru44/enva/service/api/internal/interface/controllers"
)

type pmf struct {
	Path   string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

const (
	anyMethod    = "ANY"
	upsertMethod = "UPSERT"
)

var (
	router        = http.NewServeMux()
	middlewareMap = map[string]func(http.Handler) http.Handler{}

	jp   = infra.NewJwtParser()
	base = controllers.NewBaseController(jp)
)

func main() {
	sql := infra.NewSqlHandler()
	pass := &infra.Password{}

	kv := controllers.NewKvController(sql)
	cliKv := controllers.NewCliKvController(sql)
	project := controllers.NewProjectController(sql)
	cliU := controllers.NewCliUserController(sql, pass)

	middlewareMap["login"] = base.LoginRequiredMiddleware
	middlewareMap["user"] = base.GiveUserMiddleware
	middlewareMap["cli"] = cliU.BaseMiddleware
	middlewareMap["loginCli"] = cliU.LoginRequiredMiddleware

	// no middlewares
	sv([]pmf{
		s("/", anyMethod, base.NotFoundView),
		s("/cli/user/validate", http.MethodPost, cliU.ValidateView),
	})

	// get user from ctx
	sv([]pmf{s("/test/user", anyMethod, base.UserTestView)}, "user")

	// login required
	sv(
		[]pmf{
			/* kv */
			s("/kv", http.MethodGet, kv.ListView),
			s("/kv/create", http.MethodPost, kv.CreateView),
			s("/kv/update", http.MethodPut, kv.UpdateView),
			s("/kv/delete", http.MethodDelete, kv.DeleteView),

			/* project */
			s("/project/list/user", http.MethodGet, project.ListByUserView),
			s("/project/list/org", http.MethodGet, project.ListByOrgView),
			s("/project/slugs/user", http.MethodGet, project.SlugListByUserView),
			s("/project/detail", http.MethodGet, project.ProjectDetailView),
			s("/project/create", http.MethodPost, project.CreateView),
			s("/project/delete", http.MethodDelete, project.DeleteView),

			/* cli_users */
			s("/cli/user", http.MethodGet, cliU.ExistsView),
			s("/cli/user/create", http.MethodGet, cliU.CreateView),
			s("/cli/user/update", http.MethodGet, cliU.UpdateView),
		},
		"login",
	)

	svCli(
		[]pmf{
			/* cli_kv */
			s("/cli/kv", http.MethodGet, cliKv.ListView),
			s("/cli/kv/detail", http.MethodGet, cliKv.DetailView),
			s("/cli/kv/create", http.MethodPost, cliKv.CreateView),
			s("/cli/kv/update", http.MethodPut, cliKv.UpdateView),
			s("/cli/kv/delete", http.MethodPost, cliKv.DeleteView),
		},
		"cli", "loginCli",
	)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}
}

func sv(ps []pmf, middlewares ...string) {
	keySet, err := base.GetKeySet()
	if err != nil {
		panic(err)
	}

	if middlewares == nil {
		for _, p := range ps {
			router.Handle(p.Path, base.BaseMiddleware(keySet, http.HandlerFunc(p.Func)))
		}
		return
	}

	countMiddleware := len(middlewares)
	for _, p := range ps {
		mmF := base.GetOnlyMiddleware
		switch p.Method {
		case http.MethodGet:
			mmF = base.GetOnlyMiddleware
		case http.MethodPost:
			mmF = base.PostOnlyMiddleware
		case http.MethodPut:
			mmF = base.PutOnlyMiddleware
		case http.MethodDelete:
			mmF = base.DeleteOnlyMiddleware
		case upsertMethod:
			// mm = base.Upsert
			// @TODO make upsert middleware
		case anyMethod:
			mmF = base.AnyMethodMiddleware
		}

		f := middlewareMap[middlewares[countMiddleware-1]](http.HandlerFunc(p.Func))
		for i := countMiddleware - 2; i >= 0; i-- {
			f = middlewareMap[middlewares[i]](f)
		}

		router.Handle(p.Path, base.BaseMiddleware(keySet, mmF(f)))
	}
	return
}

func svCli(ps []pmf, middlewares ...string) {
	if middlewares == nil {
		for _, p := range ps {
			router.Handle(p.Path, http.HandlerFunc(p.Func))
		}
		return
	}

	countMiddleware := len(middlewares)
	for _, p := range ps {
		mmF := base.GetOnlyMiddleware
		switch p.Method {
		case http.MethodGet:
			mmF = base.GetOnlyMiddleware
		case http.MethodPost:
			mmF = base.PostOnlyMiddleware
		case http.MethodPut:
			mmF = base.PutOnlyMiddleware
		case http.MethodDelete:
			mmF = base.DeleteOnlyMiddleware
		case upsertMethod:
			// mm = base.Upsert
			// @TODO make upsert middleware
		case anyMethod:
			mmF = base.AnyMethodMiddleware
		}

		f := middlewareMap[middlewares[countMiddleware-1]](http.HandlerFunc(p.Func))
		for i := countMiddleware - 2; i >= 0; i-- {
			f = middlewareMap[middlewares[i]](f)
		}

		router.Handle(p.Path, mmF(f))
	}
	return
}

func s(path string, method string, fun func(http.ResponseWriter, *http.Request)) pmf {
	return pmf{Path: path, Method: method, Func: fun}
}
