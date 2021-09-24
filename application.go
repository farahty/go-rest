package main

import (
	"fmt"

	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nimerfarahty/go-rest/config"
	"github.com/nimerfarahty/go-rest/controllers"

	"github.com/nimerfarahty/go-rest/graph"
	"github.com/nimerfarahty/go-rest/graph/generated"
)

func runApp() error {

	app := fiber.New()
	addr := fmt.Sprintf(":%s", config.Keys.Server.Port)

	setupMeddlewares(app)
	setupRoutes(app)

	return app.Listen(addr)
}

func setupRoutes(app *fiber.App) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := srv.Handler()
	playground := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		playground(c.Context())
		return nil
	})

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
