package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/log"
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
	}

	acc := NewGarageDoorOpener(info)
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(targetState int) {
		switch targetState {
		case characteristic.TargetDoorStateClosed:
			toggleDoorIf("open", options)
		case characteristic.TargetDoorStateOpen:
			toggleDoorIf("closed", options)
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	go pollDoorStatus(acc, options.statusPin)
	hc.OnTermination(t.Stop)
	t.Start()
}
