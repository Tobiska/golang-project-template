package sql

import (
	"context"
	"golang-project-template/internal/domains/group/entity"
	"golang-project-template/pkg/db/postgres"
)

type Repository struct {
	client postgres.Client
}

func NewRepository(client postgres.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetAll(ctx context.Context, limit, offset int) ([]*entity.Group, error) {
	panic("implement me!!!")
}

func (r *Repository) GetByUUID(ctx context.Context, uuid string) (*entity.Group, error) {
	panic("implement me!!!")
}

func (r *Repository) CreateGroup(ctx context.Context, user *entity.Group) error {
	panic("implement me!!!")
}

func (r *Repository) UpdateGroup(ctx context.Context, uuid string) (*entity.Group, error) {
	panic("implement me!!!")
}
