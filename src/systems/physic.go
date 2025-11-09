package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/wada1001/plane/src/components"
)

func Process(w *ecs.World) error {
	filter := generic.NewFilter2[components.Position, components.Velocity]()

	query := filter.Query(w)
	delta := 1.0 / float64(ebiten.TPS())
	for query.Next() {
		pos, velocity := query.Get()

		dv := velocity.Scale(delta)
		pos.X = pos.X + dv.X
		pos.Y = pos.Y + dv.Y
	}

	return nil
}
