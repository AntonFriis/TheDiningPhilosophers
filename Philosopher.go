package main

import (
	"fmt"
	"time"
)

/*
  each philosopher must include two channels (one for input and one
  for output, both usable from the outside) through which it is
  possible to make queries on the state of the philosopher (number of
  times eaten, eating or thinking)
*/

//Philosopher states
var philosopherIsThinking = -1
var philosopherIsEating = -2

//Philosopher input commands
var philosopherAskIsEating = 1
var philosopherAskTimesEaten = 2
var philosopherSetEating = 3
var philosopherSetThinking = 4

//Philosopher output commands
var philosopherRecuestEating = -3

//Philosopher values (java: fields)
type Philosopher struct {
	name       int
	hasFork    bool
	state      int //Is either eating or thinking (see top)
	timesEaten int
	input      chan int
	output     chan int
	leftFork   Fork
	rightFork  Fork
}

//Philosopher constructer
func NewPhilosopher(philosopherNumber int, intputChannel, outputChannel chan int, lf, rf Fork) Philosopher {
	var philosopher = Philosopher{philosopherNumber, false, philosopherIsThinking, 0, intputChannel, outputChannel, lf, rf}
	return philosopher

	/*philosopher := new(Philosopher)
	philosopher.name = philosopherNumber
	philosopher.hasFork = false
	philosopher.state = philosopherIsThinking
	philosopher.timesEaten = 0
	philosopher.input = intputChannel
	philosopher.output = outputChannel
	philosopher.leftFork = lf
	philosopher.rightFork = rf
	return philosopher*/
}

//Philosopher gorouting function
//Loops forever, performs commands given via input channel (see top)
//Anwers via output channel if given question

func PhilosopherStart(philosopher Philosopher) {
	for {
		if philosopher.state == philosopherIsThinking {
			go recuestEating(philosopher)
		} else if philosopher.state == philosopherAskIsEating {
			time.Sleep(time.Millisecond * 10)
			stopEating(philosopher)
		}
		command := <-philosopher.input

		//cases of the command is descriped at the top
		switch command {
		case philosopherAskIsEating:
			//will answer with either is eating or is thinking (see top)
			philosopher.output <- philosopher.state
		case philosopherAskTimesEaten:
			//Answers with the number of times the philosopher has eaten
			philosopher.output <- philosopher.timesEaten
		case philosopherSetEating:
			//Set the philosophers state to eating
			philosopherAssert(philosopher, command) //checks that the philosopher isnt already eating
			philosopher.state = philosopherIsEating
			philosopher.timesEaten++

		case philosopherSetThinking:
			//Set the philosophers state to thinking and incroments the times he has eaten
			philosopherAssert(philosopher, command) //checks that the philosopher isnt already thinking
			philosopher.state = philosopherIsThinking
		}
	}
}

func recuestEating(philosopher Philosopher) {
	philosopher.output <- philosopherRecuestEating
}

func stopEating(philosopher Philosopher) {
	philosopher.state = philosopherIsThinking
	philosopher.output <- philosopher.timesEaten
}

//Checks that philosopher wont change its state (isEating) to something that it is already doing
//Prints in Terminal if an error is detected
//Application will still continue
func philosopherAssert(philosopher Philosopher, command int) {
	if command == philosopherSetEating && philosopher.state == philosopherIsEating {
		fmt.Printf("Error: Philosopher %d is already in use", philosopher.name)
	}
	if command == philosopherSetThinking && philosopher.state == philosopherIsThinking {
		fmt.Printf("Error: Philosopher %d is already free", philosopher.name)
	}
}
