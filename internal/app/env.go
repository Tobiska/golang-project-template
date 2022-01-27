package app

import (
	userServ "golang-project-template/internal/domains/user/service"
)

type Env struct {
	//TODO ADD COMPOSITE
	UserService *userServ.Service
}

func NewEnv(us *userServ.Service) *Env {
	return &Env{
		UserService: us,
	}
}
