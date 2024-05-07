package service

import (
	"context"
	"noted/data/request"
	"noted/data/response"
)

type NoteService interface {
	Create(ctx context.Context, request request.NoteCreateRequest)
	Update(ctx context.Context, request request.NoteUpdateRequest)
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) response.NoteResponse
	FindAll(ctx context.Context) []response.NoteResponse
}