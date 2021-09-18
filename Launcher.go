package main

import (
	"fmt"
	"time"
)

var timesEatenByAll [5]int

func main() {
	fmt.Println("Pleas write the seconds the program should run for")
	fmt.Println("If a non integer it will default to 40s")
	var seconds int

	_, err := fmt.Scanf("%d", &seconds)

	if err != nil {
		seconds = 40

	}

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
		if counter%2 == 0 {
			handSide = true
		}
		philosophers[counter] = NewPhil(counter, handSide, channels[i+2], channels[i+3], channels[(i+4)%20], channels[(i+5)%20])
		counter++
		readercount = readercount + 2
		handSide = false
	}

	for i := 0; i < 5; i++ {
		go ForkStart(forks[i])

	}
	for i := 0; i < 5; i++ {
		go action(philosophers[i])
	}
	fmt.Printf("The program wil print times eaten in %d seconds", seconds)
	fmt.Println()
	time.Sleep(time.Duration(seconds) * time.Second)

	for i := 0; i < 5; i++ {
		fmt.Printf("Philosopher %d has eaten %d times", i, timesEatenByAll[i])
		fmt.Println()

	}

}
