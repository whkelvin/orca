package main

import (
	. "orca/internal/logger"
	. "orca/internal/sh_generator"
)

func main() {
	testEnv := Argument{
		IsEnvVariable: true,
		Name:          "TEST_ENV",
		ShortName:     "",
		Mask:          false,
	}
	username := Argument{
		IsEnvVariable: false,
		Name:          "username",
		ShortName:     "o",
		Mask:          true,
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
		Name:      "orca", //"build_docker_img_and_publish_to_docker_hub",
		Arguments: []Argument{testEnv, username, password, path, imageName, tag},
		Commands: []string{
			`echo "hello $_arg_username"`,
			`echo "hello $_arg_password"`,
			`echo "hello $_arg_path"`,
			`echo "hello $_arg_image_name"`,
			`echo "hello $_arg_image_tag"`,
		},
	}

	err := script.SaveToFile("./out/")
	if err != nil {
		Error(err.Error())
	}
}
