package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	currX, currY := robotgo.GetMousePos()
	afkCounter := 0

	for {
		x, y := robotgo.GetMousePos()
		if x == currX && y == currY {
			afkCounter++
		} else {
			afkCounter = 0
			currX, currY = x, y
		}

		if afkCounter > 5 {
			newX := rand.Intn(5120-2560+1) + 2560
			newY := rand.Intn(1440-1+1) + 1

			robotgo.MoveMouseSmooth(newX, newY, 0.5)
			currX, currY = robotgo.GetMousePos()
		}

		fmt.Printf("AFK COUNTER: %d\n", afkCounter)
		time.Sleep(2 * time.Second)
	}
}
