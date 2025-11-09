package components

type (
	ShapeType uint

	Primitive struct {
		Shape ShapeType
		Size  Vector
	}
)

const (
	ShapeTypeCircle ShapeType = iota
	ShapeTypeRectangle
)
