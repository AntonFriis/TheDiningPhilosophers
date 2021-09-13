package The_dining_philosophers

/*
  each philosopher must include two channels (one for input and one
  for output, both usable from the outside) through which it is
  possible to make queries on the state of the philosopher (number of
  times eaten, eating or thinking)
*/

type Philosopher struct {
	name       int
	hasFork    bool
	isEating   bool
	timesEaten int
	input      chan int
	output     chan int
	leftFork   *Fork
	rightFork  *Fork
}

func NewPhilosopher(philosopherNumber int, intputChannel, outputChannel chan int, lf, rf *Fork) *Philosopher {
	philosopher := new(Philosopher)
	philosopher.name = philosopherNumber
	philosopher.hasFork = false
	philosopher.isEating = false
	philosopher.timesEaten = 0
	philosopher.input = intputChannel
	philosopher.output = outputChannel
	philosopher.leftFork = lf
	philosopher.rightFork = rf
	return philosopher
}

func Eat(p *Philosopher) {

}
