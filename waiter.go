package main

var philIn, philOut, forkIn, forkOut []chan int
var philoInQue int = -1

func WaiterStartup(runtimes int) {

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
			x++
		case x := <-philOut[1]:
			x++
		case x := <-philOut[2]:
			x++
		case x := <-philOut[3]:
			x++
		case x := <-philOut[4]:
			x++
		}

		i++
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
