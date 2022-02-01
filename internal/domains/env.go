package domains

import (
	groupServ "golang-project-template/internal/domains/group/service"
	userServ "golang-project-template/internal/domains/user/service"
)

type Env struct {
	Services *Services
}

type Services struct {
	User  *userServ.Service
	Group *groupServ.Service
}

func NewEnv(us *userServ.Service, gs *groupServ.Service) *Env {
	return &Env{
		Services: &Services{
			User:  us,
			Group: gs,
		},
	}
}
