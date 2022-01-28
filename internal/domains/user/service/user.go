package service

import (
	"context"
	"errors"
	"golang-project-template/internal/domains/user/entity"
	"golang-project-template/pkg/auth"
	"strconv"
)

type Service struct {
	repository   Repository
	tokenManager auth.TokenManager
}

func New(repository Repository, manager auth.TokenManager) *Service {
	return &Service{
		repository:   repository,
		tokenManager: manager,
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

func (s *Service) SignIn(ctx context.Context, dto SignInDTO) (*entity.User, string, error) {
	u, err := s.repository.GetByEmail(ctx, dto.Username)
	if err != nil {
		return nil, "", err
	}

	if res := u.CheckPassword(dto.Password); res == true {
		return nil, "", errors.New("Login or Password invalid")
	}

	jwt, err := s.tokenManager.NewJWT(strconv.Itoa(u.Id))
	if err != nil {
		return nil, "", err
	}

	return u, jwt, nil
}
