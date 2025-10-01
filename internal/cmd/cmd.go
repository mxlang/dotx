package cmd

import (
	"os"
	"os/exec"
)

func Run(command string) error {
	cmd := exec.Command(command)

	cmd.Stdin = os.Stdin   // pass input from terminal/user
	cmd.Stdout = os.Stdout // print output to terminal
	cmd.Stderr = os.Stderr // print errors to terminal

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
