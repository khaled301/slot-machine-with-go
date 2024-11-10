package main

// implements formatted I/O
import (
	"fmt"
	"math/rand"
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

func generateSymbols(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(" %s ", symbol)
			if j < len(row)-1 {
				fmt.Printf("|")
			}
		}

		fmt.Println()
	}
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}

		for row := 0; row < rows; row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)

				_, exists := selected[randomIndex]

				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}

	return result
}

// entry point function "main"
func main() {
	// map | key = string, value = uint
	symbols := map[string]uint{
		"A": 3,
		"B": 7,
		"C": 13,
		"D": 21,
	}

	// multipliers := map[string]uint{
	// 	"A": 30,
	// 	"B": 20,
	// 	"C": 10,
	// 	"D": 2,
	// }

	symbolArr := generateSymbols(symbols)

	balance := uint(200) // var balance unit = 200
	getName()

	for balance > 0 {
		bet := getBet(balance)

		if bet == 0 {
			break
		}

		balance -= bet

		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)
	}

	fmt.Printf("You left with, $%d.\n", balance)
}
