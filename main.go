package main
import (
		"fmt"
		"os"
		"runtime"
		"strings"
		//"strconv"
	)



func main(){

	ipList := readFile("./cnIp.txt")
	fmt.Printf(" %T *** Len(%d) Cap(%d)\n",ipList,len(ipList),cap(ipList))
	if !isAscSorted(ipList){
		fmt.Println("Ip list error ")
	}
	bol := isLocateIn(ipList,os.Args[1])
	fmt.Println(bol)
	return 
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(os.Args);
	if len(os.Args) < 4 {
		fmt.Println("Arguments error -> c|s host port")
		return
	}
	fmt.Printf("host:%v(%T) port:%+v(%T)\n",os.Args[1],os.Args[1],os.Args[2],os.Args[2])

	if strings.Compare(os.Args[1],"C") == 0 || strings.Compare(os.Args[1],"c") == 0 {
		host := os.Args[2]
		port := os.Args[3]
		clientStart(host,port)
	}else{
		port := os.Args[3]
		svrStart(port)
	}

}
