package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type Instructions struct {
	retroFont font.Face
}

func (i *Instructions) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		gameState = MenuScreen
	}
}

func (i *Instructions) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	defaultColor := color.RGBA{108, 122, 137, 255}
	x := 50
	spacing := 40

	texts := []struct {
		Text  string
		Color color.Color
	}{
		{"Instructions", defaultColor},
		{"Move paddle: [Left] and [Right]", defaultColor},
		{"Launch the code: [Space]", defaultColor},
		{"Pause: [P]", defaultColor},
		{"Back to menu from here: [Esc]", defaultColor},
	}

	initialY := screenHeight/2 - len(texts)*spacing/2

	for n, t := range texts {
		y := initialY + n*spacing
		text.Draw(screen, t.Text, i.retroFont, x, y, t.Color)
	}
}

var instructions = &Instructions{
	retroFont: LoadFont(),
}
