package resolver

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/handlers/gql/middleware"
	"golang-project-template/internal/handlers/gql/runtime"
	"golang-project-template/pkg/auth"
)

type Router struct {
	resolver *Resolver
}

func (r *Router) Register(router *gin.Engine, manager *auth.Manager) {
	router.Use(middleware.AuthenticateUser(r.resolver.Env.UserService, manager))
	router.GET("/", r.playgroundHandler())
	router.POST("/query", r.graphqlHandler())
}

func NewRouter(resolver *Resolver) *Router {
	return &Router{
		resolver: resolver,
	}
}

func (r *Router) graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(runtime.NewExecutableSchema(runtime.Config{Resolvers: r.resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func (r *Router) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
