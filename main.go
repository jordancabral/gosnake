package main

import (
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
	})

	level.AddEntity(tl.NewRectangle(1, 1, 50, 2, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(1, 3, 2, 19, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(1, 20, 50, 2, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(49, 3, 2, 17, tl.ColorBlue))

	snake := newSnake(10, 10)
	// Set the character at position (0, 0) on the entity.
	snake.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '█'})
	level.AddEntity(&snake)

	game.Screen().SetLevel(level)

	game.Start()

}
