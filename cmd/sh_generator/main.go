package main

import (
	"fmt"
	. "orca/pkg/sh_generator"
)

func main() {
	//arg := Argument{
	//	Required:      true,
	//	IsEnvVariable: false,
	//	Name:          "somename",
	//}
	//script := NewScript("test", arg)

	//script.Write("echo %s", arg)
	//script.Write("echo %s", arg)
	//script.Write("echo %s", arg)
	//script.Write("echo %s", arg)
	//script.Write("echo %s", arg)

	//sh, err := script.Generate()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(sh)
	//script.SaveToFile("./")

	tpl := Template{
		Script: "echo \"hello world\"",
	}
	err := tpl.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("done")
}
