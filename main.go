package main

// implements formatted I/O
import (
	"fmt"
)

func generateSymbols(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func getWinResult(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]

		for _, symbol := range row {
			if checkSymbol != symbol {
				win = false
				break
			}
		}

		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}

	return lines
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

	multipliers := map[string]uint{
		"A": 30,
		"B": 20,
		"C": 10,
		"D": 2,
	}

	symbolArr := generateSymbols(symbols)

	balance := uint(200) // var balance unit = 200
	GetName()            // imported from the utils.go

	for balance > 0 {
		bet := GetBet(balance) // imported from the utils.go

		if bet == 0 {
			break
		}

		balance -= bet

		spin := GetSpin(symbolArr, 3, 3) // imported from the spin.go
		PrintSpin(spin)                  // imported from the spin.go

		winLines := getWinResult(spin, multipliers)

		for i, line := range winLines {
			win := line * bet
			balance += win
			if line > 0 {
				fmt.Printf("You won $%d, (%dx) on Line %d\n", win, line, i+1)
			}
		}
	}

	fmt.Printf("You left with, $%d.\n", balance)
}
