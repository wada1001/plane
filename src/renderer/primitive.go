package renderer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/wada1001/plane/src/components"
)

type (
	PrimitiveCommand struct {
		Position components.Vector
		Shape    components.ShapeType
		Size     components.Vector
		Order    uint
	}

	System struct {
		// Filter generic.Filter3[components.Position, components.Primitive, components.Layer]
	}
)

func (c PrimitiveCommand) Execute(screen *ebiten.Image) error {
	vector.FillCircle(
		screen,
		float32(c.Position.X),
		float32(c.Position.Y),
		float32(c.Size.Length()),
		color.White,
		false,
	)
	return nil
}

func (c PrimitiveCommand) GetOrder() uint {
	return c.Order
}

func Proccess(w *ecs.World) (CommandBuffer, error) {
	filter := generic.NewFilter3[components.Primitive, components.Position, components.Layer]()

	query := filter.Query(w)

	cb := CommandBuffer{}
	for query.Next() {
		primitive, pos, layer := query.Get()
		cb = append(cb, PrimitiveCommand{
			Position: pos.Vector,
			Shape:    primitive.Shape,
			Size:     primitive.Size,
			Order:    layer.ID,
		})
	}

	return cb, nil
}
