package config

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrConnNotFound = errors.New("connection to database not found")
	ErrUnableToConnect = errors.New("unable to connect to database")
)

type PgxConf struct {
	DB *pgxpool.Pool
}

func (p *PgxConf) GetDBConn() *pgxpool.Pool {
	return p.DB
}

func (p *PgxConf) SetDBConn(conn *pgxpool.Pool) {
	p.DB = conn
}

func InitPgsql() (*PgxConf, error) {
	var pgxConf = &PgxConf{}
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return pgxConf, ErrUnableToConnect
	}
	
	log.Println("Connected to database successfully")
	if err := runMigrations(); err != nil {
		return pgxConf, err
	}
	
	pgxConf.SetDBConn(conn)
	return pgxConf, nil
}

func runMigrations() error {
	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}