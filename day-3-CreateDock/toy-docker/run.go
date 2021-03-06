package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"toy-docker/cgroups"
	"toy-docker/cgroups/subsystems"
	"toy-docker/container"
)

func Run(tty bool, comArray []string,res *subsystems.ResourceConfig) {
	parent,writePipe := container.NewParentProcess(tty)
	if parent ==nil{
		log.Errorf("New parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	cgroupManager := cgroups.NewCgroupManager("toy-docker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(comArray,writePipe)
	parent.Wait()
	os.Exit(-1)
}

func sendInitCommand(comArray []string,writePipe *os.File){
	command := strings.Join(comArray," ")
	log.Infof("command all is %s",command)
	writePipe.WriteString(command)
	writePipe.Close()
}
