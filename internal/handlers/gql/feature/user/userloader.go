package user

import (
	"context"
	"golang-project-template/internal/domains/user/entity"
	userServ "golang-project-template/internal/domains/user/service"
	"time"
)

func NewConfiguredUserLoader(service *userServ.Service, maxBatch int) *UserLoader {
	return NewUserLoader(UserLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []int) ([]*entity.User, []error) {
			errors := make([]error, len(keys))
			users, err := service.GetByIds(context.Background(), keys...)

			for id, _ := range users {
				errors[id] = err
			}

			return users, errors
		},
	})
}
