//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/arsmn/fastgql/api"
	"github.com/arsmn/fastgql/codegen/config"
	"github.com/arsmn/fastgql/plugin/modelgen"
)

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {

			tag := ""

			//println(field.Type.String())
			println(field.Name, field.Description)
			if field.Type.String() == "github.com/google/uuid.UUID" {
				tag += "type:uuid;primaryKey;default:uuid_generate_v4()"
			}

			if field.Type.String() == "gorm.io/gorm.DeletedAt" {
				tag += "index"
			}

			if field.Type.String() == "*gorm.io/gorm.DeletedAt" {
				tag += "index"
			}

			if tag != "" {
				field.Tag += ` gorm:"` + tag + `"`
			}
		}
	}

	return b
}

func main() {
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
