package main

import "fmt"

/*
- each fork must include two channels (one for input and one for
  output, both usable from outside) through which it is possible to
  make queries on the state of the fork (number of times used, in use
  or free)
*/

//Fork states
var forkIsFree = -1
var forkInUse = -2

//Fork input commands
var forkAskInUse = 1
var forkAskTimesEaten = 2
var forkSetFree = 4

type Fork struct {
	state       bool
	timesUsed   int
	inputRight  chan int
	outputRight chan int
	inputLeft   chan int
	outputLeft  chan int
}

func NewFork(outputChannelRight, intputChannelRight, outputChannelLeft, intputChannelLeft chan int) Fork {
	var fork = Fork{true, 0, outputChannelRight, intputChannelRight, outputChannelLeft, intputChannelLeft}
	fmt.Println("FORK IS MADE")
	return fork

}

func ForkStart(fork Fork) {
	fmt.Println("hej")
	for {

		select {
		case x := <-fork.inputRight:
			fmt.Println("jeg er Right")
			if x == forkAskInUse {

				if fork.state {
					fork.state = false
					fork.outputRight <- forkIsFree
				} else {
					fork.outputRight <- forkInUse
				}
			} else {
				fork.state = true
			}

		case x := <-fork.inputLeft:
			fmt.Println("jeg er Left")
			if x == forkAskInUse {

				if fork.state {
					fork.state = false
					fork.outputLeft <- forkIsFree
				} else {
					fork.outputLeft <- forkInUse
				}
			} else {
				fork.state = true
			}

		}

	}
}
