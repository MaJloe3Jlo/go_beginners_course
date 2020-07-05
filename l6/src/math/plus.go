package math

// Plus(numbers ...int) int recieving many numbers.
// Return result plus of this numbers.
func Plus(numbers ...int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}
