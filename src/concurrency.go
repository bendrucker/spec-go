package spec

import "time"

func pong(channel chan string) {
	message := <-channel
	if message == "ping" {
		channel <- "pong"
	}
}

func PingPongChannel() string {
	channel := make(chan string)

	go pong(channel)
	channel <- "ping"

	return <-channel
}

func BlockChannel() string {
	string := ""
	done := make(chan bool)

	work := func(done chan bool) {
		time.Sleep(1)
		string = "done"
		done <- true
	}

	go work(done)
	<-done
	return string
}

func DeferredIncrement() int {
	value := 0
	increment := func() {
		value++
	}

	defer increment()
	return value
}
