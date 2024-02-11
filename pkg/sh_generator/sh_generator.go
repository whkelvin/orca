package sh_generator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Argument struct {
	Required      bool
	IsEnvVariable bool
	Name          string
}

func (arg *Argument) ToString() string {
	return "$" + arg.Name
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
		buf[i] = args[i].ToString()
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
	return builder.String(), nil
}

func (script *Script) SaveToFile(path string) error {
	content, err := script.Generate()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path+script.Name+".sh", []byte(content), 0666)
	if err != nil {
		return errors.New("Error saving file")
	}
	return nil
}
