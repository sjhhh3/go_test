package main 

import (
    "fmt"
    "net"
    "os"
)

func main() {
    addr , err := net.ResolveIPAddr("ip","www.daizhiyong.top")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(addr.IP)
}
