package pkg

import (
	"fmt"
	"strings"
)

type Job struct {
	configs          *JobConfig
	workingDirectory *string
	cmdExecutor      ICommandExecutor
}

type JobConfig struct {
	Image string
}

func NewJob(config *JobConfig, cmdExecutor ICommandExecutor) *Job {
	return &Job{config, nil, cmdExecutor}
}

func (job *Job) Init() {
	fmt.Println("initilizing job, starting a docker container")
}

func (job *Job) Connect() error {
	err := job.cmdExecutor.Connect()
	if err != nil {
		return err
	}
	return nil
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
	return out, nil
}

func (job *Job) Close() error {
	fmt.Println("finishing job, nuking docker container")
	return nil
}
