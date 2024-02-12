package fileio

import (
	"errors"
	"io/ioutil"
)

func SaveToFile(path string, content string) error {
	err := ioutil.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return errors.New("Error saving file: " + path)
	}
	return nil
}
