package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env
//
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for envvar, envvalue := range env {
		if envvalue == "" {
			os.Unsetenv(envvar)

			continue
		}

		err := os.Setenv(envvar, envvalue)
		if err != nil {
			fmt.Println(err)
		}
	}

	cmdExec := exec.Command(cmd[0], cmd[1:]...) // nolint:gosec
	cmdExec.Stdin = os.Stdin
	cmdExec.Stdout = os.Stdout
	_ = cmdExec.Run()

	return cmdExec.ProcessState.ExitCode()
}
