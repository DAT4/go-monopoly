package game

import (
	"monopoly/board"
	"monopoly/constants"
	"monopoly/player"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct {
	currentPlayer int
	Players       []*player.Player
	fields        []*board.Field
	Tiles         [][]*board.Tile
}

func (g *Game) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		if g.Players[0].X > 0 {
			g.Players[0].X = g.Players[0].X - 1
		}
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		if g.Players[0].Y < constants.ScreenHeight/constants.TileSize-1 {
			g.Players[0].Y = g.Players[0].Y + 1
		}
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		if g.Players[0].Y > 0 {
			g.Players[0].Y = g.Players[0].Y - 1
		}
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		if g.Players[0].X < constants.ScreenWidth/constants.TileSize-1 {
			g.Players[0].X = g.Players[0].X + 1
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Green)
	g.drawBoard(screen)
	g.drawPlayers(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideWidth
}

func (g *Game) drawPlayers(screen *ebiten.Image) {
	for _, e := range g.Players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.X*constants.TileSize), float64(e.Y*constants.TileSize))
		screen.DrawImage(e.Image, op)
	}
}

func (g *Game) drawBoard(screen *ebiten.Image) {
	for _, es := range g.Tiles {
		for _, e := range es {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(e.X*constants.TileSize), float64(e.Y*constants.TileSize))
			screen.DrawImage(e.Image, op)
		}
	}
}
