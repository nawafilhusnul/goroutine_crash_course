package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// variables
var seatingCapacities = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	//	seed our random number generator
	rand.New(rand.NewSource(time.Now().UnixNano())).Int63()

	//	print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	//	create channels if we need any
	clientChan := make(chan string, seatingCapacities)
	doneChannel := make(chan bool)

	//	create the barbershop data structure
	shop := &Barbershop{
		ShopCapacity:    seatingCapacities,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChannel,
		Open:            true,
	}

	color.Green("The shop is open for the day!!!")

	//	add barbers
	shop.addBarber("Elias")
	shop.addBarber("Frank")
	shop.addBarber("Gerald")
	shop.addBarber("Hamilton")
	shop.addBarber("John")

	//	start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForTheDay()
		closed <- true
	}()

	//	add clients
	i := 1

	go func() {
		for {
			//get a random number with average arrival rate
			randomMils := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMils)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	//	block until the barbershop is closed
	<-closed
}
