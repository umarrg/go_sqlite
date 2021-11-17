package connection

import (
	"database/sql"
	"fmt"

	studentModel "github.com/umarrg/go_crud/model"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./store.db3")
	HandleErr(err)

	res, err := Db.Exec(studentModel.CreateTableQuery)
	HandleErr(err)
	fmt.Println(res)
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
