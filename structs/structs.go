package main

import (
    "fmt"
)

type Human struct {
    height int
    weight int
    age int
}

func main() {
    somebody := Human{1, 2, 3}
    fmt.Println(somebody)
    fmt.Printf("somebody's age is %v.\n", somebody.age)
    somebodyage := &somebody.age
    fmt.Println("Happy birthday, somebody.")
    *somebodyage++
    fmt.Printf("somebody's age is now %v.\n", somebody.age)
}