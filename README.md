# Go Slot Machine Game

A simple slot machine game built with Go. This project allows users to spin reels of symbols and win based on matching symbols across rows. The game helps enhance programming skills and allows players to learn Go in a fun way.

## How to Play

1. Start with a balance of 100 units.
2. Place a bet for each spin.
3. The machine will display a 3x3 grid of symbols.
4. If a row has matching symbols, you win based on the symbol's multiplier.
5. Continue playing as long as you have a balance.

## Project Files

- `main.go`: The main file containing the game's core logic.
- `spin.go`: Contains functions for generating spin results and printing the grid.
- `utils.go`: Contains utility functions for user input and random number generation.

## Game Logic

- Symbols: Each symbol has a different chance of appearing and a different multiplier:

  - **A**: 4 (20x multiplier)
  - **B**: 7 (10x multiplier)
  - **C**: 12 (5x multiplier)
  - **D**: 20 (2x multiplier)

- Winning Lines: The game checks each row in the 3x3 grid. If all symbols in a row match, you win that row's multiplier.

## Important Note

> **Warning**: Gambling can be addictive. This project is created solely for learning purposes and should not be used as a real gambling tool. Please play responsibly and remember that gambling is not encouraged.

## Run the Project

To run the game:

1. Clone the repository.
2. In the terminal, navigate to the project folder and run:
   ```bash
   go run .
   ```

Enjoy the game and good luck!
