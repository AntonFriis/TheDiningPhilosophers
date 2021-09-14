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
var forkSetUse = 3
var forkSetFree = 4

//Fork values (java: fields)
type Fork struct {
	name      int
	state     int //Is either in use or free (see top)
	timesUsed int
	input     chan int
	output    chan int
}

//Fork constructer
func NewFork(forkNumber int, intputChannel, outputChannel chan int) Fork {
	var fork = Fork{forkNumber, forkIsFree, 0, intputChannel, outputChannel}
	return fork
}

//Fork gorouting function
//Loops forever, performs commands given via input channel (see top)
//Anwers via output channel if given question
func ForkStart(fork Fork) {
	for {
		//int given from input channel
		command := <-fork.input

		//cases of the command is descriped at the top
		switch command {
		case forkAskInUse:
			//will answer with either is in use or is free (see top)
			fork.output <- fork.state
		case forkAskTimesEaten:
			//Answers with the number of times the fork has been used
			fork.output <- fork.timesUsed
		case forkSetUse:
			//Set the forks state to in use
			forkAssert(fork, command) //checks that the fork isnt already in use
			fork.state = forkInUse
			fork.timesUsed++
		case forkSetFree:
			//Set the forks state to not in use and incroments the times it has been used
			forkAssert(fork, command) //checks that the fork isnt already not in use
			fork.state = forkIsFree
		}
	}
}

//Checks that forks wont change its state (inUse) to something that it is already doing
//Prints in Terminal if an error is detected
//Application will still continue
func forkAssert(fork Fork, command int) {
	if command == forkSetUse && fork.state == forkInUse {
		fmt.Printf("Error: Fork %d is already in use", fork.name)
	}
	if command == forkSetFree && fork.state == forkIsFree {
		fmt.Printf("Error: Fork %d is already free", fork.name)
	}
}
