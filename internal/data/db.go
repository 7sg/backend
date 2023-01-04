package data

import (
	"database/sql"
	"fmt"

	"backend-GuardRails/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	// postgres driver
	_ "github.com/lib/pq"
)

func NewPostgresClient(c *conf.Data, logger log.Logger) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		c.Database.GetHost(),
		c.Database.GetPort(),
		c.Database.GetUser(),
		c.Database.GetPassword(),
		c.Database.GetDbName(),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.NewHelper(logger).Errorf("error connecting to database %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.NewHelper(logger).Errorf("error pinging to database %v", err)
	}

	return db, err
}
