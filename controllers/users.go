package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nimerfarahty/go-rest/services"
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

	return c.JSON(services.UserCreateInput{Password: "test", Phone: "t"}.Validate())
}
