package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("too few arguments")
	}

	dir, err := os.Stat(args[1])
	if err != nil {
		panic(err)
	}
	if !dir.IsDir() {
		panic("First argument is not a dir")
	}

	env, err := ReadDir(args[1])
	if err != nil {
		panic(err)
	}
	rc := RunCmd(args[2:], env)

	os.Exit(rc)
}
