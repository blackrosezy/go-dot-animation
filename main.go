package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

const (
	screenWidth  = 480
	screenHeight = 300
	dotRadius    = 10
)

type Game struct {
	x, y        float64
	dx, dy      float64
	angle       float64
	trails      [][2]float64
	trailColors []color.RGBA
}

func (g *Game) Update() error {
	g.x += g.dx
	g.y += g.dy

	// Bounce off walls
	if g.x <= dotRadius || g.x >= screenWidth-dotRadius {
		g.dx = -g.dx
	}
	if g.y <= dotRadius || g.y >= screenHeight-dotRadius {
		g.dy = -g.dy
	}

	// Update angle for color cycling
	g.angle += 0.05
	if g.angle > 2*math.Pi {
		g.angle -= 2 * math.Pi
	}

	// Update trail
	g.trails = append(g.trails, [2]float64{g.x, g.y})
	if len(g.trails) > 50 {
		g.trails = g.trails[1:]
		g.trailColors = g.trailColors[1:]
	}

	// Add new trail color
	r := uint8((math.Sin(g.angle) + 1) * 127)
	b := uint8((math.Cos(g.angle) + 1) * 127)
	g.trailColors = append(g.trailColors, color.RGBA{r, 255 - r/2, b, 255})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw trails
	for i, trail := range g.trails {
		vector.DrawFilledCircle(screen, float32(trail[0]), float32(trail[1]), float32(dotRadius*float64(i)/float64(len(g.trails))), g.trailColors[i], true)
	}

	// Draw main dot
	vector.DrawFilledCircle(screen, float32(g.x), float32(g.y), dotRadius, color.White, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bouncing Dot")

	game := &Game{
		x:  screenWidth / 2,
		y:  screenHeight / 2,
		dx: 3,
		dy: 2,
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
