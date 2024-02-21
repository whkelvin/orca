package remote_executor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	IMAGE_OS_UBUNTU = iota // 0
	IMAGE_OS_NODE
)

func StopContainer(containerId string) error {
	fmt.Println(containerId)
	cmd := exec.Command("docker", "rm", "-f", containerId)
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
	stdoutString := string(respBytes)
	if stdoutString != "" {
		fmt.Println("===STDOUT===")
		fmt.Printf(stdoutString)
		fmt.Println("===STDOUT===")
	}

	respBytes, err = ioutil.ReadAll(stderr)
	if err != nil {
		panic(err)
	}
	stderrString := string(respBytes)
	if stderrString != "" {
		fmt.Println("===STDERR===")
		fmt.Printf(stderrString)
		fmt.Println("===STDERR===")
	}
	return nil
}

func StartContainer(imageOS int) (*string, error) {
	switch imageOS {
	default:
		return nil, errors.New("ImageOS not supported")
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
		stdoutString := string(respBytes)
		if stdoutString != "" {
			fmt.Println("===STDOUT===")
			fmt.Printf(stdoutString)
			fmt.Println("===STDOUT===")
		}

		respBytes, err = ioutil.ReadAll(stderr)
		if err != nil {
			panic(err)
		}
		stderrString := string(respBytes)
		if stderrString != "" {
			fmt.Println("===STDERR===")
			fmt.Printf(stderrString)
			fmt.Println("===STDERR===")
		}

		out := strings.TrimSuffix(string(stdoutString), "\n")
		return &out, nil
	}
}
