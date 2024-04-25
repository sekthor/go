package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type mysqlRepo struct {
	db   *sql.DB
	conf mysql.Config
}

func (r *mysqlRepo) Connect() error {
	db, err := sql.Open("mysql", r.conf.FormatDSN())
	if err != nil {
		return err
	}
	r.db = db
	return nil
}

func (r *mysqlRepo) Ping() error {
	return r.db.Ping()
}

func (r *mysqlRepo) Migrate() error {
	return nil
}
