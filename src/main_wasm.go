//go:build js || wasm

package main

import (
	"sampleProgram1/src/common"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("ブロック崩し")
	ebiten.RunGame(common.NewGame())
}
