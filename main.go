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
		fmt.Println("CPU Temperature:", fmt.Sprintf("%.2f", temp), "°C")

		// wait for the configured interval before reading again
		time.Sleep(time.Duration(env.GetEnvAsInt("CPU_TEMPERATURE_INTERVAL", 1000)) * time.Millisecond)
	}
}
