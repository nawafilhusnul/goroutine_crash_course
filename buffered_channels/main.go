package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch

		fmt.Println("Got", i, "from channel")

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		color.Yellow("sending %d to channel...", i)
		ch <- i
		color.Green("sent %d to channel!", i)
	}

	fmt.Println("DONE!!!")
	close(ch)
}
