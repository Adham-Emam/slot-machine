package main

import "fmt"

func GetName() string {
    var name string

    fmt.Println("Welcome to the Slot Machine")
    fmt.Println("Enter your name: ")
    _, err := fmt.Scanln(&name)

    if err != nil {
        fmt.Println("Invalid input")
        return ""
    }

    fmt.Printf("Welcome Back %s, Let's play!\n", name)
    return name
}

func GetBet(balance uint) uint {
    var bet uint
    for true {
        fmt.Printf("Enter your bet or type 0 to quit (Balance:%d): ", balance)
        fmt.Scan(&bet)

        if bet > balance {
            fmt.Println("You don't have enough balance")
        } else  {
            break
        }
    }
    return bet
}
