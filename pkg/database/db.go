package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/komugi8/clava/pkg/config"
)

func NewDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	log.Print(conn)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	xdb := sqlx.NewDb(db, "mysql")
	return xdb, nil
}
