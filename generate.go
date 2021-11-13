package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/arsmn/fastgql/api"
	"github.com/arsmn/fastgql/codegen/config"
	"github.com/arsmn/fastgql/plugin/modelgen"
)

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {

			tags := []string{}

			if strings.HasPrefix(field.Description, "#orm:") {
				tags = append(tags, getORM(field.Description)...)
				field.Description = ""
			}

			if field.Name == "deletedAt" {
				tags = append(tags, "index")
			}

			if len(tags) > 0 {
				field.Tag += ` gorm:"` + strings.Join(tags, ";") + `"`
			}
		}
	}

	return b
}

func getORM(str string) []string {

	orm := []string{}

	parts := strings.Split(strings.TrimPrefix(str, "#orm:"), ";")

	for _, item := range parts {
		switch item {
		case "pk":
			orm = append(orm, "type:serial;primaryKey")
		default:
			orm = append(orm, item)
		}
	}

	return orm
}

func Generate() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
