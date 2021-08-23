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

func (j JoystickManager) TrackEvent(e *sdl.Event) bool {

	ret := false
	// var id uint8
	// ev := *e

	// switch t := ev.(type) {
	// case *sdl.JoyAxisEvent:
	// 	id = uint8(t.Which)
	// case *sdl.JoyButtonEvent:
	// 	id = uint8(t.Which)
	// case *sdl.JoyBallEvent:
	// 	id = uint8(t.Which)
	// case *sdl.JoyDeviceAddedEvent:
	// 	id = uint8(t.Which)
	// case *sdl.JoyDeviceRemovedEvent:
	// 	id = uint8(t.Which)
	// case *sdl.JoyHatEvent:
	// 	id = uint8(t.Which)
	// default:
	// }

	// if ret {
	// 	j.Joysticks[id].id = id
	// 	j.Joysticks[id].TrackEvent()
	// }

	return ret
}

func (j JoystickManager) getJoystickState(id uint8) *JoystickState {

	pJoy, ok := j.Joysticks[id]
	if ok {
		return pJoy
	}

	return nil
}

// func (j *JoystickManager) GetStatusManager() {
// 	for _, joy := range j.Joysticks {
// 		joy.handle.GetStatusJoy()
// 	}
// }

//  void TrackEvent(SDL_Event* event);
//  inline static double AxisScale(const Sint16& value) {
// 	 return value >= 0 ? ((double)value) / 32767.0f : ((double)value) / 32768.0f;
//  }
//  inline double AxisScaled(const Uint8& axis, const double& low, const double& high, const double& deadzone = 0.0f,
// 						  const double& deadzone_at_ends = 0.0f) {
// 	 return low + (high - low) * (Axis(axis, deadzone, deadzone_at_ends) + 1.0f) / 2.0f;
//  }
//  double Axis(const Uint8& axis, const double& deadzone = 0.0f, const double& deadzone_at_ends = 0.0f);
//  bool ButtonDown(const Uint8& button);
//  Uint8 Hat(const Uint8& hat);
//  inline bool HatDir(const Uint8& hat, const Uint8& dir) { return Hat(hat) & dir; }
//  void GetStatusJoy(void);
