package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"sync"
	"time"
)

// Philosopher is a struct which stores information about a philosopher.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is list of all philosophers.
var philosophers = []Philosopher{
	{
		name:      "Plato",
		rightFork: 0,
		leftFork:  4,
	},
	{
		name:      "Socrates",
		rightFork: 1,
		leftFork:  0,
	},
	{
		name:      "Aristotle",
		rightFork: 2,
		leftFork:  1,
	},
	{
		name:      "Pascal",
		rightFork: 3,
		leftFork:  2,
	},
	{
		name:      "Locke",
		rightFork: 4,
		leftFork:  3,
	},
}

// define some activities variables
var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	// print out a welcome message
	color.Cyan("Dining Philosophers Problem")
	color.Cyan("---------------------------")
	color.Yellow("The table is empty.")

	time.Sleep(sleepTime)

	//	start the meal
	dine()

	//	print out finished message
	color.Yellow("The table is empty.")

	time.Sleep(sleepTime)
	color.Green("Order finished: %s.\n", strings.Join(orderFinished, ", "))
}

func dine() {
	//eatTime = 0 * time.Second
	//sleepTime = 0 * time.Second
	//thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}

	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//	fork is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//	start the meal
	for i := 0; i < len(philosophers); i++ {
		//	fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//	seat the philosopher at the table
	color.Blue("%s is seated at the table.\n", philosopher.name)
	seated.Done()
	seated.Wait()
	//	eat three times

	for i := hunger; i > 0; i-- {
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.x\n", philosopher.name)

			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)

		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)

			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating. \n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking. \n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
