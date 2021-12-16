package main

import (
	"log"
	"net/http"

	"github.com/maru44/ichigo/service/api/internal/infra"
	"github.com/maru44/ichigo/service/api/internal/interface/controllers"
)

type pmf struct {
	Path   string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

var (
	router        = http.NewServeMux()
	middlewareMap = map[string]func(http.Handler) http.Handler{}
)

func main() {
	jp := infra.NewJwtParser()
	sql := infra.NewSqlHandler()

	base := controllers.NewBaseController(jp)
	kv := controllers.NewKvController(sql)
	project := controllers.NewProjectController(sql)

	keySet, err := base.GetKeySet()
	if err != nil {
		panic(err)
	}

	middlewareMap["login"] = base.LoginRequiredMiddleware
	middlewareMap["user"] = base.GiveUserMiddleware

	router.Handle("/", base.BaseMiddleware(keySet, http.HandlerFunc(base.NotFoundView)))
	router.Handle("/test/user/", base.BaseMiddleware(keySet, base.GiveUserMiddleware(http.HandlerFunc(base.UserTestView))))

	/* key value */
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

func serving(path string, middlewares []func(http.Handler) http.Handler, conFunc func(http.ResponseWriter, *http.Request)) {
	jp := infra.NewJwtParser()

	base := controllers.NewBaseController(jp)

	keySet, err := base.GetKeySet()
	if err != nil {
		panic(err)
	}

	// if no middlewares
	if middlewares == nil {
		router.Handle(path, base.BaseMiddleware(keySet, http.HandlerFunc(conFunc)))
	}

	countMiddleware := len(middlewares)
	f := middlewares[countMiddleware-1](http.HandlerFunc(conFunc))

	for i := countMiddleware - 2; i >= 0; i-- {
		f = middlewares[i](f)
	}
	router.Handle(path, base.BaseMiddleware(keySet, f))
}

func sv(middlewares []string, ps []pmf) {
	jp := infra.NewJwtParser()
	base := controllers.NewBaseController(jp)
	keySet, err := base.GetKeySet()
	if err != nil {
		panic(err)
	}

	if middlewares == nil {
		for _, p := range ps {
			router.Handle(p.Path, base.BaseMiddleware(keySet, http.HandlerFunc(p.Func)))
		}
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
		case "UPSERT":
			// mm = base.Upsert
		case "Any":
			// mm = base.
		default:
		}

		f := middlewareMap[middlewares[countMiddleware-1]](http.HandlerFunc(p.Func))
		for i := countMiddleware - 2; i >= 0; i-- {
			f = middlewareMap[middlewares[i]](f)
		}

		router.Handle(p.Path, base.BaseMiddleware(keySet, mmF(f)))
	}
}

func s(path string, method string, fun func(http.ResponseWriter, *http.Request)) pmf {
	return pmf{Path: path, Method: method, Func: fun}
}
