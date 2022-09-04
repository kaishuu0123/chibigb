package chibigb

// This values are from the tests written by Shay Green
// (http://blargg.parodius.com/gb-tests/)

var OPCodeBranchMachineCycles = [256]byte{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1,
	1, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1,
	3, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1,
	3, 3, 2, 2, 3, 3, 3, 1, 3, 2, 2, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 2, 2, 2, 2, 2, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	5, 3, 4, 4, 6, 4, 2, 4, 5, 4, 4, 0, 6, 6, 2, 4,
	5, 3, 4, 0, 6, 4, 2, 4, 5, 4, 4, 0, 6, 0, 2, 4,
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4,
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4,
}

var OPCodeCBMachineCycles = [256]byte{
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
}

var OPCodeCBAccurate = [256]byte{
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0,
}

func (c *CPU) createCBTable() {
	c.OpCodeCBTable = [256]CPUOpCodeEntity{
		{opcode: 0x00, fn: c.opcodeCB0x00, name: "XXX"},
		{opcode: 0x01, fn: c.opcodeCB0x01, name: "XXX"},
		{opcode: 0x02, fn: c.opcodeCB0x02, name: "XXX"},
		{opcode: 0x03, fn: c.opcodeCB0x03, name: "XXX"},
		{opcode: 0x04, fn: c.opcodeCB0x04, name: "XXX"},
		{opcode: 0x05, fn: c.opcodeCB0x05, name: "XXX"},
		{opcode: 0x06, fn: c.opcodeCB0x06, name: "XXX"},
		{opcode: 0x07, fn: c.opcodeCB0x07, name: "XXX"},
		{opcode: 0x08, fn: c.opcodeCB0x08, name: "XXX"},
		{opcode: 0x09, fn: c.opcodeCB0x09, name: "XXX"},
		{opcode: 0x0A, fn: c.opcodeCB0x0A, name: "XXX"},
		{opcode: 0x0B, fn: c.opcodeCB0x0B, name: "XXX"},
		{opcode: 0x0C, fn: c.opcodeCB0x0C, name: "XXX"},
		{opcode: 0x0D, fn: c.opcodeCB0x0D, name: "XXX"},
		{opcode: 0x0E, fn: c.opcodeCB0x0E, name: "XXX"},
		{opcode: 0x0F, fn: c.opcodeCB0x0F, name: "XXX"},
		{opcode: 0x10, fn: c.opcodeCB0x10, name: "XXX"},
		{opcode: 0x11, fn: c.opcodeCB0x11, name: "XXX"},
		{opcode: 0x12, fn: c.opcodeCB0x12, name: "XXX"},
		{opcode: 0x13, fn: c.opcodeCB0x13, name: "XXX"},
		{opcode: 0x14, fn: c.opcodeCB0x14, name: "XXX"},
		{opcode: 0x15, fn: c.opcodeCB0x15, name: "XXX"},
		{opcode: 0x16, fn: c.opcodeCB0x16, name: "XXX"},
		{opcode: 0x17, fn: c.opcodeCB0x17, name: "XXX"},
		{opcode: 0x18, fn: c.opcodeCB0x18, name: "XXX"},
		{opcode: 0x19, fn: c.opcodeCB0x19, name: "XXX"},
		{opcode: 0x1A, fn: c.opcodeCB0x1A, name: "XXX"},
		{opcode: 0x1B, fn: c.opcodeCB0x1B, name: "XXX"},
		{opcode: 0x1C, fn: c.opcodeCB0x1C, name: "XXX"},
		{opcode: 0x1D, fn: c.opcodeCB0x1D, name: "XXX"},
		{opcode: 0x1E, fn: c.opcodeCB0x1E, name: "XXX"},
		{opcode: 0x1F, fn: c.opcodeCB0x1F, name: "XXX"},
		{opcode: 0x20, fn: c.opcodeCB0x20, name: "XXX"},
		{opcode: 0x21, fn: c.opcodeCB0x21, name: "XXX"},
		{opcode: 0x22, fn: c.opcodeCB0x22, name: "XXX"},
		{opcode: 0x23, fn: c.opcodeCB0x23, name: "XXX"},
		{opcode: 0x24, fn: c.opcodeCB0x24, name: "XXX"},
		{opcode: 0x25, fn: c.opcodeCB0x25, name: "XXX"},
		{opcode: 0x26, fn: c.opcodeCB0x26, name: "XXX"},
		{opcode: 0x27, fn: c.opcodeCB0x27, name: "XXX"},
		{opcode: 0x28, fn: c.opcodeCB0x28, name: "XXX"},
		{opcode: 0x29, fn: c.opcodeCB0x29, name: "XXX"},
		{opcode: 0x2A, fn: c.opcodeCB0x2A, name: "XXX"},
		{opcode: 0x2B, fn: c.opcodeCB0x2B, name: "XXX"},
		{opcode: 0x2C, fn: c.opcodeCB0x2C, name: "XXX"},
		{opcode: 0x2D, fn: c.opcodeCB0x2D, name: "XXX"},
		{opcode: 0x2E, fn: c.opcodeCB0x2E, name: "XXX"},
		{opcode: 0x2F, fn: c.opcodeCB0x2F, name: "XXX"},
		{opcode: 0x30, fn: c.opcodeCB0x30, name: "XXX"},
		{opcode: 0x31, fn: c.opcodeCB0x31, name: "XXX"},
		{opcode: 0x32, fn: c.opcodeCB0x32, name: "XXX"},
		{opcode: 0x33, fn: c.opcodeCB0x33, name: "XXX"},
		{opcode: 0x34, fn: c.opcodeCB0x34, name: "XXX"},
		{opcode: 0x35, fn: c.opcodeCB0x35, name: "XXX"},
		{opcode: 0x36, fn: c.opcodeCB0x36, name: "XXX"},
		{opcode: 0x37, fn: c.opcodeCB0x37, name: "XXX"},
		{opcode: 0x38, fn: c.opcodeCB0x38, name: "XXX"},
		{opcode: 0x39, fn: c.opcodeCB0x39, name: "XXX"},
		{opcode: 0x3A, fn: c.opcodeCB0x3A, name: "XXX"},
		{opcode: 0x3B, fn: c.opcodeCB0x3B, name: "XXX"},
		{opcode: 0x3C, fn: c.opcodeCB0x3C, name: "XXX"},
		{opcode: 0x3D, fn: c.opcodeCB0x3D, name: "XXX"},
		{opcode: 0x3E, fn: c.opcodeCB0x3E, name: "XXX"},
		{opcode: 0x3F, fn: c.opcodeCB0x3F, name: "XXX"},
		{opcode: 0x40, fn: c.opcodeCB0x40, name: "XXX"},
		{opcode: 0x41, fn: c.opcodeCB0x41, name: "XXX"},
		{opcode: 0x42, fn: c.opcodeCB0x42, name: "XXX"},
		{opcode: 0x43, fn: c.opcodeCB0x43, name: "XXX"},
		{opcode: 0x44, fn: c.opcodeCB0x44, name: "XXX"},
		{opcode: 0x45, fn: c.opcodeCB0x45, name: "XXX"},
		{opcode: 0x46, fn: c.opcodeCB0x46, name: "XXX"},
		{opcode: 0x47, fn: c.opcodeCB0x47, name: "XXX"},
		{opcode: 0x48, fn: c.opcodeCB0x48, name: "XXX"},
		{opcode: 0x49, fn: c.opcodeCB0x49, name: "XXX"},
		{opcode: 0x4A, fn: c.opcodeCB0x4A, name: "XXX"},
		{opcode: 0x4B, fn: c.opcodeCB0x4B, name: "XXX"},
		{opcode: 0x4C, fn: c.opcodeCB0x4C, name: "XXX"},
		{opcode: 0x4D, fn: c.opcodeCB0x4D, name: "XXX"},
		{opcode: 0x4E, fn: c.opcodeCB0x4E, name: "XXX"},
		{opcode: 0x4F, fn: c.opcodeCB0x4F, name: "XXX"},
		{opcode: 0x50, fn: c.opcodeCB0x50, name: "XXX"},
		{opcode: 0x51, fn: c.opcodeCB0x51, name: "XXX"},
		{opcode: 0x52, fn: c.opcodeCB0x52, name: "XXX"},
		{opcode: 0x53, fn: c.opcodeCB0x53, name: "XXX"},
		{opcode: 0x54, fn: c.opcodeCB0x54, name: "XXX"},
		{opcode: 0x55, fn: c.opcodeCB0x55, name: "XXX"},
		{opcode: 0x56, fn: c.opcodeCB0x56, name: "XXX"},
		{opcode: 0x57, fn: c.opcodeCB0x57, name: "XXX"},
		{opcode: 0x58, fn: c.opcodeCB0x58, name: "XXX"},
		{opcode: 0x59, fn: c.opcodeCB0x59, name: "XXX"},
		{opcode: 0x5A, fn: c.opcodeCB0x5A, name: "XXX"},
		{opcode: 0x5B, fn: c.opcodeCB0x5B, name: "XXX"},
		{opcode: 0x5C, fn: c.opcodeCB0x5C, name: "XXX"},
		{opcode: 0x5D, fn: c.opcodeCB0x5D, name: "XXX"},
		{opcode: 0x5E, fn: c.opcodeCB0x5E, name: "XXX"},
		{opcode: 0x5F, fn: c.opcodeCB0x5F, name: "XXX"},
		{opcode: 0x60, fn: c.opcodeCB0x60, name: "XXX"},
		{opcode: 0x61, fn: c.opcodeCB0x61, name: "XXX"},
		{opcode: 0x62, fn: c.opcodeCB0x62, name: "XXX"},
		{opcode: 0x63, fn: c.opcodeCB0x63, name: "XXX"},
		{opcode: 0x64, fn: c.opcodeCB0x64, name: "XXX"},
		{opcode: 0x65, fn: c.opcodeCB0x65, name: "XXX"},
		{opcode: 0x66, fn: c.opcodeCB0x66, name: "XXX"},
		{opcode: 0x67, fn: c.opcodeCB0x67, name: "XXX"},
		{opcode: 0x68, fn: c.opcodeCB0x68, name: "XXX"},
		{opcode: 0x69, fn: c.opcodeCB0x69, name: "XXX"},
		{opcode: 0x6A, fn: c.opcodeCB0x6A, name: "XXX"},
		{opcode: 0x6B, fn: c.opcodeCB0x6B, name: "XXX"},
		{opcode: 0x6C, fn: c.opcodeCB0x6C, name: "XXX"},
		{opcode: 0x6D, fn: c.opcodeCB0x6D, name: "XXX"},
		{opcode: 0x6E, fn: c.opcodeCB0x6E, name: "XXX"},
		{opcode: 0x6F, fn: c.opcodeCB0x6F, name: "XXX"},
		{opcode: 0x70, fn: c.opcodeCB0x70, name: "XXX"},
		{opcode: 0x71, fn: c.opcodeCB0x71, name: "XXX"},
		{opcode: 0x72, fn: c.opcodeCB0x72, name: "XXX"},
		{opcode: 0x73, fn: c.opcodeCB0x73, name: "XXX"},
		{opcode: 0x74, fn: c.opcodeCB0x74, name: "XXX"},
		{opcode: 0x75, fn: c.opcodeCB0x75, name: "XXX"},
		{opcode: 0x76, fn: c.opcodeCB0x76, name: "XXX"},
		{opcode: 0x77, fn: c.opcodeCB0x77, name: "XXX"},
		{opcode: 0x78, fn: c.opcodeCB0x78, name: "XXX"},
		{opcode: 0x79, fn: c.opcodeCB0x79, name: "XXX"},
		{opcode: 0x7A, fn: c.opcodeCB0x7A, name: "XXX"},
		{opcode: 0x7B, fn: c.opcodeCB0x7B, name: "XXX"},
		{opcode: 0x7C, fn: c.opcodeCB0x7C, name: "XXX"},
		{opcode: 0x7D, fn: c.opcodeCB0x7D, name: "XXX"},
		{opcode: 0x7E, fn: c.opcodeCB0x7E, name: "XXX"},
		{opcode: 0x7F, fn: c.opcodeCB0x7F, name: "XXX"},
		{opcode: 0x80, fn: c.opcodeCB0x80, name: "XXX"},
		{opcode: 0x81, fn: c.opcodeCB0x81, name: "XXX"},
		{opcode: 0x82, fn: c.opcodeCB0x82, name: "XXX"},
		{opcode: 0x83, fn: c.opcodeCB0x83, name: "XXX"},
		{opcode: 0x84, fn: c.opcodeCB0x84, name: "XXX"},
		{opcode: 0x85, fn: c.opcodeCB0x85, name: "XXX"},
		{opcode: 0x86, fn: c.opcodeCB0x86, name: "XXX"},
		{opcode: 0x87, fn: c.opcodeCB0x87, name: "XXX"},
		{opcode: 0x88, fn: c.opcodeCB0x88, name: "XXX"},
		{opcode: 0x89, fn: c.opcodeCB0x89, name: "XXX"},
		{opcode: 0x8A, fn: c.opcodeCB0x8A, name: "XXX"},
		{opcode: 0x8B, fn: c.opcodeCB0x8B, name: "XXX"},
		{opcode: 0x8C, fn: c.opcodeCB0x8C, name: "XXX"},
		{opcode: 0x8D, fn: c.opcodeCB0x8D, name: "XXX"},
		{opcode: 0x8E, fn: c.opcodeCB0x8E, name: "XXX"},
		{opcode: 0x8F, fn: c.opcodeCB0x8F, name: "XXX"},
		{opcode: 0x90, fn: c.opcodeCB0x90, name: "XXX"},
		{opcode: 0x91, fn: c.opcodeCB0x91, name: "XXX"},
		{opcode: 0x92, fn: c.opcodeCB0x92, name: "XXX"},
		{opcode: 0x93, fn: c.opcodeCB0x93, name: "XXX"},
		{opcode: 0x94, fn: c.opcodeCB0x94, name: "XXX"},
		{opcode: 0x95, fn: c.opcodeCB0x95, name: "XXX"},
		{opcode: 0x96, fn: c.opcodeCB0x96, name: "XXX"},
		{opcode: 0x97, fn: c.opcodeCB0x97, name: "XXX"},
		{opcode: 0x98, fn: c.opcodeCB0x98, name: "XXX"},
		{opcode: 0x99, fn: c.opcodeCB0x99, name: "XXX"},
		{opcode: 0x9A, fn: c.opcodeCB0x9A, name: "XXX"},
		{opcode: 0x9B, fn: c.opcodeCB0x9B, name: "XXX"},
		{opcode: 0x9C, fn: c.opcodeCB0x9C, name: "XXX"},
		{opcode: 0x9D, fn: c.opcodeCB0x9D, name: "XXX"},
		{opcode: 0x9E, fn: c.opcodeCB0x9E, name: "XXX"},
		{opcode: 0x9F, fn: c.opcodeCB0x9F, name: "XXX"},
		{opcode: 0xA0, fn: c.opcodeCB0xA0, name: "XXX"},
		{opcode: 0xA1, fn: c.opcodeCB0xA1, name: "XXX"},
		{opcode: 0xA2, fn: c.opcodeCB0xA2, name: "XXX"},
		{opcode: 0xA3, fn: c.opcodeCB0xA3, name: "XXX"},
		{opcode: 0xA4, fn: c.opcodeCB0xA4, name: "XXX"},
		{opcode: 0xA5, fn: c.opcodeCB0xA5, name: "XXX"},
		{opcode: 0xA6, fn: c.opcodeCB0xA6, name: "XXX"},
		{opcode: 0xA7, fn: c.opcodeCB0xA7, name: "XXX"},
		{opcode: 0xA8, fn: c.opcodeCB0xA8, name: "XXX"},
		{opcode: 0xA9, fn: c.opcodeCB0xA9, name: "XXX"},
		{opcode: 0xAA, fn: c.opcodeCB0xAA, name: "XXX"},
		{opcode: 0xAB, fn: c.opcodeCB0xAB, name: "XXX"},
		{opcode: 0xAC, fn: c.opcodeCB0xAC, name: "XXX"},
		{opcode: 0xAD, fn: c.opcodeCB0xAD, name: "XXX"},
		{opcode: 0xAE, fn: c.opcodeCB0xAE, name: "XXX"},
		{opcode: 0xAF, fn: c.opcodeCB0xAF, name: "XXX"},
		{opcode: 0xB0, fn: c.opcodeCB0xB0, name: "XXX"},
		{opcode: 0xB1, fn: c.opcodeCB0xB1, name: "XXX"},
		{opcode: 0xB2, fn: c.opcodeCB0xB2, name: "XXX"},
		{opcode: 0xB3, fn: c.opcodeCB0xB3, name: "XXX"},
		{opcode: 0xB4, fn: c.opcodeCB0xB4, name: "XXX"},
		{opcode: 0xB5, fn: c.opcodeCB0xB5, name: "XXX"},
		{opcode: 0xB6, fn: c.opcodeCB0xB6, name: "XXX"},
		{opcode: 0xB7, fn: c.opcodeCB0xB7, name: "XXX"},
		{opcode: 0xB8, fn: c.opcodeCB0xB8, name: "XXX"},
		{opcode: 0xB9, fn: c.opcodeCB0xB9, name: "XXX"},
		{opcode: 0xBA, fn: c.opcodeCB0xBA, name: "XXX"},
		{opcode: 0xBB, fn: c.opcodeCB0xBB, name: "XXX"},
		{opcode: 0xBC, fn: c.opcodeCB0xBC, name: "XXX"},
		{opcode: 0xBD, fn: c.opcodeCB0xBD, name: "XXX"},
		{opcode: 0xBE, fn: c.opcodeCB0xBE, name: "XXX"},
		{opcode: 0xBF, fn: c.opcodeCB0xBF, name: "XXX"},
		{opcode: 0xC0, fn: c.opcodeCB0xC0, name: "XXX"},
		{opcode: 0xC1, fn: c.opcodeCB0xC1, name: "XXX"},
		{opcode: 0xC2, fn: c.opcodeCB0xC2, name: "XXX"},
		{opcode: 0xC3, fn: c.opcodeCB0xC3, name: "XXX"},
		{opcode: 0xC4, fn: c.opcodeCB0xC4, name: "XXX"},
		{opcode: 0xC5, fn: c.opcodeCB0xC5, name: "XXX"},
		{opcode: 0xC6, fn: c.opcodeCB0xC6, name: "XXX"},
		{opcode: 0xC7, fn: c.opcodeCB0xC7, name: "XXX"},
		{opcode: 0xC8, fn: c.opcodeCB0xC8, name: "XXX"},
		{opcode: 0xC9, fn: c.opcodeCB0xC9, name: "XXX"},
		{opcode: 0xCA, fn: c.opcodeCB0xCA, name: "XXX"},
		{opcode: 0xCB, fn: c.opcodeCB0xCB, name: "XXX"},
		{opcode: 0xCC, fn: c.opcodeCB0xCC, name: "XXX"},
		{opcode: 0xCD, fn: c.opcodeCB0xCD, name: "XXX"},
		{opcode: 0xCE, fn: c.opcodeCB0xCE, name: "XXX"},
		{opcode: 0xCF, fn: c.opcodeCB0xCF, name: "XXX"},
		{opcode: 0xD0, fn: c.opcodeCB0xD0, name: "XXX"},
		{opcode: 0xD1, fn: c.opcodeCB0xD1, name: "XXX"},
		{opcode: 0xD2, fn: c.opcodeCB0xD2, name: "XXX"},
		{opcode: 0xD3, fn: c.opcodeCB0xD3, name: "XXX"},
		{opcode: 0xD4, fn: c.opcodeCB0xD4, name: "XXX"},
		{opcode: 0xD5, fn: c.opcodeCB0xD5, name: "XXX"},
		{opcode: 0xD6, fn: c.opcodeCB0xD6, name: "XXX"},
		{opcode: 0xD7, fn: c.opcodeCB0xD7, name: "XXX"},
		{opcode: 0xD8, fn: c.opcodeCB0xD8, name: "XXX"},
		{opcode: 0xD9, fn: c.opcodeCB0xD9, name: "XXX"},
		{opcode: 0xDA, fn: c.opcodeCB0xDA, name: "XXX"},
		{opcode: 0xDB, fn: c.opcodeCB0xDB, name: "XXX"},
		{opcode: 0xDC, fn: c.opcodeCB0xDC, name: "XXX"},
		{opcode: 0xDD, fn: c.opcodeCB0xDD, name: "XXX"},
		{opcode: 0xDE, fn: c.opcodeCB0xDE, name: "XXX"},
		{opcode: 0xDF, fn: c.opcodeCB0xDF, name: "XXX"},
		{opcode: 0xE0, fn: c.opcodeCB0xE0, name: "XXX"},
		{opcode: 0xE1, fn: c.opcodeCB0xE1, name: "XXX"},
		{opcode: 0xE2, fn: c.opcodeCB0xE2, name: "XXX"},
		{opcode: 0xE3, fn: c.opcodeCB0xE3, name: "XXX"},
		{opcode: 0xE4, fn: c.opcodeCB0xE4, name: "XXX"},
		{opcode: 0xE5, fn: c.opcodeCB0xE5, name: "XXX"},
		{opcode: 0xE6, fn: c.opcodeCB0xE6, name: "XXX"},
		{opcode: 0xE7, fn: c.opcodeCB0xE7, name: "XXX"},
		{opcode: 0xE8, fn: c.opcodeCB0xE8, name: "XXX"},
		{opcode: 0xE9, fn: c.opcodeCB0xE9, name: "XXX"},
		{opcode: 0xEA, fn: c.opcodeCB0xEA, name: "XXX"},
		{opcode: 0xEB, fn: c.opcodeCB0xEB, name: "XXX"},
		{opcode: 0xEC, fn: c.opcodeCB0xEC, name: "XXX"},
		{opcode: 0xED, fn: c.opcodeCB0xED, name: "XXX"},
		{opcode: 0xEE, fn: c.opcodeCB0xEE, name: "XXX"},
		{opcode: 0xEF, fn: c.opcodeCB0xEF, name: "XXX"},
		{opcode: 0xF0, fn: c.opcodeCB0xF0, name: "XXX"},
		{opcode: 0xF1, fn: c.opcodeCB0xF1, name: "XXX"},
		{opcode: 0xF2, fn: c.opcodeCB0xF2, name: "XXX"},
		{opcode: 0xF3, fn: c.opcodeCB0xF3, name: "XXX"},
		{opcode: 0xF4, fn: c.opcodeCB0xF4, name: "XXX"},
		{opcode: 0xF5, fn: c.opcodeCB0xF5, name: "XXX"},
		{opcode: 0xF6, fn: c.opcodeCB0xF6, name: "XXX"},
		{opcode: 0xF7, fn: c.opcodeCB0xF7, name: "XXX"},
		{opcode: 0xF8, fn: c.opcodeCB0xF8, name: "XXX"},
		{opcode: 0xF9, fn: c.opcodeCB0xF9, name: "XXX"},
		{opcode: 0xFA, fn: c.opcodeCB0xFA, name: "XXX"},
		{opcode: 0xFB, fn: c.opcodeCB0xFB, name: "XXX"},
		{opcode: 0xFC, fn: c.opcodeCB0xFC, name: "XXX"},
		{opcode: 0xFD, fn: c.opcodeCB0xFD, name: "XXX"},
		{opcode: 0xFE, fn: c.opcodeCB0xFE, name: "XXX"},
		{opcode: 0xFF, fn: c.opcodeCB0xFF, name: "XXX"},
	}
}

// RLC B
func (c *CPU) opcodeCB0x00() {
	c.instruction_RLC(c.BC.GetHighPointer(), false)
}

// RLC C
func (c *CPU) opcodeCB0x01() {
	c.instruction_RLC(c.BC.GetLowPointer(), false)
}

// RLC D
func (c *CPU) opcodeCB0x02() {
	c.instruction_RLC(c.DE.GetHighPointer(), false)
}

// RLC E
func (c *CPU) opcodeCB0x03() {
	c.instruction_RLC(c.DE.GetLowPointer(), false)
}

// RLC H
func (c *CPU) opcodeCB0x04() {
	c.instruction_RLC(c.HL.GetHighPointer(), false)
}

// RLC L
func (c *CPU) opcodeCB0x05() {
	c.instruction_RLC(c.HL.GetLowPointer(), false)
}

// RLC (HL)
func (c *CPU) opcodeCB0x06() {
	c.instruction_RLC_HL()
}

// RLC A
func (c *CPU) opcodeCB0x07() {
	c.instruction_RLC(c.AF.GetHighPointer(), false)
}

// RRC B
func (c *CPU) opcodeCB0x08() {
	c.instruction_RRC(c.BC.GetHighPointer(), false)
}

// RRC C
func (c *CPU) opcodeCB0x09() {
	c.instruction_RRC(c.BC.GetLowPointer(), false)
}

// RRC D
func (c *CPU) opcodeCB0x0A() {
	c.instruction_RRC(c.DE.GetHighPointer(), false)
}

// RRC E
func (c *CPU) opcodeCB0x0B() {
	c.instruction_RRC(c.DE.GetLowPointer(), false)
}

// RRC H
func (c *CPU) opcodeCB0x0C() {
	c.instruction_RRC(c.HL.GetHighPointer(), false)
}

// RRC L
func (c *CPU) opcodeCB0x0D() {
	c.instruction_RRC(c.HL.GetLowPointer(), false)
}

// RRC (HL)
func (c *CPU) opcodeCB0x0E() {
	c.instruction_RRC_HL()
}

// RRC A
func (c *CPU) opcodeCB0x0F() {
	c.instruction_RRC(c.AF.GetHighPointer(), false)
}

// RL B
func (c *CPU) opcodeCB0x10() {
	c.instruction_RL(c.BC.GetHighPointer(), false)
}

// RL C
func (c *CPU) opcodeCB0x11() {
	c.instruction_RL(c.BC.GetLowPointer(), false)
}

// RL D
func (c *CPU) opcodeCB0x12() {
	c.instruction_RL(c.DE.GetHighPointer(), false)
}

// RL E
func (c *CPU) opcodeCB0x13() {
	c.instruction_RL(c.DE.GetLowPointer(), false)
}

// RL H
func (c *CPU) opcodeCB0x14() {
	c.instruction_RL(c.HL.GetHighPointer(), false)
}

// RL L
func (c *CPU) opcodeCB0x15() {
	c.instruction_RL(c.HL.GetLowPointer(), false)
}

// RL (HL)
func (c *CPU) opcodeCB0x16() {
	c.instruction_RL_HL()
}

// RL A
func (c *CPU) opcodeCB0x17() {
	c.instruction_RL(c.AF.GetHighPointer(), false)
}

// RR B
func (c *CPU) opcodeCB0x18() {
	c.instruction_RR(c.BC.GetHighPointer(), false)
}

// RR C
func (c *CPU) opcodeCB0x19() {
	c.instruction_RR(c.BC.GetLowPointer(), false)
}

// RR D
func (c *CPU) opcodeCB0x1A() {
	c.instruction_RR(c.DE.GetHighPointer(), false)
}

// RR E
func (c *CPU) opcodeCB0x1B() {
	c.instruction_RR(c.DE.GetLowPointer(), false)
}

// RR H
func (c *CPU) opcodeCB0x1C() {
	c.instruction_RR(c.HL.GetHighPointer(), false)
}

// RR L
func (c *CPU) opcodeCB0x1D() {
	c.instruction_RR(c.HL.GetLowPointer(), false)
}

// RR (HL)
func (c *CPU) opcodeCB0x1E() {
	c.instruction_RR_HL()
}

// RR A
func (c *CPU) opcodeCB0x1F() {
	c.instruction_RR(c.AF.GetHighPointer(), false)
}

// SLA B
func (c *CPU) opcodeCB0x20() {
	c.instruction_SLA(c.BC.GetHighPointer())
}

// SLA C
func (c *CPU) opcodeCB0x21() {
	c.instruction_SLA(c.BC.GetLowPointer())
}

// SLA D
func (c *CPU) opcodeCB0x22() {
	c.instruction_SLA(c.DE.GetHighPointer())
}

// SLA E
func (c *CPU) opcodeCB0x23() {
	c.instruction_SLA(c.DE.GetLowPointer())
}

// SLA H
func (c *CPU) opcodeCB0x24() {
	c.instruction_SLA(c.HL.GetHighPointer())
}

// SLA L
func (c *CPU) opcodeCB0x25() {
	c.instruction_SLA(c.HL.GetLowPointer())
}

// SLA (HL)
func (c *CPU) opcodeCB0x26() {
	c.instruction_SLA_HL()
}

// SLA A
func (c *CPU) opcodeCB0x27() {
	c.instruction_SLA(c.AF.GetHighPointer())
}

// SRA B
func (c *CPU) opcodeCB0x28() {
	c.instruction_SRA(c.BC.GetHighPointer())
}

// SRA C
func (c *CPU) opcodeCB0x29() {
	c.instruction_SRA(c.BC.GetLowPointer())
}

// SRA D
func (c *CPU) opcodeCB0x2A() {
	c.instruction_SRA(c.DE.GetHighPointer())
}

// SRA E
func (c *CPU) opcodeCB0x2B() {
	c.instruction_SRA(c.DE.GetLowPointer())
}

// SRA H
func (c *CPU) opcodeCB0x2C() {
	c.instruction_SRA(c.HL.GetHighPointer())
}

// SRA L
func (c *CPU) opcodeCB0x2D() {
	c.instruction_SRA(c.HL.GetLowPointer())
}

// SRA (HL)
func (c *CPU) opcodeCB0x2E() {
	c.instruction_SRA_HL()
}

// SRA A
func (c *CPU) opcodeCB0x2F() {
	c.instruction_SRA(c.AF.GetHighPointer())
}

// SWAP B
func (c *CPU) opcodeCB0x30() {
	c.instruction_SWAP_Register(c.BC.GetHighPointer())
}

// SWAP C
func (c *CPU) opcodeCB0x31() {
	c.instruction_SWAP_Register(c.BC.GetLowPointer())
}

// SWAP D
func (c *CPU) opcodeCB0x32() {
	c.instruction_SWAP_Register(c.DE.GetHighPointer())
}

// SWAP E
func (c *CPU) opcodeCB0x33() {
	c.instruction_SWAP_Register(c.DE.GetLowPointer())
}

// SWAP H
func (c *CPU) opcodeCB0x34() {
	c.instruction_SWAP_Register(c.HL.GetHighPointer())
}

// SWAP L
func (c *CPU) opcodeCB0x35() {
	c.instruction_SWAP_Register(c.HL.GetLowPointer())
}

// SWAP (HL)
func (c *CPU) opcodeCB0x36() {
	c.instruction_SWAP_HL()
}

// SWAP A
func (c *CPU) opcodeCB0x37() {
	c.instruction_SWAP_Register(c.AF.GetHighPointer())
}

// SRL B
func (c *CPU) opcodeCB0x38() {
	c.instruction_SRL(c.BC.GetHighPointer())
}

// SRL C
func (c *CPU) opcodeCB0x39() {
	c.instruction_SRL(c.BC.GetLowPointer())
}

// SRL D
func (c *CPU) opcodeCB0x3A() {
	c.instruction_SRL(c.DE.GetHighPointer())
}

// SRL E
func (c *CPU) opcodeCB0x3B() {
	c.instruction_SRL(c.DE.GetLowPointer())
}

// SRL H
func (c *CPU) opcodeCB0x3C() {
	c.instruction_SRL(c.HL.GetHighPointer())
}

// SRL L
func (c *CPU) opcodeCB0x3D() {
	c.instruction_SRL(c.HL.GetLowPointer())
}

// SRL (HL)
func (c *CPU) opcodeCB0x3E() {
	c.instruction_SRL_HL()
}

// SRL A
func (c *CPU) opcodeCB0x3F() {
	c.instruction_SRL(c.AF.GetHighPointer())
}

// BIT 0 B
func (c *CPU) opcodeCB0x40() {
	c.instruction_BIT(c.BC.GetHighPointer(), 0)
}

// BIT 0 C
func (c *CPU) opcodeCB0x41() {
	c.instruction_BIT(c.BC.GetLowPointer(), 0)
}

// BIT 0 D
func (c *CPU) opcodeCB0x42() {
	c.instruction_BIT(c.DE.GetHighPointer(), 0)
}

// BIT 0 E
func (c *CPU) opcodeCB0x43() {
	c.instruction_BIT(c.DE.GetLowPointer(), 0)
}

// BIT 0 H
func (c *CPU) opcodeCB0x44() {
	c.instruction_BIT(c.HL.GetHighPointer(), 0)
}

// BIT 0 L
func (c *CPU) opcodeCB0x45() {
	c.instruction_BIT(c.HL.GetLowPointer(), 0)
}

// BIT 0 (HL)
func (c *CPU) opcodeCB0x46() {
	c.instruction_BIT_HL(0)
}

// BIT 0 A
func (c *CPU) opcodeCB0x47() {
	c.instruction_BIT(c.AF.GetHighPointer(), 0)
}

// BIT 1 B
func (c *CPU) opcodeCB0x48() {
	c.instruction_BIT(c.BC.GetHighPointer(), 1)
}

// BIT 1 C
func (c *CPU) opcodeCB0x49() {
	c.instruction_BIT(c.BC.GetLowPointer(), 1)
}

// BIT 1 D
func (c *CPU) opcodeCB0x4A() {
	c.instruction_BIT(c.DE.GetHighPointer(), 1)
}

// BIT 1 E
func (c *CPU) opcodeCB0x4B() {
	c.instruction_BIT(c.DE.GetLowPointer(), 1)
}

// BIT 1 H
func (c *CPU) opcodeCB0x4C() {
	c.instruction_BIT(c.HL.GetHighPointer(), 1)
}

// BIT 1 L
func (c *CPU) opcodeCB0x4D() {
	c.instruction_BIT(c.HL.GetLowPointer(), 1)
}

// BIT 1 (HL)
func (c *CPU) opcodeCB0x4E() {
	c.instruction_BIT_HL(1)
}

// BIT 1 A
func (c *CPU) opcodeCB0x4F() {
	c.instruction_BIT(c.AF.GetHighPointer(), 1)
}

// BIT 2 B
func (c *CPU) opcodeCB0x50() {
	c.instruction_BIT(c.BC.GetHighPointer(), 2)
}

// BIT 2 C
func (c *CPU) opcodeCB0x51() {
	c.instruction_BIT(c.BC.GetLowPointer(), 2)
}

// BIT 2 D
func (c *CPU) opcodeCB0x52() {
	c.instruction_BIT(c.DE.GetHighPointer(), 2)
}

// BIT 2 E
func (c *CPU) opcodeCB0x53() {
	c.instruction_BIT(c.DE.GetLowPointer(), 2)
}

// BIT 2 H
func (c *CPU) opcodeCB0x54() {
	c.instruction_BIT(c.HL.GetHighPointer(), 2)
}

// BIT 2 L
func (c *CPU) opcodeCB0x55() {
	c.instruction_BIT(c.HL.GetLowPointer(), 2)
}

// BIT 2 (HL)
func (c *CPU) opcodeCB0x56() {
	c.instruction_BIT_HL(2)
}

// BIT 2 A
func (c *CPU) opcodeCB0x57() {
	c.instruction_BIT(c.AF.GetHighPointer(), 2)
}

// BIT 3 B
func (c *CPU) opcodeCB0x58() {
	c.instruction_BIT(c.BC.GetHighPointer(), 3)
}

// BIT 3 C
func (c *CPU) opcodeCB0x59() {
	c.instruction_BIT(c.BC.GetLowPointer(), 3)
}

// BIT 3 D
func (c *CPU) opcodeCB0x5A() {
	c.instruction_BIT(c.DE.GetHighPointer(), 3)
}

// BIT 3 E
func (c *CPU) opcodeCB0x5B() {
	c.instruction_BIT(c.DE.GetLowPointer(), 3)
}

// BIT 3 H
func (c *CPU) opcodeCB0x5C() {
	c.instruction_BIT(c.HL.GetHighPointer(), 3)
}

// BIT 3 L
func (c *CPU) opcodeCB0x5D() {
	c.instruction_BIT(c.HL.GetLowPointer(), 3)
}

// BIT 3 (HL)
func (c *CPU) opcodeCB0x5E() {
	c.instruction_BIT_HL(3)
}

// BIT 3 A
func (c *CPU) opcodeCB0x5F() {
	c.instruction_BIT(c.AF.GetHighPointer(), 3)
}

// BIT 4 B
func (c *CPU) opcodeCB0x60() {
	c.instruction_BIT(c.BC.GetHighPointer(), 4)
}

// BIT 4 C
func (c *CPU) opcodeCB0x61() {
	c.instruction_BIT(c.BC.GetLowPointer(), 4)
}

// BIT 4 D
func (c *CPU) opcodeCB0x62() {
	c.instruction_BIT(c.DE.GetHighPointer(), 4)
}

// BIT 4 E
func (c *CPU) opcodeCB0x63() {
	c.instruction_BIT(c.DE.GetLowPointer(), 4)
}

// BIT 4 H
func (c *CPU) opcodeCB0x64() {
	c.instruction_BIT(c.HL.GetHighPointer(), 4)
}

// BIT 4 L
func (c *CPU) opcodeCB0x65() {
	c.instruction_BIT(c.HL.GetLowPointer(), 4)
}

// BIT 4 (HL)
func (c *CPU) opcodeCB0x66() {
	c.instruction_BIT_HL(4)
}

// BIT 4 A
func (c *CPU) opcodeCB0x67() {
	c.instruction_BIT(c.AF.GetHighPointer(), 4)
}

// BIT 5 B
func (c *CPU) opcodeCB0x68() {
	c.instruction_BIT(c.BC.GetHighPointer(), 5)
}

// BIT 5 C
func (c *CPU) opcodeCB0x69() {
	c.instruction_BIT(c.BC.GetLowPointer(), 5)
}

// BIT 5 D
func (c *CPU) opcodeCB0x6A() {
	c.instruction_BIT(c.DE.GetHighPointer(), 5)
}

// BIT 5 E
func (c *CPU) opcodeCB0x6B() {
	c.instruction_BIT(c.DE.GetLowPointer(), 5)
}

// BIT 5 H
func (c *CPU) opcodeCB0x6C() {
	c.instruction_BIT(c.HL.GetHighPointer(), 5)
}

// BIT 5 L
func (c *CPU) opcodeCB0x6D() {
	c.instruction_BIT(c.HL.GetLowPointer(), 5)
}

// BIT 5 (HL)
func (c *CPU) opcodeCB0x6E() {
	c.instruction_BIT_HL(5)
}

// BIT 5 A
func (c *CPU) opcodeCB0x6F() {
	c.instruction_BIT(c.AF.GetHighPointer(), 5)
}

// BIT 6 B
func (c *CPU) opcodeCB0x70() {
	c.instruction_BIT(c.BC.GetHighPointer(), 6)
}

// BIT 6 C
func (c *CPU) opcodeCB0x71() {
	c.instruction_BIT(c.BC.GetLowPointer(), 6)
}

// BIT 6 D
func (c *CPU) opcodeCB0x72() {
	c.instruction_BIT(c.DE.GetHighPointer(), 6)
}

// BIT 6 E
func (c *CPU) opcodeCB0x73() {
	c.instruction_BIT(c.DE.GetLowPointer(), 6)
}

// BIT 6 H
func (c *CPU) opcodeCB0x74() {
	c.instruction_BIT(c.HL.GetHighPointer(), 6)
}

// BIT 6 L
func (c *CPU) opcodeCB0x75() {
	c.instruction_BIT(c.HL.GetLowPointer(), 6)
}

// BIT 6 (HL)
func (c *CPU) opcodeCB0x76() {
	c.instruction_BIT_HL(6)
}

// BIT 6 A
func (c *CPU) opcodeCB0x77() {
	c.instruction_BIT(c.AF.GetHighPointer(), 6)
}

// BIT 7 B
func (c *CPU) opcodeCB0x78() {
	c.instruction_BIT(c.BC.GetHighPointer(), 7)
}

// BIT 7 C
func (c *CPU) opcodeCB0x79() {
	c.instruction_BIT(c.BC.GetLowPointer(), 7)
}

// BIT 7 D
func (c *CPU) opcodeCB0x7A() {
	c.instruction_BIT(c.DE.GetHighPointer(), 7)
}

// BIT 7 E
func (c *CPU) opcodeCB0x7B() {
	c.instruction_BIT(c.DE.GetLowPointer(), 7)
}

// BIT 7 H
func (c *CPU) opcodeCB0x7C() {
	c.instruction_BIT(c.HL.GetHighPointer(), 7)
}

// BIT 7 L
func (c *CPU) opcodeCB0x7D() {
	c.instruction_BIT(c.HL.GetLowPointer(), 7)
}

// BIT 7 (HL)
func (c *CPU) opcodeCB0x7E() {
	c.instruction_BIT_HL(7)
}

// BIT 7 A
func (c *CPU) opcodeCB0x7F() {
	c.instruction_BIT(c.AF.GetHighPointer(), 7)
}

// RES 0 B
func (c *CPU) opcodeCB0x80() {
	c.instruction_RES(c.BC.GetHighPointer(), 0)
}

// RES 0 C
func (c *CPU) opcodeCB0x81() {
	c.instruction_RES(c.BC.GetLowPointer(), 0)
}

// RES 0 D
func (c *CPU) opcodeCB0x82() {
	c.instruction_RES(c.DE.GetHighPointer(), 0)
}

// RES 0 E
func (c *CPU) opcodeCB0x83() {
	c.instruction_RES(c.DE.GetLowPointer(), 0)
}

// RES 0 H
func (c *CPU) opcodeCB0x84() {
	c.instruction_RES(c.HL.GetHighPointer(), 0)
}

// RES 0 L
func (c *CPU) opcodeCB0x85() {
	c.instruction_RES(c.HL.GetLowPointer(), 0)
}

// RES 0 (HL)
func (c *CPU) opcodeCB0x86() {
	c.instruction_RES_HL(0)
}

// RES 0 A
func (c *CPU) opcodeCB0x87() {
	c.instruction_RES(c.AF.GetHighPointer(), 0)
}

// RES 1 B
func (c *CPU) opcodeCB0x88() {
	c.instruction_RES(c.BC.GetHighPointer(), 1)
}

// RES 1 C
func (c *CPU) opcodeCB0x89() {
	c.instruction_RES(c.BC.GetLowPointer(), 1)
}

// RES 1 D
func (c *CPU) opcodeCB0x8A() {
	c.instruction_RES(c.DE.GetHighPointer(), 1)
}

// RES 1 E
func (c *CPU) opcodeCB0x8B() {
	c.instruction_RES(c.DE.GetLowPointer(), 1)
}

// RES 1 H
func (c *CPU) opcodeCB0x8C() {
	c.instruction_RES(c.HL.GetHighPointer(), 1)
}

// RES 1 L
func (c *CPU) opcodeCB0x8D() {
	c.instruction_RES(c.HL.GetLowPointer(), 1)
}

// RES 1 (HL)
func (c *CPU) opcodeCB0x8E() {
	c.instruction_RES_HL(1)
}

// RES 1 A
func (c *CPU) opcodeCB0x8F() {
	c.instruction_RES(c.AF.GetHighPointer(), 1)
}

// RES 2 B
func (c *CPU) opcodeCB0x90() {
	c.instruction_RES(c.BC.GetHighPointer(), 2)
}

// RES 2 C
func (c *CPU) opcodeCB0x91() {
	c.instruction_RES(c.BC.GetLowPointer(), 2)
}

// RES 2 D
func (c *CPU) opcodeCB0x92() {
	c.instruction_RES(c.DE.GetHighPointer(), 2)
}

// RES 2 E
func (c *CPU) opcodeCB0x93() {
	c.instruction_RES(c.DE.GetLowPointer(), 2)
}

// RES 2 H
func (c *CPU) opcodeCB0x94() {
	c.instruction_RES(c.HL.GetHighPointer(), 2)
}

// RES 2 L
func (c *CPU) opcodeCB0x95() {
	c.instruction_RES(c.HL.GetLowPointer(), 2)
}

// RES 2 (HL)
func (c *CPU) opcodeCB0x96() {
	c.instruction_RES_HL(2)
}

// RES 2 A
func (c *CPU) opcodeCB0x97() {
	c.instruction_RES(c.AF.GetHighPointer(), 2)
}

// RES 3 B
func (c *CPU) opcodeCB0x98() {
	c.instruction_RES(c.BC.GetHighPointer(), 3)
}

// RES 3 C
func (c *CPU) opcodeCB0x99() {
	c.instruction_RES(c.BC.GetLowPointer(), 3)
}

// RES 3 D
func (c *CPU) opcodeCB0x9A() {
	c.instruction_RES(c.DE.GetHighPointer(), 3)
}

// RES 3 E
func (c *CPU) opcodeCB0x9B() {
	c.instruction_RES(c.DE.GetLowPointer(), 3)
}

// RES 3 H
func (c *CPU) opcodeCB0x9C() {
	c.instruction_RES(c.HL.GetHighPointer(), 3)
}

// RES 3 L
func (c *CPU) opcodeCB0x9D() {
	c.instruction_RES(c.HL.GetLowPointer(), 3)
}

// RES 3 (HL)
func (c *CPU) opcodeCB0x9E() {
	c.instruction_RES_HL(3)
}

// RES 3 A
func (c *CPU) opcodeCB0x9F() {
	c.instruction_RES(c.AF.GetHighPointer(), 3)
}

// RES 4 B
func (c *CPU) opcodeCB0xA0() {
	c.instruction_RES(c.BC.GetHighPointer(), 4)
}

// RES 4 C
func (c *CPU) opcodeCB0xA1() {
	c.instruction_RES(c.BC.GetLowPointer(), 4)
}

// RES 4 D
func (c *CPU) opcodeCB0xA2() {
	c.instruction_RES(c.DE.GetHighPointer(), 4)
}

// RES 4 E
func (c *CPU) opcodeCB0xA3() {
	c.instruction_RES(c.DE.GetLowPointer(), 4)
}

// RES 4 H
func (c *CPU) opcodeCB0xA4() {
	c.instruction_RES(c.HL.GetHighPointer(), 4)
}

// RES 4 L
func (c *CPU) opcodeCB0xA5() {
	c.instruction_RES(c.HL.GetLowPointer(), 4)
}

// RES 4 (HL)
func (c *CPU) opcodeCB0xA6() {
	c.instruction_RES_HL(4)
}

// RES 4 A
func (c *CPU) opcodeCB0xA7() {
	c.instruction_RES(c.AF.GetHighPointer(), 4)
}

// RES 5 B
func (c *CPU) opcodeCB0xA8() {
	c.instruction_RES(c.BC.GetHighPointer(), 5)
}

// RES 5 C
func (c *CPU) opcodeCB0xA9() {
	c.instruction_RES(c.BC.GetLowPointer(), 5)
}

// RES 5 D
func (c *CPU) opcodeCB0xAA() {
	c.instruction_RES(c.DE.GetHighPointer(), 5)
}

// RES 5 E
func (c *CPU) opcodeCB0xAB() {
	c.instruction_RES(c.DE.GetLowPointer(), 5)
}

// RES 5 H
func (c *CPU) opcodeCB0xAC() {
	c.instruction_RES(c.HL.GetHighPointer(), 5)
}

// RES 5 L
func (c *CPU) opcodeCB0xAD() {
	c.instruction_RES(c.HL.GetLowPointer(), 5)
}

// RES 5 (HL)
func (c *CPU) opcodeCB0xAE() {
	c.instruction_RES_HL(5)
}

// RES 5 A
func (c *CPU) opcodeCB0xAF() {
	c.instruction_RES(c.AF.GetHighPointer(), 5)
}

// RES 6 B
func (c *CPU) opcodeCB0xB0() {
	c.instruction_RES(c.BC.GetHighPointer(), 6)
}

// RES 6 C
func (c *CPU) opcodeCB0xB1() {
	c.instruction_RES(c.BC.GetLowPointer(), 6)
}

// RES 6 D
func (c *CPU) opcodeCB0xB2() {
	c.instruction_RES(c.DE.GetHighPointer(), 6)
}

// RES 6 E
func (c *CPU) opcodeCB0xB3() {
	c.instruction_RES(c.DE.GetLowPointer(), 6)
}

// RES 6 H
func (c *CPU) opcodeCB0xB4() {
	c.instruction_RES(c.HL.GetHighPointer(), 6)
}

// RES 6 L
func (c *CPU) opcodeCB0xB5() {
	c.instruction_RES(c.HL.GetLowPointer(), 6)
}

// RES 6 (HL)
func (c *CPU) opcodeCB0xB6() {
	c.instruction_RES_HL(6)
}

// RES 6 A
func (c *CPU) opcodeCB0xB7() {
	c.instruction_RES(c.AF.GetHighPointer(), 6)
}

// RES 7 B
func (c *CPU) opcodeCB0xB8() {
	c.instruction_RES(c.BC.GetHighPointer(), 7)
}

// RES 7 C
func (c *CPU) opcodeCB0xB9() {
	c.instruction_RES(c.BC.GetLowPointer(), 7)
}

// RES 7 D
func (c *CPU) opcodeCB0xBA() {
	c.instruction_RES(c.DE.GetHighPointer(), 7)
}

// RES 7 E
func (c *CPU) opcodeCB0xBB() {
	c.instruction_RES(c.DE.GetLowPointer(), 7)
}

// RES 7 H
func (c *CPU) opcodeCB0xBC() {
	c.instruction_RES(c.HL.GetHighPointer(), 7)
}

// RES 7 L
func (c *CPU) opcodeCB0xBD() {
	c.instruction_RES(c.HL.GetLowPointer(), 7)
}

// RES 7 (HL)
func (c *CPU) opcodeCB0xBE() {
	c.instruction_RES_HL(7)
}

// RES 7 A
func (c *CPU) opcodeCB0xBF() {
	c.instruction_RES(c.AF.GetHighPointer(), 7)
}

// SET 0 B
func (c *CPU) opcodeCB0xC0() {
	c.instruction_SET(c.BC.GetHighPointer(), 0)
}

// SET 0 C
func (c *CPU) opcodeCB0xC1() {
	c.instruction_SET(c.BC.GetLowPointer(), 0)
}

// SET 0 D
func (c *CPU) opcodeCB0xC2() {
	c.instruction_SET(c.DE.GetHighPointer(), 0)
}

// SET 0 E
func (c *CPU) opcodeCB0xC3() {
	c.instruction_SET(c.DE.GetLowPointer(), 0)
}

// SET 0 H
func (c *CPU) opcodeCB0xC4() {
	c.instruction_SET(c.HL.GetHighPointer(), 0)
}

// SET 0 L
func (c *CPU) opcodeCB0xC5() {
	c.instruction_SET(c.HL.GetLowPointer(), 0)
}

// SET 0 (HL)
func (c *CPU) opcodeCB0xC6() {
	c.instruction_SET_HL(0)
}

// SET 0 A
func (c *CPU) opcodeCB0xC7() {
	c.instruction_SET(c.AF.GetHighPointer(), 0)
}

// SET 1 B
func (c *CPU) opcodeCB0xC8() {
	c.instruction_SET(c.BC.GetHighPointer(), 1)
}

// SET 1 C
func (c *CPU) opcodeCB0xC9() {
	c.instruction_SET(c.BC.GetLowPointer(), 1)
}

// SET 1 D
func (c *CPU) opcodeCB0xCA() {
	c.instruction_SET(c.DE.GetHighPointer(), 1)
}

// SET 1 E
func (c *CPU) opcodeCB0xCB() {
	c.instruction_SET(c.DE.GetLowPointer(), 1)
}

// SET 1 H
func (c *CPU) opcodeCB0xCC() {
	c.instruction_SET(c.HL.GetHighPointer(), 1)
}

// SET 1 L
func (c *CPU) opcodeCB0xCD() {
	c.instruction_SET(c.HL.GetLowPointer(), 1)
}

// SET 1 (HL)
func (c *CPU) opcodeCB0xCE() {
	c.instruction_SET_HL(1)
}

// SET 1 A
func (c *CPU) opcodeCB0xCF() {
	c.instruction_SET(c.AF.GetHighPointer(), 1)
}

// SET 2 B
func (c *CPU) opcodeCB0xD0() {
	c.instruction_SET(c.BC.GetHighPointer(), 2)
}

// SET 2 C
func (c *CPU) opcodeCB0xD1() {
	c.instruction_SET(c.BC.GetLowPointer(), 2)
}

// SET 2 D
func (c *CPU) opcodeCB0xD2() {
	c.instruction_SET(c.DE.GetHighPointer(), 2)
}

// SET 2 E
func (c *CPU) opcodeCB0xD3() {
	c.instruction_SET(c.DE.GetLowPointer(), 2)
}

// SET 2 H
func (c *CPU) opcodeCB0xD4() {
	c.instruction_SET(c.HL.GetHighPointer(), 2)
}

// SET 2 L
func (c *CPU) opcodeCB0xD5() {
	c.instruction_SET(c.HL.GetLowPointer(), 2)
}

// SET 2 (HL)
func (c *CPU) opcodeCB0xD6() {
	c.instruction_SET_HL(2)
}

// SET 2 A
func (c *CPU) opcodeCB0xD7() {
	c.instruction_SET(c.AF.GetHighPointer(), 2)
}

// SET 3 B
func (c *CPU) opcodeCB0xD8() {
	c.instruction_SET(c.BC.GetHighPointer(), 3)
}

// SET 3 C
func (c *CPU) opcodeCB0xD9() {
	c.instruction_SET(c.BC.GetLowPointer(), 3)
}

// SET 3 D
func (c *CPU) opcodeCB0xDA() {
	c.instruction_SET(c.DE.GetHighPointer(), 3)
}

// SET 3 E
func (c *CPU) opcodeCB0xDB() {
	c.instruction_SET(c.DE.GetLowPointer(), 3)
}

// SET 3 H
func (c *CPU) opcodeCB0xDC() {
	c.instruction_SET(c.HL.GetHighPointer(), 3)
}

// SET 3 L
func (c *CPU) opcodeCB0xDD() {
	c.instruction_SET(c.HL.GetLowPointer(), 3)
}

// SET 3 (HL)
func (c *CPU) opcodeCB0xDE() {
	c.instruction_SET_HL(3)
}

// SET 3 A
func (c *CPU) opcodeCB0xDF() {
	c.instruction_SET(c.AF.GetHighPointer(), 3)
}

// SET 4 B
func (c *CPU) opcodeCB0xE0() {
	c.instruction_SET(c.BC.GetHighPointer(), 4)
}

// SET 4 C
func (c *CPU) opcodeCB0xE1() {
	c.instruction_SET(c.BC.GetLowPointer(), 4)
}

// SET 4 D
func (c *CPU) opcodeCB0xE2() {
	c.instruction_SET(c.DE.GetHighPointer(), 4)
}

// SET 4 E
func (c *CPU) opcodeCB0xE3() {
	c.instruction_SET(c.DE.GetLowPointer(), 4)
}

// SET 4 H
func (c *CPU) opcodeCB0xE4() {
	c.instruction_SET(c.HL.GetHighPointer(), 4)
}

// SET 4 L
func (c *CPU) opcodeCB0xE5() {
	c.instruction_SET(c.HL.GetLowPointer(), 4)
}

// SET 4 (HL)
func (c *CPU) opcodeCB0xE6() {
	c.instruction_SET_HL(4)
}

// SET 4 A
func (c *CPU) opcodeCB0xE7() {
	c.instruction_SET(c.AF.GetHighPointer(), 4)

}

// SET 5 B
func (c *CPU) opcodeCB0xE8() {
	c.instruction_SET(c.BC.GetHighPointer(), 5)
}

// SET 5 C
func (c *CPU) opcodeCB0xE9() {
	c.instruction_SET(c.BC.GetLowPointer(), 5)
}

// SET 5 D
func (c *CPU) opcodeCB0xEA() {
	c.instruction_SET(c.DE.GetHighPointer(), 5)
}

// SET 5 E
func (c *CPU) opcodeCB0xEB() {
	c.instruction_SET(c.DE.GetLowPointer(), 5)
}

// SET 5 H
func (c *CPU) opcodeCB0xEC() {
	c.instruction_SET(c.HL.GetHighPointer(), 5)
}

// SET 5 L
func (c *CPU) opcodeCB0xED() {
	c.instruction_SET(c.HL.GetLowPointer(), 5)
}

// SET 5 (HL)
func (c *CPU) opcodeCB0xEE() {
	c.instruction_SET_HL(5)
}

// SET 5 A
func (c *CPU) opcodeCB0xEF() {
	c.instruction_SET(c.AF.GetHighPointer(), 5)
}

// SET 6 B
func (c *CPU) opcodeCB0xF0() {
	c.instruction_SET(c.BC.GetHighPointer(), 6)
}

// SET 6 C
func (c *CPU) opcodeCB0xF1() {
	c.instruction_SET(c.BC.GetLowPointer(), 6)
}

// SET 6 D
func (c *CPU) opcodeCB0xF2() {
	c.instruction_SET(c.DE.GetHighPointer(), 6)
}

// SET 6 E
func (c *CPU) opcodeCB0xF3() {
	c.instruction_SET(c.DE.GetLowPointer(), 6)
}

// SET 6 H
func (c *CPU) opcodeCB0xF4() {
	c.instruction_SET(c.HL.GetHighPointer(), 6)
}

// SET 6 L
func (c *CPU) opcodeCB0xF5() {
	c.instruction_SET(c.HL.GetLowPointer(), 6)
}

// SET 6 (HL)
func (c *CPU) opcodeCB0xF6() {
	c.instruction_SET_HL(6)
}

// SET 6 A
func (c *CPU) opcodeCB0xF7() {
	c.instruction_SET(c.AF.GetHighPointer(), 6)
}

// SET 7 B
func (c *CPU) opcodeCB0xF8() {
	c.instruction_SET(c.BC.GetHighPointer(), 7)
}

// SET 7 C
func (c *CPU) opcodeCB0xF9() {
	c.instruction_SET(c.BC.GetLowPointer(), 7)
}

// SET 7 D
func (c *CPU) opcodeCB0xFA() {
	c.instruction_SET(c.DE.GetHighPointer(), 7)
}

// SET 7 E
func (c *CPU) opcodeCB0xFB() {
	c.instruction_SET(c.DE.GetLowPointer(), 7)
}

// SET 7 H
func (c *CPU) opcodeCB0xFC() {
	c.instruction_SET(c.HL.GetHighPointer(), 7)
}

// SET 7 L
func (c *CPU) opcodeCB0xFD() {
	c.instruction_SET(c.HL.GetLowPointer(), 7)
}

// SET 7 (HL)
func (c *CPU) opcodeCB0xFE() {
	c.instruction_SET_HL(7)
}

// SET 7 A
func (c *CPU) opcodeCB0xFF() {
	c.instruction_SET(c.AF.GetHighPointer(), 7)
}
