package main

import (
	"fmt"
	"ha-rpi-monitoring/v0.1/lib/env"
	"io"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// rawConfig represents the structure of the YAML configuration file.
type rawConfig struct {
	Debug          bool           `yaml:"debug"`
	HomeAssistant  hASection      `yaml:"homeassistant"`
	Device         deviceSection  `yaml:"device"`
	CpuTemperature metricsSection `yaml:"cpu_temperature"`
	PowerSupply    metricsSection `yaml:"power_supply"`
}

type hASection struct {
	Discovery bool   `yaml:"discovery"`
	Prefix    string `yaml:"prefix"`
}

type deviceSection struct {
	Id       string `yaml:"id"`
	UniqueId string `yaml:"unique_id"`
	Name     string `yaml:"name"`
	Model    string `yaml:"model"`
}

type metricsSection struct {
	Enabled    bool   `yaml:"enabled"`
	EntityName string `yaml:"entity_name"`
	Interval   string `yaml:"interval"`
}

// Config holds the application configuration.
type Config struct {
	MQTTConfig     mqttConfig
	HomeAssistant  hAConfig
	Device         deviceConfig
	CpuTemperature metricsConfig
	PowerSupply    metricsConfig
}

type hAConfig struct {
	Discovery bool
	Prefix    string
}

type deviceConfig struct {
	Id       string
	UniqueId string
	Name     string
	Model    string
}

type mqttConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type metricsConfig struct {
	Enabled    bool   // whether this metric is enabled
	EntityName string // entity name for the metric
	Interval   int    // in milliseconds
	Topic      string // MQTT topic to publish the metric
}

func parseInterval(intervalStr string) int {
	cpuInterval := 0
	if intervalStr[len(intervalStr)-1] == 's' {
		cpuInterval, _ = strconv.Atoi(intervalStr[:len(intervalStr)-1])
		cpuInterval *= 1000
	} else if intervalStr[len(intervalStr)-2:] == "ms" {
		cpuInterval, _ = strconv.Atoi(intervalStr[:len(intervalStr)-2])
	}
	return cpuInterval
}

var config Config

func createFileFromExample(destPath, examplePath string) error {
	exampleFile, err := os.Open(examplePath)
	if err != nil {
		return fmt.Errorf("failed to open example file: %v", err)
	}
	defer exampleFile.Close()

	dir := filepath.Dir(destPath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, exampleFile)
	if err != nil {
		return fmt.Errorf("failed to copy content: %v", err)
	}

	return nil
}

func initConfig() {
	configPath := env.GetEnv("CONFIG_PATH", "~/.config/ha-rpi-monitoring/config.yml")
	usr, _ := user.Current()
	configPath = strings.Replace(configPath, "~", usr.HomeDir, 1)
	examplePath := "example.config.yml"

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		err = createFileFromExample(configPath, examplePath)
		if err != nil {
			slog.Error("Failed to create config file from example", "error", err)
			panic(err)
		}
		slog.Info("Config file not found. Created example config at " + configPath)
	} else if err != nil {
		slog.Error("Failed to stat config file", "error", err)
		panic(err)
	}

	dat, err := os.ReadFile(configPath)
	if err != nil {
		slog.Error("Failed to read config file", "error", err, "path", configPath)
		panic(err)
	}
	slog.Debug("Config file loaded from " + configPath)

	var rawConfig rawConfig
	err = yaml.Unmarshal(dat, &rawConfig)
	if err != nil {
		slog.Error("Failed to parse config file", "error", err, "path", configPath)
		panic(err)
	}
	slog.Debug("Config file parsed successfully")

	config.HomeAssistant = hAConfig{
		Discovery: rawConfig.HomeAssistant.Discovery,
		Prefix:    rawConfig.HomeAssistant.Prefix,
	}

	config.Device = deviceConfig{
		Id:       rawConfig.Device.Id,
		UniqueId: rawConfig.Device.UniqueId,
		Name:     rawConfig.Device.Name,
		Model:    rawConfig.Device.Model,
	}

	config.MQTTConfig = mqttConfig{
		Host:     env.GetEnv("MQTT_HOST", "localhost"),
		Port:     env.GetEnvAsInt("MQTT_PORT", 1883),
		User:     env.GetEnv("MQTT_USER", "user"),
		Password: env.GetEnv("MQTT_PASSWORD", "password"),
	}

	config.CpuTemperature = metricsConfig{
		Enabled:  rawConfig.CpuTemperature.Enabled,
		Interval: parseInterval(rawConfig.CpuTemperature.Interval),
		Topic:    config.HomeAssistant.Prefix + "/temperature/" + config.Device.Id + "/" + rawConfig.CpuTemperature.EntityName + "/state",
	}

	if config.CpuTemperature.Enabled {
		slog.Info("CPU Temperature monitoring enabled.", "interval_ms", config.CpuTemperature.Interval, "topic", config.CpuTemperature.Topic)
	} else {
		slog.Info("CPU Temperature monitoring disabled.")
	}

	config.PowerSupply = metricsConfig{
		Enabled:  rawConfig.PowerSupply.Enabled,
		Interval: parseInterval(rawConfig.PowerSupply.Interval),
		Topic:    config.HomeAssistant.Prefix + "/power/" + config.Device.Id + "/" + rawConfig.PowerSupply.EntityName + "/state",
	}

	if config.PowerSupply.Enabled {
		slog.Info("Power Supply monitoring enabled.", "interval_ms", config.PowerSupply.Interval, "topic", config.PowerSupply.Topic)
	} else {
		slog.Info("Power Supply monitoring disabled.")
	}
}
