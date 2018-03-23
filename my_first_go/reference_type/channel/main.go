package main

import (
	"fmt"
)

func main() {
	basic()

	sender()

	channelClose()
}

func basic() {
	ch := make(chan int, 10)
	// send
	ch <- 5
	// recieve
	i := <-ch
	fmt.Println(i)
}

func sender() {
	ch := make(chan int)

	go reciever(ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
}

func reciever(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func channelClose() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch2 := make(chan int, 2)
	ch2 <- 100
	ch2 <- 200
	close(ch2)
	var (
		i  int
		ok bool
	)
	i, ok = <-ch2
	fmt.Println(i)
	fmt.Println(ok)
	i, ok = <-ch2
	fmt.Println(i)
	fmt.Println(ok)
	i, ok = <-ch2
	fmt.Println(i)
	fmt.Println(ok)
}
