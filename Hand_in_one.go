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

	for amountDoneEating < 6 {
		if amountDoneEating >= 5 {
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
		if filoNum%2 == 1 {
			left = <-chans[forkNum1]
			right = <-chans[forkNum2]
		} else {
			right = <-chans[forkNum2]
			left = <-chans[forkNum1]
		}

		if left == 0 && right == 0 {
			chans[forkNum1] <- 1
			chans[forkNum2] <- 1
			timesEaten++
			fmt.Println("Philosopher", filoNum, "is eating")
			time.Sleep(50 * time.Millisecond)

			//Checks if the phiolsopher has eaten 3 times
			if timesEaten == 3 {
				amountDoneEating++
			}

			fmt.Println("Philosopher", filoNum, "is thinking")
			time.Sleep(50 * time.Millisecond)
		}
		chans[forkNum1] <- left
		chans[forkNum2] <- right
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
