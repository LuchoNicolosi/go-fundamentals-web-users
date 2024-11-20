package boostrap

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3336)/go_course_users")
	if err != nil {
		return nil, err
	}
	return db, nil
}
