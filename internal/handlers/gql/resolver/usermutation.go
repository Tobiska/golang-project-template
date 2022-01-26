package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-project-template/internal/handlers/gql/model"
	"golang-project-template/internal/handlers/gql/runtime"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserMutation, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserMutation returns runtime.UserMutationResolver implementation.
func (r *Resolver) UserMutation() runtime.UserMutationResolver { return &userMutationResolver{r} }

type userMutationResolver struct{ *Resolver }
