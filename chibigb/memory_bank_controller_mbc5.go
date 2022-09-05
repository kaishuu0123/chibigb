package chibigb

const MBC5RAMBanksSize = 0x20000

type MemoryBankControllerMBC5 struct {
	console *Console

	currentROMBankHigh int
	currentROMBank     int
	currentROMAddress  int
	currentRAMBank     int
	currentRAMAddress  int
	ramEnabled         bool
	ramBanks           [MBC5RAMBanksSize]byte
}

func NewMemoryBankControllerMBC5(console *Console) MemoryBankController {
	mbc := &MemoryBankControllerMBC5{
		console: console,
	}

	for i := 0; i < MBC5RAMBanksSize; i++ {
		mbc.ramBanks[i] = 0xFF
	}

	mbc.currentROMBank = 1
	mbc.currentROMAddress = 0x4000

	return mbc
}

func (m *MemoryBankControllerMBC5) PerformRead(addr uint16) byte {
	switch addr & 0xE000 {
	case 0x4000, 0x6000:
		return m.console.Cartridge.ROM[(uint64(addr)-0x4000)+uint64(m.currentROMAddress)]
	case 0xA000:
		if m.ramEnabled {
			return m.ramBanks[(int(addr)-0xA000)+m.currentRAMAddress]
		} else {
			return 0xFF
		}
	default:
		return m.console.Memory.Retrieve(addr)
	}
}

func (m *MemoryBankControllerMBC5) PerformWrite(addr uint16, value byte) {
	switch addr & 0xE000 {
	case 0x0000:
		if m.console.Cartridge.ramSize > 0 {
			m.ramEnabled = (value & 0x0F) == 0x0A
		}
	case 0x2000:
		if addr < 0x3000 {
			m.currentROMBank = int(value) | (m.currentROMBankHigh << 8)
		} else {
			m.currentROMBankHigh = int(value & 0x01)
			m.currentROMBank = (m.currentROMBank & 0xFF) | (m.currentROMBankHigh << 8)
		}
		m.currentROMBank &= m.console.Cartridge.romBankCount - 1
		m.currentROMAddress = m.currentROMBank * 0x4000
	case 0x4000:
		m.currentRAMBank = int(value & 0x0F)
		m.currentRAMBank &= (m.console.Cartridge.ramBankCount - 1)
		m.currentRAMAddress = m.currentRAMBank * 0x2000
	case 0x6000:
		// NOTHING DONE
	case 0xA000:
		if m.ramEnabled {
			m.ramBanks[(int(addr)-0xA000)+m.currentRAMAddress] = value
		} else {
			// NOTHING DONE
		}
	default:
		m.console.Memory.Load(addr, value)
	}
}
