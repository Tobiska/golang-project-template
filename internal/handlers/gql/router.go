package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/plugin/federation/testdata/entityresolver/generated"
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/handlers/gql/resolver"
)

type Router struct{}

func (r *Router) Register(router *gin.Engine) {
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler())
}

func NewRouter() *Router {
	return &Router{}
}

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
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
