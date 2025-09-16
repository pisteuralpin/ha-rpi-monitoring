package main

import (
	"fmt"
	"ha-rpi-monitoring/v0.1/lib/env"
	"time"
)

func main() {
	fmt.Println("Starting Raspberry Pi Monitoring...")

	env.LoadEnv()

	initConfig()

	fmt.Println("✅ Raspberry Pi Monitoring started.")

	// CPU Temperature
	if env.GetEnvAsBool("CPU_TEMPERATURE_ENABLED", false) {
		go monitorCPUTemperature()
	}

	// Prevent the main function from exiting
	select {}
}

func monitorCPUTemperature() {
	for {
		temp, err := readCPUTemperature()
		if err != nil {
			fmt.Println("Error reading CPU temperature:", err)
			continue
		}
		sendViaMQTT(cpuTemperatureConfig.Topic, fmt.Sprintf("%.2f", temp))

		// wait for the configured interval before reading again
		time.Sleep(time.Duration(cpuTemperatureConfig.Interval) * time.Millisecond)
	}
}
