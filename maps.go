package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["k1"]=1
    m["k2"]=2

    fmt.Println("map:",m)

    v1 := m["k1"]
    fmt.Println("v1:",v1)

    fmt.Println("len:",len(m))

    a,prs := m["k2"]
    fmt.Println("v2:",a)
    fmt.Println("prs:",prs)

    n := map[string]int{"foo":3,"bar":4}
    fmt.Println("map:",n)
}

