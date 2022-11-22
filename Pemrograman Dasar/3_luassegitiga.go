package main

import "fmt"

func main() {
    var a, t uint16
    fmt.Scan(&a, &t)
    fmt.Printf("%.2f\n", float32(a * t)/2)
}