package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func main() {
	kversion, _ := host.KernelVersion()
	fmt.Println(kversion)
	platform, family, version, _ := host.PlatformInformation()
	fmt.Println(platform)
	fmt.Println(family)
	fmt.Println(version)
	vir1, vir2, _ := host.Virtualization()
	fmt.Println(vir1)
	fmt.Println(vir2)
	info, _ := host.Info()
	fmt.Println(info)
	cpuInfo, _ := cpu.Info()
	fmt.Println(cpuInfo[0])
}
