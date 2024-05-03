package main

import (
	"fmt"
	"log"
	"net"
)

// TCP Echo Server
func main() {
	port := ":8080"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Listener failed, error: %v", err)
	}
	log.Printf("Server listening on port: %v", port)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Connection failed, error: %v", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Reading error: %v", err)
		return
	}

	// Data processing
	requestData := buf[:n]
	fmt.Println("Received: ", string(requestData))

	// Send the echo buffer back
	conn.Write(buf[:n])
}
