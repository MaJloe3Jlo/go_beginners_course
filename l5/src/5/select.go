package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func(c1 *chan string) {
		for {
			*c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}(&c1)
	go func(c2 *chan string) {
		for {
			*c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}(&c2)
	//for {
	//select {
	//case msg1 := <-c1:
	//	fmt.Println(msg1)
	//case msg2 := <-c2:
	//	fmt.Println(msg2)
	//}
	//}

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Message 1", msg1)
		case msg2 := <-c2:
			fmt.Println("Message 2", msg2)
		case <-time.After(time.Second):
			fmt.Println("timeout")
			return
		default:
			fmt.Println("nothing ready")
			time.Sleep(10 * time.Second)
		}
	}

}
