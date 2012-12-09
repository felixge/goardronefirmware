package motors

import (
	"math/rand"
	"os"
	"time"
)

func AnimateLeds(speeds *Controller) error {
	for i := 0; i < 4; i++ {
		speeds.SetLed(i, LedOff)
	}
	err := speeds.UpdateLeds()
	if err != nil {
		return err
	}

	for loop := 0; loop < 50; loop++ {
		for i := 0; i < 4; i++ {
			speeds.SetLed(i, LedColor(rand.Intn(LedOrange + 1)))
			err = speeds.UpdateLeds()
			time.Sleep(25 * time.Millisecond)

			if err != nil {
				return err
			}
		}
	}

	for i := 0; i < 4; i++ {
		speeds.SetLed(i, LedGreen)
	}
	return speeds.UpdateLeds()
}

type Controller struct {
	file   *os.File
	speeds []int
	leds   []LedColor
}

func NewController() (*Controller, error) {
	speeds := &Controller{
		speeds: make([]int, 4),
		leds:   make([]LedColor, 4),
	}
	err := speeds.Open("/dev/ttyO0")
	if err != nil {
		return nil, err
	}
	return speeds, nil
}

func (c *Controller) Open(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	c.file = file
	return nil
}

// TODO: make value a float between 0...1 instead of a 0...512 PWM value
func (c *Controller) SetSpeed(motorId int, value int) {
	c.speeds[motorId] = value
}

func (c *Controller) SetLed(motorId int, color LedColor) {
	c.leds[motorId] = color
}

type LedColor int

const (
	LedOff    LedColor = iota
	LedRed             = 1
	LedGreen           = 2
	LedOrange          = 3
)

// cmd = 011rrrrx xxxggggx (used to be 011grgrg rgrxxxxx in AR Drone 1.0)
// see: https://github.com/ardrone/ardrone/blob/master/ardrone/motorboard/motorboard.c#L243
func (m *Controller) ledCmd() []byte {
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

// see: https://github.com/ardrone/ardrone/blob/master/ardrone/motorboard/motorboard.c
func (m *Controller) pwmCmd() []byte {
	cmd := make([]byte, 5)
	cmd[0] = byte(0x20 | ((m.speeds[0] & 0x1ff) >> 4))
	cmd[1] = byte(((m.speeds[0] & 0x1ff) << 4) | ((m.speeds[1] & 0x1ff) >> 5))
	cmd[2] = byte(((m.speeds[1] & 0x1ff) << 3) | ((m.speeds[2] & 0x1ff) >> 6))
	cmd[3] = byte(((m.speeds[2] & 0x1ff) << 2) | ((m.speeds[3] & 0x1ff) >> 7))
	cmd[4] = byte(((m.speeds[3] & 0x1ff) << 1))
	return cmd
}

func (m *Controller) UpdateController() error {
	_, err := m.file.Write(m.pwmCmd())
	return err
}

func (m *Controller) UpdateLeds() error {
	_, err := m.file.Write(m.ledCmd())
	return err
}
