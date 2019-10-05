package main

import (
	"machine"
)

const (
	led    = machine.Pin(8)
	button = machine.Pin(7)
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if button.Get() == false {
			led.High()
		} else {
			led.Low()
		}

	}
}
