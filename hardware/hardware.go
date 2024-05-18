package hardware

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetSystemSection() (string, error) {
	runTimeOS := runtime.GOOS
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("HostName: %s, Totalt Memory: %.2f, Used Memory: %.2f, OS: %s",
		hostStat.Hostname, bytesToGB(vmStat.Total), bytesToGB(vmStat.Used), runTimeOS)
	return output, nil
}

func GetCpuSection() (string, error) {

	cpuStat, err := cpu.Info()
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("CPU: %s Cores: %d", cpuStat[0].ModelName, cpuStat[0].Cores)
	return output, nil
}

func GetDiskSection() (string, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("Total Disk Space: %.2f, Free Disk SPace: %.2f", bytesToGB(diskStat.Total), bytesToGB(diskStat.Free))
	return output, nil
}

func bytesToGB(bytes uint64) float64 {
	const bytesInGB = 1073741824.0
	return float64(bytes) / bytesInGB
}
