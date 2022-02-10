package entity

import "golang-project-template/internal/domains/user/entity"

type Group struct {
	Uuid    string       `json:"uuid" db:"uuid"`
	Name    string       `json:"name" db:"name"`
	Owner   *entity.User `json:"owner" db:"-"`
	OwnerId int          `json:"owner_id" db:"owner_id"`
}
