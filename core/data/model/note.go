package model

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	Id        int            `json:"id" gorm:"primaryKey,column:id"`
	UserId    int            `json:"user_id" gorm:"column:user_id"`
	Title     string         `json:"title" gorm:"column:title"`
	Content   string         `json:"content" gorm:"column:content"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Files     []NoteFile     `json:"files" gorm:"foreignKey:NoteId"`
}

type NoteFile struct {
	Id     int    `json:"id" gorm:"primaryKey,column:id"`
	NoteId int    `json:"note_id" gorm:"column:note_id;foreignKey:NoteId"`
	File   string `json:"file" gorm:"column:file"`
}
