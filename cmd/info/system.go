package info

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `json:"hostname"`
	Platform string `json:"platform"`
	CPU      string `json:"cpu"`
	RAM      uint64 `json:"ram"`
	Disk     uint64 `json:"disk"`
}

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Show system info",
	Long:  "Show system info description",
	Run: func(cmd *cobra.Command, args []string) {
		getSystemInfo()
	},
}

func init() {
	systemCmd.Flags().BoolP("help", "h", false, "show help message")
}

func getSystemInfo() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/") // If you're in Windows change this "/" for "\\"

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024
	fmt.Println("System information: ")
	fmt.Printf("Hostname: %+v\n", info.Hostname)
	fmt.Printf("Platform: %+v\n", info.Platform)
	fmt.Printf("CPU: %+v\n", info.CPU)
	fmt.Printf("RAM: %+v MB\n", info.RAM)
	fmt.Printf("DISC: %+v MB\n", info.Disk)
}
