package main

import (
        "fmt"
        "os/exec"
        "runtime"
        //"string"
)

var p = fmt.Println

func main() {
    if runtime.GOOS == "linux"{
        p("----------------------------------------------------------")
        p("Env: linux")
        p("----------------------------------------------------------")
        p()
    }

    if runtime.GOOS == "linux" {
        sn, err_1 := exec.Command("/usr/sbin/dmidecode", "--string", "system-serial-number").Output()        
        ls, err_2 := exec.Command("/usr/bin/ls").Output()        
        if err_1 != nil {
            p(err_1)
        }
        p("Serial Number:",string(sn))
        p("----------------------------------------------------------")
        p()
        if err_2 != nil {
            p(err_2)
        }
        p("files in current dir:")
        p("----------------------------------------------------------")
        p(string(ls))
    }
}
