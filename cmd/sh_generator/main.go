package main

import (
	. "orca/internal/logger"
	. "orca/pkg/sh_generator"
)

func main() {
	name := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "name",
	}

	adj := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "adj",
	}
	script := NewScript("orca", name, adj)

	script.Write("echo \"Hello %s, you are %s !!!\"", name, adj)

	err := script.SaveToFile("./out/")
	if err != nil {
		Error(err.Error())
	}
}
