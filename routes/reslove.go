package routes

import (
	"github.com/Pradumnasaraf/Shortify/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	shortUrlPath := c.Params("shortPath")

	r := database.CreateClient(0)
	defer r.Close()

	// Get the actual URL from the short URL
	value, err := r.Get(database.Ctx, shortUrlPath).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short URL not found in DB"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to database"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	// Increment the counter for the short URL
	_ = rInr.Incr(database.Ctx, "counter") 

	return c.Redirect(value, fiber.StatusMovedPermanently)
}
