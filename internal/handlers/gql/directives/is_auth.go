package directives

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"golang-project-template/internal/handlers/gql/middleware"
)

type AuthJWTDirectiveFunc = func(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver, //?
) (res interface{}, err error)

func NewAuthJWTDirective() AuthJWTDirectiveFunc {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ctxUser := ctx.Value(middleware.CtxKeyUser)
		if ctxUser != nil {
			return next(ctx)
		} else {
			return nil, errors.New("UnauthorizedError") //TODO ADD CUSTOM ERROR
		}
	}
}
