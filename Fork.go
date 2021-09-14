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
var forkSetUse = 3
var forkSetFree = 4

//Fork values (java: fields)
type Fork struct {
	name        int
	state       int //Is either in use or free (see top)
	timesUsed   int
	inputRight  chan int
	outputRight chan int
	inputLeft   chan int
	outputLeft  chan int
}

//Fork constructer
func NewFork(forkNumber int, intputChannelRight, outputChannelRight, intputChannelLeft, outputChannelLeft chan int) Fork {
	var fork = Fork{forkNumber, forkIsFree, 0, intputChannelRight, outputChannelRight, intputChannelLeft, outputChannelLeft}
	return fork

}

//Fork gorouting function
//Loops forever, performs commands given via input channel (see top)
//Anwers via output channel if given question
func ForkStart(fork Fork) {
	for {

	}
}
