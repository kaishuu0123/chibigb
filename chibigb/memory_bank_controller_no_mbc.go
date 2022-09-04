package chibigb

type MemoryBankControllerNoMBC struct {
	console *Console
}

func NewMemoryBankControllerNoMBC(console *Console) MemoryBankController {
	return &MemoryBankControllerNoMBC{
		console: console,
	}
}

func (m *MemoryBankControllerNoMBC) PerformRead(addr uint16) byte {
	if addr >= 0xA000 && addr < 0xC000 {
		if m.console.Cartridge.ramSize > 0 {
			return m.console.Memory.Retrieve(addr)
		} else {
			return 0xFF
		}
	}

	return m.console.Memory.Retrieve(addr)
}

func (m *MemoryBankControllerNoMBC) PerformWrite(addr uint16, value byte) {
	// ROM
	if addr < 0x8000 {

	} else if addr >= 0xA000 && addr < 0xC000 {
		if m.console.Cartridge.ramSize > 0 {
			m.console.Memory.Load(addr, value)
		} else {
			// Attempting to write to RAM
		}
	} else {
		m.console.Memory.Load(addr, value)
	}
}
