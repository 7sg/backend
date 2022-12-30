package data

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewPostgresClient, NewData, NewGitCloneKafkaRepo, NewFileContentKafkaRepo, NewGreeterRepo, NewRepositoryRepo, NewScanRepo)

// Data .
type Data struct {
	db *sql.DB
}

// NewData .
func NewData(db *sql.DB) (*Data, error) {
	return &Data{db: db}, nil
}
