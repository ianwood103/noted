package service

import (
	"context"
	"noted/data/request"
	"noted/data/response"
	"noted/helper"
	"noted/model"
	"noted/repository"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
}

func NewNoteRepositoryImpl(noteRepository repository.NoteRepository) NoteService {
	return &NoteServiceImpl{NoteRepository: noteRepository}
}

// Create implements NoteService.
func (n *NoteServiceImpl) Create(ctx context.Context, request request.NoteCreateRequest) {
	note := model.Note{
		Text: request.Text,
	}

	n.NoteRepository.Save(ctx,note)
}

// Delete implements NoteService.
func (n *NoteServiceImpl) Delete(ctx context.Context, noteId int) {
	note, err := n.NoteRepository.FindById(ctx, noteId)
	helper.PanicIfError(err)
	n.NoteRepository.Delete(ctx, note.Id)
}

// FindAll implements NoteService.
func (n *NoteServiceImpl) FindAll(ctx context.Context) []response.NoteResponse {
	notes := n.NoteRepository.FindAll(ctx)
	var noteResponses []response.NoteResponse

	for _, note := range notes {
		noteResponse := response.NoteResponse{Id: note.Id, Text: note.Text}
		noteResponses = append(noteResponses, noteResponse)
	}

	return noteResponses
}

// FindById implements NoteService.
func (n *NoteServiceImpl) FindById(ctx context.Context, noteId int) response.NoteResponse {
	note, err := n.NoteRepository.FindById(ctx, noteId)
	helper.PanicIfError(err)
	return response.NoteResponse{Id: note.Id, Text: note.Text}
}

// Update implements NoteService.
func (n *NoteServiceImpl) Update(ctx context.Context, request request.NoteUpdateRequest) {
	note, err := n.NoteRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	note.Text = request.Text
	n.NoteRepository.Update(ctx, note)
}
