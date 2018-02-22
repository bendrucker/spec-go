package spec

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