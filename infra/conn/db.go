package conn

import (
	"CRUD_API/infra/logger"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println(err.Error())
	}

	logger.Info("mysql connection successful...")
	return db
}
