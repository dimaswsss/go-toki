package main

import "fmt"

func main() {
    var N int32

    fmt.Scan(&N)

    if N > 0 && N <= 9  {
        fmt.Print("satuan")
    } else if N >= 10 && N <= 99 {
        fmt.Print("puluhan")
    } else if N >= 100 && N <=999 {
        fmt.Print("ratusan")
    } else if N >= 1000 && N <=9999 {
        fmt.Print("ribuan")
    } else if N >= 10000 && N <= 99999 {
        fmt.Print("puluhribuan")
    } else {
        fmt.Print("none")
    }
}