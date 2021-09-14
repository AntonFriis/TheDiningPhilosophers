package main

import (
	"fmt"
	"time"
)

func main() {

	var channels [20]chan int
	for i := 0; i < 20; i++ {
		channels[i] = make(chan int)
	}
	var forks [5]Fork
	var philosophers [5]Philosopher
	var counter int = 0

	for i := 0; i < 20; i = i + 4 {

		forks[counter] = NewFork(channels[i], channels[i+1], channels[i+2], channels[i+3])
		var x int
		if i == 0 {
			x = 19
		} else {
			x = i - 1
		}
		philosophers[counter] = NewPhil(channels[x-1], channels[x], channels[i], channels[i+1])
		counter++
	}
	fmt.Println(counter)

	for i := 0; i < 5; i++ {
		go ForkStart(forks[i])
		fmt.Println(i)
	}
	for i := 0; i < 5; i++ {
		go action(philosophers[i])
	}
	time.Sleep(10 * time.Second)

}
