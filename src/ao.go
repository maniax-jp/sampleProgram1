package main

import (
	"image/color"
	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 400
	screenHeight = 600
	paddleWidth  = 80
	paddleHeight = 10
	ballSize     = 8
	blockRows    = 5
	blockCols    = 8
	blockWidth   = 40
	blockHeight  = 20
)

type Block struct {
	x, y   float64
	active bool
}

type Game struct {
	paddleX        float64
	ballX, ballY   float64
	ballVX, ballVY float64
	blocks         []Block
	// timer          int // 経過フレーム数
}

func NewGame() *Game {
	blocks := make([]Block, blockRows*blockCols)
	for r := 0; r < blockRows; r++ {
		for c := 0; c < blockCols; c++ {
			blocks[r*blockCols+c] = Block{
				x:      float64(c * blockWidth),
				y:      float64(r * blockHeight),
				active: true,
			}
		}
	}
	return &Game{
		paddleX: screenWidth/2 - paddleWidth/2,
		ballX:   screenWidth - ballSize,
		ballY:   0,
		ballVX:  2,
		ballVY:  -2,
		blocks:  blocks,
	}
}

func (g *Game) Update() error {
	// g.timer++ // 毎フレーム加算
	// パドル操作
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.paddleX -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.paddleX += 4
	}
	if g.paddleX < 0 {
		g.paddleX = 0
	}
	if g.paddleX > screenWidth-paddleWidth {
		g.paddleX = screenWidth - paddleWidth
	}

	// 重力加速度
	gravity := 0.5 / 3
	g.ballVY += gravity
	// ボール移動（減衰なし）
	g.ballX += g.ballVX
	g.ballY += g.ballVY

	// 小数点誤差による減衰防止
	minSpeed := 0.01
	if g.ballVX > 0 && g.ballVX < minSpeed {
		g.ballVX = minSpeed
	} else if g.ballVX < 0 && g.ballVX > -minSpeed {
		g.ballVX = -minSpeed
	}
	if g.ballVY > 0 && g.ballVY < minSpeed {
		g.ballVY = minSpeed
	} else if g.ballVY < 0 && g.ballVY > -minSpeed {
		g.ballVY = -minSpeed
	}

	// 壁反射
	if g.ballX < 0 || g.ballX > screenWidth-ballSize {
		g.ballVX *= -1
	}
	if g.ballY < 0 {
		g.ballY = 0
		g.ballVY *= -1
	}

	// パドル反射
	if g.ballY+ballSize >= screenHeight-paddleHeight &&
		g.ballX+ballSize >= g.paddleX && g.ballX <= g.paddleX+paddleWidth {
		g.ballVY *= -1
		g.ballY = screenHeight - paddleHeight - ballSize
	}

	// ブロック衝突
	for i := range g.blocks {
		b := &g.blocks[i]
		if b.active &&
			g.ballX+ballSize > b.x && g.ballX < b.x+blockWidth &&
			g.ballY+ballSize > b.y && g.ballY < b.y+blockHeight {
			b.active = false
			g.ballVY *= -1
			break
		}
	}

	// 下に落ちたらリセット
	if g.ballY > screenHeight {
		g.ballX = screenWidth - ballSize // 画面の右上
		g.ballY = 0
		g.ballVX = 2
		g.ballVY = -2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 背景を白で塗りつぶす
	screen.Fill(color.White)
	// タイマー表示（左上）
	// ebitenutil.DebugPrintAt(screen, "Time: "+fmt.Sprint(g.timer), 10, 10)
	// パドル
	ebitenutil.DrawRect(screen, g.paddleX, screenHeight-paddleHeight, paddleWidth, paddleHeight, color.RGBA{0, 0, 255, 255})
	// ボール
	ebitenutil.DrawRect(screen, g.ballX, g.ballY, ballSize, ballSize, color.RGBA{255, 0, 0, 255})
	// ブロック
	for _, b := range g.blocks {
		if b.active {
			ebitenutil.DrawRect(screen, b.x, b.y, blockWidth-2, blockHeight-2, color.RGBA{0, 255, 0, 255})
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
