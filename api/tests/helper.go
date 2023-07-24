package tests

import (
	"fmt"

	"github.com/Pradumnasaraf/Shortify/database"
	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load .env file. Lodding from ENV if available")
	}

}

// LoadTestData loads test data into the database
func LoadTestData() {

	db := database.CreateClient(0)
	db.Set(database.Ctx, "shortpath1", "https://test1.com", 0)
}
