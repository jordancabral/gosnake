package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type Apple struct {
	*tl.Rectangle
	remove chan bool
}

func NewApple(x int, y int) *Apple {
	removeChan := make(chan bool)
	return &Apple{tl.NewRectangle(x, y, 1, 1, tl.ColorRed), removeChan}
}

func setApples(level *tl.BaseLevel) {
	rand.Seed(time.Now().UnixNano())
	randx := rand.Intn(48-3) + 3
	randy := rand.Intn(18-3) + 3

	apple := NewApple(randx, randy)
	level.AddEntity(apple)
	go func() {
		<-apple.remove
		level.RemoveEntity(apple)
		go setApples(level)
	}()
}
