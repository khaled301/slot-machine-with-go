package main

// implements formatted I/O
import (
	"fmt"
	"strconv"
)

func getName() string {
	name := ""

	fmt.Println("Welcome to Slot Machine Game")
	fmt.Printf("Please enter your name: ")

	_, err := fmt.Scanln(&name)

	if err != nil {
		return ""
	}

	fmt.Printf("Welcome %s, let's play\n", name)

	return name
}

func getBet(balance uint) uint {
	var bet uint
	var input string
	for true {
		fmt.Printf("Enter your bet, or O to quit (balance = $%d): ", balance)

		fmt.Scan(&input)
		inputBet, err := strconv.ParseUint(input, 10, 64) // converts string to uint64
		bet = uint(inputBet)

		if err != nil {
			fmt.Println("Invalid input. Please enter a positive number")
			continue
		} else if bet > balance {
			fmt.Println("You don't have enough balance")
		} else {
			break
		}
	}

	return bet
}

// entry point function "main"
func main() {
	balance := uint(200) // var balance unit = 200
	getName()

	for balance > 0 {
		bet := getBet(balance)

		if bet == 0 {
			break
		}

		balance -= bet
	}

	fmt.Printf("You left with, $%d.\n", balance)
}
