package platform

import (
	"database/sql"
	"fmt"
	"rental_car/config/env"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open(env.DriverDb.GetValue(), env.DbDsn.GetValue())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("connect db success")
	return db
}
