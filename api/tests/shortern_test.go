package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pradumnasaraf/Shortify/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/magiconair/properties/assert"
)

// Test Shorten functionality with no custom short path
func TestShortenRouteNoCustomShort(t *testing.T) {

	// Load .env file
	LoadEnv()

	app := fiber.New()
	app.Post("/api/v1", routes.ShortenURL)

	// Create a new HTTP request
	body := `{"url": "https://test2.com"}`
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

	var responseBody Response
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, resp.StatusCode, 201, "Status code should be 200")
	if resp.StatusCode != 201 {
		return
	}
	assert.Equal(t, responseBody.URL, "https://test2.com", "URL should be https://test2.com")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")

}

// Test Shorten functionality with custom short path
func TestShortenRouteWithCustomShort(t *testing.T) {

	// Load .env file
	LoadEnv()

	app := fiber.New()
	app.Post("/api/v1", routes.ShortenURL)

	// Create a new HTTP request
	body := `{"url": "https://test3.com", "short_path": "shortpath3"}`
	req := httptest.NewRequest("POST", "/api/v1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Use the Fiber app to create a new HTTP response
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error while testing shorten route: %v", err)
	}

	var responseBody Response
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, resp.StatusCode, 201, "Status code should be 201")
	if resp.StatusCode != 201 {
		return
	}
	assert.Equal(t, responseBody.URL, "https://test3.com", "URL should be https://test3.com")
	assert.Equal(t, responseBody.CustomShortUrl, "localhost:8080/shortpath3", "CustomShortUrl should be localhost:8080/shortpath3")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")

}

// Test Shorten functionality with same custom short path
func TestShortenRouteWithSameCustomShort(t *testing.T) {

	// Load .env file
	LoadEnv()

	app := fiber.New()
	app.Post("/api/v1", routes.ShortenURL)

	// Create a new HTTP request
	body := `{"url": "https://test4.com", "short_path": "shortpath3"}`
	req := httptest.NewRequest("POST", "/api/v1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Use the Fiber app to create a new HTTP response
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error while testing shorten route: %v", err)
	}

	var responseBody Error
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, resp.StatusCode, 403, "Status code should be 403")
	if resp.StatusCode != 201 {
		return
	}
	assert.Equal(t, responseBody.Error, "Short URL already exists", "Error should be Short path already exists")
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")

}
