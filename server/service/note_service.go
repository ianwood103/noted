package service

import (
	"context"
	"noted/data/request"
	"noted/data/response"
	"noted/helper"
	"noted/model"
	"noted/repository"
)

type NoteService struct {
	NoteRepository repository.INoteRepository
}

func NewNoteService(noteRepository repository.INoteRepository) INoteService {
	return &NoteService{NoteRepository: noteRepository}
}

type INoteService interface {
	Create(ctx context.Context, request request.NoteCreateRequest)
	Update(ctx context.Context, request request.NoteUpdateRequest)
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) response.NoteResponse
	FindAll(ctx context.Context) []response.NoteResponse
}

// Create implements NoteService.
func (n *NoteService) Create(ctx context.Context, request request.NoteCreateRequest) {
	note := model.Note{
		Text: request.Text,
	}

	n.NoteRepository.Save(ctx,note)
}

// Delete implements NoteService.
func (n *NoteService) Delete(ctx context.Context, noteId int) {
	note, err := n.NoteRepository.FindById(ctx, noteId)
	helper.PanicIfError(err)
	n.NoteRepository.Delete(ctx, note.Id)
}

// FindAll implements NoteService.
func (n *NoteService) FindAll(ctx context.Context) []response.NoteResponse {
	notes := n.NoteRepository.FindAll(ctx)
	var noteResponses []response.NoteResponse

	for _, note := range notes {
		noteResponse := response.NoteResponse{Id: note.Id, Text: note.Text}
		noteResponses = append(noteResponses, noteResponse)
	}

	return noteResponses
}

// FindById implements NoteService.
func (n *NoteService) FindById(ctx context.Context, noteId int) response.NoteResponse {
	note, err := n.NoteRepository.FindById(ctx, noteId)
	helper.PanicIfError(err)
	return response.NoteResponse{Id: note.Id, Text: note.Text}
}

// Update implements NoteService.
func (n *NoteService) Update(ctx context.Context, request request.NoteUpdateRequest) {
	note, err := n.NoteRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	note.Text = request.Text
	n.NoteRepository.Update(ctx, note)
}
