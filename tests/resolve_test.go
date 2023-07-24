package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Pradumnasaraf/Shortify/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/magiconair/properties/assert"
)

// Test Resolve Route
func TestResolveRoute(t *testing.T) {

	// Load .env file
	LoadEnv()

	// Load test data into redis
	LoadTestData()

	app := fiber.New()
	app.Get("/:shortPath", routes.ResolveURL)

	// Test for valid short path
	req := httptest.NewRequest("GET", "/shortpath1", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Error("Error while testing resolve route")
	}

	assert.Equal(t, resp.StatusCode, 301, "Status code should be 301")

}

// Test Resolve Route for a short path that doesn't exist
func TestResolveRouteForError(t *testing.T) {

	// Load .env file
	LoadEnv()

	app := fiber.New()
	app.Get("/:shortPath", routes.ResolveURL)

	// Test for valid short path
	req := httptest.NewRequest("GET", "/shortpath0", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Error("Error while testing resolve route")
	}

	var responseBody Error
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Error("Error while decoding response body")
	}

	assert.Equal(t, resp.StatusCode, 404, "Status code should be 404")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")
	assert.Equal(t, responseBody.Error, "Short URL not found in DB", "Error should be Short URL not found in DB")
}
