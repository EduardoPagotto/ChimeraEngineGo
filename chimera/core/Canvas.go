package Core

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	//gl "github.com/chsc/gogl/gl33"
	//"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
	//"github.com/engoengine/glm"
)

type ICanvas interface {
	Create(title string, width, heigh int32, fullScreen bool)
	Destroy()
	Before(indexEye uint8)
	After(indexEye uint8)
	ToggleFullScreen()
	Reshape(width, heigh int32)
}

type CanvasBase struct {
	title      string
	width      int32
	heigh      int32
	fullScreen bool
	window     *sdl.Window
}

type CanvasGL struct {
	CanvasBase
	context sdl.GLContext
	// event   sdl.Event
}

func NewCanvasGL(title string, width, heigh int32, fullScreen bool) *CanvasGL {
	c := new(CanvasGL)
	c.Create(title, width, heigh, fullScreen)
	return c
}

func (c *CanvasGL) Create(title string, width, heigh int32, fullScreen bool) {

	c.title = title
	c.width = width
	c.heigh = heigh
	c.fullScreen = fullScreen

	var err error
	runtime.LockOSThread()
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2); err != nil { // FIXME
		panic(err)
	}

	if err = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1); err != nil { // FIXME
		panic(err)
	}

	if err = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1); err != nil {
		panic(err)
	}

	if err = sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24); err != nil {
		panic(err)
	}

	c.window, err = sdl.CreateWindow(c.title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, heigh, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}

	c.context, err = c.window.GLCreateContext()
	if err != nil {
		panic(err)
	}

	if err = sdl.GLSetSwapInterval(1); err != nil {
		panic(err)
	}

	if err = gl.Init(); err != nil {
		panic(err)
	}
}

func (c *CanvasGL) Destroy() {
	c.window.Destroy()
	sdl.GLDeleteContext(c.context)
}

func (c *CanvasGL) swapWindow() { c.window.GLSwap() }

func (c *CanvasGL) Before(indexEye uint8) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func _assertErroCritico() {
	error := gl.GetError()
	if error != gl.NO_ERROR {
		msg := fmt.Sprintf("Erro GL code: %d", error)
		sdl.LogCritical(sdl.LOG_CATEGORY_ERROR, msg)
		panic(msg)
	}
}

func (c *CanvasGL) initGL() {
	//GLenum error = GL_NO_ERROR;

	// Initialize Projection Matrix
	gl.MatrixMode(gl.PROJECTION)

	gl.LoadIdentity()

	_assertErroCritico()

	// Initialize Modelview Matrix
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	//l.Enable(gl.DEPTH_TEST)

	_assertErroCritico()

	// Initialize clear color
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	_assertErroCritico()
}

func (c *CanvasGL) After(indexEye uint8)       {}
func (c *CanvasGL) ToggleFullScreen()          {}
func (c *CanvasGL) Reshape(width, heigh int32) {}
