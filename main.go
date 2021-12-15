package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

func main() {

	game := &Game{
		players: []*Player{
			NewPlayer(colornames.Red, 1, 2),
			NewPlayer(colornames.Blue, 2, 2),
			NewPlayer(colornames.Yellow, 3, 2),
			NewPlayer(colornames.Snow, 4, 2),
		},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Monopoly")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
