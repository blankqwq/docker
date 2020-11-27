package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	log "github.com/sirupsen/logrus"
)

func pivotRoot(root string) error {
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC); err != nil {
		return fmt.Errorf("Mount rootfs to itself error: %v",err)
	}
	// 生成pivot路径
	pivotDir := filepath.Join(root,".pivot_root")
	if err:=os.Mkdir(pivotDir,0777); err!=nil {
		return err
	}
	//pivot_root 到新的rootfs 老大old_root现在挂在到rootfs/.pivot_root上
	// 挂载点目前任然可以在mount命令中查看到
	if err := syscall.PivotRoot(root,pivotDir); err!=nil{
		return  fmt.Errorf("pivot_root %v",err)
	}

	// 切换目录
	if err:=syscall.Chdir("/");err!=nil {
		return fmt.Errorf("chdir / %v",err)
	}
	pivotDir = filepath.Join("/",".pivot_root")
	// unmount rootfs/.pivot_root
	if err:=syscall.UnMount(pivotDir,syscall.MNT_DETCH);err!=nil {
		return fmt.Errorf("unmount pivot_root dir %v",err)
	}
	// 删除相关目录
	return os.Remove(pivotDir)
}


// UFS 和 AUFS
func setUpMount(){
	// 获取当前路径
	pwd,err := os.Getwd()
	if err!=nil{
		log.Errorf("Get current location error %v",err);
		return
	}
	log.Infof("Current location is %s",pwd)
	pivotRoot(pwd)

	// mount proc
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID |syscall.MS_NODEV
	syscall.Mount("proc","/proc",uintptr(defaultMountFlags),"")
	syscall.Mount("tmpfs","/dev","tmpfs",syscall.MS_NOSUID | syscall.MS_STRICTATIME,"mode=755")
}


func main(){
	setUpMount()
}


