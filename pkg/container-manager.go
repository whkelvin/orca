package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

const (
	IMAGE_OS_UBUNTU = iota // 0
)

func StartContainer(imageOS int) error {
	switch imageOS {
	default:
		return errors.New("ImageOS not supported")
	case IMAGE_OS_UBUNTU:
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
		if respString != "" {
			fmt.Println("===STDOUT===")
			fmt.Printf(respString)
			fmt.Println("===STDOUT===")
		}

		respBytes, err = ioutil.ReadAll(stderr)
		if err != nil {
			panic(err)
		}
		respString = string(respBytes)
		if respString != "" {
			fmt.Println("===STDERR===")
			fmt.Printf(respString)
			fmt.Println("===STDERR===")
		}
		return nil
	}
}
