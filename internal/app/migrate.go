package app

import (
	"context"
	"golang-project-template/config"
	"golang-project-template/pkg/db/migrations"
	"log"
)

func init() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrations.NewMigration(context.TODO(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Up(); err != nil {
		log.Fatal(err)
	}
}
