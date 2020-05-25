package mysql_service

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() {

	fmt.Println("Connecting to Mysql Database")

	db, error := sql.Open("mysql", "root:ueducation@tcp(localhost:3306)/producer")

	if error != nil {
		panic(error.Error())
	}

	defer db.Close()
}

func main() {
	ConnectDB()
}
