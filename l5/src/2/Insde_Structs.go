package main

import (
	"fmt"
)

type Animal struct {
	Class string
}

func (a *Animal) Talk(voice string) {
	fmt.Println(voice)
}

type Cat struct {
	Animal Animal
	Type   string
}

type Dog struct {
	Animal
	Type string
}

func main() {
	fmt.Println("Hello, playground")
	cat := Cat{}	
	cat.Animal.Talk("meow")
	// cat.Talk("another meow")
	dog := Dog{}
	dog.Talk("woof")
	dog.Animal.Talk("another woof")
}

