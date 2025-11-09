package systems

import (
	"math/rand/v2"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/wada1001/plane/src/components"
)

func ProcessRain(w *ecs.World) error {
	filter := generic.NewFilter2[components.Position, components.Velocity]()

	query := filter.Query(w)
	for query.Next() {
		pos, velocity := query.Get()
		if pos.Y < 250 {
			continue
		}

		pos.Y -= 250
		// velocity.X = rand.Float64() * 3
		velocity.Y = rand.Float64() * 10
	}

	return nil
}
