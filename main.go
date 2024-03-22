package main

import (
	"machine"
	"time"

	"tinygoleds/twang"

	"tinygo.org/x/drivers/ws2812"
)

const leds = 60

func main() {

	machine.InitADC()

	machine.D8.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led := ws2812.New(machine.D8)

	xPin := machine.ADC{Pin: machine.ADC0}
	xPin.Configure(machine.ADCConfig{})

	machine.D0.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})

	var middleX uint16 = 65534 / 2
	var threshold uint16 = 2000
	game := twang.NewGame(leds)

	for {

		attack := !machine.D0.Get()
		game.Player.Attack(attack)

		if !attack {
			x := xPin.Get()
			if x < middleX-threshold {
				game.Player.Move(-1)
			} else if x > middleX+threshold {
				game.Player.Move(1)
			}
		}

		game.Update()
		game.Render(led)

		time.Sleep(50 * time.Millisecond)

	}
}
