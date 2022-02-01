package entity

import "golang-project-template/internal/domains/user/entity"

type Group struct {
	Uuid    string       `json:"uuid"`
	Name    string       `json:"name"`
	Owner   *entity.User `json:"owner"`
	OwnerId int          `json:"owner_id"`
}
