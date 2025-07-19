//go:build js || wasm

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("ブロック崩し")
	ebiten.RunGame(NewGame())
}
