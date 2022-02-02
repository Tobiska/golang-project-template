package service

import (
	"context"
	"golang-project-template/internal/domains/group/entity"
	userServ "golang-project-template/internal/domains/user/service"
)

type Service struct {
	repository  Repository
	userService *userServ.Service
}

func NewGroupService(repository Repository, service *userServ.Service) *Service {
	return &Service{
		repository:  repository,
		userService: service,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]*entity.Group, error) {
	groups, err := s.repository.GetAll(ctx)
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
		Name:    dto.Name,
		OwnerId: dto.GroupOwner.Id,
	}
	if err := s.repository.CreateGroup(ctx, g); err != nil {
		return nil, err
	}
	dto.GroupOwner.GroupID = g.Uuid
	if err := s.userService.AttachGroup(ctx, dto.GroupOwner); err != nil {
		return nil, err
	}

	return g, nil
}
