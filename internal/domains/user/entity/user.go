package entity

import (
	"golang-project-template/internal/domains/group/entity"
)

type User struct {
	Id                int    `json:"id,omitempty"`
	Role              string `json:"role,omitempty"`
	Email             string `json:"email,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	PasswordEncrypted string `json:"-"`
	Group             *entity.Group
}
