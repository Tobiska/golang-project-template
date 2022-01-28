package domains

import (
	groupServ "golang-project-template/internal/domains/group/service"
	userServ "golang-project-template/internal/domains/user/service"
)

type Env struct {
	UserService  *userServ.Service
	GroupService *groupServ.Service
}

func NewEnv(us *userServ.Service, gs *groupServ.Service) *Env {
	return &Env{
		UserService:  us,
		GroupService: gs,
	}
}
