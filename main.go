package main

// implements formatted I/O
import "fmt"

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

// entry point function "main"
func main() {
	name := getName()
	fmt.Println(name)
}
