package main

import (
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/log"
	"github.com/dillonhafer/garage-server/door"
)

func pollDoorStatus(acc *GarageDoorOpener, pin int) {
	for {
		if status, err := door.CheckDoorStatus(pin); err != nil {
			switch status {
			case "open":
				acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpen)
			case "closed":
				acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)
			}
		}

		time.Sleep(time.Second)
	}
}

func main() {
	log.Verbose = true
	log.Info = true

	info := accessory.Info{
		Name:         "Garage Door",
		Manufacturer: "Dillon Hafer",
		Model:        "Raspberry Pi",
	}

	acc := NewGarageDoorOpener(info)
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(targetState int) {
		switch targetState {
		case characteristic.TargetDoorStateClosed:
			closeGarage()
		case characteristic.TargetDoorStateOpen:
			openGarage()
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	go pollDoorStatus(acc, 10)
	hc.OnTermination(t.Stop)
	t.Start()
}
