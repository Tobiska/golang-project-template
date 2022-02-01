package sql

import (
	"context"
	"github.com/google/uuid"
	"golang-project-template/internal/domains/group/entity"
	"golang-project-template/pkg/db/postgres"
)

type Repository struct {
	client postgres.Client
}

func NewRepository(client postgres.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]*entity.Group, error) {
	groups := make([]*entity.Group, 0, 10) //TODO optimize
	rows, err := r.client.Query(ctx, `SELECT groups.uuid, groups.name, groups.owner_id FROM groups`)
	if err != nil {
		return nil, err
	}
	it := 0
	for rows.Next() {
		g := &entity.Group{}
		if err := rows.Scan(&g.Uuid, &g.Name, &g.OwnerId); err != nil {
			return nil, err
		}
		groups = append(groups, g)
		it++
	}

	return groups, nil
}

func (r *Repository) GetByUUID(ctx context.Context, uuid string) (*entity.Group, error) {
	g := &entity.Group{}
	row := r.client.QueryRow(ctx, `SELECT groups.uuid, groups.name, groups.owner_id FROM groups WHERE groups.uuid=$1`, uuid)
	if err := row.Scan(&g.Uuid, &g.Name, &g.OwnerId); err != nil {
		return nil, err
	}
	return g, nil
}

func (r *Repository) CreateGroup(ctx context.Context, group *entity.Group) error {
	//TODO move uuid gen to db
	q := `INSERT INTO groups 
					(uuid , name, owner_id)
		  VALUES 
					($1, $2, $3)
		  RETURNING Uuid`
	if err := r.client.QueryRow(ctx, q, uuid.New(), group.Name, group.Owner.Id).Scan(&group.Uuid); err != nil {
		return err //todo add custom errors
	}
	return nil
}

func (r *Repository) UpdateGroup(ctx context.Context, uuid string) (*entity.Group, error) {
	panic("implement me!!!")
}
