package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value :" + c.Params("value"))
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("api path :" + c.Params("*"))
	})

	app.Static("/", "./public")

	app.Listen(":8081")
}
