package webgl

type Vec3 interface {
	Floats() [3]float32
}

type Mat4 interface {
	Floats() [16]float32
}
