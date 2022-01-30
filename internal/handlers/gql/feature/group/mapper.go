package group

import (
	"golang-project-template/internal/domains/group/entity"
	"golang-project-template/internal/handlers/gql/model"
)

func MapOneTOGqlModel(group entity.Group) *model.Group {
	return &model.Group{
		UUID:    group.Uuid,
		Name:    group.Name,
		OwnerID: group.OwnerId,
	}
}
