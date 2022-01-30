package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	groupFeature "golang-project-template/internal/handlers/gql/feature/group"
	"golang-project-template/internal/handlers/gql/model"
	"golang-project-template/internal/handlers/gql/runtime"
)

func (r *queryResolver) Group(ctx context.Context) (*model.GroupQuery, error) {
	return &model.GroupQuery{}, nil
}

func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	all, err := r.Env.Services.Group.GetAll(ctx) //TODO ADD filters and pagination with cond
	if err != nil {
		return nil, err
	}
	gms := make([]*model.Group, len(all))
	for id, group := range all {
		gms[id] = groupFeature.MapOneTOGqlModel(*group)
	}
	return gms, nil
}

// GroupQuery returns runtime.GroupQueryResolver implementation.
func (r *Resolver) GroupQuery() runtime.GroupQueryResolver { return &groupQueryResolver{r} }

type groupQueryResolver struct{ *Resolver }
