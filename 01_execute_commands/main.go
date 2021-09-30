package main

import (
	"log"
	"os"
	"os/exec"
)

func excuteCommand(command string, args_arr []string) (err error) {
	args := args_arr
	cmd_object := exec.Command(command, args...)

	cmd_object.Stderr = os.Stderr
	cmd_object.Stdout = os.Stdout

	err = cmd_object.Run()

	if err != nil {
		log.Fatal(err)
		return
	}

	return nil
}

func main() {
	command := "ls"

	excuteCommand(command, []string{"-lart"})
}
