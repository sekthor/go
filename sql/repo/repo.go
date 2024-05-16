package main

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	DbType string `yaml:"dbType"`
	Mysql  mysql.Config
}

type Repo interface {
	Connect() error
	Ping() error
	Migrate() error
	CreateUser(User) (User, error)
	GetById(string) (User, error)
}

func NewRepository(conf Config) (Repo, error) {
	switch conf.DbType {
	case "mysql":
		return &mysqlRepo{
			conf: conf.Mysql,
		}, nil
	}
	return nil, fmt.Errorf("unknown db type '%s'", conf.DbType)
}
