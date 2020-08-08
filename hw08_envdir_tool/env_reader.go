package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := Environment{}

	for _, file := range files {
		fileName := dir + "/" + file.Name()
		if file.IsDir() || strings.Contains(fileName, "=") {
			continue
		}

		fd, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		value, err := ioutil.ReadAll(fd)
		if err != nil {
			return nil, err
		}

		fileLines := strings.Split(string(value), "\n")
		firstLine := strings.TrimRight(fileLines[0], " \t")

		firstLineCleaned := strings.Replace(firstLine, "\x00", "\n", -1)

		env[file.Name()] = firstLineCleaned
	}

	return env, nil
}
