package main

import "fmt"

var frin, frou, phin, phou []chan int

func InitWaiter(forkIn, forkOut, philIn, philOut []chan int) {
	frin, frou, phin, phou = forkIn, forkOut, philIn, philOut
}

func StartWaiter() {
	fmt.Println("starting")

	for i := 0; i < 20; i++ {
		for {
			select {
			case <-phou[0]:
				p := 0
				if checkForks(p) {
					setEat(p)
				}
			case <-phou[1]:
				p := 1
				if checkForks(p) {
					setEat(p)
				}
			case <-phou[2]:
				p := 2
				if checkForks(p) {
					setEat(p)
				}
			case <-phou[3]:
				p := 3
				if checkForks(p) {
					setEat(p)
				}
			case <-phou[4]:
				p := 4
				if checkForks(p) {
					setEat(p)
				}
			default:
				fmt.Println("no input")
				break
			}
		}

		/*for {
			select {
			case fork0Out := <-phou[0]:
			case fork1Out := <-phou[1]:
			case fork2Out := <-phou[2]:
			case fork3Out := <-phou[3]:
			case fork4Out := <-phou[4]:
			default:
				break
			}
		}*/
	}
}

func checkForks(philosopher int) bool {
	var fork1, fork2 = philosopher - 1, philosopher
	if fork1 < 0 {
		fork1 = 4
	}

	fmt.Printf("Fork %d will now be asked", fork1)
	frin[fork1] <- forkAskInUse
	fork1Status := <- frou[fork1]
	fmt.Printf("Fork %d is ready", fork1)
	fmt.Printf("Fork %d will now be asked", fork2)
	frin[fork2] <- forkAskInUse
	fork2Status := <- frou[fork2]
	fmt.Printf("Fork %d is ready", fork2)

	if fork1Status == forkIsFree && fork2Status == forkIsFree {
		return true
	} else {
		return false
	}
}

func setEat(philosopher int) {
	fmt.Printf("Phil %d will now eat", philosopher)
	phin[philosopher] <- philosopherSetEating
	eaten := <- phou[philosopher]
	fmt.Printf("Phil %d has eate %d times", philosopher, eaten)
}
