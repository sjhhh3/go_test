package main

import "fmt"

func val(ival int){
    ival = 0
}

func ptr(iptr *int){
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("init:",i)
    val(i)
    fmt.Println("val:",i)
    ptr(&i)
    fmt.Println("ptr:",i)
}
