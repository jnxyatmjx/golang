package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func readFile(fname string)  []uint32 {

	filefd, err := os.Open(fname)
 
	if err != nil {
		fmt.Printf("File %s Read error",fname)
		return nil
	}
	defer filefd.Close()

	scanner := bufio.NewScanner(filefd)
	scanner.Split(bufio.ScanLines)

	var  ipPairList []uint32
	for scanner.Scan(){
		if  strings.Index(scanner.Text(),"CN|ipv4") == -1 {
			continue
		}
		//scanner.Text() --> apnic|CN|ipv4|223.255.0.0|32768|20100810|allocated
		subv := strings.Split(scanner.Text(),"|")
		ipoffet,_ := strconv.ParseUint(subv[4],10,32) 	//uint64
		ipmin := ip2uint32(subv[3]) 					//uint32
		ipPairList = append(ipPairList,ipmin,ipmin + (uint32)(ipoffet) - 1)
	}
	
	return ipPairList
}

func ip2uint32(ipCIDR string) uint32{
	subuint8 := strings.Split(ipCIDR,".")
	a,_ := strconv.ParseUint(subuint8[0],10,8)
	b,_ := strconv.ParseUint(subuint8[1],10,8)
	c,_ := strconv.ParseUint(subuint8[2],10,8)
	d,_ := strconv.ParseUint(subuint8[3],10,8)
	return (uint32)(a<<24 | b<<16 | c<<8 | d)
	//val >> 24 & 255 , val >> 16 & 255,val >> 8 & 255,val & 255
}

func isAscSorted(ipList []uint32) bool{
	for i,j := 0,1 ; j < len(ipList) ; i++ {
		if(ipList[i] > ipList[j]) {
			return false
		}
		j++
	}
	return true
}

func isLocateIn(iplist []uint32,ipv string) bool{
	ip := ip2uint32(ipv)
	//fmt.Println("this is:",ip)
	l,r := 0,len(iplist)-1
	for l <= r {
		mid := l + (r-l)/2
		if ip > iplist[mid] {
			l = mid + 1
		}else if ip < iplist[mid] {
			r = mid - 1
		}else {
			return true
		}

	}//end for
	
	if (l % 2 == 0) {
		return false
	}else {
		return true
	}
}
