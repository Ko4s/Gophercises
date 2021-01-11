package utils

import (
	"io/ioutil"
)

func OpenAndReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return content, nil
}
