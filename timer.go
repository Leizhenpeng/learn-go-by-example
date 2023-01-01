package main

import "time"

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	println("Timer 1 expired")

	//If you just wanted to wait, you could have used time.
	//Sleep. One reason a timer may be useful is that you can cancel the timer before it fires.
	//Here’s an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		println("Timer 2 expired")
	}()
	stop := timer2.Stop()
	if stop {
		println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}
