package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct {
	currentPlayer int
	players       []*Player
	fields        []*Field
}

func (g *Game) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		if g.players[0].x > 0 {
			g.players[0].x = g.players[0].x - 1
		}
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		if g.players[0].y < screenHeight/tileSize-1 {
			g.players[0].y = g.players[0].y + 1
		}
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		if g.players[0].y > 0 {
			g.players[0].y = g.players[0].y - 1
		}
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		if g.players[0].x < screenWidth/tileSize-1 {
			g.players[0].x = g.players[0].x + 1
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Green)
	g.drawPlayers(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideWidth
}

func (g *Game) drawPlayers(screen *ebiten.Image) {
	for _, e := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.x*tileSize), float64(e.y*tileSize))
		screen.DrawImage(e.Image, op)
	}
}
