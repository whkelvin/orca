package sh_generator

import (
	"fmt"
	. "orca/internal/fio"
	services "orca/pkg/sh_generator/services"
	"strings"
)

type Variable interface {
	GetName() string
	ToShellString() string
}

type Argument struct {
	Required      bool
	IsEnvVariable bool
	Name          string
}

func (a *Argument) ToShellString() string {
	return "$_arg_" + a.Name
}
func (a *Argument) GetName() string {
	return a.Name
}

type Output struct {
	Name string
}

func (v *Output) ToShellString() string {
	return "$" + v.Name
}
func (a *Output) GetName() string {
	return a.Name
}

func NewResultVariable(name string) string {
	if name == "" {
		return "orca" + "_result"
	}
	return "orca_" + name + "_result"
}

type Script struct {
	Name      string
	Arguments []Argument
	cmdBuffer []string
}

type Cmd struct {
	Cmd  string
	Args []Argument
}

func NewScript(name string, args ...Argument) *Script {
	cmdBuffer := make([]string, 0)
	var script Script = Script{name, args, cmdBuffer}
	return &script
}

func (script *Script) Write(outputVariableName string, cmd string, args ...Variable) Output {
	resultVar := NewResultVariable(outputVariableName)

	if args == nil {
		out := cmd
		out = fmt.Sprintf("%s=$(%s)", resultVar, out)
		script.cmdBuffer = append(script.cmdBuffer, out)
		return Output{Name: resultVar}
	}

	buf := make([]any, len(args))

	for i := 0; i < len(args); i++ {
		buf[i] = services.ToShellVarName(args[i].GetName())
	}
	out := fmt.Sprintf(cmd, buf...)

	out = fmt.Sprintf("%s=$(%s)", resultVar, out)
	script.cmdBuffer = append(script.cmdBuffer, out)

	return Output{Name: resultVar}
}

func (script *Script) ToShellString() string {
	builder := strings.Builder{}
	for i := 0; i < len(script.cmdBuffer); i++ {
		_, err := builder.WriteString(script.cmdBuffer[i] + "\n")
		if err != nil {
			panic(err.Error())
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

//type ComparisonOperator stuct{
//  Type string
//}
//
//type Condition struct{
//  Val1 Variable
//  Operator ComparisonOperator
//  Val2 Variable
//}
//
//type IfConditional struct {
//	Condition Condition
//	Content   string
//	Else      *ElseConditional
//}
//
//type ElseConditional struct {
//	Condition Condition
//	Content   string
//}
//
//func If(condition Condition) IfConditional     {}
//func Else(condition Condition) ElseConditional {}
//
//func And(c1 Condition, c2 Condition) Condition {}
