package sh_generator

import (
	"fmt"
	. "orca/internal/fio"
	services "orca/pkg/sh_generator/services"
	"strings"
)

type Argument struct {
	Required      bool
	IsEnvVariable bool
	Name          string
}

type Script struct {
	Name      string
	Arguments []Argument
	cmdBuffer []string
}

func NewScript(name string, args ...Argument) *Script {
	cmdBuffer := make([]string, 0)
	var script Script = Script{name, args, cmdBuffer}
	return &script
}

func (script *Script) Write(cmd string, args ...Argument) error {
	if args == nil {
		script.cmdBuffer = append(script.cmdBuffer, cmd)
		return nil
	}

	buf := make([]any, len(args))

	for i := 0; i < len(args); i++ {
		buf[i] = services.ToShellVarName(args[i].Name)
	}
	out := fmt.Sprintf(cmd, buf...)
	script.cmdBuffer = append(script.cmdBuffer, out)

	return nil
}

func (script *Script) Generate() (string, error) {
	builder := strings.Builder{}
	for i := 0; i < len(script.cmdBuffer); i++ {
		_, err := builder.WriteString(script.cmdBuffer[i] + "\n")
		if err != nil {
			return "", err
		}
	}

	var args []services.Argument
	for i := 0; i < len(script.Arguments); i++ {
		var arg services.Argument = services.Argument{
			Name:          script.Arguments[i].Name,
			Required:      script.Arguments[i].Required,
			IsEnvVariable: script.Arguments[i].IsEnvVariable,
		}
		args = append(args, arg)
	}

	template := services.ArgBashTemplate{
		Script:    builder.String(),
		Arguments: args,
	}

	out, err := template.Output()
	if err != nil {
		return "", err
	}

	path := "./out/templates/test.txt" // TODO make this a guid

	err = SaveToFile(path, out)
	if err != nil {
		return "", err
	}

	out, err = services.GenerateScript(path)
	if err != nil {
		return "", err
	}

	return out, nil
}

func (script *Script) SaveToFile(path string) error {
	content, err := script.Generate()
	if err != nil {
		return err
	}
	err = SaveToFile(path+script.Name+".sh", content)
	if err != nil {
		return err
	}
	return nil
}
