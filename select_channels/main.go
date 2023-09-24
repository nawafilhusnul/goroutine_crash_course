package main

import (
	"github.com/fatih/color"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This from server 2"
	}
}

func main() {
	color.Cyan("SELECT WITH CHANNELS!")
	color.Cyan("---------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			color.Green("Case 1: %s", s1)
		case s2 := <-channel1:
			color.Green("Case 2: %s", s2)
		case s3 := <-channel2:
			color.Yellow("Case 3: %s", s3)
		case s4 := <-channel2:
			color.Yellow("Case 4: %s", s4)
			//default:
			//	//	avoiding deadlock
			//
		}
	}
}
