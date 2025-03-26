package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Failed to stablish connection with %s: %v", addr, err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from stdin: %v", err)
			continue
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
	}
}
