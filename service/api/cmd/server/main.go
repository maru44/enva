package main

import (
	"log"
	"net/http"

	"github.com/maru44/ichigo/service/api/internal/infra"
	"github.com/maru44/ichigo/service/api/internal/interface/controllers"
)

func main() {
	router := http.NewServeMux()

	jp := infra.NewJwtParser()
	sql := infra.NewSqlHandler()

	base := controllers.NewBaseController(jp)
	kv := controllers.NewKvController(sql)
	project := controllers.NewProjectController(sql)

	keySet, err := base.GetKeySet()
	if err != nil {
		panic(err)
	}

	router.Handle("/", base.BaseMiddleware(keySet, http.HandlerFunc(base.NotFoundView)))
	router.Handle("/test/user/", base.BaseMiddleware(keySet, base.GiveUserMiddleware(http.HandlerFunc(base.UserTestView))))

	/* key value */
	router.Handle("/kv", base.BaseMiddleware(keySet, base.GiveUserMiddleware(http.HandlerFunc(kv.ListView))))
	router.Handle("/kv/create", base.BaseMiddleware(keySet, base.PostOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(kv.CreateView)))))
	router.Handle("/kv/update", base.BaseMiddleware(keySet, base.PutOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(kv.UpdateView)))))

	/* projects */
	router.Handle("/project/list/user", base.BaseMiddleware(keySet, base.GetOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(project.ListByUserView)))))
	router.Handle("/project/list/org", base.BaseMiddleware(keySet, base.LoginRequiredMiddleware(http.HandlerFunc(project.ListByProjectView))))
	router.Handle("/project/slugs/user", base.BaseMiddleware(keySet, base.GetOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(project.SlugListByUserView)))))
	router.Handle("/project/detail", base.BaseMiddleware(keySet, base.GetOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(project.ProjectDetailView)))))
	router.Handle("/project/create", base.BaseMiddleware(keySet, base.PostOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(project.CreateView)))))

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}
}
