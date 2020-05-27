package main

import (
	"flag"
	"os"

	tl "github.com/JoelOtter/termloop"
	"github.com/google/logger"
)

const logPath = "logger.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {

	// Log config
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()
	defer logger.Init("LoggerExample", *verbose, true, lf).Close()

	game := tl.NewGame()

	logger.Info("Config level")
	// Level Background
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
	})

	// Game Frame
	level.AddEntity(tl.NewRectangle(1, 1, 50, 2, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(1, 3, 2, 19, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(1, 20, 50, 2, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(49, 3, 2, 17, tl.ColorBlue))

	logger.Info("Creating Snake")
	// Snake character creation
	snake := newSnake(10, 10)
	// Set the character at position (0, 0) on the entity.
	snake.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '█'})
	level.AddEntity(&snake)

	game.Screen().SetLevel(level)

	// Game over
	go func() {
		<-snake.end
		logger.Info("Game over")
		gameOverLevel := tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorRed,
			Fg: tl.ColorBlack,
		})
		game.Screen().SetLevel(gameOverLevel)
	}()

	logger.Info("Starting game..")
	game.Start()

}
