package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        switch {
            case i % 3 == 0:
                fmt.Println("Fizz")
                fallthrough
            case i % 5 == 0:
                fmt.Println("Buzz")
                fallthrough
            case i % 15 == 0:
                fmt.Println("FizzBuzz")
            default:
                fmt.Println(i)
        }
    }
}
