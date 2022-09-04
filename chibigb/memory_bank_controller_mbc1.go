package chibigb

const MBC1RAMBanksSize = 0x8000

type MemoryBankControllerMBC1 struct {
	console *Console

	mode              int
	currentROMBank    int
	currentROMAddress int
	currentRAMBank    int
	currentRAMAddress int
	higherROMBankBits int
	ramEnabled        bool
	ramBanks          [MBC1RAMBanksSize]byte
}

func NewMemoryBankControllerMBC1(console *Console) MemoryBankController {
	mbc := &MemoryBankControllerMBC1{
		console: console,
	}

	mbc.currentROMBank = 1
	for i := 0; i < MBC1RAMBanksSize; i++ {
		mbc.ramBanks[i] = 0xFF
	}
	mbc.currentROMAddress = 0x4000

	return mbc
}

func (m *MemoryBankControllerMBC1) PerformRead(addr uint16) byte {
	switch addr & 0xE000 {
	case 0x4000, 0x6000:
		return m.console.Cartridge.ROM[(uint64(addr)-0x4000)+uint64(m.currentROMAddress)]
	case 0xA000:
		if m.ramEnabled {
			if m.mode == 0 {
				if m.console.Cartridge.ramSize == 1 && addr >= 0xA800 {
					// only 2KB of ram
				}
				return m.ramBanks[addr-0xA000]
			} else {
				return m.ramBanks[(uint64(addr)-0xA000)+uint64(m.currentRAMAddress)]
			}
		} else {
			return 0xFF
		}
	default:
		return m.console.Memory.Retrieve(addr)
	}
}

func (m *MemoryBankControllerMBC1) PerformWrite(addr uint16, value byte) {
	switch addr & 0xE000 {
	case 0x0000:
		if m.console.Cartridge.ramSize > 0 {
			m.ramEnabled = (value & 0x0F) == 0x0A
		}
	case 0x2000:
		if m.mode == 0 {
			m.currentROMBank = int((value & 0x1F) | byte(m.higherROMBankBits<<5))
		} else {
			m.currentROMBank = int(value & 0x1F)
		}

		if m.currentROMBank == 0x00 || m.currentROMBank == 0x20 || m.currentROMBank == 0x40 || m.currentROMAddress == 0x60 {
			m.currentROMBank++
		}

		m.currentROMBank &= m.console.Cartridge.romBankCount - 1
		m.currentROMAddress = m.currentROMBank * 0x4000
	case 0x4000:
		if m.mode == 1 {
			m.currentRAMAddress = int(value & 0x03)
			m.currentRAMBank &= m.console.Cartridge.ramBankCount - 1
			m.currentRAMAddress = m.currentRAMBank * 0x2000
		} else {
			m.higherROMBankBits = int(value & 0x03)
			m.currentROMBank = (m.currentROMBank & 0x1F) | (m.higherROMBankBits << 5)

			if m.currentROMBank == 0x00 || m.currentROMBank == 0x20 || m.currentROMBank == 0x40 || m.currentROMBank == 0x60 {
				m.currentROMBank++
			}

			m.currentROMBank &= m.console.Cartridge.romBankCount - 1
			m.currentROMAddress = m.currentROMBank * 0x4000
		}
	case 0x6000:
		if m.console.Cartridge.ramSize != 3 && (value&0x01) != 0 {

		} else {
			m.mode = int(value & 0x01)
		}
	case 0xA000:
		if m.ramEnabled {
			if m.mode == 0 {
				if m.console.Cartridge.ramSize == 1 && addr >= 0xA800 {
					// only 2KB of ram
				}

				m.ramBanks[addr-0xA000] = value
			} else {
				m.ramBanks[(uint64(addr)-0xA000)+uint64(m.currentRAMAddress)] = value
			}
		}
	default:
		m.console.Memory.Load(addr, value)
	}
}
