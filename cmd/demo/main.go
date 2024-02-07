package main

import (
	"fmt"
	. "poc/pkg"
)

func main() {
	var job Job = *NewJob(&JobConfig{Runner: "ubuntu"})

	job.Init()

	err := job.Connect()
	defer job.Close()

	if err != nil {
		return
	}

	wd, err := job.PrintWorkingDirectory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("current directory is: " + wd)

	fmt.Println("running cmd: ls -la")
	out, err := job.Execute("ls -la")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	fmt.Println("running cmd: mkdir test2")
	out, err = job.Execute("mkdir -p test2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	fmt.Println("running cmd: cd test2")
	out, err = job.ChangeDir("test2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("running cmd: pwd")
	out, err = job.PrintWorkingDirectory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	out, err = job.GenerateScript()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nhere is your shell script!!!")
	fmt.Println("=============================")
	fmt.Print(out)
	fmt.Println("=============================")

}
