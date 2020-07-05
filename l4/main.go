package main

import "fmt"

func changeValue(x int) {
	x = x + 100
}

func changeValueByPtr(x *int) {
	*x = *x + 100
}

func main() {
	var x = 4
	fmt.Println("Value: ", x)
	changeValue(x)
	fmt.Println("Try to change value direct: ", x)
	changeValueByPtr(&x)
	fmt.Println("Try to change value by pointer: ", x)
}