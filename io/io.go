package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)
    fmt.Printf("You entered \"%v\".  It is %v characters long.\n", text, len(text))
}