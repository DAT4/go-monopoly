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
			player.NewPlayer(colornames.Salmon, 1, 2),
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
	board.NewStartField(),
	board.NewOwnableField("Rodovrevej", 1200, 0, colornames.Royalblue),
	board.NewChanceField(),
	board.NewOwnableField("Hvidovrevej", 1200, 0, colornames.Royalblue),
	board.NewTaxField(4000),
	board.NewOwnableField("Helsingoer - Helsingborg", 4000, 1, colornames.Gainsboro),
	board.NewOwnableField("Roskildevej", 2000, 2, colornames.Orange),
	board.NewChanceField(),
	board.NewOwnableField("Valby Langade", 2000, 2, colornames.Orange),
	board.NewOwnableField("Allegade", 2400, 2, colornames.Orange),
	board.NewVisitField(),
	board.NewOwnableField("Frederikberg Alle", 2800, 3, colornames.Green),
	board.NewOwnableField("Tuborg Squash", 3000, 3, colornames.Cyan),
	board.NewOwnableField("Bullowsvej", 2800, 3, colornames.Green),
	board.NewOwnableField("Gl. Kongevej", 3200, 3, colornames.Green),
	board.NewOwnableField("Mols-Linjen", 4000, 1, colornames.Gainsboro),
	board.NewOwnableField("Bernstorffsvej", 3600, 4, colornames.Gray),
	board.NewChanceField(),
	board.NewOwnableField("Hellerupvej", 3600, 4, colornames.Gray),
	board.NewOwnableField("Strandvejen", 4000, 4, colornames.Gray),
	board.NewParkingField(),
	board.NewOwnableField("Trianglen", 4400, 5, colornames.Red),
	board.NewChanceField(),
	board.NewOwnableField("Osterbrogade", 4400, 5, colornames.Red),
	board.NewOwnableField("Gronningen", 4800, 5, colornames.Red),
	board.NewOwnableField("Gedser - Rostok", 4000, 1, colornames.Gainsboro),
	board.NewOwnableField("Bredgade", 5200, 6, colornames.Ghostwhite),
	board.NewOwnableField("Kgs. Nytorv", 5200, 6, colornames.Ghostwhite),
	board.NewOwnableField("Coca cola", 3000, 3, colornames.Cyan),
	board.NewOwnableField("Ostergade", 5600, 6, colornames.Ghostwhite),
	board.NewPrisonField(),
	board.NewOwnableField("Amagertorv", 6000, 7, colornames.Yellow),
	board.NewOwnableField("Vimmelskaftet", 6000, 7, colornames.Yellow),
	board.NewChanceField(),
	board.NewOwnableField("Nygade", 6400, 7, colornames.Yellow),
	board.NewOwnableField("Rodby - Puttgarden", 4000, 1, colornames.Gainsboro),
	board.NewChanceField(),
	board.NewOwnableField("Frederiksberggade", 7000, 8, colornames.Purple),
	board.NewTaxField(2000),
	board.NewOwnableField("Raadhuspladsen", 8000, 8, colornames.Purple),
}
