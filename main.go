package main

import (
	"fmt"
)



func checkWin(spin [][]string, multipliers map[string]uint) []uint {
    lines := []uint{}

    for _, row := range spin {
        win := true
        checkSymbol := row[0]
        for _, symbol := range row[1:] {
            if symbol != checkSymbol {
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

func main() {
    symbols := map[string]uint {
        "A": 4,
        "B": 7,
        "C": 12,
        "D": 20,
    }

    symbolArr := GenerateSymbolArray(symbols)

    multipliers := map[string]uint {
            "A": 20,
            "B": 10,
            "C": 5,
            "D": 2,
        }

    var balance uint = 100
    GetName()

    for balance > 0 {
        bet := GetBet(balance)
        if bet == 0 {
            break
        }
        balance -= bet
        spin := GetSpin(symbolArr, 3, 3)
        PrintSpin(spin)
        winningLines := checkWin(spin, multipliers)
        for i, multi := range winningLines {
            win := multi * bet
            balance += win
            if multi > 0 {
                fmt.Printf("You won %d$ (%dx) on line #%d\n", win, multi, i + 1)
            }
        }
    }
    

    fmt.Printf("You are left with %d$\n", int(balance))
}