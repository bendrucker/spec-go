package spec

func BufferedChannel() chan string {
	channel := make(chan string, 2)

	channel <- "hello"
	channel <- "world"

	return channel
}
