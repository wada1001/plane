package entites

import (
	"math/rand/v2"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/wada1001/plane/src/components"
)

func Create(w *ecs.World) error {
	mapper := generic.NewMap4[components.Primitive, components.Position, components.Velocity, components.Layer](w)
	for i := 0; i < 100; i++ {
		e := mapper.New()
		primitive, pos, velocity, layer := mapper.Get(e)
		pos.X = rand.Float64()*300 + 5
		pos.Y = rand.Float64()*400 + 5
		primitive.Shape = components.ShapeTypeCircle
		defaultSize := components.Vector{X: 1, Y: 1}
		primitive.Size = defaultSize.Scale(rand.Float64() * 1.2)
		velocity.Vector = components.Vector{X: 0, Y: 10}
		layer.ID = 1000
	}

	return nil
}
