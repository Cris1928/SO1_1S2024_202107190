package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "../react/dist")
	app.Get("/data", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"datos": "Cristian Daniel Gomez Escobar - 202107190"})
	})

	app.Listen(":3000")
	fmt.Println("server en 3000")

}
