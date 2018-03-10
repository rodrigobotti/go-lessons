package math

// Div divides x by y
func Div(x int, y int) (res int) {
	res = x / y
	return
}

// DivWithRemainder returns result and remainder of x / y
func DivWithRemainder(x int, y int) (int, int) {
	return x / y, x % y
}
