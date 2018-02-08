package main

import (
    "fmt"
)

func main() {
    var i1 = 5
    var intP *int
    fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
    intP = &i1
    fmt.Printf("%p\n", intP)
    fmt.Printf("%d\n", *intP)
}
