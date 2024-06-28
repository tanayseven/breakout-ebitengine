package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"image/color"
)

type MenuState string

type MenuOption struct {
	Text   string
	Color  color.Color
	Bounds image.Rectangle
	State  MenuState
}

type Menu struct {
	selection   MenuState
	retroFont   font.Face
	menuOptions []MenuOption
	mouseX      int
	mouseY      int
}

const (
	GameStart        MenuState = "Start Game"
	GameInstructions MenuState = "Instructions"
	GameExit         MenuState = "Exit Game"
)

var (
	DefaultColor = color.RGBA{R: 108, G: 122, B: 137, A: 255}
	SelectColor  = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)

func (m *Menu) Update() error {
	x := 50
	spacing := 40

	initialY := screenHeight/2 - len(m.menuOptions)*spacing/2
	for i, t := range m.menuOptions {
		y := initialY + i*spacing
		m.menuOptions[i].Bounds = image.Rect(x, y-spacing, x+t.Bounds.Dx(), y+t.Bounds.Dy())
	}
	mouseX, mouseY = ebiten.CursorPosition()
	if m.isMouseMoved() {
		m.mouseX, m.mouseY = ebiten.CursorPosition()
		for _, t := range m.menuOptions {
			if m.mouseX > t.Bounds.Min.X && m.mouseX < t.Bounds.Max.X && m.mouseY > t.Bounds.Min.Y && m.mouseY < t.Bounds.Max.Y {
				m.selection = t.State
			}
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for _, t := range m.menuOptions {
			if m.mouseX > t.Bounds.Min.X && m.mouseX < t.Bounds.Max.X && m.mouseY > t.Bounds.Min.Y && m.mouseY < t.Bounds.Max.Y {
				if t.State == GameStart {
					gameState = Running
				} else if t.State == GameInstructions {
					gameState = InstructionsScreen
				} else if t.State == GameExit {
					gameState = ClosingScreen
				}
			}
		}
	}
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

func (m *Menu) isMouseMoved() bool {
	return mouseX != m.mouseX || mouseY != m.mouseY
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// draw menu items
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	for i, _ := range m.menuOptions {
		m.menuOptions[i].Color = DefaultColor
	}

	switch m.selection {
	case GameStart:
		m.menuOptions[0].Color = SelectColor
		break
	case GameInstructions:
		m.menuOptions[1].Color = SelectColor
		break
	case GameExit:
		m.menuOptions[2].Color = SelectColor
		break
	}
	x := 50
	spacing := 40

	initialY := screenHeight/2 - len(m.menuOptions)*spacing/2

	for i, t := range m.menuOptions {
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
		menuOptions: []MenuOption{
			{Text: "Start Breakout", Color: DefaultColor, Bounds: text.BoundString(LoadFont(), "Start Breakout"), State: GameStart},
			{Text: "Instructions", Color: DefaultColor, Bounds: text.BoundString(LoadFont(), "Instructions"), State: GameInstructions},
			{Text: "Quit", Color: DefaultColor, Bounds: text.BoundString(LoadFont(), "Quit"), State: GameExit},
		},
	}
	mouseX, mouseY = ebiten.CursorPosition()
)
