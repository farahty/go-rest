package controllers

import (
	"github.com/gofiber/fiber/v2"
	authService "github.com/nimerfarahty/go-rest/services/auth"
)

func Authintication(g fiber.Router) {
	g.Post("/login", login)
	g.Post("/logout", logout)
	g.Post("/refreshToken", refreshToken)
}

func login(c *fiber.Ctx) error {

	var loginRespons *authService.LoginRespose
	var loginInput authService.LoginInput
	var err error

	if err = c.BodyParser(&loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"status": "error", "message": "Error in reading request body", "data": err},
		)
	}

	if loginRespons, err = authService.Login(&loginInput); err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"status": "error", "message": "Authintication Error", "data": err.Error()},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"status": "ok", "message": "login success", "data": loginRespons},
	)
}

func logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "route still not ready.",
	})
}

func refreshToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "route still not ready.",
	})
}
