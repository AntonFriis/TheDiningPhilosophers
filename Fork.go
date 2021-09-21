package main

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
	number      int
	state       bool
	timesUsed   int
	inputRight  chan int
	outputRight chan int
	inputLeft   chan int
	outputLeft  chan int
}

func NewFork(counter int, outputChannelLeft, inputChannelLeft, outputChannelRight, inputChannelRight chan int) Fork {
	var fork = Fork{counter, true, 0, inputChannelRight, outputChannelRight, inputChannelLeft, outputChannelLeft}
	return fork

}

func ForkStart(fork Fork) {

	for {

		select {
		case x := <-fork.inputRight:

			if x == forkAskInUse {

				if fork.state {
					fork.state = false
					fork.outputRight <- forkIsFree
				} else {
					fork.outputRight <- forkInUse
				}
			} else {
				fork.timesUsed++
				timesUsed[fork.number] = fork.timesUsed
				fork.state = true
			}

		case x := <-fork.inputLeft:

			if x == forkAskInUse {

				if fork.state {
					fork.state = false
					fork.outputLeft <- forkIsFree

				} else {
					fork.outputLeft <- forkInUse
				}
			} else {
				fork.timesUsed++
				timesUsed[fork.number] = fork.timesUsed
				fork.state = true
			}

		}

	}
}
