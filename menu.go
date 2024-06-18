package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type MenuState string

type Menu struct {
	selection MenuState
	retroFont font.Face
}

const (
	GameStart        MenuState = "Start Game"
	GameInstructions MenuState = "Instructions"
	GameExit         MenuState = "Exit Game"
)

func (m *Menu) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		menu.NextSelection()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		menu.PrevSelection()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if m.selection == GameStart {
			gameState = Running
		} else if m.selection == GameInstructions {
			gameState = InstructionsScreen
		} else if m.selection == GameExit {
			gameState = ClosingScreen
		}
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// draw menu items
	screen.Fill(color.RGBA{0, 0, 0, 255})
	defaultColor := color.RGBA{108, 122, 137, 255}
	selectColor := color.RGBA{255, 255, 255, 255}

	startColor := defaultColor
	instructionsColor := defaultColor
	quitColor := defaultColor

	switch m.selection {
	case GameStart:
		startColor = selectColor
		break
	case GameInstructions:
		instructionsColor = selectColor
		break
	case GameExit:
		quitColor = selectColor
		break
	}
	x := 50
	spacing := 40

	texts := []struct {
		Text  string
		Color color.Color
	}{
		{"Start Breakout", startColor},
		{"Instructions", instructionsColor},
		{"Quit", quitColor},
	}

	initialY := screenHeight/2 - len(texts)*spacing/2

	for i, t := range texts {
		y := initialY + i*spacing
		text.Draw(screen, t.Text, m.retroFont, x, y, t.Color)
	}
}

func (m *Menu) NextSelection() {
	switch m.selection {
	case GameStart:
		m.selection = GameInstructions
		break
	case GameInstructions:
		m.selection = GameExit
		break
	case GameExit:
		m.selection = GameStart
		break
	}
}

func (m *Menu) PrevSelection() {
	switch m.selection {
	case GameStart:
		m.selection = GameExit
		break
	case GameInstructions:
		m.selection = GameStart
		break
	case GameExit:
		m.selection = GameInstructions
		break
	}
}

var (
	menu = &Menu{
		selection: GameStart,
		retroFont: LoadFont(),
	}
)
