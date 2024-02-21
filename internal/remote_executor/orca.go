package internal

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Job struct {
	configs          *JobConfig
	workingDirectory *string
	cmdExecutor      ICommandExecutor
	cmdHistory       []string
	containerId      *string
}

type JobConfig struct {
	Runner string
}

type JobArgument struct {
	Name       string
	ShouldMask bool
}

func NewJob(config *JobConfig) *Job {
	cmdHistory := make([]string, 0)

	var sshCmdExecutor ICommandExecutor = NewSSHCommandExecutor(&SSHCommandExecutorConfigs{
		Host:           "localhost",
		Port:           "2222",
		Username:       "root",
		PrivateKeyPath: "/home/whkelvin/.ssh/id_rsa_orca",
		Password:       "root123",
	})
	return &Job{config, nil, sshCmdExecutor, cmdHistory, nil}
}

// if you want to provide a cmd executor
func CreateJobWithCommandExecutor(config *JobConfig, cmdExecutor ICommandExecutor) *Job {
	cmdHistory := make([]string, 0)
	return &Job{config, nil, cmdExecutor, cmdHistory, nil}
}

func (job *Job) Init() error {
	fmt.Println("initilizing job, starting a docker container")
	if job.configs.Runner == "ubuntu" {
		containerId, err := StartContainer(IMAGE_OS_UBUNTU)
		if err != nil {
			return err
		}
		job.containerId = containerId
		return nil
	} else if job.configs.Runner == "node" {
		containerId, err := StartContainer(IMAGE_OS_NODE)
		if err != nil {
			return err
		}
		job.containerId = containerId
		return nil
	} else {
		return errors.New("Runner not supporter.")
	}
}

func (job *Job) Connect() error {
	retryCount := 5

	for i := 0; i < retryCount; i++ {
		err := job.cmdExecutor.Connect()
		if err != nil {
			fmt.Printf("Connecting ...")
			time.Sleep(2 * time.Second)
			continue
		}
		return nil
	}

	return errors.New("Connection Failed")
}

func (job *Job) PrintWorkingDirectory() (string, error) {
	if job.workingDirectory == nil || *job.workingDirectory == "" {
		dir, err := job.cmdExecutor.PrintWorkingDirectory()
		if err != nil {
			return "", err
		}
		return dir, nil
	}
	return *job.workingDirectory, nil
}

func (job *Job) ChangeDir(path string) (string, error) {
	res, err := job.cmdExecutor.Execute("", "cd "+path+"; pwd;")
	if err != nil {
		return "", err
	}

	currentDir := strings.Trim(res, "\n")
	job.workingDirectory = &currentDir

	job.cmdHistory = append(job.cmdHistory, "cd "+path)
	return *job.workingDirectory, nil
}

func (job *Job) wrapCmdWithWorkDir(cmd string) string {
	if job.workingDirectory == nil || *job.workingDirectory == "" {
		fmt.Println("running " + cmd)
		return cmd
	} else {
		fmt.Println("cd " + *job.workingDirectory + "; " + cmd)
		dir := strings.TrimSuffix(*job.workingDirectory, "\n")
		return "cd " + dir + "; " + cmd
	}
}

func (job *Job) Execute(cmd string) (string, error) {
	var workingDir string
	if job.workingDirectory == nil {
		workingDir = ""
	} else {
		workingDir = *job.workingDirectory
	}
	out, err := job.cmdExecutor.Execute(workingDir, cmd)
	if err != nil {
		return "", err
	}

	job.cmdHistory = append(job.cmdHistory, cmd)
	return out, nil
}

func (job *Job) GenerateScript() (string, error) {
	builder := strings.Builder{}
	for i := 0; i < len(job.cmdHistory); i++ {
		_, err := builder.WriteString(job.cmdHistory[i] + "\n")
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}

func (job *Job) Close() error {
	fmt.Println("finishing job, nuking docker container")
	if job.containerId != nil {
		err := StopContainer(*job.containerId)
		if err != nil {
			return err
		}
	}

	return nil
}
