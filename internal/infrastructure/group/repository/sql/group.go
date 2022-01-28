package sql

import "golang-project-template/pkg/db/postgres"

type Repository struct {
	client postgres.Client
}

func NewRepository(client postgres.Client) *Repository {
	return &Repository{
		client: client,
	}
}
