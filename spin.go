package main

import (
	"fmt"
	"math/rand"
)

func GenerateSymbolArray(symbols map[string]uint) []string{
    symbolArr := []string{}
    for symbol, count := range symbols {
        for i := 0; i < int(count); i++ {
            symbolArr = append(symbolArr, symbol)
        }
    }
    return symbolArr
}

func getRandomNumber(min int, max int) int {
    return rand.Intn(max - min + 1) + min
}

func GetSpin(reel []string, rows int, cols int) [][]string {
    results := [][]string{}
    
    for i :=0; i < rows; i++ {
        results = append(results, []string{ })
    }

    for col := 0; col < cols; col++ {
        selected := map[int]bool{}
        for row := 0; row < rows; row++ {
            for true {
                randomIndex := getRandomNumber(0, len(reel) - 1)
                _, exists := selected[randomIndex] 
                if !exists {
                    selected[randomIndex] = true
                    results[row] = append(results[row], reel[randomIndex])
                    break
                }
            }
        }
    }
    return results
}

func PrintSpin(spin [][]string) {
    for _, row := range spin {
        for j, symbol := range row {
            print(symbol)
            if j != len(row) - 1 {
                fmt.Print(" | ")
            }
        }
        fmt.Println()
    }
}
