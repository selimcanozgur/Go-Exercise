package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var balance, err = fileops.ReadToBalance(accountBalanceFile)

	if err != nil {
		fmt.Println(err)
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to banking app")
	fmt.Println("Reach us 24/7",randomdata.PhoneNumber())
	
	optionFunc()

	for {
		fmt.Print("Enter a value: ")
		var value float64
		fmt.Scan(&value)

		switch value {
		case 1:
			fmt.Println("Your balance: ", balance)
		case 2:
			fmt.Print("The amount you want to deposit: ")
			var depositCount float64
			fmt.Scan(&depositCount)
			if depositCount <= 0 {
				fmt.Print("The entered value is incorrect")
				continue
			}
			balance += depositCount
			fileops.WriteToBalance(balance, accountBalanceFile)
		case 3:
			fmt.Print("The amount you want to withdraw:: ")
			var withdrawCount float64
			fmt.Scan(&withdrawCount)
			if balance < withdrawCount {
				fmt.Print("Insufficient funds")
				continue
			}
			balance -= withdrawCount
			fileops.WriteToBalance(balance, accountBalanceFile)
		default:
			fmt.Println("Have a nice day")
			fmt.Println("Thank you for choosing us")
			return
		}

	}

}