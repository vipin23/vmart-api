package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vipin23/vmart-api/api_router"
	"github.com/vipin23/vmart-api/database"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}

func main() {
	fmt.Println("Welcome to V-Mart API server")
	LoadEnv()
	database.SetupDB()
	api_router.SetupRouter()
	database.InitData()
	api_router.Run(os.Getenv("SERVER_PORT"))
}
