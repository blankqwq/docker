package main

import (
	_ "fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// fork a thread to run sh
	cmd := exec.Command("sh")
	// UTS NAMESPACE and IPC NAMESPACE and PIC NAMESPACE and USER NAMESPACE and NETWORK
	//centos 7 use USER NAMESPACE must run grubby --args="user_namespace.enable=1" --update-kernel="$(grubby --default-kernel)"
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET| syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      syscall.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      syscall.Getgid(),
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
