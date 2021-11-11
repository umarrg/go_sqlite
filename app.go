package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type Student struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./store.db3")
	handleErr(err)
	createTableQuery := `create table if not exists students(
		id integer not null primary key autoincrement,
		name text,
		age integer
	)`

	res, err := Db.Exec(createTableQuery)
	handleErr(err)
	fmt.Println(res)
}

func main() {
	router := httprouter.New()

	router.GET("/", fetchAll)
	router.GET("/:id", fetchById)
	router.POST("/", createStudent)
	router.PUT("/:id", updateStudent)
	router.DELETE("/:id", deleteStudent)
	fmt.Println("Server listening to port 3000")
	http.ListenAndServe(":3000", router)
}

func fetchAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	rows, err := fetchAllFromDb()
	handleErr(err)
	students := []Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		handleErr(err)
		newStudent := Student{id, name, age}
		students = append(students, newStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	handleErr(err1)

}

func createStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	newStudent := Student{}
	json.NewDecoder(req.Body).Decode(&newStudent)
	insertQuery := fmt.Sprintf("insert into students(name,age) values('%v', '%v')", newStudent.Name, newStudent.Age)
	result, err := Db.Exec(insertQuery)
	handleErr(err)
	lid, err := result.LastInsertId()
	handleErr(err)
	fmt.Println("insert result >", lid)
	newStudent.Id = lid
	json.NewEncoder(res).Encode(newStudent)

}

func updateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	id, _ := strconv.Atoi(params[0].Value)
	student := Student{}
	json.NewDecoder(req.Body).Decode(&student)
	rows, err := updateOneFromDb(student.Name, student.Age, int64(id))
	handleErr(err)
	students := []Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		handleErr(err)
		updatedStudent := Student{id, name, age}
		students = append(students, updatedStudent)

	}

	err1 := json.NewEncoder(res).Encode(students)
	handleErr(err1)
	json.NewEncoder(res).Encode(student)

}

func deleteStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	id, _ := strconv.Atoi(params[0].Value)

	rows, err := deleteOneFromDb(int64(id))
	handleErr(err)
	students := []Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		handleErr(err)
		deletedStudent := Student{id, name, age}
		students = append(students, deletedStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	handleErr(err1)

}

func fetchById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")

	id, _ := strconv.Atoi(params[0].Value)

	rows, err := fetchOneFromDb(int64(id))
	handleErr(err)
	students := []Student{}
	for rows.Next() {
		var id int64
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		handleErr(err)
		newStudent := Student{id, name, age}
		students = append(students, newStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	handleErr(err1)

}

func fetchOneFromDb(id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from students where id='%v'", id)
	rows, err := Db.Query(query)
	handleErr(err)
	return rows, err
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func fetchAllFromDb() (*sql.Rows, error) {
	query := "select * from students"
	rows, err := Db.Query(query)
	handleErr(err)
	return rows, err
}

func updateOneFromDb(name string, age int, id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("update students set name='%v', age='%v'  where id='%v'", name, age, id)
	rows, err := Db.Query(query)
	handleErr(err)
	return rows, err
}

func deleteOneFromDb(id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("delete from students where id='%v'", id)
	rows, err := Db.Query(query)
	handleErr(err)
	return rows, err
}
