package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/plugin/federation/testdata/entityresolver/generated"
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/handlers/gql/resolver"
)

type Router struct {
	resolver *resolver.Resolver
}

func (r *Router) Register(router *gin.Engine) {
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler())
}

func NewRouter(resolver *resolver.Resolver) *Router {
	return &Router{
		resolver: resolver,
	}
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
