package main

import (
    "fmt"
)

// underscores or camelcase?

func main() {
    fmt.Println("Pointers!")

    somenumber := 9000
    pointynumber := &somenumber
    fmt.Printf("somenumber's literal value is %v.  pointynumber's literal value is %v.\n", somenumber, pointynumber)
    fmt.Printf("pointynumber's \"through\" value is %v.\n", *pointynumber)
    *pointynumber += 1
    fmt.Printf("the numbers' new values are %v and %v.\n", somenumber, *pointynumber)

    fmt.Println("")

    fmt.Println("Slices!")
    intarr := []int{} // slices vs arrays?  I think I like slices
    intarr = append(intarr, 1337) // kind of annoying... not a function of a slice?
    fmt.Println(intarr)

    for i := 0; i < 500; i++ {
        intarr = append(intarr, i * 50)
    }

    fmt.Printf("intarr is %v integers long.  It should be 501 integers long.\n", len(intarr))

    fmt.Println("")
}