package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/log"
	"github.com/dillonhafer/garage-server/door"
)

func main() {
	log.Verbose = true
	log.Info = true

	info := accessory.Info{
		Name:         "Garage Door",
		Manufacturer: "Dillon Hafer",
		Model:        "Raspberry Pi",
	}

	acc := NewGarageDoorOpener(info)
	acc.GarageDoorOpener.CurrentDoorState.OnValueRemoteUpdate(func(int) {
		status, _ := door.CheckDoorStatus(25)
		switch status {
		case "open":
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpen)
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)
		case "closed":
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)
		}
	})

	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(targetState int) {
		switch targetState {
		case characteristic.TargetDoorStateClosed:
			closeGarage()
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)
		case characteristic.TargetDoorStateOpen:
			openGarage()
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpen)
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
