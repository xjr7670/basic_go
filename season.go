package main

import "fmt"

func main() {
    var m int
    fmt.Printf("Please input a number(1-12):")
    fmt.Scanf("%d", &m)
    switch {
        case m <= 3:
            fmt.Println("Spring")
        case m <= 6:
            fmt.Println("Summer")
        case m <= 9:
            fmt.Println("Autumn")
        case m <= 12:
            fmt.Println("Winter")
        default:
            fmt.Println("not 1 - 12")
    }
}
