package subsystems

type CpusetSubSystem struct {

}


func (s *CpusetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	return nil
}

func (s *CpusetSubSystem) Apply(cgroupPath string, pid int) error {
	return nil
}

func (s *CpusetSubSystem) Name() string {
	return "cpuset"
}

func (s *CpusetSubSystem) Remove(cgoupPath string) error {
	return nil
}
