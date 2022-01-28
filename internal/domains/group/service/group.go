package service

import (
	"context"
	"golang-project-template/internal/domains"
	"golang-project-template/internal/domains/group/entity"
	"golang-project-template/internal/domains/user/service"
)

type Service struct {
	repository Repository
	env        *domains.Env
}

func NewGroupService(repository Repository, env *domains.Env) *Service {
	return &Service{
		repository: repository,
		env:        env,
	}
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) ([]*entity.Group, error) {
	groups, err := s.repository.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *Service) GetByUuid(ctx context.Context, uuid string) (*entity.Group, error) {
	group, err := s.repository.GetByUUID(ctx, uuid)
	if err != nil {
		return group, err
	}

	return group, err
}

func (s *Service) Create(ctx context.Context, dto CreateDTO) (*entity.Group, error) {
	g := &entity.Group{
		Name: dto.Name,
	}
	if err := s.repository.CreateGroup(ctx, g); err != nil {
		return nil, err
	}

	updateUserDto := service.UpdateDTO{
		Username: dto.GroupOwner.Username,
		Group:    g,
		Email:    dto.GroupOwner.Email,
	}
	u, err := s.env.UserService.UpdateUser(ctx, updateUserDto)
	dto.GroupOwner.Id = u.Id
	if err != nil {
		return nil, err
	}

	return g, nil
}
