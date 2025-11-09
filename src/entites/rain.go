package entites

import (
	"math/rand/v2"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/wada1001/plane/src/components"
)

func Create(w *ecs.World) error {
	mapper := generic.NewMap3[components.Primitive, components.Position, components.Layer](w)
	for i := 0; i < 100; i++ {
		e := mapper.New()
		primitive, pos, layer := mapper.Get(e)

		pos.X = rand.Float64()*300 + 5
		pos.Y = rand.Float64()*400 + 5
		primitive.Shape = components.ShapeTypeCircle
		primitive.Size = components.Vector{X: 3, Y: 3}
		layer.ID = 1000
	}

	return nil
}
