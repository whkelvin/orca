package main

import (
	"fmt"
	. "poc/pkg"
	"time"
)

func main() {
	var job Job = *NewJob(&JobConfig{Image: "ubuntu"})

	job.Init()

	retryCount := 5

	for i := 0; i < retryCount; i++ {
		err := job.Connect()
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Retry: %v / %v \n", i, retryCount)
			time.Sleep(2 * time.Second)
			continue
		}
		defer job.Close()
		break
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
