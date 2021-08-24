package Core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type IEvents interface {
	start()
	stop()
	render()
	keyboardEvent(k *sdl.KeyboardEvent)
	mouseMotion(m *sdl.MouseMotionEvent)
	mouseButton(m *sdl.MouseButtonEvent)
	userEvent(u *sdl.UserEvent)
	joystickStatusUpdate(j *JoystickState)
	newFPS(fps uint32)
	windowEvent(w *sdl.WindowEvent)
	paused() bool
}

type FlowControl struct {
	pGameClientEvents IEvents
	joystickManager   JoystickManager
	timerFPS          Timer
	event             sdl.Event
}

func NewflowControl(pGameClientEvents IEvents) *FlowControl {

	f := new(FlowControl)
	f.pGameClientEvents = pGameClientEvents
	f.timerFPS.elapsedCount = 1000
	f.timerFPS.start()
	f.joystickManager.Initialized = false

	return f
}

func (f *FlowControl) Open() {

	f.joystickManager.Initialize()
	f.joystickManager.joySearchAll()
	f.pGameClientEvents.start()

}
func (f *FlowControl) Close() {

	f.joystickManager.joyReleaseAll()

	eventQuit := sdl.QuitEvent{
		Type:      sdl.QUIT,
		Timestamp: 0,
	}
	sdl.PushEvent(&eventQuit) // FIXME: ler se erro

}

func (f *FlowControl) countFrame() {

	if f.timerFPS.stepCount() {
		fps := f.timerFPS.countStep
		f.pGameClientEvents.newFPS(fps)
	}
}

func (f *FlowControl) processaGame() {

	f.countFrame()
	f.pGameClientEvents.render()
}

func (f *FlowControl) Loop() {

	running := true

	var fpsMin uint32 = 60
	var minimumFrameTime uint32 = 1000 / fpsMin
	var frameTime, deltaTime, lastFrameTime, timeElapsed, tot_delay uint32

	for running {

		frameTime = sdl.GetTicks()

		for f.event = sdl.PollEvent(); f.event != nil; f.event =
			sdl.PollEvent() {
			switch t := f.event.(type) {
			case *sdl.KeyboardEvent:
				f.pGameClientEvents.keyboardEvent(t)
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				f.pGameClientEvents.mouseMotion(t)
			case *sdl.MouseButtonEvent:
				f.pGameClientEvents.mouseButton(t)
			case *sdl.UserEvent:
				f.pGameClientEvents.userEvent(t)
			case *sdl.WindowEvent:
				f.pGameClientEvents.windowEvent(t)
			default:
				if joy := f.joystickManager.TrackEvent(&f.event); joy != nil {
					f.pGameClientEvents.joystickStatusUpdate(joy)
				}
			}
		}

		if !f.pGameClientEvents.paused() {
			f.processaGame()
		}

		// inicio contadores
		deltaTime = frameTime - lastFrameTime
		lastFrameTime = frameTime

		timeElapsed = (sdl.GetTicks() - frameTime)
		if timeElapsed < minimumFrameTime {
			tot_delay = minimumFrameTime - timeElapsed
			sdl.LogDebug(sdl.LOG_CATEGORY_RENDER, "DeltaTime: %d  Delay: %d", deltaTime, tot_delay)
			sdl.Delay(tot_delay)
		} else {
			sdl.LogDebug(sdl.LOG_CATEGORY_RENDER, "DeltaTime: %d TimeElapsed: %d", deltaTime, timeElapsed)
		}
	}

	f.pGameClientEvents.stop()
}
