package main

func initDiscovering() {
	if config.HomeAssistant.Discovery {
		prefix := config.HomeAssistant.Prefix
		if prefix[len(prefix)-1] != '/' {
			prefix += "/"
		}

		device := `{
			"name": "` + config.Device.Name + `",
			"model": "` + config.Device.Model + `",
			"identifiers": ["` + config.Device.UniqueId + `"]
		}`

		// Device Discovery
		topic := prefix + "device/" + config.Device.Id + "/config"
		payload := device
		sendViaMQTT(topic, payload)

		// CPU Temperature Discovery
		if config.CpuTemperature.Enabled {
			topic := prefix + "sensor/" + config.Device.Id + "_cpu_temperature/config"
			payload := `{
				"device_class": "temperature",
				"name": "CPU Temperature",
				"state_topic": "` + config.CpuTemperature.Topic + `",
				"unit_of_measurement": "°C",
				"unique_id": "cpu_temperature",
				"device": ` + device + `
			}`
			sendViaMQTT(topic, payload)
		}

		// Power Supply Discovery
		if config.PowerSupply.Enabled {
			topic := prefix + "sensor/" + config.Device.Id + "_power/config"
			payload := `{
				"device_class": "power",
				"name": "Power Supply",
				"state_topic": "` + config.PowerSupply.Topic + `",
				"unit_of_measurement": "W",
				"unique_id": "power",
				"device": ` + device + `
			}`
			sendViaMQTT(topic, payload)
		}
	}
}
