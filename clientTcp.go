package main

import (
	"net"
	"fmt"
	"time"
	"bufio"
	"os"
)

func receiveCon(cxnn net.Conn) {
	for {
		netData, err := bufio.NewReader(cxnn).ReadString('\n')
		if err != nil {
			fmt.Println("Client Read Error: ",err)
			break
		}
		fmt.Printf("Replay %s\n",netData)
	}
}


func responseCon(cxnn net.Conn,val []byte){

	_,err := cxnn.Write(val)
	if err != nil{
		fmt.Println("Client hearbeat error:",err)
		return
	}
}

func scaninput(inpu chan string) {

	for{
		input,_ := bufio.NewReader(os.Stdin).ReadString('\n')
		if len(input) > 1 {
			fmt.Println("size:",len(input));
			inpu <- input
		}
	}
}

func clientStart(host,port string){
	dstsvr,cerror := net.Dial("tcp",host+":"+port)
	if cerror != nil {
		panic(cerror)
	}
	
	defer dstsvr.Close()
	fmt.Printf("LocalAddr %s %T\n",dstsvr.LocalAddr().String(),dstsvr)	
	go receiveCon(dstsvr)
	defer dstsvr.Close()

	input := make(chan string )
	go scaninput(input)
	// var inrer chan error
	for {
		select {
			case  vax := <-input :
				fmt.Printf("inp :%s %v\n",vax,time.Now())
			case <- time.After(60*time.Second):
				//responseCon(dstsvr,makePack0x1())
				fmt.Println("this is heart beat")
		}
		
	}//end for

}