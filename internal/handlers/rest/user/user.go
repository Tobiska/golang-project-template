package user

import (
	"golang-project-template/internal/domains/user/service"
)

type Handler struct {
	service service.Service
}
