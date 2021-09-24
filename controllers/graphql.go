package controllers

import (
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	"github.com/arsmn/fastgql/plugin/modelgen"
	"github.com/gofiber/fiber/v2"
	"github.com/nimerfarahty/go-rest/graph"
	"github.com/nimerfarahty/go-rest/graph/generated"
)

func Graphql(router fiber.Router) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := srv.Handler()
	playground := playground.Handler("GraphQL playground", "/query")

	router.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	router.All("/", func(c *fiber.Ctx) error {
		playground(c.Context())
		return nil
	})

}

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			field.Tag += ` orm_binding:"` + model.Name + `.` + field.Name + `"`
		}
	}

	return b
}
