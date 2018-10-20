package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	fmt.Println("1")
	channel <- "h1"
	fmt.Println("2")
	channel <- "w2"

	fmt.Println("3")
	select {

	case channel <- "c3":
		fmt.Println("ok")
	default:
		fmt.Println("channel is full !")
	}

	fmt.Println("...")
	msg1 := <-channel
	fmt.Println(msg1)
}
