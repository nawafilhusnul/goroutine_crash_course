package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMessage(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func Challenge1() {
	msg = "Hello World"

	wg.Add(1)
	go updateMessage(&wg, "Hello Universe !!!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage(&wg, "Hello Cosmos !!!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage(&wg, "Hello World !!!")
	wg.Wait()
	printMessage()

}
