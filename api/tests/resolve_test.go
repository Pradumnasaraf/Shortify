package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/Pradumnasaraf/Shortify/routes"
	"github.com/gofiber/fiber/v2"
)

// Test Redirect functionality
func TestResolveRoute(t *testing.T) {


	// Load .env file
	LoadEnv()

	// Load test data into redis
	LoadTestData()

	app := fiber.New()
	app.Get("/:shortPath", routes.ResolveURL)

	// Test for valid short path
	req := httptest.NewRequest("GET", "/test-short", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Error("Error while testing resolve route")
	}

	if resp.StatusCode != 301 {
		t.Error("Expected status code 301, got ", resp.StatusCode)
	}

}

// Test for Error
func TestResolveRouteForError(t *testing.T) {

	// Load .env file
	LoadEnv()

	app := fiber.New()
	app.Get("/:shortPath", routes.ResolveURL)

	// Test for valid short path
	req := httptest.NewRequest("GET", "/test-short1", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Error("Error while testing resolve route")
	}

	if resp.StatusCode != 404 {
		t.Error("Expected status code 301, got ", resp.StatusCode)
	}

}
