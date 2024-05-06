package repository

import (
	"context"
	"database/sql"
	"errors"
	"noted/helper"
	"noted/model"
)

type NoteRepositoryImpl struct {
	Db *sql.DB
}

func NewNoteRepository(Db *sql.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: Db}
}

// Delete implements NoteRepository.
func (n *NoteRepositoryImpl) Delete(ctx context.Context, noteId int) {
	tx, err := n.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM note WHERE id=$1"
	_, errExec := tx.ExecContext(ctx, SQL, noteId)
	helper.PanicIfError(errExec)
}

// FindAll implements NoteRepository.
func (n *NoteRepositoryImpl) FindAll(ctx context.Context) []model.Note {
	tx, err := n.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, text FROM note"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var notes []model.Note

	for result.Next() {
		note := model.Note{}
		err := result.Scan(&note.Id, &note.Text)
		helper.PanicIfError(err)

		notes = append(notes, note)
	}

	return notes
}

// FindById implements NoteRepository.
func (n *NoteRepositoryImpl) FindById(ctx context.Context, noteId int) (model.Note, error) {
	tx, err := n.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, text FROM note where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, noteId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	note := model.Note{}

	if result.Next() {
		err := result.Scan(&note.Id, &note.Text)
		helper.PanicIfError(err)
		return note, nil
	} else {
		return note, errors.New("note id not found")
	}
}

// Save implements NoteRepository.
func (n *NoteRepositoryImpl) Save(ctx context.Context, note model.Note) {
	tx, err := n.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO note(text) value($1)"
	_, err = tx.ExecContext(ctx, SQL, note.Text)
	helper.PanicIfError(err)
}

// Update implements NoteRepository.
func (n *NoteRepositoryImpl) Update(ctx context.Context, note model.Note) {
	tx, err := n.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE note SET text=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, SQL, note.Text, note.Id)
	helper.PanicIfError(err)
}


