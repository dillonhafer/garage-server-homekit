package main

import (
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	"github.com/dillonhafer/garage-server/door"
)

type GarageDoorOpener struct {
	Accessory        *accessory.Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

func NewGarageDoorOpener(info accessory.Info) *GarageDoorOpener {
	acc := GarageDoorOpener{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.Accessory.AddService(acc.GarageDoorOpener.Service)
	return &acc
}

func toggleDoorIf(target string, o Options) {
	if status, err := door.CheckDoorStatus(o.statusPin); err != nil {
		if status == target {
			door.ToggleSwitch(o.relayPin, o.sleepTimeout)
		}
	}
}

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
