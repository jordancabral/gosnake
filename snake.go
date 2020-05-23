package main

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Directions. All snake possible directions
const (
	DirectionRight = "RIGHT"
	DirectionLeft  = "LEFT"
	DirectionUp    = "UP"
	DirectionDown  = "DOWN"
)

type Snake struct {
	*tl.Entity
	prevX     int
	prevY     int
	velocity  int // milliseconds
	direction chan string
}

// Snake constructor
func newSnake(x int, y int) Snake {
	directionChan := make(chan string)
	snake := Snake{tl.NewEntity(x, y, 2, 2), x, y, 200, directionChan}
	go snake.start()
	return snake
}

// Tick function with keyboard events
func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		snake.prevX, snake.prevY = snake.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			snake.advance(DirectionRight)
		case tl.KeyArrowLeft:
			snake.advance(DirectionLeft)
		case tl.KeyArrowUp:
			snake.advance(DirectionUp)
		case tl.KeyArrowDown:
			snake.advance(DirectionDown)
		}
	}
}

func (snake *Snake) advance(direction string) {
	snake.direction <- direction
}

func (snake *Snake) start() {

	stop := make(chan bool)
	starting := true
	for direction := range snake.direction {

		// Don't send stop message in first run
		if !starting {
			// Stop the advance in previous direction
			stop <- true
		}

		// Start to advance in received direction
		go func(dir string) {
			timer := time.Tick(time.Millisecond * time.Duration(snake.velocity))
			for {
				select {
				case <-timer:
					snake.goToDirection(dir)
				case <-stop:
					return
				}
			}
		}(direction)

		starting = false
	}
}

func (snake *Snake) goToDirection(direction string) {
	snake.prevX, snake.prevY = snake.Position()
	switch direction {
	case DirectionRight:
		snake.SetPosition(snake.prevX+1, snake.prevY)
	case DirectionLeft:
		snake.SetPosition(snake.prevX-1, snake.prevY)
	case DirectionUp:
		snake.SetPosition(snake.prevX, snake.prevY-1)
	case DirectionDown:
		snake.SetPosition(snake.prevX, snake.prevY+1)
	}
}

// Collide function to reset position when hit frame
func (snake *Snake) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		snake.SetPosition(snake.prevX, snake.prevY)
	}
}
