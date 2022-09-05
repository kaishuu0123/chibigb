package chibigb

import (
	"log"
)

type MemoryBankController interface {
	PerformRead(addr uint16) byte
	PerformWrite(addr uint16, value byte)
}

type Memory struct {
	console *Console
	mbc     MemoryBankController

	maps                    [65536]byte
	currentWRAMBank         int
	currentLCDRAMBank       int
	wramBanks               [0x8000]byte
	lcdRAMBank1             [0x2000]byte
	hdmaEnabled             bool
	hdma                    [5]byte
	hdmaSource              uint16
	hdmaDestination         uint16
	bootromDMGEnabled       bool
	bootromGBCEnabled       bool
	bootromDMGLoaded        bool
	bootromGBCLoaded        bool
	bootromDMG              []byte
	bootromGBC              []byte
	bootromRegistryDisabled bool
}

// From Gambatte emulator
var InitialValuesForFFXX = [256]byte{
	0xCF, 0x00, 0x7E, 0xFF, 0xD3, 0x00, 0x00, 0xF8, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE1,
	0x80, 0xBF, 0xF3, 0xFF, 0xBF, 0xFF, 0x3F, 0x00, 0xFF, 0xBF, 0x7F, 0xFF, 0x9F, 0xFF, 0xBF, 0xFF,
	0xFF, 0x00, 0x00, 0xBF, 0x77, 0xF3, 0xF1, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0x71, 0x72, 0xD5, 0x91, 0x58, 0xBB, 0x2A, 0xFA, 0xCF, 0x3C, 0x54, 0x75, 0x48, 0xCF, 0x8F, 0xD9,
	0x91, 0x80, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFC, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0x2B, 0x0B, 0x64, 0x2F, 0xAF, 0x15, 0x60, 0x6D, 0x61, 0x4E, 0xAC, 0x45, 0x0F, 0xDA, 0x92, 0xF3,
	0x83, 0x38, 0xE4, 0x4E, 0xA7, 0x6C, 0x38, 0x58, 0xBE, 0xEA, 0xE5, 0x81, 0xB4, 0xCB, 0xBF, 0x7B,
	0x59, 0xAD, 0x50, 0x13, 0x5E, 0xF6, 0xB3, 0xC1, 0xDC, 0xDF, 0x9E, 0x68, 0xD7, 0x59, 0x26, 0xF3,
	0x62, 0x54, 0xF8, 0x36, 0xB7, 0x78, 0x6A, 0x22, 0xA7, 0xDD, 0x88, 0x15, 0xCA, 0x96, 0x39, 0xD3,
	0xE6, 0x55, 0x6E, 0xEA, 0x90, 0x76, 0xB8, 0xFF, 0x50, 0xCD, 0xB5, 0x1B, 0x1F, 0xA5, 0x4D, 0x2E,
	0xB4, 0x09, 0x47, 0x8A, 0xC4, 0x5A, 0x8C, 0x4E, 0xE7, 0x29, 0x50, 0x88, 0xA8, 0x66, 0x85, 0x4B,
	0xAA, 0x38, 0xE7, 0x6B, 0x45, 0x3E, 0x30, 0x37, 0xBA, 0xC5, 0x31, 0xF2, 0x71, 0xB4, 0xCF, 0x29,
	0xBC, 0x7F, 0x7E, 0xD0, 0xC7, 0xC3, 0xBD, 0xCF, 0x59, 0xEA, 0x39, 0x01, 0x2E, 0x00, 0x69, 0x00,
}

func NewMemory(console *Console) *Memory {
	return &Memory{
		console: console,
	}
}

func (m *Memory) SetMemoryBankController() {
	var memoryBankController MemoryBankController
	switch m.console.Cartridge.cartridgeType {
	case CARTRIDGE_NO_MBC:
		memoryBankController = NewMemoryBankControllerNoMBC(m.console)
	case CARTRIDGE_MBC1:
		memoryBankController = NewMemoryBankControllerMBC1(m.console)
	case CARTRIDGE_MBC2:
		memoryBankController = NewMemoryBankControllerMBC2(m.console)
	case CARTRIDGE_MBC3:
		memoryBankController = NewMemoryBankControllerMBC3(m.console)
	case CARTRIDGE_MBC5:
		memoryBankController = NewMemoryBankControllerMBC5(m.console)
	default:
		log.Fatalf("not supported cartridge type: %d\n", m.console.Cartridge.cartridgeType)
	}

	m.mbc = memoryBankController
}

func (m *Memory) Reset() {
	m.currentWRAMBank = 1

	for i := 0; i < 65536; i++ {
		m.maps[i] = 0x00

		if i >= 0x8000 && i < 0xA000 {
			m.maps[i] = 0x00
		} else if i >= 0xC000 && i < 0xE000 {
			if ((i & 0x08) ^ ((i & 0x0800) >> 8)) != 0 {
				m.maps[i] = 0x0F
			} else {
				m.maps[i] = 0xFF
				if i >= 0xD000 {
					for a := 0; a < 8; a++ {
						if a != 2 {
							m.wramBanks[(i-0xD000)+(0x1000*a)] = m.maps[i-0x1000]
						} else {
							m.wramBanks[(i-0xD000)+(0x1000*a)] = 0x00
						}
					}
				}
			}
		} else if i >= 0xFF00 {
			m.maps[i] = InitialValuesForFFXX[i-0xFF00]
		} else {
			m.maps[i] = 0xFF
		}
	}
}

func (m *Memory) Read(addr uint16) byte {
	switch addr & 0xE000 {
	case 0x0000:
		return m.mbc.PerformRead(addr)
	case 0x2000, 0x4000, 0x6000:
		return m.mbc.PerformRead(addr)
	case 0x8000:
		return m.CommonPerformRead(addr)
	case 0xA000:
		return m.mbc.PerformRead(addr)
	case 0xC000, 0xE000:
		if addr < 0xFF00 {
			return m.CommonPerformRead(addr)
		} else {
			return m.IORegistersPerformRead(addr)
		}
	default:
		return m.Retrieve(addr)
	}
}

func (m *Memory) Write(addr uint16, value byte) {
	switch addr & 0xE000 {
	case 0x0000, 0x2000, 0x4000, 0x6000:
		m.mbc.PerformWrite(addr, value)
	case 0x8000:
		m.CommonPerformWrite(addr, value)
	case 0xA000:
		m.mbc.PerformWrite(addr, value)
	case 0xC000, 0xE000:
		if addr < 0xFF00 {
			m.CommonPerformWrite(addr, value)
		} else {
			m.IORegistersPerformWrite(addr, value)
		}
	default:
		m.Load(addr, value)
	}
}

func (m *Memory) Retrieve(addr uint16) byte {
	return m.maps[addr]
}

func (m *Memory) Load(addr uint16, value byte) {
	m.maps[addr] = value
}

func (m *Memory) CommonPerformRead(addr uint16) byte {
	if addr >= 0xFEA0 && addr < 0xFF00 {
		if (((addr + ((addr >> 4) - 0x0FEA)) >> 2) & 1) != 0 {
			return 0x00
		} else {
			return 0xFF
		}
	}

	return m.Retrieve(addr)
}

func (m *Memory) CommonPerformWrite(addr uint16, value byte) {
	switch addr & 0xE000 {
	case 0x8000:
		m.Load(addr, value)
	case 0xC000:
		if addr < 0xDE00 {
			m.Load(addr, value)
			m.Load(addr+0x2000, value)
		} else {
			m.Load(addr, value)
		}
	case 0xE000:
		if addr < 0xFE00 {
			m.Load(addr-0x2000, value)
			m.Load(addr, value)
		} else {
			m.Load(addr, value)
		}
	default:
		m.Load(addr, value)
	}
}

func (m *Memory) IORegistersPerformRead(addr uint16) byte {
	if addr >= 0xFF10 && addr <= 0xFF3F {
		return m.console.APU.Read(addr)
	}

	switch addr {
	// P1
	case 0xFF00:
		return m.console.Controller.Read()
	// UNDOCUMENTED
	case 0xFF03:
		return 0xFF
	// TAC
	case 0xFF07:
		return m.Retrieve(0xFF07) | 0xF8
	// UNDOCUMENTED
	case 0xFF08, 0xFF09, 0xFF0A, 0xFF0B, 0xFF0C, 0xFF0D, 0xFF0E:
		return 0xFF
	// IF
	case 0xFF0F:
		return m.Retrieve(0xFF0F) | 0xE0
	// STAT
	case 0xFF41:
		return m.Retrieve(0xFF41) | 0x80
	// LY
	case 0xFF44:
		if m.console.PPU.screenEnabled {
			return m.Retrieve(0xFF44)
		} else {
			return 0x00
		}
	// UNDOCUMENTED
	case 0xFF4C:
		return 0xFF
	// VBK
	case 0xFF4F:
		return m.Retrieve(0xFF4F) | 0xFE
	// HDMA1
	case 0xFF51:
		return m.Retrieve(addr)
	// HDMA2
	case 0xFF52:
		return m.Retrieve(addr)
	// HDMA3
	case 0xFF53:
		return m.Retrieve(addr)
	// HDMA4
	case 0xFF54:
		return m.Retrieve(addr)
	// DMA CGB
	case 0xFF55:
		return m.Retrieve(addr)
	// BCPS, OCPS
	case 0xFF68, 0xFF6A:
		return 0xC0
	// BCPD, OCPD
	case 0xFF69, 0xFF6B:
		return 0xFF
	// SVBK
	case 0xFF70:
		return 0xFF
	// UNDOCUMENTED
	case 0xFF76:
		return 0xFF
	case 0xFF77:
		return 0xFF
	}

	return m.Retrieve(addr)
}

func (m *Memory) IORegistersPerformWrite(addr uint16, value byte) {
	if addr >= 0xFF10 && addr <= 0xFF3F {
		m.console.APU.Write(addr, value)
		return
	}

	switch addr {
	case 0xFF00:
		m.console.Controller.Write(value)
	case 0xFF04:
		// DIV
		m.console.CPU.ResetDIVCycles()
	case 0xFF07:
		// TAC
		v := value & 0x07
		currentTAC := m.Retrieve(0xFF07)
		if currentTAC&0x03 != v&0x03 {
			m.console.CPU.ResetTIMACycles()
		}
		m.Load(addr, value)
	case 0xFF0F:
		// IF
		m.Load(addr, value&0x1F)
	case 0xFF40:
		// LCDC
		currentLCDC := m.Retrieve(0xFF40)
		newLCDC := value
		m.Load(addr, newLCDC)
		if CheckBit(currentLCDC, 5) == false && CheckBit(newLCDC, 5) {
			m.console.PPU.ResetWindowLine()
		}
		if CheckBit(newLCDC, 7) {
			m.console.PPU.EnableScreen()
		} else {
			m.console.PPU.DisableScreen()
		}
	case 0xFF41:
		// STAT
		currentSTAT := m.Retrieve(0xFF41) & 0x07
		newSTAT := (value & 0x78) | (currentSTAT & 0x07)

		m.Load(addr, newSTAT)

		lcdc := m.Retrieve(0xFF40)
		signal := m.console.PPU.GetIRQ48Signal()
		mode := m.console.PPU.GetCurrentStatusMode()

		signal &= (newSTAT >> 3) & 0x0F
		m.console.PPU.SetIRQ48Signal(signal)

		if CheckBit(lcdc, 7) {
			if CheckBit(newSTAT, 3) && mode == 0 {
				if signal == 0 {
					m.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
				}
				signal = SetBit(signal, 0)
			}
			if CheckBit(newSTAT, 4) && mode == 1 {
				if signal == 0 {
					m.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
				}
				signal = SetBit(signal, 1)
			}
			if CheckBit(newSTAT, 5) && mode == 2 {
				if signal == 0 {
					m.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
				}
				// signal = SetBit(signal, 2)
			}
			m.console.PPU.CompareLYToLYC()
		}
	case 0xFF44:
		// LY
		currentLY := m.Retrieve(0xFF44)
		if CheckBit(currentLY, 7) && CheckBit(value, 7) == false {
			m.console.PPU.DisableScreen()
		}
	case 0xFF45:
		// LYC
		currentLYC := m.Retrieve(0xFF45)
		if currentLYC != value {
			m.Load(0xFF45, value)
			lcdc := m.Retrieve(0xFF40)
			if CheckBit(lcdc, 7) {
				m.console.PPU.CompareLYToLYC()
			}
		}
	case 0xFF46:
		// DMA
		m.Load(addr, value)
		m.PerformDMA(value)
	case 0xFF4D:
		// KEY1
		m.Load(addr, value)
	case 0xFF4F:
		// VBK
		m.Load(addr, value)
	case 0xFF50:
		// BOOT
		if (value & 0x01) > 0 {
			// DisableBootromRegistry()
		}
	case 0xFF51:
		// HDMA1
		m.Load(addr, value)
	case 0xFF52:
		// HDMA2
		m.Load(addr, value)
	case 0xFF53:
		// HDMA3
		m.Load(addr, value)
	case 0xFF54:
		// HDMA4
		m.Load(addr, value)
	case 0xFF55:
		// DMA CGB
		m.Load(addr, value)
	case 0xFF68:
		// BCPS
		m.Load(addr, value)
	case 0xFF69:
		// BCPD
		m.Load(addr, value)
	case 0xFF6A:
		// OCPS
		m.Load(addr, value)
	case 0xFF6B:
		// OCPD
		m.Load(addr, value)
	case 0xFF6C:
		// UNDOCUMENTED
		m.Load(0xFF6C, value|0xFE)
	case 0xFF70:
		// SVBK
		m.Load(addr, value)
	case 0xFF75:
		// UNDOCUMENTED
		m.Load(0xFF75, value|0x8F)
	case 0xFFFF:
		// IE
		m.Load(addr, value&0x1F)
	default:
		m.Load(addr, value)
	}
}

func (m *Memory) PerformDMA(value byte) {
	var addr uint16 = uint16(value) << 8
	if addr >= 0x8000 && addr < 0xE000 {
		for i := 0; i < 0xA0; i++ {
			m.Load(0xFE00+uint16(i), m.Read(addr+uint16(i)))
		}
	}
}
