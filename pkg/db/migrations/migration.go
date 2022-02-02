package migrations

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang-project-template/config"
	"golang-project-template/pkg/utils"
	"time"
)

type Migration struct {
	migrate *migrate.Migrate
}

func NewMigration(ctx context.Context, cfg *config.Config) (*Migration, error) {
	var m *migrate.Migrate
	var err error
	databaseURL := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`, cfg.PG.Username, cfg.PG.Password, cfg.PG.Host, cfg.PG.Port, cfg.PG.DbName)
	err = utils.DoWithTries(func() error {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		m, err = migrate.New(cfg.PG.MigrationsSourceURL, databaseURL)
		if err != nil {
			return err
		}

		return nil
	}, cfg.PG.AttemptToConnect, 3*time.Second)
	if err != nil {
		return nil, err
	}
	return &Migration{
		migrate: m,
	}, nil
}

func (m *Migration) Up() error {
	if err := m.migrate.Up(); err != nil {
		return err
	}
	return nil
}

func (m *Migration) Down() error {
	if err := m.migrate.Down(); err != nil {
		return err
	}
	return nil
}
