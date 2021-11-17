package genfunctions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/umarrg/go_crud/connection"
	studentModel "github.com/umarrg/go_crud/model"
	"github.com/umarrg/go_crud/studentDao"
)

func CreateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	newStudent := studentModel.Student{}
	json.NewDecoder(req.Body).Decode(&newStudent)
	insertQuery := fmt.Sprintf("insert into students( firstName, lastName, email, gender, age) values('%v', '%v', '%v', '%v', '%v' )", newStudent.FirstName, newStudent.LastName, newStudent.Email, newStudent.Gender, int(newStudent.Age))
	result, err := connection.Db.Exec(insertQuery)
	HandleErr(err)
	lid, err := result.LastInsertId()
	HandleErr(err)
	fmt.Println("insert result >", lid)
	newStudent.Id = lid

}

func UpdateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params[0].Value)
	student := studentModel.Student{}
	json.NewDecoder(req.Body).Decode(&student)
	rows, err := studentDao.UpdateOneFromDb(student.FirstName, student.LastName, student.Email, student.Gender, student.Age, int64(id))
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var firstName string
		var lastName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &firstName, &age, &lastName, &email, &gender)
		HandleErr(err)
		updatedStudent := studentModel.Student{id, firstName, lastName, email, gender, age}
		students = append(students, updatedStudent)

	}

	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)
	json.NewEncoder(res).Encode(student)

}

func DeleteStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params[0].Value)
	rows, err := studentDao.DeleteOneFromDb(int64(id))
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var firstName string
		var lastName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &firstName, &age, &lastName, &email, &gender)
		HandleErr(err)
		deletedStudent := studentModel.Student{id, firstName, lastName, email, gender, age}
		students = append(students, deletedStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)

}

func FetchById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params[0].Value)
	rows, err := studentDao.DeleteOneFromDb(int64(id))
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var firstName string
		var lastName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &firstName, &age, &lastName, &email, &gender)
		HandleErr(err)
		newStudent := studentModel.Student{id, firstName, lastName, email, gender, age}
		students = append(students, newStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)
	HandleErr(err1)

}

func FetchAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	rows, err := studentDao.FetchAllFromDb()
	HandleErr(err)
	students := []studentModel.Student{}
	for rows.Next() {
		var id int64
		var firstName string
		var lastName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &firstName, &lastName, &email, &gender, &age)
		HandleErr(err)
		newStudent := studentModel.Student{id, firstName, lastName, email, gender, age}
		students = append(students, newStudent)
	}
	err1 := json.NewEncoder(res).Encode(students)

	HandleErr(err1)

}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
