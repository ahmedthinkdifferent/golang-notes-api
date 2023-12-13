package note

import (
	"noteapp/core/data/db"
	"noteapp/core/data/model"
)

func UserNotes(userId int, repo INoteRepo) []model.Note {
	notes, _ := repo.UserNotes(userId)
	return notes
}

func SaveUserNote(userId int, title string, content string, files []string, repo INoteRepo) model.Note {
	note := model.Note{
		UserId:  userId,
		Title:   title,
		Content: content,
	}
	db.Database.Create(&note)
	var noteFiles = make([]model.NoteFile, 0)
	if len(files) > 0 {
		// save files
		for _, file := range files {
			noteFile := model.NoteFile{
				NoteId: note.Id,
				File:   file,
			}
			db.Database.Create(&noteFile)
			noteFiles = append(noteFiles, noteFile)
		}
	}
	note.Files = noteFiles
	return note
}
