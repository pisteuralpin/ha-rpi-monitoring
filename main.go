package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Raspberry Pi Monitoring...")

	for {
		temp, err := readCPUTemperature()
		if err != nil {
			fmt.Println("Error reading CPU temperature:", err)
			continue
		}
		fmt.Println("CPU Temperature:", fmt.Sprintf("%.2f", temp), "°C")

		// wait for the configured interval before reading again
		time.Sleep(1 * time.Second)
	}
}
