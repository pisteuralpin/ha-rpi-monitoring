package main

import (
	"os"
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
