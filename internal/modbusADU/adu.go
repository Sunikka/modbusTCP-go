package modbusADU

// Constants
const RegAreaLen = 0x00FF
const MinMsgLen = 12

// Area offsets
const DOAreaOffset = 0
const DIAreaOffset = 10000
const AOAreaOffset = 30000
const AIAreaOffset = 40000

// FC
const FC01 = 0x01
const FC02 = 0x02
const FC03 = 0x03
const FC04 = 0x04
const FC05 = 0x05
const FC06 = 0x06
const FC15 = 0x0f
const FC16 = 0x10

type MBAPHeader struct {
	transactionID uint16
	protocolID    uint16
	length        uint16
	unitID        uint8
}

type MBTCPADU struct {
	Header   MBAPHeader
	FuncCode uint8
	RegAddr  uint16
	Regp1    uint16

	// pdata*
	// exception_code uint8
}

// Fills the Modbus ADU with parameter values
func (m MBTCPADU) Fill(fc uint8, regAddr uint16, regp1 uint16) {
	mbHeader := MBAPHeader{
		transactionID: 0x0001,
		protocolID:    0x0000,
		length:        0x000c,
		unitID:        0x0001,
	}

	m.Header = mbHeader
	m.FuncCode = fc
	m.RegAddr = regAddr
	m.Regp1 = regp1
}
