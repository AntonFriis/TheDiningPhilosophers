package main

func main() {
	var channels [20]chan int
	for i := 0; i < 20; i++ {
		channels[i] = make(chan int)
	}
	var forks []Fork
	var philosophers []Philosopher
	for i := 0; i < 5; i = i + 4 {
		forks = append(forks, NewFork(channels[i], channels[i+1], channels[i+2], channels[i+3]))
		var x int
		if i == 0 {
			x = 19
		} else {
			x = i - 1
		}
		philosophers = append(philosophers, NewPhil(channels[x-1], channels[x], channels[i], channels[i+1]))

	}

}
