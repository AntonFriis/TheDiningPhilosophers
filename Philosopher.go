package main

import "time"

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

	var gotFork bool = true
	for {
		for gotFork {
			if phil.handSide && checkRight(phil.rightIN, phil.rightOUT) == forkIsFree {
				gotFork = false

			} else if checkLeft(phil.leftIN, phil.leftOUT) == forkIsFree {
				gotFork = false
			}
		}
		gotFork = true
		for gotFork {
			if phil.handSide && checkLeft(phil.leftIN, phil.leftOUT) == forkIsFree {
				gotFork = false

			} else if checkRight(phil.rightIN, phil.rightOUT) == forkIsFree {
				gotFork = false

			}
		}
		phil.timesEaten++
		timesEatenByAll[phil.number] = phil.timesEaten
		phil.rightOUT <- forkSetFree
		phil.leftOUT <- forkSetFree
		time.Sleep(time.Nanosecond)

	}

}
