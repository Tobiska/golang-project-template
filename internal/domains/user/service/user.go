package service

import (
	"context"
	"golang-project-template/internal/domains/user/entity"
)

type Service struct {
	repository *Repository
}

func New(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	panic("implement me!!!")
}

func (s *Service) GetById(ctx context.Context, id int) (*entity.User, error) {
	panic("implement me!!!")
}

func (s *Service) CreateUser(ctx context.Context, dto CreateDTO) (*entity.User, error) {
	panic("implement me!!!")
}

func (s *Service) UpdateUser(ctx context.Context, dto UpdateDTO) (*entity.User, error) {
	panic("implement me!!!")
}
