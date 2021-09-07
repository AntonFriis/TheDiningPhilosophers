package The_dining_philosophers

/*
- each fork must include two channels (one for input and one for
  output, both usable from outside) through which it is possible to
  make queries on the state of the fork (number of times used, in use
  or free)
*/

type Fork struct {
	name      int
	timesUsed int
	input     chan int
	output    chan int
	inUse     bool
	f         func()
}

func NewFork() *Fork {
	var fork *Fork = new(Fork)

	fork.input = make(chan int)
	fork.output = make(chan int)
	return fork
}

func Forkfunc(fork Fork) {
	var comand int

	for {
		comand = <-fork.input

		switch comand {
		case 1:
			fork.output <- fork.timesUsed

		case 2:
			if fork.inUse {
				//Fork is taken
				fork.output <- -1
			} else {
				//Fork is free
				fork.output <- -2
			}
		case 3:
			switchstate(fork)
		}
	}

}

func switchstate(fork Fork) {

	if fork.inUse {
		fork.inUse = false
		//Fork is free
		fork.output <- -2
	} else {
		fork.inUse = true
		fork.timesUsed++
		//Fork is taken
		fork.output <- -1
	}

}
