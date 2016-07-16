package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/log"
	"github.com/dillonhafer/garage-server/door"
)

const Version = "0.0.1"

type Options struct {
	pin          string
	relayPin     int
	statusPin    int
	sleepTimeout int
	version      bool
}

var options Options
var serialNumber = os.Getenv("SERIAL_NUMBER")

func toggleDoor(o Options) func(int) {
	return func(targetState int) {
		nextState := "closed"
		if targetState == characteristic.TargetDoorStateOpen {
			nextState = "open"
		}

		if currentDoorState, err := door.CheckDoorStatus(o.statusPin); err != nil {
			fmt.Printf("ERROR: Could not read status pin %v\n", err)
		} else {
			if currentDoorState != nextState {
				door.ToggleSwitch(o.relayPin, o.sleepTimeout)
			}
		}
	}
}

func pollDoorStatus(acc *GarageDoorOpener, pin int) {
	for {
		if status, err := door.CheckDoorStatus(pin); err != nil {
			fmt.Printf("ERROR: Could not read status pin %v\n", err)
		} else {
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&options.pin, "pin", "", "8-digit Pin for securing garage door")
	flag.IntVar(&options.relayPin, "relay-pin", 25, "GPIO pin of relay")
	flag.IntVar(&options.statusPin, "status-pin", 10, "GPIO pin of reed switch")
	flag.IntVar(&options.sleepTimeout, "sleep", 500, "Time in milliseconds to keep switch closed")
	flag.BoolVar(&options.version, "version", false, "print version and exit")
	flag.Parse()

	if options.version {
		fmt.Printf("garage-server-homekit v%v\n", Version)
		os.Exit(0)
	}

	if serialNumber == "" {
		println("You did not set SERIAL_NUMBER env var")
		os.Exit(1)
	}

	if options.pin == "" || len(options.pin) != 8 {
		println("Pin must be and 8 digit number")
		os.Exit(0)
	}

	log.Verbose = true
	log.Info = true

	info := accessory.Info{
		Name:         "Garage Door",
		Manufacturer: "Dillon Hafer",
		Model:        "Raspberry Pi",
		SerialNumber: serialNumber,
	}

	acc := NewGarageDoorOpener(info)
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(toggleDoor(options))

	t, err := hc.NewIPTransport(hc.Config{Pin: options.pin}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	go pollDoorStatus(acc, options.statusPin)
	hc.OnTermination(t.Stop)
	t.Start()
}
