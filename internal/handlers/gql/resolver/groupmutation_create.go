package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	groupServ "golang-project-template/internal/domains/group/service"
	"golang-project-template/internal/domains/user/entity"
	groupFeature "golang-project-template/internal/handlers/gql/feature/group"
	"golang-project-template/internal/handlers/gql/middleware"
	"golang-project-template/internal/handlers/gql/model"
)

func (r *groupMutationResolver) Create(ctx context.Context, obj *model.GroupMutation, input model.GroupCreateInput) (model.GroupCreateResult, error) {
	u := ctx.Value(middleware.CtxKeyUser).(*entity.User)
	dto := groupServ.CreateDTO{
		Name:       input.Name,
		GroupOwner: u,
	}
	group, err := r.Env.Services.Group.Create(ctx, dto)
	fmt.Println("Group: ", group)
	if err != nil {
		return nil, err
	}

	return model.GroupCreateOk{
		Group: groupFeature.MapOneTOGqlModel(*group),
	}, nil
}
