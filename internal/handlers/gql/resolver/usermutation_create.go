package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-project-template/internal/domains/user/service"
	"golang-project-template/internal/handlers/gql/feature/user"
	"golang-project-template/internal/handlers/gql/model"
)

func (r *userMutationResolver) Create(ctx context.Context, obj *model.UserMutation, input model.UserCreateInput) (model.UserCreateResult, error) {
	dto := service.CreateDTO{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}
	u, err := r.Env.UserService.CreateUser(ctx, dto)
	if err != nil {
		return nil, err
	}

	um := user.MapOneTOGqlModel(*u)

	return model.UserCreateOk{
		User: um,
	}, nil
}
