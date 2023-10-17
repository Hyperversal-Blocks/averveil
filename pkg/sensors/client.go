package sensors

import (
	"averveil/pkg/store"
)

type client struct {
	store store.Store
}

type Client interface {
}

func InitClient() Client {
	return &client{}
}

// TODO: ref: https://github.com/arduino/iot-client-go/blob/master/example/main.go
//	https://github.com/arduino/iot-client-go
//	https://www.arduino.cc/reference/en/iot/api/
