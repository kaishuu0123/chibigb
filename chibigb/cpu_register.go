package chibigb

type CPURegister struct {
	high byte
	low  byte
}

func NewCPURegister() *CPURegister {
	return &CPURegister{
		high: 0x00,
		low:  0x00,
	}
}

func (cr *CPURegister) GetHigh() byte {
	return cr.high
}

func (cr *CPURegister) GetLow() byte {
	return cr.low
}

func (cr *CPURegister) GetHighPointer() *byte {
	return &cr.high
}

func (cr *CPURegister) GetLowPointer() *byte {
	return &cr.low
}

func (cr *CPURegister) SetHigh(value byte) {
	cr.high = value
}

func (cr *CPURegister) SetLow(value byte) {
	cr.low = value
}

func (cr *CPURegister) Get() uint16 {
	return (uint16(cr.high) << 8) | uint16(cr.low)
}

func (cr *CPURegister) Set(value uint16) {
	cr.high = byte(value >> 8)
	cr.low = byte(value & 0xFF)
}

func (cr *CPURegister) Increment() {
	cr.Set(cr.Get() + 1)
}

func (cr *CPURegister) Decrement() {
	cr.Set(cr.Get() - 1)
}
