package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/umarrg/go_crud/routes"
)

func main() {
	router := httprouter.New()
	router.GET("/student", routes.FetchAll)
	router.GET("/student/:id", routes.FetchById)
	// router.POST("/student", routes.CreateStudent)
	router.PUT("/student/:id", routes.UpdateStudent)
	router.DELETE("/student/:id", routes.DeleteStudent)
	fmt.Println("Server listening to port 3000")
	handler := cors.AllowAll().Handler(router)
	err := http.ListenAndServe(":3000", handler)
	fmt.Println(err)

}
