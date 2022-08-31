package Core

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Client struct {
	isPaused bool
	canvas   *CanvasGL
}

func (c *Client) start() {
	c.canvas.initGL()
}

func (c *Client) stop() {
	c.canvas.Destroy()
}

func (c *Client) render() {

	c.canvas.Before(0)

	c.canvas.After(0)

	c.canvas.swapWindow()
}

func (c *Client) keyboardEvent(k *sdl.KeyboardEvent) {

	switch k.Keysym.Mod {
	case sdl.KMOD_CAPS:
		fmt.Println("Caps Lock")
	}

	switch k.Keysym.Sym { //sdl.Keycode
	case sdl.K_ESCAPE:
		eventQuit := sdl.QuitEvent{
			Type:      sdl.QUIT,
			Timestamp: 0,
		}
		sdl.PushEvent(&eventQuit)
	}
}

func (c *Client) mouseButton(m *sdl.MouseButtonEvent) {

	if (m.Type == sdl.MOUSEBUTTONDOWN) && (m.Button == sdl.BUTTON_LEFT) {
		fmt.Println("botao ok")
	}
}

func (c *Client) mouseMotion(m *sdl.MouseMotionEvent) {
	//xrot := float32(m.Y) / 2
	//yrot := float32(m.X) / 2
	fmt.Printf("[%dms]MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", m.Timestamp, m.Which, m.X, m.Y, m.XRel, m.YRel)

}

func (c *Client) windowEvent(w *sdl.WindowEvent) {}

func (c *Client) userEvent(u *sdl.UserEvent) {
	fmt.Println(" Code", u.Code)
}

func (c *Client) newFPS(fps uint32) {}

func (c *Client) paused() bool {
	return c.isPaused
}

func (c *Client) joystickStatusUpdate(j *JoystickState) {

}

func NewClient(c *CanvasGL) *Client {
	client := new(Client)
	client.isPaused = false
	client.canvas = c
	return client
}
