package sql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"golang-project-template/internal/domains/group/entity"
	"golang-project-template/internal/infrastructure"
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

func (r *Repository) GetAll(ctx context.Context, opts ...infrastructure.Option) ([]*entity.Group, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	var groups []*entity.Group
	qb := psql.Select("groups.uuid, groups.name, groups.owner_id").From("groups")
	for _, opt := range opts {
		qb = opt.Apply(qb)
	}
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, r.client, &groups, sql, args...); err != nil {
		return nil, err
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
	if err := r.client.QueryRow(ctx, q, uuid.New(), group.Name, group.OwnerId).Scan(&group.Uuid); err != nil {
		return err //todo add custom errors
	}
	return nil
}

func (r *Repository) UpdateGroup(ctx context.Context, uuid string) (*entity.Group, error) {
	panic("implement me!!!")
}
