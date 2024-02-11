package sh_generator

import (
	"os"
	"text/template"
)

type Template struct {
	Script string
}

func (t *Template) Output() error {
	txt, err := os.ReadFile("./assets/templates/argbash.txt")
	if err != nil {
		return err
	}

	tmpl, err := template.New("test").Parse(string(txt))
	if err != nil {
		return err
	}

	filename := "./" + "orca" + ".sh"
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, t)
	if err != nil {
		return err
	}

	return nil
}
