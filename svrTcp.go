package main

import (
        "bufio"
        "fmt"
        "net"
        "strings"
)
func handClose(c net.Conn){
	fmt.Printf("%+v Close! : %v\n",c,c.RemoteAddr().String())
	c.Close()
}
func handConn(c net.Conn){
        defer handClose(c)
        fmt.Printf("Handle new connection:%v\n",c);
        for{
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println("Server read error ",err)
                        return
                }
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("Server receive : ", string(netData))
                replay := c.RemoteAddr().String() + "\n"
                c.Write([]byte(replay))
        }
}

func svrStart(port string){

        l, err := net.Listen("tcp", ":"+ port)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()
        fmt.Println("Tcp server start success")

        for{
                c, err := l.Accept()
                if err == nil {
                        go handConn(c)
                }
        }//end for
}
    
