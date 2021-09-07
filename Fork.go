package The_dining_philosophers

/*
- each fork must include two channels (one for input and one for
  output, both usable from outside) through which it is possible to
  make queries on the state of the fork (number of times used, in use
  or free)
*/

var inputFork = make(chan int)
var outputFork = make(chan int)

type Fork struct {
	name      int
	timesUsed int
}

func NewFork(forkNumber int) *Fork {
	fork := new(Fork)
	fork.name = forkNumber
	return fork
}

func forkfunc(input chan int, output chan int) {
	var comand int
	for {
		comand = <-input

		switch comand {
		case 1:

		case 2:

		case 3:
		}
	}

}
