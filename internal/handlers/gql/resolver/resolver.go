package resolver

import (
	"golang-project-template/internal/domains"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Env *domains.Env
}

func NewResolver(env *domains.Env) *Resolver {
	return &Resolver{
		Env: env,
	}
}
