package auth

import (
	"noteapp/core/data/db"
	"noteapp/core/data/model"
)

type IUserRepo interface {
	FindByEmail(email string) *model.User
	RegisterUser(name string, email string, password string, age uint) *model.User
}

type UserRepository struct{}

var userRepository *UserRepository

func NewUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = new(UserRepository)
	}
	return userRepository
}

func (receiver *UserRepository) FindByEmail(email string) *model.User {
	var user = new(model.User)
	db.Database.Where("email = ?", email).Find(&user)
	if user.Id > 0 {
		return user
	}
	return nil
}

func (receiver *UserRepository) RegisterUser(name string, email string, password string, age uint) *model.User {
	var user = &model.User{Email: email, Name: name, Password: password, Age: age}
	db.Database.Create(&user) // pass pointer of data to Create
	return user
}
