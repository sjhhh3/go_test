package main

import "encoding/json"
import "fmt"
//import "os"
import "reflect"

type Response1 struct {
    Page int
    Fruits []string
}


// json自定义输出的json数据键
// json解析的字段首字母必须大写，否则无法正常输出
type Response2 struct {
    Page int          `json:"page_json"`
    Fruits []string   `json:"fruits_json"`
    meet string       
}

func main(){
    bolB,_ := json.Marshal(true)
    fmt.Println(string(bolB))
    fmt.Println(reflect.ValueOf(1))
    fmt.Println(reflect.ValueOf(bolB))

    intB,_ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB,_ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    slcD := []string{"a","b","c"}
    slcB,_ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"a":5,"b":6,"c":7}
    mapB,_ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res1D := Response1{
        Page: 1,
        Fruits: []string{"a","b","c"},
    }
    res1B,_ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := Response2{
        Page: 1,
        Fruits: []string{"a","b","c"},
        meet: "pig",
    }
    res2B,_ := json.Marshal(res2D)
    fmt.Println(string(res2B))


}
