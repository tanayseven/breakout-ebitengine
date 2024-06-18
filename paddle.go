package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Paddle struct {
	x         int
	y         int
	leftKeys  []ebiten.Key
	rightKeys []ebiten.Key
}

var (
	paddle = &Paddle{
		x:         screenWidth / 2,
		y:         screenHeight - paddleHeight,
		leftKeys:  []ebiten.Key{ebiten.KeyLeft, ebiten.KeyA},
		rightKeys: []ebiten.Key{ebiten.KeyRight, ebiten.KeyD},
	}
)

func (p *Paddle) Update(g GameState) {
	if g == Paused {
		return
	}
	isLeftKeyPressed := false
	for _, key := range p.leftKeys {
		if ebiten.IsKeyPressed(key) {
			isLeftKeyPressed = true
			break
		}
	}
	isRightKeyPressed := false
	for _, key := range p.rightKeys {
		if ebiten.IsKeyPressed(key) {
			isRightKeyPressed = true
			break
		}
	}
	if isLeftKeyPressed && p.x > 0 {
		p.x -= playerSpeed
	}
	if isRightKeyPressed && p.x+paddleWidth < screenWidth {
		p.x += playerSpeed
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.x), float32(p.y), float32(paddleWidth), float32(paddleHeight), color.White, true)
}
