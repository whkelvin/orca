package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

func main() {
	host := "192.168.2.121"
	port := "22"
	user := "whkelvin"
	privateKeyPath := "/home/whkelvin/.ssh/id_rsa_orca"
	password := "your password here"

	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create the SSH configuration
	signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, []byte(password))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Insecure, use for testing purposes only
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", host+":"+port, sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer client.Close()

	// Example: Run a command on the remote server
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Replace "your_command" with the command you want to execute
	output, err := session.CombinedOutput("ls -la")
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	fmt.Println(string(output))

}
