//go:build js || wasm

package main

import (
   "github.com/hajimehoshi/ebiten/v2"
   "sampleProgram1/src/common"
)

func main() {
   ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
   ebiten.SetWindowTitle("ブロック崩し")
   ebiten.RunGame(common.NewGame())
}
