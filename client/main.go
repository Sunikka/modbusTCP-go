package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	baseServerAddr := "127.0.0.1:"
	//testData := "Hello World!"

	var serverAddr string
	var modbusMsg MBTCPADU
	var port string
	var FC uint8
	var regAddr uint16
	var regCount uint16

	// CLI arguments
	// Example command: ./goTCPClient "Hello world!"
	// 	if len(os.Args) == 2 {
	// port = "8080"
	// msg = os.Args[1]
	// } else if len(os.Args) == 3 {
	// port = string(os.Args[1])
	// msg = os.Args[2]
	// } else {
	// log.Printf("No arguments provided")
	// return
	// }

	// CLI arguments
	// Example: ./modbusTCPClient 502 04 0000 0000
	// 1. Port
	// 2. FC
	// 3. first reg address
	// 4. number of registers
	if len(os.Args) == 4 {
		port = os.Args[1]
		FC = os.Args[2]
		regAddr = os.Args[3]
		regCount = os.Args[4]
	}

	serverAddr = baseServerAddr + port
	log.Println("Host address: ", serverAddr)

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Printf("Connection failed, %v", err)
		return
	}
	defer conn.Close()

	buf_in := []byte(msg)
	_, err = conn.Write(buf_in)
	if err != nil {
		log.Printf("Writing to connection failed, error: %v", err)
		return
	}

	// Reading the server Response
	buf_out := make([]byte, 1024)
	n, err := conn.Read(buf_out)
	if err != nil {
		log.Printf("Reading the server response failed, error: %v", err)
		return
	}

	response := string(buf_out[:n])
	fmt.Println("Server response: ", response)
}
