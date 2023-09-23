package main

import (
	"fmt"
	"sync"
)

func main() {

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

func printSomething(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	fmt.Println(s)
}
