package main

 

import (

       "fmt"

       "encoding/json"

)

 

type User struct {

       Username string

        Password string

        FriendName []string

}

 

func main() {

 

         user:=User{}

         user.Username="Tom"

         user.Password="123456"

         user.FriendName=[]string{"Li","Fei"}

 

       //将struct转成json字符串，注意：结构体中的字段首字母必须大写，否则无法解析

       if userJSON,err:=json.Marshal(user);err==nil{

             fmt.Println(userJSON)   //打印结果：{"Username":"Tom","Password":"123456","FriendName":["Li","Fei"]}

        }

 

        //将slice转成json字符串

        arr:=[]string{"Apple","Orange","Banana"}

       if arrJSON,err:=json.Marshal(arr);err==nil{

               fmt.Println(string(arrJSON))  //打印结果：["Apple","Orange","Banana"]

        }

 

        //将map转成json字符串

         m:=map[string]string{"浙江":"杭州","湖南":"长沙"}

         if mJSON,err:=json.Marshal(m);err==nil{

                fmt.Println(string(mJSON))   //打印结果：{"浙江":"杭州","湖南":"长沙"}

          }

 

         //json转成struct

         jsonStr:=`{"Username":"Tom","Password":"123456","FriendName":["Li","Fei"]}`

         var userJSON User

         if err:=json.Unmarshal([]byte(jsonStr),&userJSON);err==nil{

                   fmt.Println(userJSON)   //打印结果：{Tom 123456 [Li Fei]}

         }

 

         //json转成slice

         jsonFruit:=`["Apple","Orange","Banana"]`

         var arrFruit  []string

         if err:=json.Unmarshal([]byte(jsonFruit),&arrFruit);err==nil{

               fmt.Println(arrFruit)   //打印结果：[Apple Orange Banana]

          }

 

         //json转成map

         jsonCity:=`{"浙江":"杭州","湖南":"长沙"}`

         var mapCity map[string]string

         if err:=json.Unmarshal([]byte(jsonCity),&mapCity);err==nil{

                  fmt.Println(mapCity)  //打印结果: map[浙江:杭州 湖南:长沙]

         }

}
