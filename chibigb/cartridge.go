package chibigb

import (
	"log"
)

type CartridgeType int

const (
	CARTRIDGE_NO_MBC        CartridgeType = 0
	CARTRIDGE_MBC1                        = 1
	CARTRIDGE_MBC2                        = 2
	CARTRIDGE_MBC3                        = 3
	CARTRIDGE_MBC5                        = 4
	CARTRIDGE_MBC1_MULTI                  = 5
	CARTRIDGE_NOT_SUPPORTED               = 6
)

type Cartridge struct {
	console *Console

	ROM []byte

	cartridgeType CartridgeType
	romSize       int
	ramSize       int
	version       int
	ramBankCount  int
	romBankCount  int

	isGameboyColor bool
	isSuperGameboy bool
	hasBattery     bool
	RTCPresent     bool
	RumblePresent  bool
}

func NewCartridge(console *Console) *Cartridge {
	return &Cartridge{
		console: console,
	}
}

func (c *Cartridge) LoadFromBuffer(buffer []byte) {
	c.ROM = make([]byte, len(buffer))
	copy(c.ROM, buffer)
	c.collectMetadata()
}

func (c *Cartridge) collectMetadata() {
	c.isGameboyColor = c.ROM[0x0143] == 0x80 || c.ROM[0x0143] == 0xC0
	c.isSuperGameboy = c.ROM[0x0146] == 0x03
	c.romSize = int(c.ROM[0x0148])
	c.ramSize = int(c.ROM[0x0149])
	c.version = int(c.ROM[0x014C])

	c.collectCartridgeType(c.ROM[0x0147])

	switch c.ramSize {
	case 0x00:
		if c.cartridgeType == CARTRIDGE_MBC2 {
			c.ramBankCount = 1
		} else {
			c.ramBankCount = 0
		}
	case 0x01, 0x02:
		c.ramBankCount = 1
	case 0x04:
		c.ramBankCount = 16
	default:
		c.ramBankCount = 4
	}

	c.romBankCount = len(c.ROM) / 0x4000
	if c.romBankCount < 2 {
		c.romBankCount = 2
	}

	log.Printf("Cartridge Size %d\n", len(c.ROM))
	log.Printf("ROM Version %d\n", c.version)
	log.Printf("ROM Type %X\n", c.cartridgeType)
	log.Printf("ROM Size %X\n", c.romSize)
	log.Printf("ROM Bank Count %d\n", c.romBankCount)
	log.Printf("RAM Size %X\n", c.ramSize)
	log.Printf("RAM Bank Count %X\n", c.ramBankCount)
	log.Printf("Has Battery: %v\n", c.hasBattery)

	if c.ROM[0x0143] == 0xC0 {
		log.Println("Game Boy Color only")
	} else if c.isGameboyColor {
		log.Println("Game Boy Color supported")
	}

	if c.isSuperGameboy {
		log.Println("Super Game Boy supported")
	}

	checksum := 0
	for j := 0x0134; j < 0x014E; j++ {
		checksum += int(c.ROM[j])
	}
	isValidROM := ((checksum + 25) & 0xFF) == 0
	if isValidROM {
		log.Println("Checksum OK")
	} else {
		log.Println("Checksum FAILED")
	}
}

func (c *Cartridge) collectCartridgeType(cartType byte) {
	t := cartType
	if cartType != 0xEA && c.romSize == 0 {
		t = 0
	}

	switch t {
	case 0x00, 0x08, 0x09:
		c.cartridgeType = CARTRIDGE_NO_MBC
	case 0x01, 0x02, 0x03, 0xEA, 0xFF:
		c.cartridgeType = CARTRIDGE_MBC1
	case 0x05, 0x06:
		c.cartridgeType = CARTRIDGE_MBC2
	case 0x0F, 0x10, 0x11, 0x12, 0x13, 0xFC:
		c.cartridgeType = CARTRIDGE_MBC3
	case 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
		c.cartridgeType = CARTRIDGE_MBC5
	case 0x0B, 0x0D, 0x15, 0x16, 0x17, 0x22, 0x55, 0x56, 0xFD, 0xFE:
		c.cartridgeType = CARTRIDGE_NOT_SUPPORTED
	default:
		c.cartridgeType = CARTRIDGE_NOT_SUPPORTED
	}

	switch t {
	case 0x03, 0x06, 0x09, 0x0D, 0x0F, 0x10, 0x13, 0x17, 0x1B, 0x1E, 0x22, 0xFD, 0xFF:
		c.hasBattery = true
	default:
		c.hasBattery = false
	}

	switch t {
	case 0x0F, 0x10:
		c.RTCPresent = true
	default:
		c.RTCPresent = false
	}

	switch t {
	case 0x1C, 0x1D, 0x1E:
		c.RumblePresent = true
	default:
		c.RumblePresent = false
	}
}
