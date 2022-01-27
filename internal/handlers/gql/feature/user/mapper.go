package user

import (
	"golang-project-template/internal/domains/user/entity"
	"golang-project-template/internal/handlers/gql/model"
)

func MapOneTOGqlModel(user entity.User) *model.User {
	return &model.User{
		ID:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}
