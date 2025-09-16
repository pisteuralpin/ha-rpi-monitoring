package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func sendViaMQTT(topic string, payload string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + mqttCredentials.Host + ":" + fmt.Sprintf("%d", mqttCredentials.Port))
	// opts.SetUsername(mqttCredentials.User)
	// opts.SetPassword(mqttCredentials.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to MQTT broker:", token.Error())
		panic(token.Error())
	}

	token := client.Publish(topic, 0, false, payload)
	token.Wait()

	client.Disconnect(250)
}
