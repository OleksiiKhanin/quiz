package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDB(host, user, pass, dbname string, port int) *sqlx.DB {
	db := sqlx.MustConnect("postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, pass, dbname,
		),
	)
	return db
}
