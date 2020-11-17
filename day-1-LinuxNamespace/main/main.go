package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// fork a thread to run sh
	cmd := exec.Command("sh")
	// UTS NAMESPACE and IPC NAMESPACE and PIC NAMESPACE
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags :syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |syscall.CLONE_NEWPID |syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
