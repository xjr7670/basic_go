package main

import "fmt"
import "math"

var sqrt float64

func main() {
    var input float64
    fmt.Println("input a number: ")
    fmt.Scanf("%d", &input)
    sqrt, err = MySqrt(input)
    if err != nil {
        fmt.Println(sqrt)
    }
    sqrt, err = MySqrt1(input)
    if err != nil {
        fmt.Println(sqrt)
    }
}

func MySqrt(input float64) (float64) {
    if input < 0 {
        return 1.0
    }
    return math.Sqrt(input), nil
}

func MySqrt1(input float64) (sqrt float64) {
    if input < 0 {
        return 1.0
    }
    sqrt = math.Sqrt(input)
    return sqrt, nil
}
