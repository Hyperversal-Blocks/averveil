package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type broker struct {
	client *mqtt.Client
}

type Broker interface {
}

func New() error {
	c := mqtt.NewClientOptions()
	client := mqtt.NewClient(c)
	token := client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		return fmt.Errorf("error bootstrapping MQTT broker: %w", token.Error())
	}
	return nil
}
