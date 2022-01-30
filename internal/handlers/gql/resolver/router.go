package resolver

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/handlers/gql/directives"
	"golang-project-template/internal/handlers/gql/middleware"
	"golang-project-template/internal/handlers/gql/runtime"
	"golang-project-template/pkg/auth"
)

type Router struct {
	resolver *Resolver
}

func (r *Router) Register(router *gin.Engine, manager *auth.Manager) {
	router.Use(middleware.DataLoadersInjector(r.resolver.Env))
	router.Use(middleware.AuthenticateUser(r.resolver.Env.Services.User, manager))
	router.GET("/", r.playgroundHandler())
	router.POST("/query", r.graphqlHandler())
}

func NewRouter(resolver *Resolver) *Router {
	return &Router{
		resolver: resolver,
	}
}

func (r *Router) graphqlHandler() gin.HandlerFunc {
	cfg := newSchemaConfig(r.resolver)
	h := handler.NewDefaultServer(runtime.NewExecutableSchema(cfg))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func newSchemaConfig(resolver *Resolver) runtime.Config {
	cfg := runtime.Config{Resolvers: resolver}

	cfg.Directives.IsAuthJWT = directives.NewAuthJWTDirective()
	cfg.Directives.HasRole = directives.NewHasRoleDirective()
	return cfg
}

func (r *Router) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
