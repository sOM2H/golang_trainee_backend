package main

import (
	"github.com/sOM2H/golang_trainee_backend/config"
	"github.com/sOM2H/golang_trainee_backend/db"
	"github.com/sOM2H/golang_trainee_backend/handler"
	"github.com/sOM2H/golang_trainee_backend/router"
	"github.com/sOM2H/golang_trainee_backend/store"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	appConfig := config.New("")
	r := router.New()

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")

	d := db.New(appConfig.DSN)
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	ps := store.NewPostStore(d)
	cs := store.NewCommentStore(d)

	h := handler.NewHandler(us, ps, cs)
	h.Register(v1)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
