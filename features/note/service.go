package note

import (
	"noteapp/core/data/db"
	"noteapp/core/data/model"
)

func UserNotes(userId int, repo INoteRepo) []model.Note {
	notes, _ := repo.UserNotes(userId)
	return notes
}

func SaveUserNote(userId int, title string, content string, repo INoteRepo) model.Note {
	note := model.Note{
		UserId:  userId,
		Title:   title,
		Content: content,
	}
	db.Database.Create(&note)
	return note
}
