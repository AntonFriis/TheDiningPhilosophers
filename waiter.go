package main

import "fmt"

var philIn, philOut, forkIn, forkOut []chan int
var philoInQue int = -1

func WaiterStartup(runtimes int) {
	//sets philosopher 0 and 3 to eat
	beginEating()
	var i = 0
	for i < runtimes {
		if philoInQue != -1 {
			var other int
			if philoInQue == 0 {
				other = 4
			} else {
				other = philoInQue - 1
			}
			if askFork(philoInQue) && askFork(other) {
				lockFork(philoInQue)
				lockFork(other)
				philIn[philoInQue] <- philosopherSetEating
				philoInQue = -1
			}
		}

		select {
		case x := <-philOut[0]:
			fmt.Println("Philosopher 0 is thinking")
			fmt.Printf("He has eaten %d", x)
			changeEater(0)
		case x := <-philOut[1]:
			fmt.Println("Philosopher 1 is thinking")
			fmt.Printf("He has eaten %d", x)
			changeEater(1)
		case x := <-philOut[2]:
			fmt.Println("Philosopher 2 is thinking")
			fmt.Printf("He has eaten %d", x)
			changeEater(2)
		case x := <-philOut[3]:
			fmt.Println("Philosopher 3 is thinking")
			fmt.Printf("He has eaten %d", x)
			changeEater(3)
		case x := <-philOut[4]:
			fmt.Println("Philosopher 4 is thinking")
			fmt.Printf("He has eaten %d", x)
			changeEater(4)
		}

		i++
	}

}
func beginEating() {
	lockFork(0)
	lockFork(4)
	philIn[0] <- philosopherSetEating

	lockFork(3)
	lockFork(2)
	philIn[3] <- philosopherSetEating

}

func changeEater(curret int) {
	var other int
	if curret == 0 {
		other = 4
	} else {
		other = curret - 1
	}
	unlockForks(curret, other)
	other = curret
	if curret == 4 {
		curret = 0
	} else {
		curret++
	}
	if askFork(curret) && askFork(other) {
		lockFork(curret)
		lockFork(other)
		philIn[curret] <- philosopherSetEating

	} else {
		philoInQue = curret
	}

}

func unlockForks(chanl1, chanl2 int) {
	forkIn[chanl1] <- forkSetFree
	forkIn[chanl2] <- forkSetFree
}
func askFork(chanl int) bool {
	forkIn[chanl] <- forkAskInUse
	answer := <-forkOut[chanl]
	return answer == -1
}
func lockFork(chanl int) {
	forkIn[chanl] <- forkSetUse
}
