package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-project-template/internal/handlers/gql/feature/user"
	"golang-project-template/internal/handlers/gql/model"
)

func (r *userQueryResolver) FindByID(ctx context.Context, obj *model.UserQuery, id int) (*model.User, error) {
	u, err := r.Env.UserService.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	um := user.MapOneTOGqlModel(*u)

	return um, nil
}
