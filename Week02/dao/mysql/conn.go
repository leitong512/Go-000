package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:170512@tcp(127.0.0.1:3306)/ceshi")
	err := db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to mysql, err : %s", err))
	}
}
func DBConn() *sql.DB {
	return db
}
