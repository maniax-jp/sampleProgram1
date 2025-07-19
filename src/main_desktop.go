//go:build !js && !wasm

package main

import (
   "github.com/hajimehoshi/ebiten/v2"
   "sampleProgram1/src/common"
)

func main() {
   game := common.NewGame()
   ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
   ebiten.SetWindowTitle("Block Breaker")
   if err := ebiten.RunGame(game); err != nil {
	   panic(err)
   }
}
