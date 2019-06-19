package sync

import "sync"

func OnceIncrement(value int) int {
	var once sync.Once

	once.Do(func() {
		value++
	})

	once.Do(func() {
		value++
	})

	return value
}
