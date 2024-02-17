package main

import "fmt"

const (
	GROUP_IP     = "234.0.0.1"
	INTERFACE_IP = "192.168.1.2"
	PORT         = 22000
)

func main() {
	sendMessage("multicast message")

	go server()

	var input string
	fmt.Scan(&input)
}
