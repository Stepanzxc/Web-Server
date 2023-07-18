package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	pool *sql.DB
}

var Connect *MySQL

func NewMySQL() {
	Connect = &MySQL{
		pool: initConnect(),
	}
}

// Pool ...
func (m *MySQL) Pool() *sql.DB {
	return m.pool
}

func initConnect() *sql.DB {
	db, err := sql.Open("mysql", "stepan:secret2@tcp(127.0.0.1:3306)/market_place")
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}
