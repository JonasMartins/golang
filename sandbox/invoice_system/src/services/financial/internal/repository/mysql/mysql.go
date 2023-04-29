package mysql

import (
	"database/sql"
	"project/src/services/financial/configs"
)

type Repository struct {
	Db *sql.DB
}

func New() (*Repository, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.Conn)
	if err != nil {
		return nil, err
	}
	return &Repository{Db: db}, nil
}
