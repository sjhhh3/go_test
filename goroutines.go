package main

import "fmt"

func f(a string) {
    for i := 0; i<3; i++ {
        fmt.Println(a,":",i)
    }
}

func main() {
    go f("dai")

    go f("zhi")

    go func (msg string) {
        fmt.Println(msg)
    }("yong")

    var input string
    fmt.Scanln(&input)
    fmt.Println("Done")
}
