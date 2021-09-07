package The_dining_philosophers

/*
  each philosopher must include two channels (one for input and one
  for output, both usable from the outside) through which it is
  possible to make queries on the state of the philosopher (number of
  times eaten, eating or thinking)
*/

var inputPhilosopher = make(chan int)
var outputPhilosopher = make(chan int)

type Philosopher struct {
	name     int
	hasFork  bool
	isEating bool
}

func NewPhilosopher(philosopherNumber int) *Philosopher {
	philosopher := new(Philosopher)
	philosopher.name = philosopherNumber
	return philosopher
}
