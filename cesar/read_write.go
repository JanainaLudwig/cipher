package main

import (
	"io/ioutil"
	"os"
	//"unicode"

	//"unicode"

	//"golang.org/x/text/secure/precis"
	//"golang.org/x/text/transform"
	//"golang.org/x/text/unicode/norm"
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

	//t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//result, _, _ := transform.String(t, "žůžo")

	return string(all), nil
}

func Write(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

