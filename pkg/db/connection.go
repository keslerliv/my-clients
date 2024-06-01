package db

import (
	"database/sql"
	"fmt"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/keslerliv/my-clients/config"
	_ "github.com/lib/pq"
)

// OpenConnection opens a connection to the database
func OpenConnection() (*sql.DB, error) {
	conf := config.Config

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.SSLMode)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

// create the database migrations
func MakeMigrations() {

	// open connection
	conn, err := OpenConnection()
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres driver: %v", err)
	}

	// create migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	// run migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	} else if err == migrate.ErrNoChange {
		log.Println("no changes in migrations")
	} else {
		log.Println("migrations applied successfully")
	}
}
