package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/Sunikka/modbusTCP-go/internal/modbusADU"
)

// MODBUS TCP CLIENT
func main() {
	baseServerAddr := "127.0.0.1:"

	var serverAddr string
	mbCmd := modbusADU.MBTCPADU{}
	var port string
	var fc uint8
	var regAddr uint16
	var regCount uint16

	// CLI arguments
	// Example: ./modbusTCPClient 502 04 0000 0000
	// 1. Port
	// 2. FC
	// 3. first reg address
	// 4. number of registers
	log.Print(len(os.Args))
	if len(os.Args) == 5 {
		port = os.Args[1]

		fcNum, err := strconv.ParseUint(os.Args[2], 16, 8)
		if err != nil {
			log.Println("fcNum error: ", err)
			return
		}

		regAddrNum, err := strconv.ParseUint(os.Args[3], 16, 16)
		if err != nil {
			log.Println("regAddrNum error:", err)
			return
		}

		regCountNum, err := strconv.ParseUint(os.Args[4], 16, 16)
		if err != nil {
			log.Println("regCountNum error: ", err)
		}

		fc = uint8(fcNum)
		regAddr = uint16(regAddrNum)
		regCount = uint16(regCountNum)
		log.Print(regCount)
		mbCmd.Fill(fc, regAddr, regCount)

	} else {
		log.Println("Invalid amount of arguments, 5 is needed")
	}
	// Marshal the modbus command into binary
	buf_in, err := json.Marshal(mbCmd)
	if err != nil {
		log.Printf("Error marshalling the modbus message: %v", err)
		return
	}
	//	log.Println("buf in: ", buf_in)

	serverAddr = baseServerAddr + port
	log.Println("Host address: ", serverAddr)

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Printf("Connection failed, %v", err)
		return
	}
	defer conn.Close()

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
