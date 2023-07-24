package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pradumnasaraf/Shortify/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/magiconair/properties/assert"
)

// Test Shorten functionality
func TestShortenRoute(t *testing.T) {

	// Load .env file
	LoadEnv()

	app := fiber.New()
	app.Post("/api/v1", routes.ShortenURL)

	// Create a new HTTP request
	body := `{"url": "https://test.dev", "short": "test-dev"}`
	req := httptest.NewRequest("POST", "/api/v1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Use the Fiber app to create a new HTTP response
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error while testing shorten route: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	assert.Equal(t, resp.StatusCode, 201, "Status code should be 200")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")

}
