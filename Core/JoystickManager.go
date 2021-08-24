package Core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type JoystickManager struct {
	Joysticks   map[uint8]*JoystickState
	Initialized bool
}

func (j *JoystickManager) Initialize() {
	if !j.Initialized {
		sdl.InitSubSystem(sdl.INIT_JOYSTICK)
		j.Joysticks = make(map[uint8]*JoystickState)
	}

	sdl.JoystickEventState(sdl.ENABLE)
	j.Initialized = true
}

func (j *JoystickManager) joySearchAll() {

	if !j.Initialized {
		return
	}

	tot := sdl.NumJoysticks()
	if tot > 0 {
		for i := 0; i < tot; i++ {
			handle := sdl.JoystickOpen(i)
			pJoy, ok := j.Joysticks[uint8(i)]
			if handle != nil {
				if !ok {
					j.Joysticks[uint8(i)] = NewJoystickState(i, handle)
				} else {
					if pJoy.handle == nil {
						ReloadJoystickState(pJoy, i, handle)
					}
				}
			}
		}
	}
}

func (j *JoystickManager) joyReleaseAll() {

	if !j.Initialized {
		return
	}

	for _, joy := range j.Joysticks {
		ReleaseJoystickState(joy)
		delete(j.Joysticks, joy.id)
	}

}

func (j JoystickManager) TrackEvent(e *sdl.Event) *JoystickState {

	var sel *JoystickState = nil

	switch t := (*e).(type) {
	case *sdl.JoyButtonEvent:
		sel = j.Joysticks[uint8(t.Which)]
		sel.ButtonEvent(t)

	case *sdl.JoyAxisEvent:
		sel = j.Joysticks[uint8(t.Which)]
		sel.AxisEvent(t)

	case *sdl.JoyHatEvent:
		sel = j.Joysticks[uint8(t.Which)]
		sel.HatEvent(t)

	case *sdl.JoyBallEvent:
		sel = j.Joysticks[uint8(t.Which)]
		sel.BallEvent(t)

	case *sdl.JoyDeviceAddedEvent:
		sel = j.Joysticks[uint8(t.Which)]
		sel.DeviceAddEvent(t)

	case *sdl.JoyDeviceRemovedEvent:
		sel = j.Joysticks[uint8(t.Which)]
		sel.DeviceRemovedEvent(t)
	}

	return sel
}

// func (j JoystickManager) getJoystickState(id uint8) *JoystickState {

// 	pJoy, ok := j.Joysticks[id]
// 	if ok {
// 		return pJoy
// 	}

// 	return nil
// }
