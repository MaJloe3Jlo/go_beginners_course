package main

import "fmt"

const (
	womanRetireAge = 60 + iota*5
	manRetireAge
)

func main() {
	var (
		name = "Dima"
		age  int
	)
	fmt.Scan(&age)
	retiteAge := manRetireAge - age
	fmt.Printf("%s will retire in %v years", name, retiteAge)
	Age()
}

func Age() {
	fmt.Println(manRetireAge)
}
