package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	genfunctions "github.com/umarrg/go_crud/genFunctions"
)

func main() {
	router := httprouter.New()
	router.GET("/student", genfunctions.FetchAll)
	router.GET("/student/:id", genfunctions.FetchById)
	router.POST("/student", genfunctions.CreateStudent)
	router.PUT("/student/:id", genfunctions.UpdateStudent)
	router.DELETE("/student/:id", genfunctions.DeleteStudent)
	fmt.Println("Server listening to port 3000")
	handler := cors.AllowAll().Handler(router)
	err := http.ListenAndServe(":3000", handler)
	fmt.Println(err)

}
