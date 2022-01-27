package resolver

import (
	"golang-project-template/internal/composites"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Env *composites.Env
}

func NewResolver(env *composites.Env) *Resolver {
	return &Resolver{
		Env: env,
	}
}
