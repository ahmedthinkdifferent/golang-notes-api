package auth

import (
	"github.com/gofiber/fiber/v2"
	"noteapp/core/data/model"
	"noteapp/core/http"
	"noteapp/core/http/res"
	"noteapp/core/util"
)

type RegisterDto struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      uint   `json:"age" validate:"required,numeric"`
}

func RegisterUser(ctx *fiber.Ctx, repository IUserRepo) error {
	registerDto := new(RegisterDto)
	_ = ctx.BodyParser(&registerDto)
	result := util.Validate(registerDto)
	if result.HasErrors() {
		return res.ValidationError(ctx, result.Errors)
	}
	var user = repository.FindByEmail(registerDto.Email)
	if user != nil {
		return res.BadRequestError(ctx, "User with email already exists")
	}
	user = repository.RegisterUser(registerDto.Name, registerDto.Email, registerDto.Password, registerDto.Age)
	return res.Success(ctx, user)
}

func LoginUser(dto LoginDto, repository *UserRepository) (*model.User, *http.AppError) {
	appError := util.Validate(dto)
	if appError.HasErrors() {
		return nil, &appError
	}
	var user = repository.FindByEmail(dto.Email)
	if user == nil {
		return nil, &http.AppError{StatusCode: 400, Error: "User with email does not exist"}
	}
	// check password
	var samePassword = util.IsPasswordCorrect(dto.Password, user.Password)
	if !samePassword {
		return nil, &http.AppError{StatusCode: 400, Error: "Password is incorrect"}
	}
	return user, nil
}
