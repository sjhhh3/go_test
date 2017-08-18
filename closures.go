package main

import "fmt"

func inSqe() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main(){
    nextInt := inSqe()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInt := inSqe()
    fmt.Println(newInt())
}
