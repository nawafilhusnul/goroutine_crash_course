package mutex

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func Main() {
	//	variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	//	print out starting value
	fmt.Printf("initial ccount balance: %d.00\n", bankBalance)

	//	define weekly revenue
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 500,
		},
		{
			Source: "Gifts",
			Amount: 100,
		},
		{
			Source: "Part Time Job",
			Amount: 300,
		},
		{
			Source: "Investments",
			Amount: 200,
		},
	}

	wg.Add(len(incomes))

	//	loop through 52 weeks and print out how much is made; keep a running total
	for idx, income := range incomes {
		go func(i int, incomeDetail Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += incomeDetail.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned %d.00 from %s\n", week, incomeDetail.Amount, incomeDetail.Source)
			}
		}(idx, income)
	}

	wg.Wait()

	//	print out final balance
	fmt.Printf("Final bank balance is %d.00\n", bankBalance)

	//	command: go run -race .
}
