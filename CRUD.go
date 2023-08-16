package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var users = []User{
	{ID: 1, FirstName: "Turter", LastName: "Dev"},
	{ID: 2, FirstName: "Jamey", LastName: "Dev"},
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}
func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Invalid"})
	}
	user.ID = len(users) + 1
	users = append(users, *user)
	return c.JSON(user)
}
func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			updateUser := new(User)
			if err := c.BodyParser(updateUser); err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "valid"})
			}
			updateUser.ID = user.ID
			users[i] = *updateUser
			return c.JSON(updateUser)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
}

func main() {
	app := fiber.New()

	//Define routes
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Get("/users", createUser)
	app.Get("/users/:id", updateUser)
	app.Get("/users/:id", deleteUser)

	//Listen on Port 6000
	err := app.Listen(":6000")
	if err != nil {
		panic(err)
	}

}
