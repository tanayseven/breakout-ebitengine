package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"strconv"
)

type BallState string

const (
	initialTimeSincePreviousClick = 10
)

const (
	Moving BallState = "moving"
	Reset  BallState = "reset"
)

type Ball struct {
	speed                  int
	x                      int
	y                      int
	deltaX                 int
	deltaY                 int
	remainingBalls         int
	state                  BallState
	timeSincePreviousClick int
}

const ballSpeedMax = 5
const ballSpeedMin = 3
const initialBalls = 3

func randInRange(min, max int) int {
	number := rand.Intn(max-min) + min
	for number < ballSpeedMin && number > -ballSpeedMin {
		number = rand.Intn(max-min) + min
	}
	return number
}

var (
	ball = Ball{
		speed:          ballSpeedMax,
		x:              screenWidth / 2,
		y:              screenHeight / 2,
		deltaX:         randInRange(-ballSpeedMax, ballSpeedMax),
		deltaY:         randInRange(-ballSpeedMax, ballSpeedMax),
		state:          Reset,
		remainingBalls: initialBalls,
	}
)

func (b *Ball) reset() {
	b.state = Reset
	b.remainingBalls--
	b.deltaX = randInRange(-ballSpeedMax, ballSpeedMax)
	b.deltaY = randInRange(-ballSpeedMax, ballSpeedMax)
}

func (b *Ball) PaddlePosition(p Paddle) {
	if b.state == Reset {
		b.x = p.x + paddleWidth/2 - ballSize/2
		b.y = p.y - ballSize - 3
	}
}

func (b *Ball) Update(g GameState) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && b.state == Reset {
		b.state = Moving
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if b.timeSincePreviousClick != 0 && b.state == Reset {
			b.timeSincePreviousClick--
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if b.timeSincePreviousClick != 0 && b.state == Reset {
			b.state = Moving
		}
		b.timeSincePreviousClick = initialTimeSincePreviousClick
	}
	if b.state == Moving {
		b.x += b.deltaX
		b.y += b.deltaY
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	// display remaining balls
	mplusNormalFont := LoadFont()
	DrawTopLeftText(screen, "Balls: "+strconv.Itoa(b.remainingBalls), 20, mplusNormalFont, color.White)
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(ballSize), float32(ballSize), color.White, true)
}

func (b *Ball) Lost() bool {
	return ball.y+ballSize >= screenHeight
}
