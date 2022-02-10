package cache

import "golang-project-template/pkg/cache/redis"

type UserCache struct {
	cacheClient *redis.Client
}

type (
	uc *UserCache
)
