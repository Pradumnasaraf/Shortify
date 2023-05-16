package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/Pradumnasaraf/Shortify/database"
	"github.com/Pradumnasaraf/Shortify/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type request struct {
	URL         	string        `json:"url"`
	CustomShortPath string        `json:"short_path,omitempty"`
	Expiry      	time.Duration `json:"expiry"`
}

type response struct {
	URL             	string        `json:"url"`
	CustomShortUrl     	string     	  `json:"short_url"`
	Expiry          	time.Duration `json:"expiry"`
	XRateRemaining  	int           `json:"rate_limit"`
	XRatelimitReset 	time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request) //	var body request
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	// Get the rate limit for the IP address
	value, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		// Set the rate limit for the IP address. IP will be the key and the value will be the rate limit
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err() //
	} else {
		valueInt, _ := strconv.Atoi(value)
		if valueInt <= 0 {
			// Get the time left for the rate limit to reset
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result() // limit is in nanoseconds
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded", "retry_after": limit / time.Nanosecond / time.Minute})
		}
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "The URL is invalid"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Cannot shorten URLs from this domain"})
	}

	body.URL = helpers.EnforceHTTPS(body.URL)

	var id string

	// If the user doesn't provide a custom short URL, generate a random one
	if body.CustomShortPath == "" {

		id = uuid.New().String()[:5]
	} else {
		id = body.CustomShortPath
	}

	r := database.CreateClient(0)
	defer r.Close()

	// Check if the short URL already exists
	val, _ := r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Short URL already exists"})
	}
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	// Set the short URL in the database with the long URL as the key
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*60*60*time.Second).Err() // it will expire after 24 hours
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to connect to database"})
	}

	resp := response{
		URL:             body.URL,
		CustomShortUrl:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRatelimitReset: 30,
	}

	r2.Decr(database.Ctx, c.IP()) // Increment the rate limit for the IP address

	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRatelimitReset = ttl / time.Nanosecond / time.Minute // Convert the time left to minutes

	// If the user doesn't provide a custom short URL, generate a random one
	resp.CustomShortUrl = os.Getenv("APP_DOMAIN") + "/" + id

	return c.Status(fiber.StatusCreated).JSON(resp)
}
