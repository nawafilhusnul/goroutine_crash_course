package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			// do something here
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	//create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput

		//	wait for the response
		response := <-pong

		color.Green("Response: %s", response)
	}

	color.Cyan("All done, closing the channels.")
	close(pong)
	close(ping)
}
