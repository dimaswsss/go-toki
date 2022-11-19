package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Scan(&a , &b)
    fmt.Println("masing-masing", a / b)
    fmt.Println("bersisa", a % b)
}