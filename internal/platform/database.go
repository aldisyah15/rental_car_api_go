package platform

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	//godotenv.Load()
	db, err := sql.Open("mysql", "root:buatakun123@tcp(127.0.0.1:3306)/rental_car?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("connect db success")
	return db
}
