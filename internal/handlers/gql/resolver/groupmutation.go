package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-project-template/internal/handlers/gql/model"
	"golang-project-template/internal/handlers/gql/runtime"
)

func (r *mutationResolver) Group(ctx context.Context) (*model.GroupMutation, error) {
	return &model.GroupMutation{}, nil
}

// GroupMutation returns runtime.GroupMutationResolver implementation.
func (r *Resolver) GroupMutation() runtime.GroupMutationResolver { return &groupMutationResolver{r} }

type groupMutationResolver struct{ *Resolver }
