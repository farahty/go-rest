package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nimerfarahty/go-rest/models"
	usersService "github.com/nimerfarahty/go-rest/services/users"
)

func UsersController(g fiber.Router) {
	g.Get("/", index)
	g.Get("/create", create)
	g.Get("/:id", findOne)

}

func index(c *fiber.Ctx) error {

	return c.
		Status(fiber.StatusBadGateway).
		JSON(fiber.Map{
			"message": "route still not ready",
		})
}

func findOne(c *fiber.Ctx) error {

	return c.
		Status(fiber.StatusBadGateway).
		JSON(fiber.Map{
			"message": "route still not ready",
		})
}

func create(c *fiber.Ctx) error {

	var input usersService.CreateInput
	var user *models.User
	var err error

	if err = c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"status": "error", "message": "Error in reading request body", "data": err},
		)
	}

	if user, err = usersService.CreateUser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"status": "error", "message": "Validation Error", "data": err},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{"status": "ok", "message": "User created successfully", "data": &user},
	)
}
