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

	"github.com/rs/cors"
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

	// setup handler with CORS enabled
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// setup server using router
	server := http.Server{Addr: "localhost:8888", Handler: handler}

	// start server
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
