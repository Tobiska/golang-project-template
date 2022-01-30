package service

import (
	group "golang-project-template/internal/domains/group/entity"
	user "golang-project-template/internal/domains/user/entity"
)

type (
	SignInDTO struct {
		Email    string
		Password string
	}

	CreateDTO struct {
		Username string
		Password string
		Email    string
		Role     string
	}

	UpdateDTO struct {
		Username string
		Email    string
		Group    *group.Group
	}

	AttachGroupDTO struct {
		User  *user.User
		Group *group.Group
	}
)
