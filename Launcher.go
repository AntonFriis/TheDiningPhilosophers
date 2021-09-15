package main

import (
	"fmt"
	"time"
)

var timesEatenByAll [5]int

func main() {

	var channels [20]chan int

	for i := 0; i < 20; i++ {
		channels[i] = make(chan int)
	}

	var forks [5]Fork
	var philosophers [5]Philosopher
	var counter int = 0
	var readercount int = 0
	var handSide bool = false
	for i := 0; i < 20; i = i + 4 {

		forks[counter] = NewFork(channels[i], channels[i+1], channels[i+2], channels[i+3])
		if counter == 4 {
			handSide = true
		}
		philosophers[counter] = NewPhil(counter, handSide, channels[i+2], channels[i+3], channels[(i+4)%20], channels[(i+5)%20])
		counter++
		readercount = readercount + 2
	}

	for i := 0; i < 5; i++ {
		go ForkStart(forks[i])

	}
	for i := 0; i < 5; i++ {
		go action(philosophers[i])
	}
	time.Sleep(10 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Printf("Philosopher %d has eaten %d times", i, timesEatenByAll[i])
		fmt.Println()

	}

}
