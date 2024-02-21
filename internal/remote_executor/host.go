package pkg

import (
	"fmt"
	"os/exec"
)

func CreateDockerContainer() (string, error) {
	cmd := exec.Command("docker", "run", "-p", "22:22", "-d", "ssh-ubuntu")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return "", err
	}

	return string(output), nil
}
