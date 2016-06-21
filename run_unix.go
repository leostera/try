//+build !windows

package main

import (
	"os/exec"
)

func _makeCmd(command []string) *exec.Cmd {
	path, _ := exec.LookPath(command[0])
	return exec.Command(path, command[1:]...)
}
