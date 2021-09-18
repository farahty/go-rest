package main

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nimerfarahty/go-rest/config"
	"github.com/nimerfarahty/go-rest/controllers"
)

func runApp() error {

	app := fiber.New()
	addr := fmt.Sprintf(":%s", config.Keys.Server.Port)

	setupMeddlewares(app)
	setupRoutes(app)

	return app.Listen(addr)
}

func setupRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.Handler)
	api := app.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	auth := v1.Group("/auth")

	controllers.UsersController(users)
	controllers.Authintication(auth)
}

func setupMeddlewares(app *fiber.App) {

	app.Use(requestid.New())
	app.Use(etag.New())

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} ${locals:requestid} ${latency}\n",
	}))

	app.Get("/monitor", monitor.New())

}
