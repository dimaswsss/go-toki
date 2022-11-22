package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h, i uint8
    fmt.Scan(&a, &b, &c)
    fmt.Scan(&d, &e, &f)
    fmt.Scan(&g, &h, &i)
    fmt.Println(a, d, g)
    fmt.Println(b, e, h)
    fmt.Println(c, f, i)
}