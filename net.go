package main
import "fmt"
import "net"

func main() {

    nets, _ := net.Interfaces()
    fmt.Println(nets)
    //[{1 65536 lo  up|loopback} {2 1500 eth0 52:54:00:0f:21:7d up|broadcast|multicast}]
    // MTU 最大传输单元  网络接口  支持状态


    addr, _ := net.InterfaceAddrs()
    fmt.Println(addr)
    //[127.0.0.1/8 10.30.66.168/32 10.30.66.23/24 ::1/128 fe80::5054:ff:fe0f:217d/64]
    //本地地址   VIP   ipv4  ipv6


    it, _ := net.LookupAddr("127.0.0.1")
    fmt.Println(it)
    //[localhost localhost.localdomain. localhost4 localhost4.localdomain4.]
    // 根据ip查找映射列表


    cname, _ := net.LookupCNAME("www.baidu.com")  
    fmt.Println(cname) 
    //www.a.shifen.com,查找规范的dns主机名字 


    ip, _ := net.LookupIP("www.daizhiyong.top")
    fmt.Println(ip)
    // [192.241.168.109] 域名解析
}
