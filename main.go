package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Println("start")

	rotors, err := NewRotors("/dev/ttyO0")
	if err != nil {
		panic(err)
	}

	start := time.Now()

	for i := 0; i < 4; i++ {
		rotors.SetLed(i, LedGreen)
	}

	for {
		err = rotors.UpdateLeds()
		time.Sleep(100 * time.Millisecond)

		if time.Since(start) > 2*time.Second {
			break
		}
	}
	log.Println("quit")
}

// TODO: Can I make rotors/leds arrays?
type Rotors struct {
	file   *os.File
	rotors []int
	leds []LedColor
}

func NewRotors(path string) (*Rotors, error) {
	rotors := &Rotors{
		rotors: make([]int, 4),
		leds: make([]LedColor, 4),
	}
	err := rotors.Open(path)
	if err != nil {
		return nil, err
	}
	return rotors, nil
}

func (m *Rotors) Open(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	m.file = file
	return nil
}

// TODO: make value a float between 0...1 instead of a 0...512 PWM value
func (m *Rotors) SetSpeed(rotorId int, value int) {
	m.rotors[rotorId] = value
}

func (m *Rotors) SetLed(rotorId int, color LedColor) {
	m.leds[rotorId] = color
}

//7: 0
//6: 1
//5: 1
//4: 0 <- LED4
//3: 0 <- LED3
//2: 0 <- LED2
//1: 0 <- LED1
//0: 0 <- ?

type LedColor int

const (
	LedOff LedColor = iota
	LedRed = 1
	LedGreen = 2
	LedOrange = 3
)


// cmd = 011rrrrx xxxggggx (used to be 011grgrg rgrxxxxx in AR Drone 1.0)
// see: https://github.com/ardrone/ardrone/blob/master/ardrone/motorboard/motorboard.c#L243
func (m *Rotors) ledCmd() []byte {
	cmd := make([]byte, 2)
	cmd[0] = 0x60

	for i, color := range m.leds {
		if color == LedRed || color == LedOrange {
			cmd[0] = cmd[0] | (1 << (byte(i) + 1))
		}

		if color == LedGreen || color == LedOrange {
			cmd[1] = cmd[1] | (1 << (byte(i) + 1))
		}
	}
	return cmd
}

// see: https://github.com/ardrone/ardrone/blob/master/ardrone/rotorboard/rotorboard.c
func (m *Rotors) pwmCmd() []byte {
	cmd := make([]byte, 5)
	cmd[0] = byte(0x20 | ((m.rotors[0] & 0x1ff) >> 4))
	cmd[1] = byte(((m.rotors[0] & 0x1ff) << 4) | ((m.rotors[1] & 0x1ff) >> 5))
	cmd[2] = byte(((m.rotors[1] & 0x1ff) << 3) | ((m.rotors[2] & 0x1ff) >> 6))
	cmd[3] = byte(((m.rotors[2] & 0x1ff) << 2) | ((m.rotors[3] & 0x1ff) >> 7))
	cmd[4] = byte(((m.rotors[3] & 0x1ff) << 1))
	return cmd
}

func (m *Rotors) UpdateMotors() error {
	_, err := m.file.Write(m.pwmCmd())
	return err
}

func (m *Rotors) UpdateLeds() error {
	_, err := m.file.Write(m.ledCmd())
	return err
}
