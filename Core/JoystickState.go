package Core

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type JoystickState struct {
	Axes        map[uint8]float64
	ButtonsDown map[uint8]bool
	Hats        map[uint8]uint8
	BallsX      map[uint8]int
	BallsY      map[uint8]int

	id     uint8
	name   string
	handle *sdl.Joystick
}

func NewJoystickState(indice int, handle *sdl.Joystick) *JoystickState {

	pJoy := new(JoystickState)
	pJoy.id = uint8(indice)
	pJoy.handle = handle
	pJoy.name = sdl.JoystickNameForIndex(indice)
	if len(pJoy.name) == 0 {
		pJoy.name = "JoystickDefault_" + strconv.Itoa(indice)
	}

	return pJoy
}

func ReloadJoystickState(pJoy *JoystickState, indice int, handle *sdl.Joystick) {
	pJoy.id = uint8(indice)
	pJoy.handle = handle
	pJoy.name = sdl.JoystickNameForIndex(indice)
	if len(pJoy.name) == 0 {
		pJoy.name = "JoystickDefault_" + strconv.Itoa(indice)
	}
}

func ReleaseJoystickState(pJoy *JoystickState) {
	if pJoy.handle != nil {
		pJoy.id = 0
		pJoy.handle = nil
		pJoy.name = "Disconnected"
		pJoy.handle.Close()

	}
}
