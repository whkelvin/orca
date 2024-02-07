package main

import (
	"fmt"
	//"github.com/labstack/echo/v4"
	//"net/http"
	. "poc/pkg"
)

//func initHandler(c echo.Context) error {
//	out, err := CreateDockerContainer()
//	if err != nil {
//		return c.String(http.StatusBadRequest, out)
//	}
//	return c.String(http.StatusOK, out)
//}

func main() {
	//e := echo.New()
	//e.POST("/init", initHandler)
	//e.Logger.Fatal(e.Start(":1323"))
	var sshCommandExecutor ICommandExecutor = NewSSHCommandExecutor(&SSHCommandExecutorConfigs{
		Host:           "localhost",
		Port:           "2222",
		Username:       "root",
		PrivateKeyPath: "/home/whkelvin/.ssh/id_rsa_orca",
		Password:       "root123",
	})

	var job Job = *CreateJobWithCommandExecutor(&JobConfig{Runner: "ubuntu"}, sshCommandExecutor)

	job.Init()
	err := job.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer job.Close()

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
