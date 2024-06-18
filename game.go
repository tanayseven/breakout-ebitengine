package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"os"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleWidth  = 100
	paddleHeight = 15
	ballSize     = 15
	playerSpeed  = 5
)

type GameState string

const (
	MenuScreen         GameState = "menu"
	InstructionsScreen GameState = "instructions"
	Running            GameState = "running"
	Paused             GameState = "paused"
	Over               GameState = "over"
	ClosingScreen      GameState = "blank"
)

var gameState = MenuScreen

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type GameObjects interface {
	Update(g GameState)
	Draw(screen *ebiten.Image)
}

func (g *Game) Update() error {
	if gameState == MenuScreen {
		menu.Update()
		currentDisplayedMessage = ""
		return nil
	}

	if gameState == InstructionsScreen {
		instructions.Update()
		currentDisplayedMessage = ""
		return nil
	}

	if gameState == Paused {
		currentDisplayedMessage = gamePausedMessage
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			gameState = Running
			currentDisplayedMessage = ""
		}
		return nil
	}

	if gameState == Over {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gameState = MenuScreen
			currentDisplayedMessage = gameStartMessage
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		gameState = Paused
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyAlt) && inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		return ebiten.Termination
	}

	paddle.Update(gameState)

	// Ball movement
	ball.PaddlePosition(*paddle)
	ball.Update(gameState)

	// Ball collision with walls
	if ball.y <= 0 || ball.y+ballSize >= screenHeight {
		ball.deltaY = -ball.deltaY
	}
	if ball.x <= 0 || ball.x+ballSize >= screenWidth {
		ball.deltaX = -ball.deltaX
	}

	// Ball collision with paddle
	if ball.y+ballSize >= paddle.y && ball.x >= paddle.x && ball.x <= paddle.x+paddleWidth {
		ball.deltaY = -ball.deltaY
	}

	// Ball is lost
	if ball.y+ballSize >= screenHeight {
		ball.reset()
		return nil
	}

	if ball.remainingBalls == 0 {
		gameState = Over
		currentDisplayedMessage = gameOverMessage
		return nil
	}

	if gameState == ClosingScreen {
		currentDisplayedMessage = ""
		time.Sleep(200 * time.Millisecond)
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if gameState == ClosingScreen {
		screen.Fill(color.RGBA{0, 0, 0, 0xff})
		return
	}

	if gameState == MenuScreen {
		menu.Draw(screen)
		return
	}

	if gameState == InstructionsScreen {
		instructions.Draw(screen)
		return
	}

	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	mplusNormalFont := LoadFont()
	DrawCenteredText(screen, currentDisplayedMessage, screenWidth/2, screenHeight/2, mplusNormalFont, color.White)

	paddle.Draw(screen)

	// Draw ball
	ball.Draw(screen)
}
