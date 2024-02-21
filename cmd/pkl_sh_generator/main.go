package main

import (
	"context"
	"orca/internal/gen/pkl/orca"
	. "orca/internal/logger"
	. "orca/pkg/sh_generator"
)

func main() {
	cfg, err := orca.LoadFromPath(context.Background(), "./assets/inputs/hello-world.pkl")
	if err != nil {
		Error(err.Error())
	}
	var args []Argument
	for i := 0; i < len(cfg.Arguments); i++ {
		var shortName string = ""
		if cfg.Arguments[i].ShortName != nil {
			shortName = *cfg.Arguments[i].ShortName
		}

		args = append(args, Argument{
			Name:          cfg.Arguments[i].Name,
			Mask:          cfg.Arguments[i].Mask,
			ShortName:     shortName,
			IsEnvVariable: cfg.Arguments[i].IsEnvVariable,
		})
	}

	script := Script{
		Name:      cfg.Name,
		Arguments: args,
		Commands:  cfg.Script,
	}

	err = script.SaveToFile("./out/")
	if err != nil {
		Error(err.Error())
	}
}
