package Core

import "github.com/engoengine/glm"

type Transform struct {
	model glm.Mat4
}

func NewTransform() *Transform {
	t := new(Transform)
	t.model = glm.Mat4{1.0}
	return t
}

func (t *Transform) set(m glm.Mat4) {
	t.model = m
}

func (t *Transform) getPosition() glm.Vec3 {
	return glm.Vec3{t.model[3]}
}

// func (t *Transform) setPosition(pos glm.Vec3) {
// 	t.model = glm.Translate3D()//glm.Translate3D(pos, t.model)
// }
