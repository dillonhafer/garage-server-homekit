package main

import (
	"github.com/brutella/hc/accessory"
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

func toggleDoorIf(target string) {
	if status, err := door.CheckDoorStatus(15); err != nil {
		if status == target {
			door.ToggleSwitch(15, 3)
		}
	}
}

func openGarage() {
	println("OPEN")
	toggleDoorIf("closed")
}

func closeGarage() {
	println("CLOSE")
	toggleDoorIf("open")
}
