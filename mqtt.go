package main

import (
	"fmt"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func sendViaMQTT(topic string, payload string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + config.MQTTConfig.Host + ":" + fmt.Sprintf("%d", config.MQTTConfig.Port))
	opts.SetUsername(config.MQTTConfig.User)
	opts.SetPassword(config.MQTTConfig.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		slog.Error("Error connecting to MQTT broker:", "error", token.Error())
		return
	}
	token := client.Publish(topic, 0, false, payload)
	if token.Error() != nil {
		slog.Error("Error publishing to MQTT broker:", "error", token.Error())
		return
	}
	token.Wait()

	slog.Debug("Data sent on " + topic)

	client.Disconnect(250)
}
