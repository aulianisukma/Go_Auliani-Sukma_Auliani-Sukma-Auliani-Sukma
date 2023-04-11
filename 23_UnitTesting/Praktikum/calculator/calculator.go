package calculator

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}

func Division(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
