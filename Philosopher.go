package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	leftIN   chan int
	leftOUT  chan int
	rightIN  chan int
	rightOUT chan int
}

func NewPhil(leftIN, leftOUT, rightIN, rightOUT chan int) Philosopher {
	var phil = Philosopher{leftIN, leftOUT, rightIN, rightOUT}
	return phil
}

func checkLeft(leftIN, leftOUT chan int) int {
	leftOUT <- forkAskInUse
	if <-leftIN == forkIsFree {
		fmt.Println("Left fork is available")
		return forkIsFree
	} else {
		return forkInUse
	}

}

func checkRight(rightIN, rightOUT chan int) int {
	rightOUT <- forkAskInUse
	if <-rightIN == forkIsFree {
		fmt.Println("Right fork is available")
		return forkIsFree
	} else {
		fmt.Println("Right fork is in use")
		return forkInUse
	}
}

func action(phil Philosopher) {

	if checkLeft(phil.leftIN, phil.leftOUT) == forkIsFree && checkRight(phil.rightIN, phil.rightOUT) == forkIsFree {
		fmt.Println("Philosopher is eating")
		d := 0.50
		s := time.Duration(float64(time.Hour.Seconds()*1) * d)
		time.Sleep(s)

		phil.leftOUT <- forkSetFree
		phil.rightOUT <- forkSetFree
	} else if checkLeft(phil.leftIN, phil.leftOUT) == forkIsFree && checkRight(phil.rightIN, phil.rightOUT) == forkInUse {
		phil.leftOUT <- forkIsFree
		fmt.Printf("Philosopher is thinking - right in use")
		d := 0.33
		s := time.Duration(float64(time.Hour.Seconds()*1) * d)
		time.Sleep(s)
	} else if checkLeft(phil.leftIN, phil.leftOUT) == forkInUse && checkRight(phil.rightIN, phil.rightOUT) == forkIsFree {
		phil.rightOUT <- forkSetFree
		fmt.Printf("Philosopher is thinking - left in use")
		d := 0.33
		s := time.Duration(float64(time.Hour.Seconds()*1) * d)
		time.Sleep(s)
	} else {
		fmt.Printf("Philosopher is thinking both in use")
		d := 0.33
		s := time.Duration(float64(time.Hour.Seconds()*1) * d)
		time.Sleep(s)
	}

}
