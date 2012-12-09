package goardronefirmware

import (
	"github.com/felixge/goardronefirmware/motors"
	"log"
)

const Version = "0.0.1"

type Firmware struct {
	mc *motors.Controller
}

func Start() (*Firmware, error) {
	f := &Firmware{}
	return f, f.Init()
}

func (f *Firmware) Init() error {
	log.Printf("Initializing firmware v%s ...\n", Version)

	mc, err := motors.NewController()
	if err != nil {
		return err
	}
	f.mc = mc
	return nil
}
