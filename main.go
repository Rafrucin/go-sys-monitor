package main

import (
	"fmt"
	"time"

	"github.com/Rafrucin/sys-monitor/hardware"
)

func main() {
	fmt.Println("Starting system monitor...")
	go func() {
		for {
			sys, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
			}
			disk, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println(err)
			}
			cpu, err := hardware.GetCpuSection()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(sys)
			fmt.Println(disk)
			fmt.Println(cpu)

			time.Sleep(3*time.Second)
		}
	}()
	time.Sleep(5*time.Minute)
}