package service

import (
	"context"
	"golang-project-template/internal/domains/group/entity"
	"golang-project-template/internal/infrastructure"
)

type (
	Repository interface {
		GetAll(ctx context.Context, opt ...infrastructure.Option) ([]*entity.Group, error)
		GetByUUID(ctx context.Context, uuid string) (*entity.Group, error)
		CreateGroup(ctx context.Context, user *entity.Group) error
		UpdateGroup(ctx context.Context, uuid string) (*entity.Group, error)
	}
)
