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

func (a *Argument) ToShellVarReference() string {
	tmp := a.Name
	if a.IsEnvVariable == false {
		tmp = strings.ToLower(tmp)
		tmp = strings.ReplaceAll(tmp, "-", "_")
		return "$_arg_" + tmp
	} else {
		return "$" + tmp
	}
}

func (a *Argument) ToShellVarName() string {
	tmp := a.Name
	if a.IsEnvVariable == false {
		tmp = strings.ToLower(tmp)
		tmp = strings.ReplaceAll(tmp, "-", "_")
		return "_arg_" + tmp
	} else {
		return ""
	}
}

type ScriptTemplate struct {
	Arguments []Argument
	Script    string
}

func (t *ScriptTemplate) Output() (string, error) {
	txt, err := os.ReadFile("./assets/templates/script.txt")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("script").Parse(string(txt))
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
