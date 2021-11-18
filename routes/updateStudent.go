package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/umarrg/go_crud/model"
	"github.com/umarrg/go_crud/myError"
	"github.com/umarrg/go_crud/studentDao"
)

func UpdateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params[0].Value)
	student := model.Student{}
	json.NewDecoder(req.Body).Decode(&student)
	rows, err := studentDao.UpdateOneFromDb(student.FirstName, student.LastName, student.Email, student.Gender, student.Age, int64(id))
	myError.HandleErr(err)
	students := []model.Student{}
	for rows.Next() {
		var id int64
		var firstName string
		var lastName string
		var email string
		var gender string
		var age int
		err := rows.Scan(&id, &firstName, &age, &lastName, &email, &gender)
		myError.HandleErr(err)
		updatedStudent := model.Student{id, firstName, lastName, email, gender, age}
		students = append(students, updatedStudent)

	}

	err1 := json.NewEncoder(res).Encode(students)
	myError.HandleErr(err1)
	json.NewEncoder(res).Encode(student)

}
