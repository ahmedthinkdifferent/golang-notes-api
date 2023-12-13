package model

import (
	"gorm.io/gorm"
	"noteapp/core/util"
	"time"
)

type User struct {
	Id        int            `json:"id" gorm:"primaryKey,column:id"`
	Name      string         `json:"name" gorm:"column:name"`
	Email     string         `json:"email" gorm:"column:email"`
	Password  string         `json:"-" gorm:"column:password"`
	Age       uint           `json:"age" gorm:"column:age"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Notes     []Note         `json:"notes" gorm:"foreignKey:UserId"`
	JwtToken  string         `json:"jwt_token" gorm:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, _ = util.HashPassword(u.Password)
	return nil
}
