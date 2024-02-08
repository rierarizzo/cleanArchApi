package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"myclothing/config"
	"time"
)

type postgresDatabase struct {
	Db *sql.DB
}

func NewPostgresDatabase(cfg *config.Db) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
		cfg.TimeZone,
	)

	const maxAttempts = 10
	var counts int64

	for {
		db, err := openDB(dsn)

		if err != nil {
			slog.Info("Postgres not yet ready")
			counts++
		} else {
			slog.Info("Connected to Postgres!")

			return &postgresDatabase{Db: db}
		}

		if counts > maxAttempts {
			panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
		}

		slog.Info("Backing off for 3 seconds...")
		time.Sleep(3 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *postgresDatabase) GetDb() *sql.DB {
	return p.Db
}
