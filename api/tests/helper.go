package tests

import (
	"fmt"

	"github.com/Pradumnasaraf/Shortify/database"
	"github.com/joho/godotenv"
)

func LoadTestData() {

	db := database.CreateClient(0)
	db.Set(database.Ctx, "test-short", "https://www.google.com", 0)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load .env file. Lodding from ENV if available")
	}

}
