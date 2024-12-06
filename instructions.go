package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type Instructions struct {
	retroFont font.Face
	mouseX    int
	mouseY    int
}

var (
	backButtonColor = DefaultColor
)

func (i *Instructions) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		gameState = MenuScreen
	}
	i.mouseX, i.mouseY = ebiten.CursorPosition()
	if i.isMouseMoved() {
		if i.mouseHoverOnBackButton() {
			backButtonColor = SelectColor
		} else {
			backButtonColor = DefaultColor
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if i.mouseHoverOnBackButton() {
			gameState = MenuScreen
		}
	}
}

func (i *Instructions) isMouseMoved() bool {
	return mouseX != i.mouseX || mouseY != i.mouseY
}

func (i *Instructions) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	x := 50
	spacing := 40

	texts := []struct {
		Text  string
		Color color.Color
	}{
		{"Instructions: Keyboard", DefaultColor},
		{"Move paddle: [Left] and [Right]", DefaultColor},
		{"Launch the ball: [Space]", DefaultColor},
		{"Pause: [P]", DefaultColor},
		{"Back to menu from here: [Esc]", DefaultColor},
		{"", DefaultColor},
		{"Instructions: Mouse/Touch", DefaultColor},
		{"Move paddle: Side [Click]/Tap", DefaultColor},
		{"Launch the ball: Quick [Click]/Tap", DefaultColor},
		{"Back/[Esc]", backButtonColor},
	}

	initialY := screenHeight/2 - len(texts)*spacing/2

	for n, t := range texts {
		y := initialY + n*spacing
		text.Draw(screen, t.Text, i.retroFont, x, y, t.Color)
	}
}

func (i *Instructions) mouseHoverOnBackButton() bool {
	return i.mouseX >= 50 && i.mouseX <= 50+150 && i.mouseY >= screenHeight-100 && i.mouseY <= screenHeight-50
}

var instructions = &Instructions{
	retroFont: LoadFont(),
}
