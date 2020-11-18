package main

import (
	"fmt"
	"os/exec"
	"path"
	"os"
	"io/ioutil"
	"syscall"
	"strconv"
)

const cgroupMemoryHierachyMount = "/sys/fs/cgroup/memory"

func main() {
	fmt.Println(os.Args)
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 50m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{

		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}else{
		fmt.Printf("%d",cmd.Process.Pid)
		os.Mkdir(path.Join(cgroupMemoryHierachyMount,"testmemorylimit"),0755)
		// add current pid to limit
		ioutil.WriteFile(path.Join(cgroupMemoryHierachyMount,"testmemorylimit","tasks"),[]byte(strconv.Itoa(cmd.Process.Pid)),0644)
		// add 100m limit to this cgroup
		ioutil.WriteFile(path.Join(cgroupMemoryHierachyMount,"testmemorylimit","memory.limit_in_bytes"),[]byte("100m"),0644)
	}
	cmd.Process.Wait()
}
