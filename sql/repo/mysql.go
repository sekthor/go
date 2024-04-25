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

func (r *mysqlRepo) CreateUser(user User) (User, error) {
	_, err := r.db.Exec("INSERT INTO users (uuid, username, email) VALUES (?,?,?)", user.UUID, user.Username, user.Email)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *mysqlRepo) GetById(id string) error {
	return nil
}
