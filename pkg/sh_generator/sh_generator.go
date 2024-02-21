package sh_generator

import (
	. "orca/internal/fio"
	services "orca/pkg/sh_generator/services"
	"strings"
)

type Argument struct {
	IsEnvVariable bool
	Name          string
	ShortName     string
	Mask          bool
}

func (a *Argument) GetName() string {
	return a.Name
}

type Script struct {
	Name      string
	Arguments []Argument
	Commands  []string
}

func (script *Script) ToShellString() string {
	builder := strings.Builder{}
	for i := 0; i < len(script.Commands); i++ {
		_, err := builder.WriteString(script.Commands[i] + "\n")
		if err != nil {
			panic(err.Error())
		}
	}

	var args []services.Argument
	for i := 0; i < len(script.Arguments); i++ {
		var arg services.Argument = services.Argument{
			Name:          script.Arguments[i].Name,
			IsEnvVariable: script.Arguments[i].IsEnvVariable,
			ShortName:     script.Arguments[i].ShortName,
			Mask:          script.Arguments[i].Mask,
		}
		args = append(args, arg)
	}

	template := services.ArgBashTemplate{
		Script:    builder.String(),
		Arguments: args,
	}

	out, err := template.Output()
	if err != nil {
		panic(err.Error())
	}

	path := "./out/templates/test.txt" // TODO make this a guid

	err = SaveToFile(path, out)
	if err != nil {
		panic(err.Error())
	}

	out, err = services.GenerateScript(path)
	if err != nil {
		panic(err.Error())
	}

	return out
}

func (script *Script) SaveToFile(path string) error {
	content := script.ToShellString()
	err := SaveToFile(path+script.Name+".sh", content)
	if err != nil {
		return err
	}
	return nil
}
