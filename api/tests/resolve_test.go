package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/Pradumnasaraf/Shortify/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/magiconair/properties/assert"
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

	assert.Equal(t, resp.StatusCode, 301, "Status code should be 301")

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

	assert.Equal(t, resp.StatusCode, 404, "Status code should be 404")

}
