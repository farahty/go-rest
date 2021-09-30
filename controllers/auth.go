package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nimerfarahty/go-rest/models"
	authService "github.com/nimerfarahty/go-rest/services/auth"
)

func Authintication(g fiber.Router) {
	g.Post("/login", login)
	g.Post("/logout", logout)
	g.Post("/refreshToken", refreshToken)
}

func login(c *fiber.Ctx) error {

	var loginRespons *models.LoginResponse
	var loginInput models.LoginInput
	var err error

	if err = c.BodyParser(&loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONResponse{Status: "error", Message: "Error in reading request body", Data: err},
		)
	}

	if loginRespons, err = authService.Login(&loginInput); err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(
			JSONResponse{Status: "error", Message: "Authintication Error", Data: err.Error()},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		JSONResponse{Status: "ok", Message: "login success", Data: loginRespons},
	)
}

func logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(
		JSONResponse{Status: "error", Message: "route still not ready.", Data: ""},
	)
}

func refreshToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(
		JSONResponse{Status: "error", Message: "route still not ready.", Data: ""},
	)
}
