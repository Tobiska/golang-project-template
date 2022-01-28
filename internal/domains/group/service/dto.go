package service

import "golang-project-template/internal/domains/user/entity"

type (
	CreateDTO struct {
		Name       string
		GroupOwner *entity.User
	}
)
