package internal

import (
	"golang.org/x/crypto/ssh"
	"strings"
)

type ICommandExecutor interface {
	PrintWorkingDirectory() (string, error)
	Connect() error
	Execute(workingDirectory string, cmd string) (string, error)
}

type SSHCommandExecutor struct {
	configs   *SSHCommandExecutorConfigs
	sshClient *ssh.Client
}

type SSHCommandExecutorConfigs struct {
	Host           string
	Port           string
	Username       string
	PrivateKeyPath string
	Password       string
}

func NewSSHCommandExecutor(configs *SSHCommandExecutorConfigs) *SSHCommandExecutor {
	return &SSHCommandExecutor{configs, nil}
}

func (ce *SSHCommandExecutor) Connect() error {
	sshConfig := &ssh.ClientConfig{
		User: ce.configs.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(ce.configs.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Insecure, use for testing purposes only
	}

	client, err := ssh.Dial("tcp", ce.configs.Host+":"+ce.configs.Port, sshConfig)
	if err != nil {
		return err
	}

	ce.sshClient = client
	return nil
}

func (ce *SSHCommandExecutor) Execute(workingDirectory string, cmd string) (string, error) {
	session, err := ce.sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	out, err := session.CombinedOutput(prefixCmdWithWorkDir(workingDirectory, cmd))
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (ce *SSHCommandExecutor) PrintWorkingDirectory() (string, error) {
	session, err := ce.sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	out, err := session.Output("pwd")
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

func prefixCmdWithWorkDir(workingDirectory string, cmd string) string {
	if workingDirectory == "" {
		return cmd
	} else {
		return "cd " + workingDirectory + "; " + cmd
	}
}
