package mysql

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Driver struct {
	Db *sql.DB
}

func New(user, pass, host, dbname string) *Driver {
	config := mysql.NewConfig()

	config.User = user
	config.Passwd = pass
	config.Addr = host
	config.DBName = dbname
	config.Net = "tcp"

	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return &Driver{
		Db: db,
	}
}
