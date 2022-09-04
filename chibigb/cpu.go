package chibigb

const (
	CPUFLAG_NONE  = 0x00
	CPUFLAG_CARRY = 0x10
	CPUFLAG_HALF  = 0x20
	CPUFLAG_SUB   = 0x40
	CPUFLAG_ZERO  = 0x80
)

type CPUInterrupt int

const (
	INTERRUPT_NONE    CPUInterrupt = 0x00
	INTERRUPT_VBLANK               = 0x01
	INTERRUPT_LCDSTAT              = 0x02
	INTERRUPT_TIMER                = 0x04
	INTERRUPT_SERIAL               = 0x08
	INTERRUPT_JOYPAD               = 0x10
)

type CPU struct {
	console *Console

	AF *CPURegister
	BC *CPURegister
	DE *CPURegister
	HL *CPURegister

	PC *CPURegister
	SP *CPURegister

	OpCodeTable   [256]CPUOpCodeEntity
	OpCodeCBTable [256]CPUOpCodeEntity

	IME         bool
	Halt        bool
	BranchTaken bool
	SkipPCBug   bool

	CurrentClockCycles   uint64
	DIVCycles            uint32
	TIMACycles           uint32
	SerialBit            int
	SerialCycles         int
	IMECycles            int
	UnhaltCycles         int
	InterruptDelayCycles int
	AccurateOpeCodeState int
	SpeedMultiplier      int
	ReadCache            byte
}

func NewCPU(console *Console) *CPU {
	cpu := &CPU{
		console: console,
		AF:      NewCPURegister(),
		BC:      NewCPURegister(),
		DE:      NewCPURegister(),
		HL:      NewCPURegister(),
		PC:      NewCPURegister(),
		SP:      NewCPURegister(),
	}

	cpu.createTable()
	cpu.createCBTable()

	return cpu
}

func (c *CPU) Reset() {
	c.PC.Set(0x0100)
	c.SP.Set(0xFFFE)

	c.AF.Set(0x01B0)
	c.BC.Set(0x0013)
	c.DE.Set(0x00D8)
	c.HL.Set(0x014D)
}

func (c *CPU) Step(count byte) byte {
	var executed byte = 0

	for executed < count {
		c.CurrentClockCycles = 0

		if c.AccurateOpeCodeState == 0 && c.Halt {
			c.CurrentClockCycles += uint64(c.adjustedCycles(4))

			if c.UnhaltCycles > 0 {
				c.UnhaltCycles -= int(c.CurrentClockCycles)
				if c.UnhaltCycles <= 0 {
					c.UnhaltCycles = 0
					c.Halt = false
				}
			}

			if c.Halt && c.InterruptPending() != INTERRUPT_NONE && c.UnhaltCycles == 0 {
				c.UnhaltCycles = c.adjustedCycles(12)
			}
		}

		var interruptServed bool = false
		if c.Halt == false {
			var interrupt CPUInterrupt = c.InterruptPending()

			if c.IME && interrupt != INTERRUPT_NONE && c.AccurateOpeCodeState == 0 {
				c.ServeInterrupt(interrupt)
				interruptServed = true
			} else {
				var opcode byte = c.MemoryRead(c.PC.Get())
				c.PC.Increment()

				if c.SkipPCBug {
					c.SkipPCBug = false
					c.PC.Decrement()
				}

				var accurateOpCodes *[256]byte
				var machineCycles *[256]byte
				var opcodeTable *[256]CPUOpCodeEntity
				var isCB = (opcode == 0xCB)

				if isCB {
					accurateOpCodes = &OPCodeCBAccurate
					machineCycles = &OPCodeCBMachineCycles
					opcodeTable = &c.OpCodeCBTable

					opcode = c.MemoryRead(c.PC.Get())
					c.PC.Increment()

					if c.SkipPCBug {
						c.SkipPCBug = false
						c.PC.Decrement()
					}
				} else {
					accurateOpCodes = &OPCodeAccurate
					machineCycles = &OPCodeMachineCycles
					opcodeTable = &c.OpCodeTable
				}

				if accurateOpCodes[opcode] != 0 && c.AccurateOpeCodeState == 0 {
					var leftCycles int
					if accurateOpCodes[opcode] < 3 {
						leftCycles = 2
					} else {
						leftCycles = 3
					}
					c.CurrentClockCycles += (uint64(machineCycles[opcode]) - uint64(leftCycles)) * uint64(c.adjustedCycles(4))
					c.AccurateOpeCodeState = 1
					c.PC.Decrement()
					if isCB {
						c.PC.Decrement()
					}
				} else {
					opcodeTable[opcode].fn()

					if c.BranchTaken {
						c.BranchTaken = false
						c.CurrentClockCycles += uint64(OPCodeBranchMachineCycles[opcode]) * uint64(c.adjustedCycles(4))
					} else {
						switch c.AccurateOpeCodeState {
						case 0:
							c.CurrentClockCycles += uint64(machineCycles[opcode]) * uint64(c.adjustedCycles(4))
						case 1:
							if accurateOpCodes[opcode] == 3 {
								c.CurrentClockCycles += 1 * uint64(c.adjustedCycles(4))
								c.AccurateOpeCodeState = 2
								c.PC.Decrement()
								if isCB {
									c.PC.Decrement()
								}
							} else {
								c.CurrentClockCycles += 2 * uint64(c.adjustedCycles(4))
								c.AccurateOpeCodeState = 0
							}
						case 2:
							c.CurrentClockCycles += 2 * uint64(c.adjustedCycles(4))
							c.AccurateOpeCodeState = 0
						}
					}
				}
			}
		}

		if interruptServed == false && c.InterruptDelayCycles > 0 {
			c.InterruptDelayCycles -= int(c.CurrentClockCycles)
		}

		if interruptServed == false && c.AccurateOpeCodeState == 0 && c.IMECycles > 0 {
			c.IMECycles -= int(c.CurrentClockCycles)

			if c.IMECycles <= 0 {
				c.IMECycles = 0
				c.IME = true
			}
		}

		executed += byte(c.CurrentClockCycles)
	}

	return executed
}

func (c *CPU) UpdateTimers(clockCycles byte) {
	c.DIVCycles += uint32(clockCycles)

	var divCycles uint32 = uint32(c.adjustedCycles(256))
	for c.DIVCycles >= divCycles {
		c.DIVCycles -= divCycles
		var div byte = c.MemoryRetrieve(0xFF04)
		div++
		c.MemoryLoad(0xFF04, div)
	}

	var tac byte = c.MemoryRetrieve(0xFF07)
	// if tima is running
	if (tac & 0x04) != 0 {
		c.TIMACycles += uint32(clockCycles)
		var freq uint32 = 0
		switch tac & 0x03 {
		case 0:
			freq = uint32(c.adjustedCycles(1024))
		case 1:
			freq = uint32(c.adjustedCycles(16))
		case 2:
			freq = uint32(c.adjustedCycles(64))
		case 3:
			freq = uint32(c.adjustedCycles(256))
		}

		for c.TIMACycles >= freq {
			c.TIMACycles -= freq
			var tima byte = c.MemoryRetrieve(0xFF05)
			if tima == 0xFF {
				tima = c.MemoryRetrieve(0xFF06)
				c.RequestInterrupt(INTERRUPT_TIMER)
			} else {
				tima++
			}

			c.MemoryLoad(0xFF05, tima)
		}
	}
}

func (c *CPU) UpdateSerial(clockCycles byte) {
	var sc byte = c.MemoryRetrieve(0xFF02)

	if CheckBit(sc, 7) && CheckBit(sc, 0) {
		c.SerialCycles += int(clockCycles)
		if c.SerialBit < 0 {
			c.SerialBit = 0
			c.SerialCycles = 0
			return
		}

		var serialCycles int = c.adjustedCycles(512)
		if c.SerialCycles >= serialCycles {
			if c.SerialBit > 7 {
				c.MemoryLoad(0xFF02, sc&0x7F)
				c.RequestInterrupt(INTERRUPT_SERIAL)
				c.SerialBit = -1

				return
			}

			var sb byte = c.MemoryRetrieve(0xFF01)
			sb <<= 1
			sb |= 0x01
			c.MemoryLoad(0xFF01, sb)

			c.SerialCycles -= serialCycles
			c.SerialBit++
		}
	}
}

func (c *CPU) InterruptPending() CPUInterrupt {
	var ieReg byte = c.MemoryRetrieve(0xFFFF)
	var ifReg byte = c.MemoryRetrieve(0xFF0F)
	var ieIf byte = ifReg & ieReg

	switch {
	case (ieIf & 0x1F) == 0:
		return INTERRUPT_NONE
	case (ieIf&0x01) != 0 && c.InterruptDelayCycles <= 0:
		return INTERRUPT_VBLANK
	case (ieIf & 0x02) != 0:
		return INTERRUPT_LCDSTAT
	case (ieIf & 0x04) != 0:
		return INTERRUPT_TIMER
	case (ieIf & 0x08) != 0:
		return INTERRUPT_SERIAL
	case (ieIf & 0x10) != 0:
		return INTERRUPT_JOYPAD
	}

	return INTERRUPT_NONE
}

func (c *CPU) RequestInterrupt(interrupt CPUInterrupt) {
	c.MemoryLoad(0xFF0F, c.MemoryRetrieve(0xFF0F)|byte(interrupt))

	if interrupt == INTERRUPT_VBLANK {
		c.InterruptDelayCycles = 4
	}
}

func (c *CPU) ServeInterrupt(interrupt CPUInterrupt) {
	var ifReg byte = c.MemoryRetrieve(0xFF0F)
	c.IME = false
	c.instruction_StackPush(c.PC)
	c.CurrentClockCycles += uint64(c.adjustedCycles(20))

	switch interrupt {
	case INTERRUPT_VBLANK:
		c.InterruptDelayCycles = 0
		c.MemoryLoad(0xFF0F, ifReg&0xFE)
		c.PC.Set(0x0040)
		// UpdateGameShark()
	case INTERRUPT_LCDSTAT:
		c.MemoryLoad(0xFF0F, ifReg&0xFD)
		c.PC.Set(0x0048)
	case INTERRUPT_TIMER:
		c.MemoryLoad(0xFF0F, ifReg&0xFB)
		c.PC.Set(0x0050)
	case INTERRUPT_SERIAL:
		c.MemoryLoad(0xFF0F, ifReg&0xF7)
		c.PC.Set(0x0058)
	case INTERRUPT_JOYPAD:
		c.MemoryLoad(0xFF0F, ifReg&0xEF)
		c.PC.Set(0x0060)
	case INTERRUPT_NONE:
		// NOTHING DONE
	}
}

func (c *CPU) InterruptIsAboutToRaise() bool {
	var ieReg byte = c.MemoryRetrieve(0xFFFF)
	var ifReg byte = c.MemoryRetrieve(0xFF0F)

	return (ifReg & ieReg & 0x1F) != 0
}

func (c *CPU) ResetTIMACycles() {
	c.TIMACycles = 0
	c.MemoryLoad(0xFF05, c.MemoryRetrieve(0xFF06))
}

func (c *CPU) ResetDIVCycles() {
	c.DIVCycles = 0
	c.MemoryLoad(0xFF04, 0x00)
}

func (c *CPU) Halted() bool {
	return c.Halt
}

func (c *CPU) SetZeroFlag(result byte) {
	if result == 0 {
		c.SetFlag(CPUFLAG_ZERO)
	}
}

// XXX: naming
func (c *CPU) KeepCarryFlag() {
	if c.CheckFlag(CPUFLAG_CARRY) {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
}

// XXX: naming
func (c *CPU) KeepZeroFlag() {
	if c.CheckFlag(CPUFLAG_ZERO) {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_ZERO)
	} else {
		c.ClearAllFlags()
	}
}

func (c *CPU) SetFlag(flag byte) {
	c.AF.SetLow(c.AF.GetLow() | flag)
}

func (c *CPU) ClearFlag(flag byte) {
	c.AF.SetLow(c.AF.GetLow() & (^flag))
}

func (c *CPU) FlipFlag(flag byte) {
	c.AF.SetLow(c.AF.GetLow() ^ flag)
}

func (c *CPU) ClearAllFlags() {
	c.AF.SetLow(CPUFLAG_NONE)
}

func (c *CPU) CheckFlag(flag byte) bool {
	return (c.AF.GetLow() & flag) == flag
}

func (c *CPU) MemoryRead(addr uint16) byte {
	return c.console.Memory.Read(addr)
}

func (c *CPU) MemoryWrite(addr uint16, value byte) {
	c.console.Memory.Write(addr, value)
}

func (c *CPU) MemoryRetrieve(addr uint16) byte {
	return c.console.Memory.Retrieve(addr)
}

func (c *CPU) MemoryLoad(addr uint16, value byte) {
	c.console.Memory.Load(addr, value)
}
