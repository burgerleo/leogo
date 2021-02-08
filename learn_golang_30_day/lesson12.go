package main

import "fmt"

func main() {

	messages := make(chan string)

	go getName("leo", messages)
	go getName("leo2", messages)

	msg1 := <-messages
	msg2 := <-messages

	fmt.Println(msg1)
	fmt.Println(msg2)
}

func getName(name string, c chan string) {
	c <- name + "-gogogo"
}
