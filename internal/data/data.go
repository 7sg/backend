package data

import (
	"database/sql"

	"github.com/google/wire"
	// postgres driver
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewPostgresClient,
	NewData,
	NewGitCloneKafkaRepo,
	NewFileContentKafkaRepo,
	NewRepositoryRepo,
	NewScanRepo,
)

// Data .
type Data struct {
	db *sql.DB
}

// NewData .
func NewData(db *sql.DB) (*Data, error) {
	return &Data{db: db}, nil
}
