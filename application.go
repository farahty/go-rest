package main

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
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

	api := app.Group("/api")

	users := api.Group("/users")
	controllers.UsersController(users)

	auth := api.Group("/auth")
	controllers.Authintication(auth)

	docs := api.Group("/docs")
	controllers.Swagger(docs)

	graphql := app.Group("/")
	controllers.Graphql(graphql)

}

func setupMeddlewares(app *fiber.App) {

	app.Use(requestid.New())
	app.Use(etag.New())

	app.Use(logger.New(logger.Config{
		//Format: "[${time}] ${status} - ${method} ${path} ${locals:requestid} ${latency}\n",
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))

	app.Get("/monitor", monitor.New())

}
