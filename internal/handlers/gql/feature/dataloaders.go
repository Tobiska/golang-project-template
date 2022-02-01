package feature

import (
	"golang-project-template/internal/domains"
	"golang-project-template/internal/handlers/gql/feature/user"
)

type DataLoaders struct {
	UserDataloader *user.UserLoader
	//GroupDataloader *group.
}

const (
	userDataloaderMaxBatch int = 10
)

func NewDataLoaders(services *domains.Services) *DataLoaders {
	return &DataLoaders{
		UserDataloader: user.NewConfiguredUserLoader(services.User, userDataloaderMaxBatch),
	}
}
