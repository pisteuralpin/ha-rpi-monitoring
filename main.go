package main

import (
	"fmt"
	"ha-rpi-monitoring/v0.1/lib/env"
	"log/slog"
	"time"
)

func main() {
	env.LoadEnv()

	initLogger()

	slog.Info("Starting Raspberry Pi Monitoring...")

	initConfig()

	fmt.Println("✅ Raspberry Pi Monitoring started.")

	if config.HomeAssistant.Discovery {
		initDiscovering()
		slog.Info("Home Assistant discovery messages sent.")
	}

	// CPU Temperature
	if config.CpuTemperature.Enabled {
		go monitorCPUTemperature()
	}

	// Power Supply Voltage
	if config.PowerSupply.Enabled {
		go monitorPowerSupply()
	}

	// Prevent the main function from exiting
	select {}
}

func monitorCPUTemperature() {
	for {
		temp, err := readCPUTemperature()
		if err != nil {
			slog.Warn("Error reading CPU Temperature:", "error", err)
			continue
		}
		sendViaMQTT(config.CpuTemperature.Topic, fmt.Sprintf("%.2f", temp))

		// wait for the configured interval before reading again
		time.Sleep(time.Duration(config.CpuTemperature.Interval) * time.Millisecond)
	}
}

func monitorPowerSupply() {
	for {
		power, err := readPowerSupply()
		if err != nil {
			slog.Warn("Error reading Power Supply:", "error", err)
			continue
		}
		power *= 1.1451
		power += 0.5879
		sendViaMQTT(config.PowerSupply.Topic, fmt.Sprintf("%.2f", power))

		// wait for the configured interval before reading again
		time.Sleep(time.Duration(config.PowerSupply.Interval) * time.Millisecond)
	}
}
