package services

import (
	"bytes"
	"os"
	"text/template"
)

type Argument struct {
	Required      bool
	IsEnvVariable bool
	Name          string
}

func (arg *Argument) ToShellVarName() string {
	return ToShellVarName(arg.Name)
}

func ToShellVarName(name string) string {
	return "$_arg_" + name
}

type ArgBashTemplate struct {
	Arguments []Argument
	Script    string
}

func (t *ArgBashTemplate) Output() (string, error) {
	txt, err := os.ReadFile("./assets/templates/argbash.txt")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("orca").Parse(string(txt))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, t)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
