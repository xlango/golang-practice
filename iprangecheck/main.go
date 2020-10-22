package main

import (
	"bytes"
	"fmt"
	"net"
)

//var (
//	ip1 = net.ParseIP("216.14.49.184")
//	ip2 = net.ParseIP("216.14.49.191")
//)

func check(ip string,beginIp string,endIp string) bool {
	ip1 := net.ParseIP(beginIp)
	ip2 := net.ParseIP(endIp)
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
		fmt.Printf("%v is between %v and %v\n", trial, ip1, ip2)
		return true
	}
	fmt.Printf("%v is NOT between %v and %v\n", trial, ip1, ip2)
	return false
}

func main() {
	check("10.34.11.90","10.34.11.1/24","10.34.11.225")
}
