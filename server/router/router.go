package router

import (
	"fmt"
	"net/http"
	"noted/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(noteController *controller.NoteController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/note", noteController.FindAll)
	router.GET("/api/note/:noteId", noteController.FindById)
	router.POST("/api/note", noteController.Create)
	router.PATCH("/api/note/:noteId", noteController.Update)
	router.DELETE("/api/note/:noteId", noteController.Delete)

	return router
}