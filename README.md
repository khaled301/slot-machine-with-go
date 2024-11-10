
# Slot Machine Game

This is a simple console-based Slot Machine Game written in Go. The project demonstrates basic concepts in Go, such as formatted I/O, functions, loops, and maps, and is a great starting point for beginners exploring Go.

## Project Overview

In this game, the user can place bets, spin the slot machine, and potentially win based on symbol alignment. The game continues until the player runs out of balance or chooses to quit.

### Key Features
- **Betting System**: Users can place bets and the balance is updated accordingly.
- **Symbol Generation**: A random set of symbols is generated for each spin.
- **Winning Conditions**: Players win if symbols align in a row, with different multipliers for each symbol.
- **Interactive Console Interface**: The game prompts the player for input, such as their name and bet amount.

### Download and Install GO
https://go.dev/dl/

### Project Files
- `main.go`: Contains the core logic for the game, including symbol generation and checking win conditions.
- `utils.go`: Contains helper functions for user input, such as getting the player’s name and bet amount.
- `spin.go`: Contains helper functions for spin simulation.

### Usage

1. Clone the repository.
2. Run the project using:
   ```bash
    go mod init <your-project-name> && go run .

    #or
    go run main.go utils.go spin.go
