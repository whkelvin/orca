package main

import (
	. "orca/internal/logger"
	. "orca/pkg/sh_generator"
)

func main() {
	testEnv := Argument{
		IsEnvVariable: true,
		Name:          "TEST_ENV",
		ShortName:     "",
		Mask:          true,
	}
	username := Argument{
		IsEnvVariable: false,
		Name:          "username",
		ShortName:     "o",
		Mask:          false,
	}

	password := Argument{
		IsEnvVariable: false,
		Name:          "password",
		Mask:          false,
	}

	path := Argument{
		IsEnvVariable: false,
		Name:          "path",
		Mask:          false,
	}

	imageName := Argument{
		IsEnvVariable: false,
		Name:          "image-name",
		Mask:          false,
	}

	tag := Argument{
		IsEnvVariable: false,
		Name:          "image-tag",
		Mask:          false,
	}

	script := Script{
		Name:      "build_docker_img_and_publish_to_docker_hub",
		Arguments: []Argument{testEnv, username, password, path, imageName, tag},
		Commands: []string{
			`docker image build -t "$_arg_image_name":"$_arg_tag" -f $_arg_path .`,
			`docker login -u "$_arg_username" -p "$_arg_password"`,
			`docker tag "$_arg_image_name" "$_arg_username"/"$_arg_image_name"`,
			`docker image push "$_arg_username"/"$_arg_image_name"`,
		},
	}

	err := script.SaveToFile("./out/")
	if err != nil {
		Error(err.Error())
	}
}
