package main

import "fmt"

// Goroutines
func main()  {
	// A 'channel'
	ch := make(chan string) // make 是啥: 类似 new

	// Start concurrent routines
	go push("moe", ch) // go 是啥
	go push("larry", ch)
	go push("curly", ch)

	// read 3 results, since our goroutines are concurrent the order isn't guaranteed

	fmt.Println(<-ch, <-ch, <-ch) // <- 操作符是用来做什么的？
}

func push(name string, ch chan string)  {
	msg := "hey, " + name
	ch <- msg
}

// Buffered channels limit the amount of messages it can keep.
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3

// Closing channels
ch <- 1
ch <- 2
ch <- 3

close(ch)

for i := range ch {
	// ...
}

v, ok := <- ch
