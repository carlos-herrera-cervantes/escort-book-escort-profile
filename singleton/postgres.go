package singleton

import (
	"database/sql"
	"log"
	"sync"

	"escort-book-escort-profile/config"

	_ "github.com/lib/pq"
)

var postgresInstance *PostgresClient
var singlePostgresClient sync.Once

type PostgresClient struct {
	EscortProfileDB *sql.DB
}

func initPostgresClient() {
	escortProfileDB, err := sql.Open(
		"postgres",
		config.InitPostgres().Databases.EscortProfile,
	)

	if err != nil {
		log.Fatalf("Error connecting with escort_profile DB: %s", err.Error())
	}

	postgresInstance = &PostgresClient{
		EscortProfileDB: escortProfileDB,
	}
}

func NewPostgresClient() *PostgresClient {
	singlePostgresClient.Do(initPostgresClient)
	return postgresInstance
}
