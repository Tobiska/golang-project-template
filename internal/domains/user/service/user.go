package service

import (
	"context"
	"golang-project-template/internal/domains/user/entity"
)

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	users, err := s.repository.GetAll(ctx, limit, offset)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *Service) GetById(ctx context.Context, id int) (*entity.User, error) {
	u, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) CreateUser(ctx context.Context, dto CreateDTO) (*entity.User, error) {
	u := &entity.User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
	}
	if err := s.repository.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) UpdateUser(ctx context.Context, dto UpdateDTO) (*entity.User, error) {
	panic("implement me!!!")
}
