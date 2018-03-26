package main

import (
	"fmt"
	"time"
)

func main() {
	basic()

	sender()

	channelClose()

	goroutineChannel()

	//loop()

	selectChannel()
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

func goroutineChannel() {
	ch := make(chan int, 20)

	go recieve("1st goroutine", ch)
	go recieve("2nd goroutine", ch)
	go recieve("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	time.Sleep(3 + time.Second)
}

func recieve(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func loop() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	for i := range ch {
		fmt.Println(i)
	}
}

func selectChannel() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	// ch1 -> ch2
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()
	// ch2 -> ch3
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("recieved", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}
