package database

import (
	"database/sql"
	"fmt"
	"quiz3/migrations"

	migrate "github.com/rubenv/sql-migrate"
)

func DBMigrate(dbParam *sql.DB) {
	DBMigrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations.MigrationsDirectory,
		Root:       ".",
	}

	migrateCount, err := migrate.Exec(dbParam, "postgres", DBMigrations, migrate.Up)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration success!")
	fmt.Println(migrateCount, "migrations applied!")
}
