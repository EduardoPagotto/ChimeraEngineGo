package Core

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type JoystickState struct {
	Axes         map[uint8]int16
	ButtonsDown  map[uint8]bool
	ButtonsState map[uint8]uint8
	Hats         map[uint8]uint8
	BallsX       map[uint8]int16
	BallsY       map[uint8]int16

	id     uint8
	name   string
	handle *sdl.Joystick
}

func (j *JoystickState) ButtonEvent(b *sdl.JoyButtonEvent) {

	switch b.Type {
	case sdl.JOYBUTTONDOWN:
		j.ButtonsDown[b.Button] = true
	case sdl.JOYBUTTONUP:
		j.ButtonsDown[b.Button] = false
	}

	j.ButtonsState[b.Button] = b.State
}

func (j *JoystickState) AxisEvent(a *sdl.JoyAxisEvent) {
	j.Axes[a.Axis] = a.Value
}

func (j *JoystickState) HatEvent(h *sdl.JoyHatEvent) {
	j.Hats[h.Hat] = h.Value
}

func (j *JoystickState) BallEvent(b *sdl.JoyBallEvent) {
	j.BallsX[b.Ball] += b.XRel
	j.BallsY[b.Ball] += b.YRel
}

func (j *JoystickState) DeviceAddEvent(d *sdl.JoyDeviceAddedEvent) {
	// TODO
}

func (j *JoystickState) DeviceRemovedEvent(d *sdl.JoyDeviceRemovedEvent) {
	// TODO
}

func (j *JoystickState) AxisZone(indice uint8, deadzone, deadzone_at_ends int16) int16 {

	axi, ok := j.Axes[indice]
	if ok {

		if _int16abs(axi) < deadzone {
			return 0
		} else {
			if axi > 0 {
				var teste int32 = int32(axi + deadzone_at_ends)
				if teste > 32767 {
					return 32767
				}
			} else {
				var teste int32 = int32(axi - deadzone_at_ends)
				if teste < -32768 {
					return -32768
				}
			}
		}

		return axi
	}
	return 0
}

func NewJoystickState(indice int, handle *sdl.Joystick) *JoystickState {

	pJoy := new(JoystickState)
	pJoy.id = uint8(indice)
	pJoy.handle = handle
	pJoy.name = sdl.JoystickNameForIndex(indice)
	if len(pJoy.name) == 0 {
		pJoy.name = "JoystickDefault_" + strconv.Itoa(indice)
	}

	pJoy.Axes = make(map[uint8]int16)
	pJoy.ButtonsDown = make(map[uint8]bool)
	pJoy.ButtonsState = make(map[uint8]uint8)
	pJoy.Hats = make(map[uint8]uint8)
	pJoy.BallsX = make(map[uint8]int16)
	pJoy.BallsY = make(map[uint8]int16)

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

func _int16abs(val int16) int16 {
	if val < 0 {
		return -val
	}
	return val
}
