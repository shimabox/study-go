package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	list := []string{}
	for _, v := range ip {
		// fmt.Printf("%T", v) => uint8
		// vはuint8なのでstringに変換
		// int(v)でintにしてからstrconv.Itoa()してappend
		list = append(list, strconv.Itoa(int(v)))
	}
	// . で結合して返す
	return fmt.Sprint(strings.Join(list, "."))
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
