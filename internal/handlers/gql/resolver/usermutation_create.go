package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"golang-project-template/internal/domains/apperror"
	"golang-project-template/internal/domains/user/service"
	"golang-project-template/internal/handlers/gql/feature/user"
	"golang-project-template/internal/handlers/gql/model"
)

func (r *userMutationResolver) Create(ctx context.Context, obj *model.UserMutation, input model.UserCreateInput) (model.UserCreateResult, error) {
	dto := service.CreateDTO{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Role:     input.Role.String(),
	}
	u, err := r.Env.Services.User.CreateUser(ctx, dto)

	if errors.Is(err, apperror.InternalError) {
		return model.InternalErrorProblem{
			Message: err.Error(),
		}, nil
	}

	if errors.Is(err, service.UserValidationError) {
		return model.ValidationErrorProblem{
			Message: err.Error(),
		}, nil
	}

	if err != nil {
		return nil, err
	}

	um := user.MapOneTOGqlModel(*u)

	return model.UserCreateOk{
		User: um,
	}, nil
}
