package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Check if there are enough command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <number>")
        return
    }

    // Convert the first argument to an integer
    numTx, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Use the converted integer
    fmt.Printf("The number you entered is: %d\n", numTx)
}
