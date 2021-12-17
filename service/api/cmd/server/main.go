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

	kv := controllers.NewKvController(sql)
	project := controllers.NewProjectController(sql)

	middlewareMap["login"] = base.LoginRequiredMiddleware
	middlewareMap["user"] = base.GiveUserMiddleware

	// no middleware
	sv(nil, []pmf{s("/", anyMethod, base.NotFoundView)})

	// get user from ctx
	sv([]string{"user"}, []pmf{s("/test/user", anyMethod, base.UserTestView)})

	// login required
	sv([]string{"login"},
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
		},
	)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}
}

func sv(middlewares []string, ps []pmf) {
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
		mmF := base.AnyMethodMiddleware
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
		default:
			mmF = base.GetOnlyMiddleware
		}

		f := middlewareMap[middlewares[countMiddleware-1]](http.HandlerFunc(p.Func))
		for i := countMiddleware - 2; i >= 0; i-- {
			f = middlewareMap[middlewares[i]](f)
		}

		router.Handle(p.Path, base.BaseMiddleware(keySet, mmF(f)))
	}
	return
}

func s(path string, method string, fun func(http.ResponseWriter, *http.Request)) pmf {
	return pmf{Path: path, Method: method, Func: fun}
}
