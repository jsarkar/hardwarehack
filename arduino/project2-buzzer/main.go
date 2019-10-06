package main

import (
	"machine"
)

const (
	buzzer = machine.Pin(9)
	button = machine.Pin(7)
)

func main() {
	buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if button.Get() == true {
			buzzer.High()
		} else {
			buzzer.Low()
		}
	}
}
