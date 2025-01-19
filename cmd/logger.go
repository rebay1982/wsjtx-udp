package main

import (
	"encoding/json"
	"fmt"
	"github.com/rebay1982/wsjtx-udp/pkg/wsjtxudp"
	"net"
	"os"
)

func main() {
	addr := net.UDPAddr{
		Port: 2237,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Can't listen on UDP: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	var msgNb int = 0
	buff := make([]byte, 1024)

	fmt.Println("Listening to WSJT-X...")
	for {
		_, clientAddr, err := conn.ReadFromUDP(buff)

		if err != nil {
			fmt.Printf("Failed to read from UDP: %v\n", err)
			continue
		}

		parser := wsjtxudp.WSJTXParser{}
		message, _ := parser.Parse(buff)

		msgJSON, _ := json.Marshal(message)

		fmt.Printf("%d: Got message of type [%s] from %s\n%s\n", msgNb, message.Header.MsgType, clientAddr, msgJSON)
		msgNb++
	}
}
