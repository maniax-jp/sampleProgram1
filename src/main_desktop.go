//go:build !js && !wasm

package main

import (
	"sampleProgram1/src/common"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := common.NewGame()
	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Block Breaker")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
