package main

import "fmt"

func main() {
    messages := make(chan string, 4)

    messages <- "bufferd"
    messages <- "channel"
    messages <- "1"
    messages <- "2"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
