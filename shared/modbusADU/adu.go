package modbusADU

type MBAPHeader struct {
	transactionID uint16
	protocolID    uint16
	length        uint16
	unitID        uint8
}

type MBTCPADU struct {
	header  MBAPHeader
	fc      uint8
	regAddr uint16
	regp1   uint16 // first param after reg

	// pdata*
	exception_code uint8
}

func (m MBTCPADU) fill(fc uint8, regAddr uint16, regp1 uint16) *MBTCPADU {
	mbHeader := MBAPHeader{
		transactionID: 0x0001,
		protocolID:    0x0000,
		length:        0x000c,
		unitID:        0x0001,
	}

	mbMsg := MBTCPADU{
		header:  mbHeader,
		fc:      fc,
		regAddr: regAddr,
		regp1:   regp1,
	}

	return &mbMsg
}
