package main

import (
    "fmt"
)

func main() {
    hello := "Hello, cliched world!"
    fmt.Printf("%s\n", hello) // must use printf to interpolate?
    fmt.Printf("Variable hello is a %T.\n", hello)
    fmt.Println("Counting down from 10.")
    for i := 10; i > 0; i-- {
        if i != 1 {
            fmt.Printf("%v... ", i)
        } else {
            fmt.Printf("%v!", i)
        }
    }
}