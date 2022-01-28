package service

import (
	"context"
	"golang-project-template/internal/domains/group/entity"
)

type (
	Repository interface {
		GetAll(ctx context.Context, limit, offset int) ([]*entity.Group, error)
		GetByUUID(ctx context.Context, uuid string) (*entity.Group, error)
		CreateGroup(ctx context.Context, user *entity.Group) error
		UpdateGroup(ctx context.Context, uuid string) (*entity.Group, error)
	}
)
