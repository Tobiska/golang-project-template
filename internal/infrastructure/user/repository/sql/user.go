package sql

import (
	"context"
	"golang-project-template/internal/domains/user/entity"
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

func (r *Repository) GetAll(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	users := make([]*entity.User, 0, limit)
	rows, err := r.client.Query(ctx, `SELECT users.id, users.name, users.email, users.role FROM users LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := &entity.User{}
		if err := rows.Scan(u.Id, u.Username, u.Role, u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *Repository) GetById(ctx context.Context, id int) (*entity.User, error) {
	u := &entity.User{}
	row := r.client.QueryRow(ctx, `SELECT users.id, users.name, users.email FROM users WHERE users.id=$1`, id)
	if err := row.Scan(u.Id, u.Username, u.Email); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *Repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	row := r.client.QueryRow(ctx, `SELECT users.id, users.name, users.email FROM users WHERE users.email=$1`, email)
	if err := row.Scan(&u.Id, &u.Username, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *Repository) CreateUser(ctx context.Context, user *entity.User) error {
	q := `INSERT INTO users 
					(name, email, password_encrypted)
		  VALUES 
					($1, $2, $3)
		  RETURNING id`
	if err := r.client.QueryRow(ctx, q, user.Username, user.Email, user.PasswordEncrypted).Scan(&user.Id); err != nil {
		return err //todo add custom errors
	}
	return nil
}
func (r *Repository) UpdateUser(ctx context.Context, id int) (*entity.User, error) {
	panic("implement me !!!")
}
