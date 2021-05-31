package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return string(strconv.Itoa(int(ip[0]))) + "." + string(strconv.Itoa(int(ip[1]))) + "." + string(strconv.Itoa(int(ip[2]))) + "." + string(strconv.Itoa(int(ip[3])))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {

		fmt.Printf("%v: %v\n", name, ip)
	}
}
