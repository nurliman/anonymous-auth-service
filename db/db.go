package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nurliman/anonymous-auth-service/env"
)

func New() *sqlx.DB {
	driver := "mysql"
	host := env.Get("DB_HOST", "127.0.0.1")
	port := env.Get("DB_PORT", "3306")
	user := env.Get("DB_USER", "root")
	pass := env.Get("DB_PASS", "password")
	schema := env.Get("DB_NAME", "schema")

	source := fmt.Sprint(user, ":", pass, "@tcp(", host, ":", port, ")/", schema)

	db := sqlx.MustConnect(driver, source)

	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(100)

	return db
}
