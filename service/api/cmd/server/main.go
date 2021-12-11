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
	post := controllers.NewKvController(sql)

	keySet, err := base.GetKeySet()
	if err != nil {
		panic(err)
	}

	router.Handle("/", base.BaseMiddleware(keySet, http.HandlerFunc(base.NotFoundView)))
	router.Handle("/test/user/", base.BaseMiddleware(keySet, base.GiveUserMiddleware(http.HandlerFunc(base.UserTestView))))

	router.Handle("/post/", base.BaseMiddleware(keySet, base.GiveUserMiddleware(http.HandlerFunc(post.ListView))))
	router.Handle("/post/create/", base.BaseMiddleware(keySet, base.PostOnlyMiddleware(base.LoginRequiredMiddleware(http.HandlerFunc(post.CreateView)))))

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}
}
