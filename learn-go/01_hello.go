package main

import "fmt"


func main() {
	message := greetMe("world")
	fmt.Println(message)
}

func greetMe(name string) (string) {
	msg := "hello, "
	return msg + name + "!"
}
