package main

import (
	"fmt"
	"sync"
	"time"
)

var arbiter sync.Mutex
var arbiter2 sync.Mutex
var chans [5]chan int
var amountDoneEating int

func main() {
	for i := range chans {
		chans[i] = make(chan int, 2)
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
			fmt.Println("Terminate")
			return
		}
	}

}

func filo(forkNum1 int, forkNum2 int, filoNum int) {
	var timesEaten int = 0
	for {
		left := <-chans[forkNum1]
		right := <-chans[forkNum2]

		if left == 0 && right == 0 {
			chans[forkNum1] <- 1
			chans[forkNum2] <- 1
			timesEaten++
			fmt.Println("Philosopher", filoNum, "is eating")
			time.Sleep(5 * time.Second)
			if timesEaten == 3 {
				amountDoneEating++
			}
			fmt.Println("Philosopher", filoNum, "is thinking")
			time.Sleep(3 * time.Second)
		} else {
			chans[forkNum1] <- left
			chans[forkNum2] <- right
		}
	}
}

func fork(forkNum int) {
	for {
		x := <-chans[forkNum]
		if x != 0 {
			time.Sleep(6 * time.Second)
			chans[forkNum] <- 0
		} else {
			chans[forkNum] <- 0
		}
	}
}
