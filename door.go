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

func toggleDoorIf(target string, o Options) {
	println(o.sleepTimeout, o.statusPin, o.relayPin)
	if status, err := door.CheckDoorStatus(o.statusPin); err != nil {
		if status == target {
			door.ToggleSwitch(o.relayPin, o.sleepTimeout)
		}
	}
}

func openGarage(o Options) {
	toggleDoorIf("closed", o)
}

func closeGarage(o Options) {
	toggleDoorIf("open", o)
}
