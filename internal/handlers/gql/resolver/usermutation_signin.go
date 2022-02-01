package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-project-template/internal/domains/user/service"
	"golang-project-template/internal/handlers/gql/feature/user"
	"golang-project-template/internal/handlers/gql/model"
)

func (r *userMutationResolver) SignIn(ctx context.Context, obj *model.UserMutation, input model.UserSignInInput) (model.UserSignInResult, error) {
	dto := service.SignInDTO{
		Email:    input.Email,
		Password: input.Password,
	}
	u, token, err := r.Env.Services.User.SignIn(ctx, dto)
	if err != nil {
		return nil, err //TODO add custom errors
	}

	um := user.MapOneTOGqlModel(*u)
	return model.UserSignOk{
		User: um,
		Token: &model.Token{
			AccessToken: token,
		},
	}, nil
}
