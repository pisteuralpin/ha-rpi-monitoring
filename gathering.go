package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func readCPUTemperature() (float32, error) {

	// Read temperature from system file
	var dat, err = os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	// Clean data from whitespace and newlines
	var cleanedData = strings.ReplaceAll(string(dat), "\n", "")
	cleanedData = strings.ReplaceAll(cleanedData, " ", "")

	// Convert to float and scale
	var temperature, err2 = strconv.ParseInt(cleanedData, 10, 32)
	if err2 != nil {
		return 0, err2
	}

	return float32(temperature) / 1000, nil
}

func readPowerSupply() (float32, error) {
	totalPower := float32(0)

	// Read using pmic_read_adc command
	var output, err = exec.Command("vcgencmd", "pmic_read_adc").Output()
	if err != nil {
		return 0, err
	}

	// Parse output
	// Example output: "VDD_CPU: 1.200000 V\n"
	var parts = strings.Split(string(output), "\n")
	for i := range 12 {
		ampStr := strings.Split(strings.Split(parts[i], "=")[1], "A")[0]
		amperage, err2 := strconv.ParseFloat(ampStr, 32)
		if err2 != nil {
			return 0, err2
		}
		voltStr := strings.Split(strings.Split(parts[12+i], "=")[1], "V")[0]
		voltage, err3 := strconv.ParseFloat(voltStr, 32)
		if err3 != nil {
			return 0, err3
		}
		totalPower += float32(amperage) * float32(voltage)
	}

	return totalPower, nil
}
