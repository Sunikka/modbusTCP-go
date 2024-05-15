package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/Sunikka/modbusTCP-go/internal/modbusADU"
)

const RegLen = modbusADU.RegAreaLen

// Modbus TCP Server
var RegArea0xxxx [RegLen]uint8
var RegArea1xxxx [RegLen]uint8
var RegArea3xxxx [RegLen]uint16
var RegArea4xxxx [RegLen]uint16

func main() {
	// initRegisters()

	port := "3000"
	addr := "127.0.0.1:" + port

	ln, err := net.Listen("tcp", addr)
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

	var buf_in []uint8
	// var buf_out []uint8

	//	n, err := conn.Read(buf_in)
	//if err != nil {
	//log.Printf("Reading error: %v", err)
	//return
	//}

	// Data processing
	var reqData *modbusADU.MBTCPADU
	err := json.Unmarshal(buf_in, &reqData)
	if err != nil {
		log.Println("Error unmarshalling the received data: ", err)
		return
	}
	fmt.Println("Function code: ", reqData.FuncCode)

	// Send the echo buffer back
	conn.Write(buf_in)
}

// func initRegisters() {
// for i := 0; modbusADU.RegAreaLen < i; i++ {
// RegArea0xxxx[i] = 0
// RegArea1xxxx[i] = 0
// RegArea3xxxx[i] = 0
// RegArea4xxxx[i] = 0
// }
// }
