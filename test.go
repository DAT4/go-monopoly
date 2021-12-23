package main

import (
	"log"
	"monopoly/board"
	"monopoly/constants"
	"monopoly/game"
	"monopoly/player"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

func main() {
	g := &game.Game{
		Players: []*player.Player{
			player.NewPlayer(colornames.Red, 1, 2),
		},
		Tiles: board.MakeMatrix(fields, 2),
	}

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Monopoly")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

var fields = []board.Field{
	&board.StartField{},
	&board.OwnableField{},
	&board.ChanceField{},
	&board.OwnableField{},
	&board.TaxField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.ChanceField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.VisitField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.ChanceField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.ParkingField{},
	&board.OwnableField{},
	&board.ChanceField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.PrisonField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.ChanceField{},
	&board.OwnableField{},
	&board.OwnableField{},
	&board.ChanceField{},
	&board.OwnableField{},
	&board.TaxField{},
	&board.OwnableField{},
}
