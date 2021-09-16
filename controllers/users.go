package controllers

import "github.com/gofiber/fiber/v2"

func UsersController(g fiber.Router) {
	g.Get("/", index)
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
