package Core

import "github.com/veandco/go-sdl2/sdl"

type Timer struct {
	started      bool
	paused       bool
	startTicks   uint32
	lastTicks    uint32
	pausedTicks  uint32
	step         uint32
	countStep    uint32
	elapsedCount uint32
}

func (t *Timer) start() {

	t.started = true
	t.paused = false
	t.startTicks = sdl.GetTicks()
	t.lastTicks = t.startTicks

}

func (t *Timer) stop() {
	t.started = false
	t.paused = false
}

func (t *Timer) pause() {

	if t.started && !t.paused {
		t.paused = true
		t.pausedTicks = sdl.GetTicks() - t.startTicks
	}

}

func (t *Timer) resume() {
	if t.paused {
		t.paused = false
		t.startTicks = sdl.GetTicks() - t.pausedTicks
		t.lastTicks = t.startTicks
		t.pausedTicks = 0
	}

}

func (t Timer) restart() uint32 {
	elapsedTicks := t.ticks()
	t.start()
	return elapsedTicks
}

func (t Timer) ticks() uint32 {

	if t.started {
		if !t.paused {
			return sdl.GetTicks() - t.startTicks
		} else {
			return t.pausedTicks
		}
	}
	return 0

}

func (t *Timer) stepCount() bool {

	temp := t.ticks()
	if temp < t.elapsedCount {
		t.step++
	} else {
		t.countStep = t.step
		t.step = 0
		t.start()
		return true
	}

	return false

}
