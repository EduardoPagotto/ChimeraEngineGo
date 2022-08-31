module github.com/EduardoPagotto/ChimeraEngineGo

go 1.19

replace github.com/EduardoPagotto/ChimeraEngineGo/chimera/core => ./chimera/core

require (
	github.com/EduardoPagotto/ChimeraEngineGo/chimera/core v0.0.0-00010101000000-000000000000
	github.com/veandco/go-sdl2 v0.4.25
)

require (
	github.com/EngoEngine/math v1.0.4 // indirect
	github.com/engoengine/glm v0.0.0-20170725114841-9c08f4d1f668 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
)
