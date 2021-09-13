package The_dining_philosophers

func main() {
	//Instantiate 5 Forks and their channels
	var forkInputChannels = make([]chan int, 5)
	var forkOutputChannels = make([]chan int, 5)
	var forks = make([]*Fork, 5)
	for f := 0; f < 5; f++ {
		forkInputChannels[f] = make(chan int)
		forkOutputChannels[f] = make(chan int)
		forks[f] = NewFork(f, forkInputChannels[f], forkOutputChannels[f])
	}

	//Instantiate 5 Philosophers and their channels
	var philosopherInputChannels = make([]chan int, 5)
	var philosopherOutputChannels = make([]chan int, 5)
	var philosophers = make([]Philosopher, 5)
	for p := 0; p < 5; p++ {
		philosopherInputChannels[p] = make(chan int)
		philosopherOutputChannels[p] = make(chan int)

		lf, rf := p-1, p //Index of the philosophers left (lf) and right (rf) fork
		if lf < 0 {      //The first philosophers left fork is the last index of the fork array
			lf = len(forks) - 1
		}
		philosophers[p] = *NewPhilosopher(p, philosopherInputChannels[p], philosopherOutputChannels[p], forks[lf], forks[rf])
	}

	//Starts goroutines for all forks
	for f := 0; f < 5; f++ {
		go ForkStart(forks[f])
	}
	//Starts goroutines for all philosophers
	for p := 0; p < 5; p++ {
		go ForkStart(forks[p])
	}

	//INSTANTIATE WAITER AND START WAITER HERE!
	philIn = philosopherInputChannels
	philOut = philosopherOutputChannels
	forkIn = forkInputChannels
	forkOut = forkInputChannels
	WaiterStartup(5000000)
}
