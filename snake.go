package main

import (
	"time"

	tl "github.com/JoelOtter/termloop"
	"github.com/google/logger"
)

// Directions. All snake possible directions
const (
	DirectionRight = "RIGHT"
	DirectionLeft  = "LEFT"
	DirectionUp    = "UP"
	DirectionDown  = "DOWN"
	DirectionStop  = "STOP"
)

type Snake struct {
	*tl.Entity
	prevX     int
	prevY     int
	velocity  int // milliseconds
	direction chan string
	end       chan bool
}

// Snake constructor
func newSnake(x int, y int) Snake {
	directionChan := make(chan string)
	endChan := make(chan bool)
	snake := Snake{tl.NewEntity(x, y, 2, 2), x, y, 200, directionChan, endChan}
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
	logger.Infof("Advance %s", direction)
	snake.direction <- direction
}

func (snake *Snake) stop() {
	logger.Info("Stop snake")
	snake.direction <- DirectionStop
	snake.end <- true
}

func (snake *Snake) start() {

	logger.Info("Start snake")

	stop := make(chan bool)
	starting := true
	for direction := range snake.direction {

		if direction == DirectionStop {
			// Stop the advance and exit
			stop <- true
			close(snake.direction)
			close(stop)
			return
		}

		// Don't send stop message in first run
		if !starting {
			// Stop the advance in previous direction
			stop <- true
		}

		// Start to advance in received direction
		if direction != DirectionStop {
			go func(dir string) {
				timer := time.Tick(time.Millisecond * time.Duration(snake.velocity))
				for {
					select {
					case <-timer:
						logger.Infof("Tick to direction %s. Current position %d, %d", direction, snake.prevX, snake.prevY)
						snake.goToDirection(dir)
					case <-stop:
						logger.Info("Stop tick")
						return
					}
				}
			}(direction)
		}

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
		logger.Infof("Collide with rectangle. Prev position %d %d", snake.prevX, snake.prevY)
		snake.SetPosition(snake.prevX, snake.prevY)
		snake.stop()
	}
}
