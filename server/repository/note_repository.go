package repository

import (
	"context"
	"noted/model"
)

type NoteRepository interface {
	Save(ctx context.Context, note model.Note)
	Update(ctx context.Context, note model.Note)
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) (model.Note, error)
	FindAll(ctx context.Context) []model.Note
}