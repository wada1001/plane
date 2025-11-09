package main

import (
	"log"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/wada1001/plane/src/entites"
	"github.com/wada1001/plane/src/renderer"
	"github.com/wada1001/plane/src/systems"
)

type System func(w *ecs.World) error
type Renderer func(w *ecs.World) (renderer.CommandBuffer, error)

type Game struct {
	World    ecs.World
	Systems  []System
	Renderer []Renderer
}

func (g *Game) Update() error {
	for _, r := range g.Systems {
		if err := r(&g.World); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cb := renderer.CommandBuffer{}
	for _, r := range g.Renderer {
		t, err := r(&g.World)
		if err != nil {
			panic(err)
		}

		cb = append(cb, t...)
	}

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", len(cb)))
	sort.Slice(cb, func(i, j int) bool {

		return cb[i].GetOrder() < cb[j].GetOrder()
	})

	for _, c := range cb {

		c.Execute(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	game := &Game{
		World: ecs.NewWorld(),
		Systems: []System{
			systems.Process,
			systems.ProcessRain,
		},
		Renderer: []Renderer{
			renderer.Proccess,
		},
	}

	entites.Create(&game.World)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
