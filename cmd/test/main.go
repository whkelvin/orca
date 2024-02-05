package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// docker run -p 80:80  --start a docker container
	// docker exec CONTAINER sh -c "echo hello"
	cmd := exec.Command("docker", "run", "-d", "-p", "2222:22", "orca-ubuntu")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	respBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}
	respString := string(respBytes)
	fmt.Println("===STDOUT===")
	fmt.Printf(respString)
	fmt.Println("===STDOUT===")

	respBytes, err = ioutil.ReadAll(stderr)
	if err != nil {
		panic(err)
	}
	respString = string(respBytes)
	fmt.Println("===STDERR===")
	fmt.Printf(respString)
	fmt.Println("===STDERR===")

	//cmd = exec.Command("ls", "-la")

	//stdout, err = cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//stderr, err = cmd.StderrPipe()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//if err := cmd.Start(); err != nil {
	//	fmt.Println(err)
	//}

	//respBytes, err = ioutil.ReadAll(stdout)
	//if err != nil {
	//	panic(err)
	//}
	//respString = string(respBytes)
	//fmt.Println("===STDOUT===")
	//fmt.Printf(respString)
	//fmt.Println("===STDOUT===")

	//respBytes, err = ioutil.ReadAll(stderr)
	//if err != nil {
	//	panic(err)
	//}
	//respString = string(respBytes)
	//fmt.Println("===STDERR===")
	//fmt.Printf(respString)
	//fmt.Println("===STDERR===")
}
