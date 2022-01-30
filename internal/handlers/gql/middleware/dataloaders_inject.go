package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/domains"
	"golang-project-template/internal/handlers/gql/feature"
)

func DataLoadersInjector(env *domains.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), CtxKeyDataloader, feature.NewDataLoaders(env.Services))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
