package main

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
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
