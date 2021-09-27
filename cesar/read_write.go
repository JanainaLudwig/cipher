package main

import (
	"io/ioutil"
	"os"
)

func Read(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(all), nil
}

func Write(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

