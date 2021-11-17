package studentDao

import (
	"database/sql"
	"fmt"

	conn "github.com/umarrg/go_crud/connection"
)

func FetchOneFromDb(id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from students where id='%v'", id)
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func FetchAllFromDb() (*sql.Rows, error) {
	query := "select * from students"
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func UpdateOneFromDb(firstName string, lastName string, email string, gender string, age int, id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("update students set firstName='%v', lastName='%v', email='%v', gender='%v', age='%v'  where id='%v'", firstName, lastName, email, gender, age, id)
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func DeleteOneFromDb(id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("delete from students where id='%v'", id)
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
