package main

import "fmt"

func main() {
    message := make(chan string)
    
    //键盘录入信息通过通道传给msg
    go func() {
        // 键盘录入存储在input中
        var input string
        fmt.Scanln(&input) 
        message <- input
    }()

    msg := <- message
    fmt.Println(msg)
}
