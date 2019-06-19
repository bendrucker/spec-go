package math

func Add(a, b int) int {
	return a + b
}

func AddVariadic(addends ...int) int {
	sum := 0

	for _, addend := range addends {
		sum += addend
	}

	return sum
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}
