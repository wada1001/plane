package renderer

import "github.com/hajimehoshi/ebiten/v2"

type (
	Command interface {
		GetOrder() uint
		Execute(screen *ebiten.Image) error
	}

	CommandBuffer []Command
)
