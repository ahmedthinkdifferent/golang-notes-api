package auth

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"noteapp/core/data/model"
	"noteapp/features/auth/mocks"
	"testing"
)

var app = fiber.New()

func TestRegisterUserReturnValidationError(t *testing.T) {
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.Request().SetBody([]byte("{}"))
	_ = RegisterUser(c, nil)
	response := map[string]any{}
	json.Unmarshal(c.Response().Body(), &response)
	assert.Contains(t, response["error"].(map[string]any)["Name"], "required")
}

func TestRegisterUserReturnEmailExist(t *testing.T) {
	// create an instance of our test object
	userRepository := mocks.NewIUserRepo(t)
	userRepository.On("FindByEmail", "ahmed2@gmail.com").Return(&model.User{Id: 3})
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.Request().SetBody([]byte(`{
    "name":"ahmed agamy",
    "email":"ahmed2@gmail.com",
    "password":"ahmed1",
    "age":35
}`))
	c.Request().Header.Add("Content-Type", "application/json")
	_ = RegisterUser(c, userRepository)
	response := map[string]any{}
	json.Unmarshal(c.Response().Body(), &response)
	assert.Contains(t, response["error"], "User with email already exists")
}
