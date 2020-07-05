package math

// Plus(numbers ...int) recieve numbers.Plus
// Return sum of numbers.
func Plus(numbers ...int) int {
	var result int
	for _, value := range numbers {
		result += value
	}
	return result
}
