package service

import (
	"context"
	"golang-project-template/internal/domains/user/entity"
)

type (
	Repository interface {
		GetAll(ctx context.Context, limit, offset int) ([]*entity.User, error)
		GetByIds(ctx context.Context, ids ...int) ([]*entity.User, error)
		GetById(ctx context.Context, id int) (*entity.User, error)
		GetUsersByGroupId(ctx context.Context, uuid string) ([]*entity.User, error)
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		CreateUser(ctx context.Context, user *entity.User) error
		UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	}
)
