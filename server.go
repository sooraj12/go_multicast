package main

import (
	"fmt"
	"log"
	"net"
)

func server() {
	groupAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", GROUP_IP, PORT))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp4", nil, groupAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	fmt.Printf("Listening for UDP multicast messages on %s:%d\n", GROUP_IP, PORT)

	buf := make([]byte, 1024)
	for {
		n, srcAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("Error reading from UDP connection: %v\n", err)
			continue
		}

		fmt.Printf("Received message from %s: %s\n", srcAddr.String(), string(buf[:n]))
	}

}
