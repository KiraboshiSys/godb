package database

import (
	"fmt"

	"github.com/KiraboshiSys/godb/config"
	"github.com/KiraboshiSys/godb/internal/database/mysql"
	"github.com/KiraboshiSys/godb/internal/database/postgres"
)

func New(cfg config.Database) (Database, error) {
	switch cfg.Driver {
	case "postgres":
		return postgres.New(cfg), nil
	case "mysql":
		return mysql.New(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}
}

var (
	_ Database = (*mysql.MySQL)(nil)
	_ Database = (*postgres.Postgres)(nil)
)
