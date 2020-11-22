package subsystems

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct {
}

func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				log.Infof("limit error: %v  value: %v", err, []byte(res.MemoryLimit))
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}

}

func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	if SubsystemCgoupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if err := ioutil.WriteFile(path.Join(SubsystemCgoupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644);
			err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}
}

func (s *MemorySubSystem) Name() string {
	return "memory"
}

func (s *MemorySubSystem) Remove(cgoupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgoupPath, false); err == nil {
		return os.Remove(subsysCgroupPath)
	} else {
		return err
	}
}
