package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

var AppEnvironment *AppEnv

type App struct {
	Port     int    `json:"port"`
	JwtToken string `json:"jwt_token"`
}

type AppEnv struct {
	App
}

func LoadEnv() {
	data, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	AppEnvironment = &AppEnv{App: App{Port: ToInt(data["PORT"]), JwtToken: data["JWT_TOKEN"]}}
	fmt.Println("env loaded successfully")
}

func ToInt(data string) int {
	val, _ := strconv.Atoi(data)
	return val
}
