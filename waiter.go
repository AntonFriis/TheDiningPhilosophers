package The_dining_philosophers

func WaiterStartup(philIn, philOut, forkIn, forkOut [5]chan int, runtimes int) {
	var philoInQue = -1

	var i = 0
	for i < runtimes {
		if philoInQue != -1 {

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

}

func fockForks(chanl int) {

}

func changeEater(curret int) {
	var other int
	if curret == 0 {
		other = 4
	} else {
		other = curret - 1
	}
	unlockForks(curret, other)

}
