package directives

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"golang-project-template/internal/domains/user/entity"
	"golang-project-template/internal/handlers/gql/middleware"
	"golang-project-template/internal/handlers/gql/model"
)

type HasRoleDirective = func(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
	role *model.Role,
) (res interface{}, err error)

func NewHasRoleDirective() HasRoleDirective {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role *model.Role) (res interface{}, err error) {
		ctxUser := ctx.Value(middleware.CtxKeyUser).(*entity.User)
		if ctxUser != nil || ctxUser.Role == obj {
			return next(ctx)
		} else {
			return nil, errors.New("Permission denied") //TODO ADD CUSTOM ERROR
		}
	}
}
