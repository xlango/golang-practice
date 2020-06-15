package main

import (
	"fmt"
	"os"

	"github.com/yumaojun03/dmidecode"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dmi, err := dmidecode.New()
	checkError(err)

	infos, err := dmi.Onboard()
	// 支持以下类型的解析
	// dmi.BaseBoard()
	// dmi.Chassis()
	// dmi.MemoryArray()
	// dmi.MemoryDevice()
	// dmi.Onboard()
	// dmi.PortConnector()
	// dmi.Processor()
	// dmi.ProcessorCache()
	// dmi.Slot()
	// dmi.System()
	checkError(err)

	for _, v := range infos {
		fmt.Println(v, "=============")
	}
}
