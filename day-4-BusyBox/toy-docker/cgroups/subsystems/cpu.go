package subsystems

type CpuSubSystem struct {

}


func (s *CpuSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	return nil
}

func (s *CpuSubSystem) Apply(cgroupPath string, pid int) error {
	return nil
}

func (s *CpuSubSystem) Name() string {
	return "cpu"
}

func (s *CpuSubSystem) Remove(cgoupPath string) error {
	return nil
}
