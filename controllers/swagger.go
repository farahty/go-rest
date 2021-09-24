package controllers

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Swagger(router fiber.Router) {

	router.Get("/*", swagger.Handler)
}
