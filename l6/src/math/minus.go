package math

// Minus(numbers ...int) int recieving many numbers.
// Return result of of this numbers.
func Minus(numbers ...int) int {
	var result int
	for _, value := range numbers {
		result -= value
	}
	return result
}
