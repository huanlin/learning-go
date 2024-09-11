package hardware

import (
	"fmt"
	"log"
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetSystemInfo() (string, error) {
	os := runtime.GOOS
	osArch := runtime.GOARCH

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}

	output := fmt.Sprintf(
		"OS: %s\nArch: %s\nTotal Memory: %v\nFree Memory: %v\nUsed Memory: %v\n",
		os, osArch, vmStat.Total, vmStat.Free, vmStat.Used)

	return output, nil
}

func GetCPUInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("CPU Info:")
	for _, info := range cpuInfo {
		fmt.Println("Model:", info.ModelName)
		fmt.Println("Cores:", info.Cores)
	}
}
