package The_dining_philosophers

/*
  each philosopher must include two channels (one for input and one
  for output, both usable from the outside) through which it is
  possible to make queries on the state of the philosopher (number of
  times eaten, eating or thinking)
*/

type Philosopher struct {
	name      int
	hasFork   bool
	isEating  bool
	input     chan int
	output    chan int
	leftFork  *Fork
	rightFork *Fork

	Eat func(t bool)
}

func NewPhilosopher(philosopherNumber int, lf, rf *Fork) *Philosopher {
	p := new(Philosopher)
	p.name = philosopherNumber
	p.hasFork = false
	p.isEating = false
	p.input = make(chan int)
	p.output = make(chan int)
	p.leftFork = lf
	p.rightFork = rf
	return p
}

func Eat()
