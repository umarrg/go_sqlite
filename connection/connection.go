package connection

import (
	"database/sql"
	"fmt"

	studentModel "github.com/umarrg/go_crud/model"
	"github.com/umarrg/go_crud/myError"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./store.db3")
	myError.HandleErr(err)

	res, err := Db.Exec(studentModel.CreateTableQuery)
	myError.HandleErr(err)
	fmt.Println(res)
}
