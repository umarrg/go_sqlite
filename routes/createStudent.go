package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/umarrg/go_crud/connection"
	"github.com/umarrg/go_crud/model"
	"github.com/umarrg/go_crud/myError"
)

func CreateStudent(res http.ResponseWriter, req *http.Request, params httprouter.Params){
	newStudent := model.Student{}
	json.NewDecoder(req.Body).Decode(&newStudent)
	insertQuery := fmt.Sprintf("insert into students( firstName, lastName, email, gender, age) values('%v', '%v', '%v', '%v', '%v' )", newStudent.FirstName, newStudent.LastName, newStudent.Email, newStudent.Gender, int(newStudent.Age))
	result, err := connection.Db.Exec(insertQuery)
	myError.HandleErr(err)
	lid, err := result.LastInsertId()
	myError.HandleErr(err)
	fmt.Println("insert result >", lid)
	newStudent.Id = lid

}
