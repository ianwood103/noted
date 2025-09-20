package controller

import (
	"net/http"
	"noted/data/request"
	"noted/data/response"
	"noted/helper"
	"noted/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type NoteController struct {
	NoteService service.INoteService
}

func NewNoteController(noteService service.INoteService) *NoteController {
	return &NoteController{NoteService: noteService}
}

func (controller *NoteController) Create(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	noteCreateRequest := request.NoteCreateRequest{}
	helper.ReadRequestBody(r, &noteCreateRequest)

	controller.NoteService.Create(r.Context(), noteCreateRequest)
	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *NoteController) Update(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	noteUpdateRequest := request.NoteUpdateRequest{}
	helper.ReadRequestBody(r, &noteUpdateRequest)

	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)
	noteUpdateRequest.Id = id

	controller.NoteService.Update(r.Context(), noteUpdateRequest)
	webResponse := response.WebResponse{
		Code: 200, 
		Status: "Ok",
		Data: nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *NoteController) Delete(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)

	controller.NoteService.Delete(r.Context(), id)
	webResponse := response.WebResponse{
		Code: 200, 
		Status: "Ok",
		Data: nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *NoteController) FindAll(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result := controller.NoteService.FindAll(r.Context())
	webResponse := response.WebResponse{
		Code: 200, 
		Status: "Ok",
		Data: result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *NoteController) FindById(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)

	result := controller.NoteService.FindById(r.Context(), id)
	webResponse := response.WebResponse{
		Code: 200, 
		Status: "Ok",
		Data: result,
	}
	helper.WriteResponseBody(writer, webResponse)
}