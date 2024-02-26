package orca_cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"orca/internal/gen/pkl/orca"
	. "orca/internal/logger"
	. "orca/internal/sh_generator"
	"os"
)

var (
	// Command definition
	rootCmd = &cobra.Command{
		Use:   "orca [-o|--out] <pkl file>",
		Short: "orca is a tool that helps you write better shell script.",
		Long:  "",
	}
)

var argValue string

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&argValue, "out", "o", "", "output path")
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if argValue != "" {
			Info(fmt.Sprintf("--out: %s", argValue))
		}

		if args[0] != "" {
			Info(fmt.Sprintf("input file: %s", args[0]))
		}

		if argValue == "" && args[0] == "" {
			Info("No arguments provided.")
		}

		Info("Generating Script...")
		err := GenerateShellScriptFromPkl(argValue, args[0])
		if err != nil {
			panic(err.Error())
		}
		Info("Done!")
	}

	if err := rootCmd.Execute(); err != nil {
		Error(err.Error())
		os.Exit(1)
	}
}

func GenerateShellScriptFromPkl(outputPath string, inputFilePath string) error {
	cfg, err := orca.LoadFromPath(context.Background(), inputFilePath)
	if err != nil {
		return err
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
			IsRequired:    cfg.Arguments[i].IsRequired,
		})
	}

	script := Script{
		Name:      cfg.Name,
		Arguments: args,
		Commands:  cfg.Script,
	}

	err = script.SaveToFile(outputPath)
	if err != nil {
		return err
	}
	return nil
}
