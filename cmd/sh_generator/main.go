package main

import (
	. "orca/internal/logger"
	. "orca/pkg/sh_generator"
)

func main() {
	username := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "username",
	}

	password := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "password",
	}

	path := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "path",
	}

	imageName := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "image_name",
	}

	tag := Argument{
		Required:      true,
		IsEnvVariable: false,
		Name:          "image_tag",
	}

	script := NewScript("build_docker_img_and_publish_to_docker_hub", username, password, path, imageName, tag)

	script.Write("docker_build_result", `docker image build -t %s:%s -f %s .`, &imageName, &tag, &path)
	script.Write("docker_login", `docker login -u %s -p %s`, &username, &password)
	script.Write("docker_tag", `docker tag %s %s/%s`, &imageName, &username, &imageName)
	script.Write("docker_push", `docker image push %s/%s`, &username, &imageName)

	err := script.SaveToFile("./out/")
	if err != nil {
		Error(err.Error())
	}
}
