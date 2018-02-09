package main

import "fmt"

func main() {
    for r := 1; r <= 10; r++ {
        for c := 1; c <= 20; c++  {
            fmt.Printf("%s", "*")
        }
        fmt.Println()
    }
}
