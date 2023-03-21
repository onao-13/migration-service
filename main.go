package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"migration-service/internals/config"
	migration "migration-service/internals/migration"
)

func main() {
	var log = logrus.New()

	log.Infoln("Server started")

	cfg := config.UploadProdConfig()
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, cfg.GetDbUrl())
	if err != nil {
		log.Fatalln("Fatal error connecting to database: ", err)
	}

	m := migration.New(pool, ctx)

	log.Infoln("Migrating starting")

	m.CreateTables()

	if cfg.LoadDevData == true {
		m.UploadDevData()
	}

	log.Infoln("Migrating finished")
}
