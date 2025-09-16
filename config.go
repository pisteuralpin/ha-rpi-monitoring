package main

import (
	"ha-rpi-monitoring/v0.1/lib/env"
	"strconv"
)

type MqttConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Prefix   string
}

type MetricsConfig struct {
	Enabled  bool   // whether this metric is enabled
	Interval int    // in milliseconds
	Topic    string // MQTT topic to publish the metric
}

var mqttCredentials MqttConfig
var cpuTemperatureConfig MetricsConfig

func initConfig() {
	mqttCredentials = MqttConfig{
		Host:     env.GetEnv("MQTT_HOST", "localhost"),
		Port:     env.GetEnvAsInt("MQTT_PORT", 1883),
		User:     env.GetEnv("MQTT_USER", "user"),
		Password: env.GetEnv("MQTT_PASSWORD", "password"),
		Prefix:   env.GetEnv("MQTT_PREFIX", "homeassistant/"),
	}

	var cpuInterval int

	if env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5s")[len(env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5s"))-1] == 's' {
		cpuInterval, _ = strconv.Atoi(env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5s")[:len(env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5s"))-1])
		cpuInterval *= 1000
	} else if env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5000ms")[len(env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5000ms"))-2:] == "ms" {
		cpuInterval, _ = strconv.Atoi(env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5000ms")[:len(env.GetEnv("CPU_TEMPERATURE_INTERVAL", "5000ms"))-2])
	}

	cpuTemperatureConfig = MetricsConfig{
		Enabled:  env.GetEnvAsBool("CPU_TEMPERATURE_ENABLED", false),
		Interval: cpuInterval,
		Topic:    mqttCredentials.Prefix + "temperature/" + env.GetEnv("CPU_ENTITY_NAME", "cpu") + "/state",
	}
}
