package db

import "database/sql"
import _ "github.com/lib/pq"

var Client *sql.DB

func Init(connString string) (err error) {
	Client, err = sql.Open("postgres", connString)
	if err != nil {
		return
	}

	return
}
