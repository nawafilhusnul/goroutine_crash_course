package main

import (
	"fmt"
	"goroutine/challenge_1"
	"goroutine/mutex"
	"sync"
)

func main() {
	var operation = "mutex"
	switch operation {
	case "challenge_1":
		challenge_1.Challenge1()
	case "mutex":
		mutex.Main()
	default:
		words := []string{
			"a",
			"b",
			"c",
			"d",
			"e",
			"f",
		}
		wg := sync.WaitGroup{}

		wg.Add(len(words))
		for i, x := range words {
			go printSomething(&wg, fmt.Sprintf("%d : %s", i, x))
		}

		wg.Wait()
	}

}

func printSomething(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	fmt.Println(s)
}
