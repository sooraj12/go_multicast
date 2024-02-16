package main

import (
	"fmt"
	"log"
	"net"
)

func sendMessage(msg string) {
	// create socket
	groupAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", GROUP_IP, PORT))
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := net.DialUDP("udp4", nil, groupAddr)
	if err != nil {
		log.Fatal((err))
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Message sent to %s:%d: %s\n", GROUP_IP, PORT, msg)
}
