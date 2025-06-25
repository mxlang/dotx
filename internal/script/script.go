package script

import (
	"os"
	"os/exec"
)

func Run(scriptPath string) error {
	cmd := exec.Command("bash", scriptPath)

	cmd.Stdin = os.Stdin   // pass input from terminal/user
	cmd.Stdout = os.Stdout // print output to terminal
	cmd.Stderr = os.Stderr // print errors to terminal

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
