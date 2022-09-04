package chibigb

const MBC2RAMBanksSize = 0x8000

type MemoryBankControllerMBC2 struct {
	console *Console

	mode              int
	currentROMBank    int
	currentROMAddress int
	ramEnabled        bool
}

func NewMemoryBankControllerMBC2(console *Console) MemoryBankController {
	mbc := &MemoryBankControllerMBC2{
		console: console,
	}

	mbc.currentROMBank = 1
	mbc.currentROMAddress = 0x4000

	return mbc
}

func (m *MemoryBankControllerMBC2) PerformRead(addr uint16) byte {
	switch addr & 0xE000 {
	case 0x4000, 0x6000:
		return m.console.Cartridge.ROM[(uint64(addr)-0x4000)+uint64(m.currentROMAddress)]
	case 0xA000:
		if addr < 0xA200 {
			if m.ramEnabled {
				return m.console.Memory.Retrieve(addr)
			} else {
				return 0xFF
			}
		} else {
			return 0x00
		}
	default:
		return m.console.Memory.Retrieve(addr)
	}
}

func (m *MemoryBankControllerMBC2) PerformWrite(addr uint16, value byte) {
	switch addr & 0xE000 {
	case 0x0000:
		if (addr & 0x0100) != 0 {
			m.ramEnabled = (value & 0x0F) == 0x0A
		}
	case 0x2000:
		if (addr & 0x0100) != 0 {
			m.currentROMBank = int(value & 0x0F)
			if m.currentROMBank == 0 {
				m.currentROMBank = 1
			}
			m.currentROMBank &= m.console.Cartridge.romBankCount - 1
			m.currentROMAddress = m.currentROMBank * 0x4000
		} else {
			// NOTHING DONE
		}
	case 0x4000, 0x6000:
		// NOTHING DONE
	case 0xA000:
		if addr < 0xA200 {
			if m.ramEnabled {
				m.console.Memory.Load(addr, value&0x0F)
			} else {
				// NOTHING DONE
			}
		} else {
			// NOTHING DONE
		}
	default:
		m.console.Memory.Load(addr, value)
	}
}
