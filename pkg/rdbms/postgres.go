package rdbms

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(cfg *Config) (RDBMS, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("Error openning connection:\n%v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Error ping database:\n%v", err)
	}

	return &rdbms{db: db}, nil
}
