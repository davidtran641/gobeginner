package main

import (
"fmt"
"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return strconv.Itoa(int(ip[0])) + "." + strconv.Itoa(int(ip[1])) + "." + strconv.Itoa(int(ip[2])) + "." + strconv.Itoa(int(ip[3]))
}

func main() {
	hosts := map[string]IPAddr {
		"lookback": {127,0,0,1},
		"google": {8,8,8,8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}