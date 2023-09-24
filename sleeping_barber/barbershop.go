package main

import (
	"github.com/fatih/color"
	"time"
)

type Barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (b *Barbershop) addBarber(barber string) {
	b.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			//	if there are no clients, the barber goes to sleep.
			if len(b.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-b.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}

				//	cut hair
				b.cutHair(barber, client)
			} else {
				//	shop is closed, so send the barber home and close the goroutine
				b.sendBarberHome(barber)
				return
			}

		}
	}()
}

func (b *Barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(b.HairCutDuration)
	color.Green("%s is finished cutting %s' hair.", barber, client)
}

func (b *Barbershop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	b.BarbersDoneChan <- true
}

func (b *Barbershop) closeShopForTheDay() {
	color.Cyan("Closing shop for the day!!!")

	close(b.ClientsChan)
	b.Open = false

	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChan
	}

	close(b.BarbersDoneChan)

	color.Red("---------------------------------------------------------------------")
	color.Red("The Barbershop is now closed for the day!! And everyone has gone home")
}

func (b *Barbershop) addClient(client string) {
	//	print out the message
	color.Green("**** %s arrives!", client)

	if b.Open {
		select {
		case b.ClientsChan <- client:
			color.Yellow("\t%s takes a seat in the waiting room!!", client)
		default:
			color.Red("The waiting room is full so the %s leaves!!", client)
		}
	} else {
		color.Red("The shop is already closed, %s leaves!", client)
	}
}
