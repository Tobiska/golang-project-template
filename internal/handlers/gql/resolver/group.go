package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-project-template/internal/handlers/gql/feature"
	"golang-project-template/internal/handlers/gql/feature/user"
	"golang-project-template/internal/handlers/gql/middleware"
	"golang-project-template/internal/handlers/gql/model"
	"golang-project-template/internal/handlers/gql/runtime"
)

func (r *groupResolver) Owner(ctx context.Context, obj *model.Group) (*model.User, error) {
	dataloader := ctx.Value(middleware.CtxKeyDataloader).(*feature.DataLoaders).UserDataloader
	owner, err := dataloader.Load(obj.OwnerID)
	if err != nil {
		return nil, err
	}

	return user.MapOneTOGqlModel(*owner), nil
}

// Group returns runtime.GroupResolver implementation.
func (r *Resolver) Group() runtime.GroupResolver { return &groupResolver{r} }

type groupResolver struct{ *Resolver }
