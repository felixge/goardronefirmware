package goardronefirmware

import (
	"github.com/felixge/goardronefirmware/rotors"
	"github.com/felixge/goardronefirmware/sensors"
	"log"
	"time"
)

const Version = "0.0.1"

type Firmware struct {
	sensors *sensors.Sensors
	rotors  *rotors.Rotors
}

func Start() (*Firmware, error) {
	f := &Firmware{}
	return f, f.Init()
}

func (f *Firmware) Init() error {
	log.Printf("Initializing firmware v%s ...\n", Version)
	f.sensors = sensors.Start()

	r, err := rotors.NewRotors()
	if err != nil {
		return err
	}
	f.rotors = r
	return nil
}

func (f *Firmware) Burnout() (interface{}, error) {
	start := time.Now()
	f.rotors.SetSpeed(0, 511)
	f.rotors.SetSpeed(1, 511)
	f.rotors.SetSpeed(2, 511)
	f.rotors.SetSpeed(3, 511)

	for {
		f.rotors.UpdateMotors()
		time.Sleep(30 * time.Millisecond)

		if time.Since(start) > 10 * time.Second {
			break
		}
	}

	return struct{ Ok bool }{true}, nil
}

func (f *Firmware) AnimateLeds() (interface{}, error) {
	err := rotors.AnimateLeds(f.rotors)
	if err != nil {
		return nil, err
	}

	return struct{ Ok bool }{true}, nil
}

func (f *Firmware) Sensors() (interface{}, error) {
	return struct{ Hello string }{"fuck"}, nil
}
