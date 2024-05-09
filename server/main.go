package main

import (
	"fmt"
	"net/http"
	"noted/config"
	"noted/controller"
	"noted/helper"
	"noted/repository"
	"noted/router"
	"noted/service"
)

func main() {
	fmt.Println("start server")

	// setup database connection
	db := config.DatabaseConnection()

	// setup note repository using database
	noteRepository := repository.NewNoteRepository(db)

	// setup note service using note repository
	noteService := service.NewNoteRepositoryImpl(noteRepository)

	// setup note controller using note service
	noteController := controller.NewNoteController(noteService)

	// setup router using note controller
	router := router.NewRouter(noteController)

	// setup server using router
	server := http.Server{Addr: "localhost:8888", Handler: router}

	// start server
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}