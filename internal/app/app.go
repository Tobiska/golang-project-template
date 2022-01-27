package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-project-template/config"
	userService "golang-project-template/internal/domains/user/service"
	"golang-project-template/internal/handlers/gql"
	userRepo "golang-project-template/internal/infrastructure/user/repository/sql"
	"golang-project-template/pkg/db/postgres"
	"golang-project-template/pkg/httpserver"
	"log"
)

func Run(cfg *config.Config) {
	//TODO add logger

	router := gin.Default()

	postgreSQLClient, err := postgres.NewClient(context.TODO(), cfg.AttemptToConnect, *cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := userRepo.NewRepository(postgreSQLClient)
	userService.New(userRepository)

	gqlRouter := gql.NewRouter()
	gqlRouter.Register(router)

	srv := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	select {
	case n := <-srv.Notify():
		log.Println(n)
	}

	if err := srv.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
