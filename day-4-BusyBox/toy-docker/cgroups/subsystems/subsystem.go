package subsystems

type ResourceConfig struct{
	MemoryLimit string
	CpuShare string
	CpuSet string
}

type Subsystem interface {
	// 返回名称
	Name() string
	// 设置某个cgroup在这个subsystem中的资源限制
	Set(path string,res *ResourceConfig) error
	// 将某个进程添加到某个
	Apply(path string,pid int) error
	// 移除
	Remove(path string) error
}

// 通过不同的subsystem初始化实力创建资源限制处理链数组
var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)