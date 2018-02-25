package main

import "fmt"

func main() {
    p := fabnacci()
    for i := 0; i < 10; i++ {
        fmt.Printf("%d -> %d\n", i, p())
    }
}

func fabnacci() func() int {
    var y int = 1
    return func() int {
        x = x + y
        y = x
        return x
    }
}
