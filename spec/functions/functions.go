package functions

func ReturnMultiple() (string, string) {
	return "foo", "bar"
}

func CreateIncrementerClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var initialTerms = []int{0, 1}

func Fibonacci(n int) int {
	if n+1 <= len(initialTerms) {
		return initialTerms[n]
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}
