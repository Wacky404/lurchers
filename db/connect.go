package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"

	dbSql "github.com/Wacky404/lurchers/db"
	"github.com/Wacky404/lurchers/util"
)

type Database struct {
	conn     *sql.DB
	ctx      *context.Context
	host     string
	port     int64
	user     string
	password string
	name     string
	sql      *dbSql.DbStatements
}

// creates an instance of the Database struct and loads in env vars
func LoadConfig(ctx *context.Context) (*Database, error) {
	port, err := strconv.ParseInt(util.GetVar("DB_PORT", ""), 10, 64)
	if err != nil {
		slog.Error("error loading .env var port", slog.Any("error", err))
		return nil, err
	}
	db := &Database{
		ctx:      ctx,
		host:     util.GetVar("DB_HOST", ""),
		port:     port,
		user:     util.GetVar("DB_USER", ""),
		password: util.GetVar("DB_PASSWORD", ""),
		name:     util.GetVar("DB_NAME", ""),
		sql:      dbSql.NewDbStatements(),
	}

	return db, nil
}

// connecting to the precious
func (d *Database) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.name)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error("error connecting to the database", slog.Any("error", err))
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	d.conn = conn
	return nil
}
