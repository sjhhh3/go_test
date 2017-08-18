package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{"Bob",20})
    fmt.Println(person{name:"Alice",age:21})
    fmt.Println(person{name:"dai"})
    fmt.Println(&person{name:"zhi",age:22})
    s := person{name:"yong",age:23}
    fmt.Println(s.name)

    ss := &s
    fmt.Println(ss.age)
}
