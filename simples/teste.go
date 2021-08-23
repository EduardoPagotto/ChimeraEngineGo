package main

import (
	"github.com/EduardoPagotto/ChimeraEngineGo/Core"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	//sdl.LogSetAllPriority(sdl.LOG_PRIORITY_DEBUG)
	sdl.Log("app iniciado")

	video := Core.NewCanvasGL("teste", 640, 580, false)
	control := Core.NewflowControl(Core.NewClient(video))
	control.Open()
	control.Loop()

	sdl.Log("app finalizado")
	//video.Destroy()

}
