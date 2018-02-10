package main

import "fmt"

var s, t, g int

func main() {
    var a, b int
    fmt.Println("input a and b:")
    fmt.Scanf("%d %d", &a, &b)
    s, t, g = get_result1(a, b)
    fmt.Println(s, t, g)
    s, t, g = get_result2(a, b)
    fmt.Println(s, t, g)
}

func get_result1(a int, b int) (int, int, int) {
    return a + b, a * b, a - b
}

func get_result2(a int, b int) (s int, t int, g int) {
    s = a + b
    t = a * b
    g = a - b
    return
}
