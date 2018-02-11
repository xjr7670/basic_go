package main

import "fmt"

func main() {
    arr := []int{9, 7, 5, 3, 1}
    pprint(arr...)
}

func pprint(arr ...int) {
    for _, a := range arr {
        fmt.Println(a)
    }
}
