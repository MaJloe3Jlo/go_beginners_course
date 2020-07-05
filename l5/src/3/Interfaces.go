package main

import (
	"fmt"
)

type Dog struct {
	Voice string
}
type Cat struct {
	Voice string
}

func (c Cat) Say() {
	fmt.Println(c.Voice)
}
func (d Dog) Say() {
	fmt.Println(d.Voice)
}

type Sayer interface {
	Say()
}

func main() {
	cat := Cat{Voice: "Meow"}
	dog := Dog{Voice: "Woof"}
	pets := []Sayer{cat, dog}
	pets[0].Say()
	
	var pet Sayer = Cat{Voice: "Another meow"}
	pet.Say()
}
