package The_dining_philosophers

func Main() {
	var forkInputChannels = make([]chan int, 5)
	var forkOutputChannels = make([]chan int, 5)
	var forks = make([]*Fork, 5)
	for f := 0; f < 5; f++ {
		forkInputChannels[f] = make(chan int)
		forkOutputChannels[f] = make(chan int)
		forks[f] = NewFork(f, forkInputChannels[f], forkOutputChannels[f])
	}

	var philosopherInputChannels = make([]chan int, 5)
	var philosopherOutputChannels = make([]chan int, 5)
	var philosophers = make([]Philosopher, 5)
	for p := 0; p < 5; p++ {
		philosopherInputChannels[p] = make(chan int)
		philosopherOutputChannels[p] = make(chan int)

		lf, rf := p-1, p
		if lf < 0 {
			lf = 4
		}
		philosophers[p] = *NewPhilosopher(p, philosopherInputChannels[p], philosopherOutputChannels[p], forks[lf], forks[rf])
	}

}
