package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"strings"
)

type Brick struct {
	x      int
	y      int
	width  int
	height int
	active bool
}

const (
	brickWidth   = 40
	brickHeight  = 20
	brickPadding = 2
	topOffset    = 50
)

var (
	bricks     []*Brick
	leftOffset = 0
)

func BrickInit(level int) {
	currentLevel := Levels[level]
	brickRows := strings.Split(currentLevel, "\n")
	brickCols := strings.Split(brickRows[0], "")
	leftOffset = (screenWidth - len(brickCols)*(brickWidth+brickPadding)) / 2
	totalBricks := 0
	for _, row := range brickRows {
		cleanedRow := strings.TrimSpace(row)
		removeInvisibleBricks := strings.ReplaceAll(cleanedRow, "-", "")
		totalBricks += len(removeInvisibleBricks)
	}
	bricks = make([]*Brick, totalBricks)

	countBricks := 0
	for i, row := range brickRows {
		for j, brick := range strings.TrimSpace(row) {
			if brick == '-' {
				continue
			}
			bricks[countBricks] = &Brick{
				x:      leftOffset + j*(brickWidth+brickPadding),
				y:      topOffset + i*(brickHeight+brickPadding),
				width:  brickWidth,
				height: brickHeight,
				active: true,
			}
			countBricks++
		}
	}
}

func (b *Brick) Draw(screen *ebiten.Image) {
	if !b.active {
		return
	}
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), color.White, true)
}

func (b *Brick) Update() {
	if !b.active {
		return
	}
}

func (b *Brick) CollidesWithHorizontalLine(ball Ball) bool {
	if !b.active {
		return false
	}
	return ball.x < b.x+b.width && ball.x+ball.deltaX > b.x && ball.y < b.y+b.height && ball.y+ball.deltaY > b.y
}

func (b *Brick) CollidesWithVerticalLine(ball Ball) bool {
	if !b.active {
		return false
	}
	return ball.y < b.y+b.height && ball.y+ball.deltaY > b.y && ball.x < b.x+b.width && ball.x+ball.deltaX > b.x
}

func AreBricksOver() bool {
	for _, brick := range bricks {
		if brick.active {
			return false
		}
	}
	return true
}

func (b *Brick) hit() {
	b.active = false
}
