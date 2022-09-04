package chibigb

import (
	"log"
)

type CPUOpCodeEntity struct {
	opcode byte
	name   string
	fn     func()
}

func (c *CPU) instruction_LD_r8_r8(reg1 *byte, reg2 byte) {
	*reg1 = reg2
}

func (c *CPU) instruction_LD_r8_d8(reg1 *byte, addr uint16) {
	*reg1 = c.MemoryRead(addr)
}

func (c *CPU) instruction_LD_d8_r8(addr uint16, reg1 byte) {
	c.MemoryWrite(addr, reg1)
}

func (c *CPU) instruction_OR(number byte) {
	result := c.AF.GetHigh() | number
	c.AF.SetHigh(result)
	c.ClearAllFlags()
	c.SetZeroFlag(result)
}

func (c *CPU) instruction_XOR(number byte) {
	result := c.AF.GetHigh() ^ number
	c.AF.SetHigh(result)
	c.ClearAllFlags()
	c.SetZeroFlag(result)
}

func (c *CPU) instruction_AND(number byte) {
	result := c.AF.GetHigh() & number
	c.AF.SetHigh(result)
	c.ClearAllFlags()
	c.SetFlag(CPUFLAG_HALF)
	c.SetZeroFlag(result)
}

func (c *CPU) instruction_CP(number byte) {
	c.ClearAllFlags()
	c.SetFlag(CPUFLAG_SUB)

	if c.AF.GetHigh() < number {
		c.SetFlag(CPUFLAG_CARRY)
	}
	if c.AF.GetHigh() == number {
		c.SetFlag(CPUFLAG_ZERO)
	}

	if ((c.AF.GetHigh() - number) & 0xF) > (c.AF.GetHigh() & 0xF) {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_INC(reg *byte) {
	result := *reg + 1
	*reg = result
	c.KeepCarryFlag()
	c.SetZeroFlag(result)
	if (result & 0x0F) == 0x00 {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_INC_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get()) + 1
		return
	}
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.KeepCarryFlag()
	c.SetZeroFlag(c.ReadCache)
	if (c.ReadCache & 0x0F) == 0x00 {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_DEC(reg *byte) {
	result := *reg - 1
	*reg = result
	c.KeepCarryFlag()
	c.SetFlag(CPUFLAG_SUB)
	c.SetZeroFlag(result)
	if (result & 0x0F) == 0x0F {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_DEC_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get()) - 1
		return
	}
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.KeepCarryFlag()
	c.SetFlag(CPUFLAG_SUB)
	c.SetZeroFlag(c.ReadCache)
	if (c.ReadCache & 0x0F) == 0x0F {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_ADD(number byte) {
	var result int = int(c.AF.GetHigh()) + int(number)
	var carrybits int = int(c.AF.GetHigh()) ^ int(number) ^ result
	c.AF.SetHigh(byte(result))
	c.ClearAllFlags()
	c.SetZeroFlag(byte(result))
	if (carrybits & 0x100) != 0 {
		c.SetFlag(CPUFLAG_CARRY)
	}
	if (carrybits & 0x10) != 0 {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_ADC(number byte) {
	var carry int
	if c.CheckFlag(CPUFLAG_CARRY) {
		carry = 1
	} else {
		carry = 0
	}
	var result int = int(c.AF.GetHigh()) + int(number) + carry
	c.ClearAllFlags()
	c.SetZeroFlag(byte(result))
	if result > 0xFF {
		c.SetFlag(CPUFLAG_CARRY)
	}
	if ((c.AF.GetHigh() & 0x0F) + (number & 0x0F) + byte(carry)) > 0x0F {
		c.SetFlag(CPUFLAG_HALF)
	}
	c.AF.SetHigh(byte(result))
}

func (c *CPU) instruction_SUB(number byte) {
	var result int = int(c.AF.GetHigh()) - int(number)
	var carrybits int = int(c.AF.GetHigh()) ^ int(number) ^ result
	c.AF.SetHigh(byte(result))
	c.ClearAllFlags()
	c.SetFlag(CPUFLAG_SUB)
	c.SetZeroFlag(byte(result))
	if (carrybits & 0x100) != 0 {
		c.SetFlag(CPUFLAG_CARRY)
	}
	if (carrybits & 0x10) != 0 {
		c.SetFlag(CPUFLAG_HALF)
	}
}

func (c *CPU) instruction_SBC(number byte) {
	var carry int
	if c.CheckFlag(CPUFLAG_CARRY) {
		carry = 1
	} else {
		carry = 0
	}
	var result int = int(c.AF.GetHigh()) - int(number) - carry
	c.ClearAllFlags()
	c.SetFlag(CPUFLAG_SUB)
	c.SetZeroFlag(byte(result))
	if result < 0 {
		c.SetFlag(CPUFLAG_CARRY)
	}
	if ((int(c.AF.GetHigh()) & 0x0F) - (int(number) & 0x0F) - int(carry)) < 0 {
		c.SetFlag(CPUFLAG_HALF)
	}
	c.AF.SetHigh(byte(result))
}

func (c *CPU) instruction_ADD_HL(number uint16) {
	var result int = int(c.HL.Get()) + int(number)
	c.KeepZeroFlag()
	if (result & 0x10000) != 0 {
		c.SetFlag(CPUFLAG_CARRY)
	}
	half := (c.HL.Get() ^ number ^ uint16(result&0xFFFF)) & 0x1000
	if half != 0 {
		c.SetFlag(CPUFLAG_HALF)
	}
	c.HL.Set(uint16(result))
}

func (c *CPU) instruction_ADD_SP(number int8) {
	var result int = int(c.SP.Get()) + int(number)
	c.ClearAllFlags()
	carry := (int(c.SP.Get()) ^ int(number) ^ (result & 0xFFFF)) & 0x100
	if carry != 0 {
		c.SetFlag(CPUFLAG_CARRY)
	}
	half := (int(c.SP.Get()) ^ int(number) ^ (result & 0xFFFF)) & 0x10
	if half != 0 {
		c.SetFlag(CPUFLAG_HALF)
	}
	c.SP.Set(uint16(result))
}

func (c *CPU) instruction_SWAP_Register(reg *byte) {
	low := *reg & 0x0F
	high := (*reg >> 4) & 0x0F
	*reg = (low << 4) + high
	c.ClearAllFlags()
	c.SetZeroFlag(*reg)
}

func (c *CPU) instruction_SWAP_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	low := c.ReadCache & 0x0F
	high := (c.ReadCache >> 4) & 0x0F
	c.ReadCache = (low << 4) + high
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.ClearAllFlags()
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_SLA(reg *byte) {
	if (*reg & 0x80) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	result := *reg << 1
	*reg = result
	c.SetZeroFlag(result)
}

func (c *CPU) instruction_SLA_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	if (c.ReadCache & 0x80) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	c.ReadCache <<= 1
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_SRA(reg *byte) {
	result := *reg
	if (*reg & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	if (result & 0x80) != 0 {
		result >>= 1
		result |= 0x80
	} else {
		result >>= 1
	}
	*reg = result
	c.SetZeroFlag(result)
}

func (c *CPU) instruction_SRA_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	if (c.ReadCache & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	if (c.ReadCache & 0x80) != 0 {
		c.ReadCache >>= 1
		c.ReadCache |= 0x80
	} else {
		c.ReadCache >>= 1
	}
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_SRL(reg *byte) {
	result := *reg
	if (result & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	result >>= 1
	*reg = result
	c.SetZeroFlag(result)
}

func (c *CPU) instruction_SRL_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	if (c.ReadCache & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	c.ReadCache >>= 1
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_RLC(reg *byte, isRegisterA bool) {
	result := *reg
	if (result & 0x80) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
		result <<= 1
		result |= 0x01
	} else {
		c.ClearAllFlags()
		result <<= 1
	}
	*reg = result
	if isRegisterA == false {
		c.SetZeroFlag(result)
	}
}

func (c *CPU) instruction_RLC_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}

	if (c.ReadCache & 0x80) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
		c.ReadCache <<= 1
		c.ReadCache |= 0x01
	} else {
		c.ClearAllFlags()
		c.ReadCache <<= 1
	}
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_RL(reg *byte, isRegisterA bool) {
	var carry byte
	if c.CheckFlag(CPUFLAG_CARRY) {
		carry = 1
	} else {
		carry = 0
	}
	result := *reg
	if (result & 0x80) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	result <<= 1
	result |= carry
	*reg = result
	if isRegisterA == false {
		c.SetZeroFlag(result)
	}
}

func (c *CPU) instruction_RL_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	var carry byte
	if c.CheckFlag(CPUFLAG_CARRY) {
		carry = 1
	} else {
		carry = 0
	}
	if (c.ReadCache & 0x80) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	c.ReadCache <<= 1
	c.ReadCache |= carry
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_RRC(reg *byte, isRegisterA bool) {
	result := *reg
	if (result & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
		result >>= 1
		result |= 0x80
	} else {
		c.ClearAllFlags()
		result >>= 1
	}
	*reg = result
	if isRegisterA == false {
		c.SetZeroFlag(result)
	}
}

func (c *CPU) instruction_RRC_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	if (c.ReadCache & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
		c.ReadCache >>= 1
		c.ReadCache |= 0x80
	} else {
		c.ClearAllFlags()
		c.ReadCache >>= 1
	}
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_RR(reg *byte, isRegisterA bool) {
	var carry byte
	if c.CheckFlag(CPUFLAG_CARRY) {
		carry = 0x80
	} else {
		carry = 0x00
	}
	result := *reg
	if (result & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	result >>= 1
	result |= carry
	*reg = result
	if isRegisterA == false {
		c.SetZeroFlag(result)
	}
}

func (c *CPU) instruction_RR_HL() {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	var carry byte
	if c.CheckFlag(CPUFLAG_CARRY) {
		carry = 0x80
	} else {
		carry = 0x00
	}
	if (c.ReadCache & 0x01) != 0 {
		c.ClearAllFlags()
		c.SetFlag(CPUFLAG_CARRY)
	} else {
		c.ClearAllFlags()
	}
	c.ReadCache >>= 1
	c.ReadCache |= carry
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
	c.SetZeroFlag(c.ReadCache)
}

func (c *CPU) instruction_BIT(reg *byte, bit int) {
	if ((*reg >> bit) & 0x01) == 0 {
		c.SetFlag(CPUFLAG_ZERO)
	} else {
		c.ClearFlag(CPUFLAG_ZERO)
	}
	c.SetFlag(CPUFLAG_HALF)
	c.ClearFlag(CPUFLAG_SUB)
}

func (c *CPU) instruction_BIT_HL(bit int) {
	if ((c.MemoryRead(c.HL.Get()) >> bit) & 0x01) == 0x00 {
		c.SetFlag(CPUFLAG_ZERO)
	} else {
		c.ClearFlag(CPUFLAG_ZERO)
	}
	c.SetFlag(CPUFLAG_HALF)
	c.ClearFlag(CPUFLAG_SUB)
}

func (c *CPU) instruction_SET(reg *byte, bit int) {
	*reg = (*reg | (0x01 << bit))
}

func (c *CPU) instruction_SET_HL(bit int) {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	c.ReadCache |= byte(0x01 << bit)
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
}

func (c *CPU) instruction_RES(reg *byte, bit int) {
	*reg = (*reg & (^(0x01 << bit)))
}

func (c *CPU) instruction_RES_HL(bit int) {
	if c.AccurateOpeCodeState == 1 {
		c.ReadCache = c.MemoryRead(c.HL.Get())
		return
	}
	c.ReadCache &= ^(0x01 << bit)
	c.MemoryWrite(c.HL.Get(), c.ReadCache)
}

func (c *CPU) instruction_StackPush(reg *CPURegister) {
	c.SP.Decrement()
	c.MemoryWrite(c.SP.Get(), reg.GetHigh())
	c.SP.Decrement()
	c.MemoryWrite(c.SP.Get(), reg.GetLow())
}

func (c *CPU) instruction_StackPop(reg *CPURegister) {
	reg.SetLow(c.MemoryRead(c.SP.Get()))
	c.SP.Increment()
	reg.SetHigh(c.MemoryRead(c.SP.Get()))
	c.SP.Increment()
}

func (c *CPU) instruction_InvalidOpcode() {
	log.Println("INVALID OP Code")
}

func (c CPU) adjustedCycles(cycles int) int {
	if cycles == 0 {
		return cycles
	}
	return cycles >> c.SpeedMultiplier
}
