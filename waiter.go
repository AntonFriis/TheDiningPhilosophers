package main

import "fmt"

var frin, frou, phin, phou []chan int

func InitWaiter(forkIn, forkOut, philIn, philOut []chan int) {
	frin, frou, phin, phou = forkIn, forkOut, philIn, philOut
}

func status(philosopher, timesEaten int) {
	fmt.Printf("Phil %d has eaten %d times\n", philosopher, timesEaten)
}

func surroundingForks(philosopher int) (int, int) {
	var fork1, fork2 = philosopher - 1, philosopher //Index of the philosophers left (fork1) and right (fork2) fork
	if fork1 < 0 {                                  //The first philosophers left fork is the last index of the fork array
		fork1 = 4
	}
	return fork1, fork2
}

func StartWaiter() {
	fmt.Println("starting")

	for {
		select {
		case philosopherOutput := <-phou[0]:
			doPhilosopher(0, philosopherOutput)
		case philosopherOutput := <-phou[1]:
			doPhilosopher(1, philosopherOutput)
		case philosopherOutput := <-phou[2]:
			doPhilosopher(2, philosopherOutput)
		case philosopherOutput := <-phou[3]:
			doPhilosopher(3, philosopherOutput)
		case philosopherOutput := <-phou[4]:
			doPhilosopher(4, philosopherOutput)
			//default:
			//fmt.Println("no input")
			//continue
		}
	}
}

func doPhilosopher(pNumber, command int) {
	fmt.Printf("phil %d got command %d", pNumber, command)
	if command >= 0 {
		status(pNumber, command)
		//unlockForks(pNumber)
	} else if command == philosopherRecuestEating && checkForks(pNumber) {
		setEat(pNumber)
	}
}

func checkForks(philosopher int) bool {
	fork1, fork2 := surroundingForks(philosopher)
	fmt.Printf("Cheking fork %d and %d\n", fork1, fork2)
	frin[fork1] <- forkAskInUse
	frin[fork2] <- forkAskInUse

	fork1Status := <-frou[fork1]
	fork2Status := <-frou[fork2]
	fmt.Printf("fork %d and %d is ready\n", fork1, fork2)
	return fork1Status == forkIsFree && fork2Status == forkIsFree
}

func setEat(philosopher int) {
	fmt.Printf("Phil %d will now eat\n", philosopher)

	//lockForks(philosopher)

	phin[philosopher] <- philosopherSetEating
	//eaten := <- phou[philosopher]

	//status(philosopher, eaten)
	fmt.Printf("Phil %d has started eating\n", philosopher)
}

func lockForks(philosopher int) {
	fork1, fork2 := surroundingForks(philosopher)
	frin[fork1] <- forkSetUse
	frin[fork2] <- forkSetUse
}

func unlockForks(philosopher int) {
	fork1, fork2 := surroundingForks(philosopher)
	frin[fork1] <- forkSetFree
	frin[fork2] <- forkSetFree
}
