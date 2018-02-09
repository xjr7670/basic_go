package main

import "fmt"

func main() {
    for i := 1; i <= 25; i++ {
        str := "GGGGGGGGGGGGGGGGGGGGGGGGG"
        fmt.Printf("%s\n", str[:i])
    }
}
