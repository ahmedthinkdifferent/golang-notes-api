package note

import (
	"noteapp/core/data/db"
	"noteapp/core/data/model"
)

type INoteRepo interface {
	Create(note *model.Note) (*model.Note, error)
	UserNotes(userId int) ([]model.Note, error)
	UserWithNotes(userId int) *model.User
}

type NoteRepo struct {
}

func (n *NoteRepo) UserWithNotes(userId int) *model.User {
	var user model.User
	db.Database.Preload("Notes").Select("id,name").Where("id = ?", userId).First(&user)
	return &user
}

func (n *NoteRepo) Create(note *model.Note) (*model.Note, error) {
	result := db.Database.Create(note)
	return note, result.Error
}

func (n *NoteRepo) UserNotes(userId int) ([]model.Note, error) {
	var notes []model.Note
	result := db.Database.Where("user_id = ?", userId).Find(&notes)
	return notes, result.Error
}
