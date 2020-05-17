package main

import (
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func beep(led *gpio.LedDriver) {
	led.On()
	time.Sleep(100 * time.Millisecond)
	led.Off()
	time.Sleep(200 * time.Millisecond)
	led.On()
	time.Sleep(100 * time.Millisecond)
	led.Off()
}

func main() {
	e := raspi.NewAdaptor()
	e.Connect()

	led := gpio.NewLedDriver(e, "13")
	led.Start()

	defer led.Halt()
	beep(led)

	// var level byte = 128
	// if err := led.Brightness(level); err != nil {
	// 	panic("no PWM support")
	// } else {
	// 	time.Sleep(1000 * time.Millisecond)
	// 	led.Off()
	// }

	beep(led)

	// for {
	// 	fmt.Println("toggle")
	// 	led.Toggle()
	// 	time.Sleep(1000 * time.Millisecond)
	// }
}
