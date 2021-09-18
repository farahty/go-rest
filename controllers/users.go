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

// Index is a function to get all books data from database
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept json
// @Produce json
func index(c *fiber.Ctx) error {

	users, err := usersService.FindAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONResponse{Status: "error", Message: "Error in retriving users list", Data: err},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		JSONResponse{Status: "ok", Message: "users list retrived successfully", Data: &users},
	)

}

func findOne(c *fiber.Ctx) error {

	return c.
		Status(fiber.StatusBadGateway).
		JSON(
			JSONResponse{Status: "error", Message: "route still not ready.", Data: ""},
		)
}

func create(c *fiber.Ctx) error {

	var input usersService.CreateInput
	var user *models.User
	var err error

	if err = c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONResponse{Status: "error", Message: "Error in reading request body", Data: &err},
		)
	}

	if user, err = usersService.CreateUser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONResponse{Status: "error", Message: "Validation Error", Data: &err},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		JSONResponse{Status: "ok", Message: "User created successfully", Data: &user},
	)
}
