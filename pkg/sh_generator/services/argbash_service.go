package services

import (
	"errors"
	"io/ioutil"
	"os/exec"
)

func GenerateScript(templatePath string) (string, error) {
	cmd := exec.Command("argbash", templatePath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	respBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	resErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		panic(err)
	}

	if string(resErr) != "" {
		return "", errors.New(string(resErr))
	}
	return string(respBytes), nil
}
