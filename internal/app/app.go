package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-project-template/config"
	"golang-project-template/internal/domains"
	res "golang-project-template/internal/handlers/gql/resolver"
	userRepo "golang-project-template/internal/infrastructure/user/repository/sql"
	"golang-project-template/pkg/auth"
	"golang-project-template/pkg/db/postgres"
	"golang-project-template/pkg/httpserver"
	"log"
)

import (
	groupServ "golang-project-template/internal/domains/group/service"
	groupRepo "golang-project-template/internal/infrastructure/group/repository/sql"
)

import (
	userServ "golang-project-template/internal/domains/user/service"
)

const (
	signingKey = "demoApiQAZWSXEDC" //TODO FIX move to env variables
)

func Run(cfg *config.Config) {
	//TODO add logger

	router := gin.Default()

	jwtTokenManager, err := auth.NewManager(signingKey, cfg)
	if err != nil {
		log.Fatal(err)
	}

	postgreSQLClient, err := postgres.NewClient(context.TODO(), cfg.PG.AttemptToConnect, *cfg)
	if err != nil {
		log.Fatal(err)
	}

	var env *domains.Env //TODO think about it!!!

	//User

	userRepository := userRepo.NewRepository(postgreSQLClient)
	userService := userServ.New(userRepository, jwtTokenManager)

	//Group

	groupRepository := groupRepo.NewRepository(postgreSQLClient)
	groupService := groupServ.NewGroupService(groupRepository, userService)

	//Session

	//Task
	env = domains.NewEnv(userService, groupService)
	resolver := res.NewResolver(env)

	gqlRouter := res.NewRouter(resolver)
	gqlRouter.Register(router, jwtTokenManager)

	srv := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	select {
	case n := <-srv.Notify():
		log.Println(n)
	}

	if err := srv.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
