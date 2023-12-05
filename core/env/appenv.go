package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"noteapp/core/util"
)

var AppEnvironment *AppEnv

type App struct {
	Port int `json:"port"`
}

type AppEnv struct {
	App
}

func LoadEnv() {
	data, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppEnvironment = &AppEnv{App: App{Port: util.ToInt(data["PORT"])}}
	fmt.Println("env loaded successfully")
}
