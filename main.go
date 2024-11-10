package main

// implements formatted I/O
import (
	"fmt"
	"math/rand"
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

		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)

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
