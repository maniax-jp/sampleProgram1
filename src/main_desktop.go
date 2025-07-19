package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Block Breaker")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
