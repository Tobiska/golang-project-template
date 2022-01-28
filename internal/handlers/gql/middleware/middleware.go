package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/domains/user/service"
	"golang-project-template/pkg/auth"
	"net/http"
	"strconv"
	"strings"
)

const (
	ctxKeyUser ctxKey = iota
)

type ctxKey int8

func AuthenticateUser(service *service.Service, manager auth.TokenManager) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		header := c.GetHeader("Authentication")

		headerParts := strings.Split(header, " ")
		if header == "" || c == nil || len(headerParts) > 2 || headerParts[0] != "Bearer" {
			c.Next()
			return
		}

		payload, err := manager.Parse(headerParts[1])
		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid token")
			return
		}

		userId, err := strconv.Atoi(payload)
		if err != nil {
			c.String(http.StatusUnauthorized, "Error convert string to int")
			return
		}

		user, err := service.GetById(c.Request.Context(), userId)

		ctx := context.WithValue(c.Request.Context(), ctxKeyUser, user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})
}
