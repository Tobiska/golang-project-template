package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	groupFeature "golang-project-template/internal/handlers/gql/feature/group"
	userFeature "golang-project-template/internal/handlers/gql/feature/user"
	"golang-project-template/internal/handlers/gql/model"
	"golang-project-template/internal/handlers/gql/runtime"
)

func (r *groupFindOkResolver) Users(ctx context.Context, obj *model.GroupFindOk) ([]*model.User, error) {
	users, err := r.Env.Services.User.GetUsersByGroupId(ctx, obj.Group.UUID)
	ums := make([]*model.User, len(users))
	for id, user := range users {
		ums[id] = userFeature.MapOneTOGqlModel(*user)
	}
	if err != nil {
		return nil, err
	}
	return ums, nil
}

func (r *groupQueryResolver) FindByUUID(ctx context.Context, obj *model.GroupQuery, uuid string) (model.GroupFindResult, error) {
	group, err := r.Env.Services.Group.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, nil //TODO add custom variables
	}
	gm := groupFeature.MapOneTOGqlModel(*group)
	return &model.GroupFindOk{
		Group: gm,
	}, nil
}

// GroupFindOk returns runtime.GroupFindOkResolver implementation.
func (r *Resolver) GroupFindOk() runtime.GroupFindOkResolver { return &groupFindOkResolver{r} }

type groupFindOkResolver struct{ *Resolver }
