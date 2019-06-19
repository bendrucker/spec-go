package spec

import "time"

func SetTimeout(fn func(), delay int) {
	<-time.After(time.Duration(delay) * time.Millisecond)
	fn()
}
