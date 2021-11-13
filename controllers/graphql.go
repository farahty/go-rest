package controllers

import (
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/nimerfarahty/go-rest/controllers/directives"
	"github.com/nimerfarahty/go-rest/generated"
	"github.com/nimerfarahty/go-rest/resolvers"
)

func Graphql(router fiber.Router) {

	config := generated.Config{Resolvers: &resolvers.Resolver{}}

	config.Directives.HasRole = directives.HasRole

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	gqlHandler := srv.Handler()

	playground := playground.Handler("GraphQL playground", "/query")

	router.All("/query", func(c *fiber.Ctx) error {
		setUserData(c)
		gqlHandler(c.Context())
		return nil
	})

	router.All("/", func(c *fiber.Ctx) error {
		playground(c.Context())
		return nil
	})

}
