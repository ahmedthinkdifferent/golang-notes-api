package note

import (
	"noteapp/core/data/db"
	"noteapp/core/data/model"
	"time"
)

type INoteRepo interface {
	Create(note *model.Note) (*model.Note, error)
	UserNotes(userId int) ([]model.Note, error)
	UserWithNotes(userId int) *model.User
}

type Repo struct {
}

func (n *Repo) UserWithNotes(userId int) *model.User {
	var user model.User
	user.Notes = make([]model.Note, 0)
	rows, _ := db.Database.Debug().Model(&model.User{}).
		Where("users.id = ?", userId).
		InnerJoins("Notes").
		Select("users.*, notes.id as notes__id, notes.title as notes__title, notes.content as notes__content, notes.created_at as notes__created_at, notes.updated_at as notes__updated_at, notes.deleted_at as notes__deleted_at").
		Rows()
	defer rows.Close()
	var data = make([]map[string]interface{}, 0)
	for rows.Next() {
		err := db.Database.ScanRows(rows, &data)
		if err != nil {
			return nil
		}
		user.Id = int(data[0]["id"].(int64))
		user.Name = data[0]["name"].(string)
		user.Email = data[0]["email"].(string)
		user.Age = uint(data[0]["age"].(int64))
		user.CreatedAt = data[0]["created_at"].(time.Time)
		user.UpdatedAt = data[0]["updated_at"].(time.Time)
		for _, row := range data {
			var note model.Note
			note.Id = int(row["notes__id"].(int64))
			note.Title = row["notes__title"].(string)
			note.Content = row["notes__content"].(string)
			note.CreatedAt = row["notes__created_at"].(time.Time)
			note.UpdatedAt = row["notes__updated_at"].(time.Time)
			user.Notes = append(user.Notes, note)
		}
	}

	return &user
}

func (n *Repo) Create(note *model.Note) (*model.Note, error) {
	result := db.Database.Create(note)
	return note, result.Error
}

func (n *Repo) UserNotes(userId int) ([]model.Note, error) {
	var notes []model.Note
	result := db.Database.Where("user_id = ?", userId).Find(&notes)
	return notes, result.Error
}
