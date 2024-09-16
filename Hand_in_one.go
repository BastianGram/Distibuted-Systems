package main

import (
	"fmt"
	"time"
)

var chans [5]chan int
var amountDoneEating int

func main() {
	for i := range chans {
		chans[i] = make(chan int, 1)
		chans[i] <- 0 // 0 means available, 1 means in use
	}

	go filo(0, 1, 1)
	go filo(1, 2, 2)
	go filo(2, 3, 3)
	go filo(3, 4, 4)
	go filo(4, 0, 5)

	go fork(0)
	go fork(1)
	go fork(2)
	go fork(3)
	go fork(4)

	// The main function loops until all philosophers have eaten 3 times or more
	for amountDoneEating < 6 {
		if amountDoneEating >= 5 { // Once 5 philosophers have eaten 3 times
			fmt.Println("Everyone has now eaten 3 or more times")
			return
		}
	}
}

func filo(forkNum1 int, forkNum2 int, filoNum int) {
	var timesEaten int = 0
	for {
		left := 0
		right := 0
		// Alternate fork acquisition order to prevent deadlock
		// Odd-numbered philosophers pick up the left fork first, then the right fork
		// Even-numbered philosophers pick up the right fork first, then the left fork
		// This breaks the circular wait condition that can cause deadlock
		if filoNum%2 == 1 {
			left = <-chans[forkNum1]  // Pick up left fork
			right = <-chans[forkNum2] // Pick up right fork
		} else {
			right = <-chans[forkNum2] // Pick up right fork
			left = <-chans[forkNum1]  // Pick up left fork
		}

		if left == 0 && right == 0 {
			chans[forkNum1] <- 1
			chans[forkNum2] <- 1
			timesEaten++
			fmt.Println("Philosopher", filoNum, "is eating")
			time.Sleep(100 * time.Millisecond)

			//Checks if the phiolsopher has eaten 3 times
			if timesEaten == 3 {
				amountDoneEating++ // Increase the global counter if philosopher eats 3 times
			}

			fmt.Println("Philosopher", filoNum, "is thinking")
			time.Sleep(100 * time.Millisecond)
		}
		// Release the forks by putting them back in the channels
		chans[forkNum1] <- left  // Release left fork
		chans[forkNum2] <- right // Release right fork
	}
}

func fork(forkNum int) {
	for {
		x := <-chans[forkNum]
		if x == 0 {
			chans[forkNum] <- 0
		}
	}
}
