package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var text string
    scanner.Scan()
    text = scanner.Text()
    fmt.Println(text)
}