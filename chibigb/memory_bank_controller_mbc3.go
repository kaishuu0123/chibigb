package chibigb

const MBC3RAMBanksSize = 0x8000

type RTCRegisters struct {
	seconds        int32
	minutes        int32
	hours          int32
	days           int32
	control        int32
	latchedSeconds int32
	latchedMinutes int32
	latchedHours   int32
	latchedDays    int32
	latchedControl int32
	lastTime       int32
	padding        int32
}

type MemoryBankControllerMBC3 struct {
	console *Console

	currentROMBank    int
	currentROMAddress int
	currentRAMBank    int
	currentRAMAddress int
	ramEnabled        bool
	rtcEnabled        bool
	ramBanks          [MBC3RAMBanksSize]byte
	rtcLatch          int32
	rtcRegister       byte
	rtcLastTimeCache  int32
	rtc               RTCRegisters
}

func NewMemoryBankControllerMBC3(console *Console) MemoryBankController {
	mbc := &MemoryBankControllerMBC3{
		console: console,
	}

	for i := 0; i < MBC3RAMBanksSize; i++ {
		mbc.ramBanks[i] = 0xFF
	}

	mbc.currentROMBank = 1
	mbc.currentROMAddress = 0x4000
	mbc.rtc.lastTime = int32(console.Cartridge.rtcCurrentTime)

	return mbc
}

func (m *MemoryBankControllerMBC3) PerformRead(addr uint16) byte {
	switch addr & 0xE000 {
	case 0x4000, 0x6000:
		return m.console.Cartridge.ROM[(uint64(addr)-0x4000)+uint64(m.currentROMAddress)]
	case 0xA000:
		if m.currentRAMBank >= 0 {
			if m.ramEnabled {
				return m.ramBanks[(int(addr)-0xA000)+m.currentRAMAddress]
			} else {
				return 0xFF
			}
		} else if m.console.Cartridge.IsRTCPresent() && m.rtcEnabled {
			switch m.rtcRegister {
			case 0x08:
				return byte(m.rtc.latchedSeconds)
			case 0x09:
				return byte(m.rtc.latchedMinutes)
			case 0x0A:
				return byte(m.rtc.latchedHours)
			case 0x0B:
				return byte(m.rtc.latchedDays)
			case 0x0C:
				return byte(m.rtc.latchedControl)
			default:
				return 0xFF
			}
		} else {
			return 0xFF
		}
	default:
		return m.console.Memory.Retrieve(addr)
	}
}

func (m *MemoryBankControllerMBC3) PerformWrite(addr uint16, value byte) {
	switch addr & 0xE000 {
	case 0x0000:
		if m.console.Cartridge.ramSize > 0 {
			m.ramEnabled = (value & 0x0F) == 0x0A
		}
		m.rtcEnabled = (value & 0x0F) == 0x0A
	case 0x2000:
		m.currentROMBank = int(value & 0x7F)
		if m.currentROMBank == 0 {
			m.currentROMBank = 1
		}
		m.currentROMBank &= m.console.Cartridge.romBankCount - 1
		m.currentROMAddress = m.currentROMBank * 0x4000
	case 0x4000:
		if (value >= 0x08) && (value <= 0x0C) {
			// RTC
			if m.console.Cartridge.IsRTCPresent() && m.rtcEnabled {
				m.rtcRegister = value
				m.currentRAMBank = -1
			} else {
				// NOTHING DONE
			}
		} else if value <= 0x03 {
			m.currentRAMBank = int(value)
			m.currentRAMBank &= m.console.Cartridge.ramBankCount - 1
			m.currentRAMAddress = m.currentRAMBank * 0x2000
		} else {
			// NOTHING DONE
		}
	case 0x6000:
		if m.console.Cartridge.IsRTCPresent() {
			// RTC Latch
			if (m.rtcLatch == 0x00) && (value == 0x01) {
				m.UpdateRTC()
				m.rtc.latchedSeconds = m.rtc.seconds
				m.rtc.latchedMinutes = m.rtc.minutes
				m.rtc.latchedHours = m.rtc.hours
				m.rtc.latchedDays = m.rtc.days
				m.rtc.latchedControl = m.rtc.control
			}

			m.rtcLatch = int32(value)
		}
	case 0xA000:
		if m.currentRAMBank >= 0 {
			if m.ramEnabled {
				m.console.Cartridge.ROM[(uint64(addr)-0x4000)+uint64(m.currentROMAddress)] = value
			} else {
				// NOTHING DONE
			}
		} else if m.console.Cartridge.IsRTCPresent() && m.rtcEnabled {
			switch m.rtcRegister {
			case 0x08:
				m.rtc.seconds = int32(value)
			case 0x09:
				m.rtc.minutes = int32(value)
			case 0x0A:
				m.rtc.hours = int32(value)
			case 0x0B:
				m.rtc.days = int32(value)
			case 0x0C:
				m.rtc.control = (m.rtc.control & 0x80) | (int32(value) & 0x01)
			}
		} else {
			// NOTHING DONE
		}
	default:
		m.console.Memory.Load(addr, value)
	}
}

func (m *MemoryBankControllerMBC3) UpdateRTC() {
	now := int32(m.console.Cartridge.rtcCurrentTime)

	if CheckBit(byte(m.rtc.control), 6) == false && m.rtcLastTimeCache != now {
		m.rtcLastTimeCache = now
		difference := int64(now - m.rtc.lastTime)
		m.rtc.lastTime = now

		if difference > 0 {
			m.rtc.seconds += int32(difference % 60)
			if m.rtc.seconds > 59 {
				m.rtc.seconds -= 60
				m.rtc.minutes++
			}

			difference /= 60
			m.rtc.minutes += int32(difference % 60)
			if m.rtc.minutes > 59 {
				m.rtc.minutes -= 60
				m.rtc.hours++
			}

			difference /= 60
			m.rtc.hours += int32(difference % 24)
			if m.rtc.hours > 23 {
				m.rtc.hours -= 24
				m.rtc.days++
			}

			difference /= 24
			m.rtc.days += int32(difference & 0xFFFFFFFF)

			if m.rtc.days > 0xFF {
				m.rtc.control = (m.rtc.control & 0xC1) | 0x01
				if m.rtc.days > 511 {
					m.rtc.days %= 512
					m.rtc.control |= 0x80
					m.rtc.control &= 0xC0
				}
			}
		}
	}
}
