package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/umarrg/go_crud/model"
	"github.com/umarrg/go_crud/myError"
	"github.com/umarrg/go_crud/routes"
)

func Requst(method, url string, body *bytes.Buffer, handle httprouter.Handle) (*httptest.ResponseRecorder, error) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	myError.HandleErr(err)
	router := httprouter.New()
	router.Handle(method, url, handle)
	router.ServeHTTP(resp, req)
	return resp, nil

}

func TestCreateStudent(t *testing.T) {
	student := model.Student{
		Id:        0,
		FirstName: "Umar",
		LastName:  "rg",
		Email:     "umarrg@gmail.com",
		Gender:    "male",
		Age:       10,
	}
	body, err := json.Marshal(student)
	myError.HandleErr(err)
	resp, err := Requst("POST", "/student", bytes.NewBuffer(body), routes.CreateStudent)
	myError.HandleErr(err)
	if resp.Code != 200 {
		t.Error("the status code is not ok")
	}

}
