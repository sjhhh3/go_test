package main

import (
        "fmt"
        "os/exec"
        "runtime"
        //"string"
)

var p = fmt.Println

func main() {
    if runtime.GOOS == "linux" {
        s, err := exec.Command("/usr/sbin/dmidecode", "--string", "system-serial-number").Output()        
        if err != nil {
            p(err)
        }
        p(string(s))
    }
}
