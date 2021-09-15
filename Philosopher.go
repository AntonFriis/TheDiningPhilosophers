package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	handSide   bool
	leftIN     chan int
	leftOUT    chan int
	rightIN    chan int
	rightOUT   chan int
	number     int
	timesEaten int
}

func NewPhil(number int, handSide bool, leftIN, leftOUT, rightIN, rightOUT chan int) Philosopher {
	var phil = Philosopher{handSide, leftIN, leftOUT, rightIN, rightOUT, number, 0}
	return phil
}

func checkLeft(leftIN, leftOUT chan int) int {

	leftOUT <- forkAskInUse

	var x int
	x = <-leftIN

	if x == forkIsFree {
		return forkIsFree
	} else {
		return forkInUse
	}

}

func checkRight(rightIN, rightOUT chan int) int {
	rightOUT <- forkAskInUse

	var x int = <-rightIN
	if x == forkIsFree {
		return forkIsFree
	} else {
		return forkInUse
	}
}

func action(phil Philosopher) {
	for {
		if phil.handSide {
			if checkRight(phil.rightIN, phil.rightOUT) == forkIsFree {

				if checkLeft(phil.leftIN, phil.leftOUT) == forkIsFree {
					phil.timesEaten++
					timesEatenByAll[phil.number] = phil.timesEaten
					fmt.Printf("Philosopher %d has eaten %d*******", phil.number, phil.timesEaten)
					fmt.Println()
					s := time.Millisecond * 500
					time.Sleep(s)
					phil.rightOUT <- forkSetFree
					phil.leftOUT <- forkSetFree
				} else {
					phil.rightOUT <- forkSetFree
					fmt.Println("Philosopher is thinking - left not available")
					s := time.Millisecond * 333
					time.Sleep(s)
				}
			}
		} else {
			if checkLeft(phil.leftIN, phil.leftOUT) == forkIsFree {

				if checkRight(phil.rightIN, phil.rightOUT) == forkIsFree {
					phil.timesEaten++
					timesEatenByAll[phil.number] = phil.timesEaten
					fmt.Printf("Philosopher %d has eaten %d", phil.number, phil.timesEaten)
					fmt.Println()
					s := time.Millisecond * 500
					time.Sleep(s)
					phil.rightOUT <- forkSetFree
					phil.leftOUT <- forkSetFree
				} else {
					phil.leftOUT <- forkSetFree
					fmt.Println("Philosopher is thinking - right not available")
					s := time.Millisecond * 333
					time.Sleep(s)
				}
			}
		}

	}

}
