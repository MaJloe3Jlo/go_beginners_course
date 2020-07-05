package math

// Minus(numbers ...int) recieve numbers.
// Return result of minus all numbers.
func Minus(numbers ...int) int {
	var result int = numbers[0]
	for _, value := range numbers[1:] {
		result -= value
	}
	return result
}
