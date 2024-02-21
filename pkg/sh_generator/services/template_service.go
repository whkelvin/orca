package services

import (
	"bytes"
	"os"
	"strings"
	"text/template"
)

type Argument struct {
	Required      bool
	DefaultValue  string
	IsEnvVariable bool
	Name          string
	ShortName     string
	Mask          bool
}

func (a *Argument) ToShellString() string {
	if a.IsEnvVariable == false {
		a.Name = strings.ToLower(a.Name)
		a.Name = strings.ReplaceAll(a.Name, "-", "_")
		return "$_arg_" + a.Name
	} else {
		return "$" + a.Name
	}
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
