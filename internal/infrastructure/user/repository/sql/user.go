package sql

import (
	"context"
	"fmt"
	"golang-project-template/internal/domains/user/entity"
	"golang-project-template/pkg/db/postgres"
	"strconv"
	"strings"
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

func (r *Repository) GetUsersByGroupId(ctx context.Context, uuid string) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	q := `SELECT
			users.id, users.name, users.email
		  FROM
			users
 		  WHERE users.group_id=$1`
	rows, err := r.client.Query(ctx, q, uuid)
	if err != nil {
		return nil, err
	}
	it := 0
	for rows.Next() {
		u := &entity.User{}
		if err := rows.Scan(&u.Id, &u.Username, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
		it++
	}
	return users, nil
}

func (r *Repository) GetByIds(ctx context.Context, ids ...int) ([]*entity.User, error) {
	placeholders, _ := makeArgsAndPlaceholders(ids)
	users := make([]*entity.User, len(ids))
	q := `SELECT
			users.id, users.name, users.email
		  FROM
			users
 		  WHERE users.id IN (` + strings.Join(placeholders, ",") + `)`
	tmp := make([]interface{}, len(ids)) //TODO fix
	for i := 0; i < len(ids); i++ {
		tmp[i] = ids[i]
	}
	rows, err := r.client.Query(ctx, q, tmp...)
	if err != nil {
		return nil, err
	}
	it := 0
	for rows.Next() {
		u := &entity.User{}
		if err := rows.Scan(&u.Id, &u.Username, &u.Email); err != nil {
			return nil, err
		}
		users[it] = u
		it++
	}
	return users, nil
}

func (r *Repository) GetById(ctx context.Context, id int) (*entity.User, error) {
	u := &entity.User{}
	row := r.client.QueryRow(ctx, `SELECT users.id, users.name, users.email FROM users WHERE users.id=$1`, id)
	if err := row.Scan(&u.Id, &u.Username, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *Repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	row := r.client.QueryRow(ctx, `SELECT users.id, users.name, users.email, users.password_encrypted FROM users WHERE users.email=$1`, email)
	if err := row.Scan(&u.Id, &u.Username, &u.Email, &u.PasswordEncrypted); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *Repository) CreateUser(ctx context.Context, user *entity.User) error {
	q := `INSERT INTO users 
					(name, email, password_encrypted, role)
		  VALUES 
					($1, $2, $3, $4)
		  RETURNING id`
	if err := r.client.QueryRow(ctx, q, user.Username, user.Email, user.PasswordEncrypted, user.Role).Scan(&user.Id); err != nil {
		return err //todo add custom errors
	}
	return nil
}
func (r *Repository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	q := `UPDATE users SET 
			name=$1, email=$2, group_id=$3
		 WHERE users.id=$4`
	exec, err := r.client.Exec(ctx, q, user.Username, user.Email, user.GroupID, user.Id)
	fmt.Println(exec)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func makeArgsAndPlaceholders(ids []int) (placeholders []string, args []interface{}) {
	placeholders = make([]string, len(ids))
	args = make([]interface{}, len(ids))
	for i := 0; i < len(ids); i++ {
		placeholders[i] = fmt.Sprintf(`$%s`, strconv.Itoa(i+1))
		args[i] = i
	}
	return
}
