package chibigb

// This values are from the tests written by Shay Green
// (http://blargg.parodius.com/gb-tests/)

var OPCodeMachineCycles = [256]byte{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1,
	1, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1,
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1,
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 2, 2, 2, 2, 2, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4,
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4,
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4,
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4,
}

var OPCodeAccurate = [256]byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
}

func (c *CPU) createTable() {
	c.OpCodeTable = [256]CPUOpCodeEntity{
		{opcode: 0x00, fn: c.opcode0x00, name: "XXX"},
		{opcode: 0x01, fn: c.opcode0x01, name: "XXX"},
		{opcode: 0x02, fn: c.opcode0x02, name: "XXX"},
		{opcode: 0x03, fn: c.opcode0x03, name: "XXX"},
		{opcode: 0x04, fn: c.opcode0x04, name: "XXX"},
		{opcode: 0x05, fn: c.opcode0x05, name: "XXX"},
		{opcode: 0x06, fn: c.opcode0x06, name: "XXX"},
		{opcode: 0x07, fn: c.opcode0x07, name: "XXX"},
		{opcode: 0x08, fn: c.opcode0x08, name: "XXX"},
		{opcode: 0x09, fn: c.opcode0x09, name: "XXX"},
		{opcode: 0x0A, fn: c.opcode0x0A, name: "XXX"},
		{opcode: 0x0B, fn: c.opcode0x0B, name: "XXX"},
		{opcode: 0x0C, fn: c.opcode0x0C, name: "XXX"},
		{opcode: 0x0D, fn: c.opcode0x0D, name: "XXX"},
		{opcode: 0x0E, fn: c.opcode0x0E, name: "XXX"},
		{opcode: 0x0F, fn: c.opcode0x0F, name: "XXX"},
		{opcode: 0x10, fn: c.opcode0x10, name: "XXX"},
		{opcode: 0x11, fn: c.opcode0x11, name: "XXX"},
		{opcode: 0x12, fn: c.opcode0x12, name: "XXX"},
		{opcode: 0x13, fn: c.opcode0x13, name: "XXX"},
		{opcode: 0x14, fn: c.opcode0x14, name: "XXX"},
		{opcode: 0x15, fn: c.opcode0x15, name: "XXX"},
		{opcode: 0x16, fn: c.opcode0x16, name: "XXX"},
		{opcode: 0x17, fn: c.opcode0x17, name: "XXX"},
		{opcode: 0x18, fn: c.opcode0x18, name: "XXX"},
		{opcode: 0x19, fn: c.opcode0x19, name: "XXX"},
		{opcode: 0x1A, fn: c.opcode0x1A, name: "XXX"},
		{opcode: 0x1B, fn: c.opcode0x1B, name: "XXX"},
		{opcode: 0x1C, fn: c.opcode0x1C, name: "XXX"},
		{opcode: 0x1D, fn: c.opcode0x1D, name: "XXX"},
		{opcode: 0x1E, fn: c.opcode0x1E, name: "XXX"},
		{opcode: 0x1F, fn: c.opcode0x1F, name: "XXX"},
		{opcode: 0x20, fn: c.opcode0x20, name: "XXX"},
		{opcode: 0x21, fn: c.opcode0x21, name: "XXX"},
		{opcode: 0x22, fn: c.opcode0x22, name: "XXX"},
		{opcode: 0x23, fn: c.opcode0x23, name: "XXX"},
		{opcode: 0x24, fn: c.opcode0x24, name: "XXX"},
		{opcode: 0x25, fn: c.opcode0x25, name: "XXX"},
		{opcode: 0x26, fn: c.opcode0x26, name: "XXX"},
		{opcode: 0x27, fn: c.opcode0x27, name: "XXX"},
		{opcode: 0x28, fn: c.opcode0x28, name: "XXX"},
		{opcode: 0x29, fn: c.opcode0x29, name: "XXX"},
		{opcode: 0x2A, fn: c.opcode0x2A, name: "XXX"},
		{opcode: 0x2B, fn: c.opcode0x2B, name: "XXX"},
		{opcode: 0x2C, fn: c.opcode0x2C, name: "XXX"},
		{opcode: 0x2D, fn: c.opcode0x2D, name: "XXX"},
		{opcode: 0x2E, fn: c.opcode0x2E, name: "XXX"},
		{opcode: 0x2F, fn: c.opcode0x2F, name: "XXX"},
		{opcode: 0x30, fn: c.opcode0x30, name: "XXX"},
		{opcode: 0x31, fn: c.opcode0x31, name: "XXX"},
		{opcode: 0x32, fn: c.opcode0x32, name: "XXX"},
		{opcode: 0x33, fn: c.opcode0x33, name: "XXX"},
		{opcode: 0x34, fn: c.opcode0x34, name: "XXX"},
		{opcode: 0x35, fn: c.opcode0x35, name: "XXX"},
		{opcode: 0x36, fn: c.opcode0x36, name: "XXX"},
		{opcode: 0x37, fn: c.opcode0x37, name: "XXX"},
		{opcode: 0x38, fn: c.opcode0x38, name: "XXX"},
		{opcode: 0x39, fn: c.opcode0x39, name: "XXX"},
		{opcode: 0x3A, fn: c.opcode0x3A, name: "XXX"},
		{opcode: 0x3B, fn: c.opcode0x3B, name: "XXX"},
		{opcode: 0x3C, fn: c.opcode0x3C, name: "XXX"},
		{opcode: 0x3D, fn: c.opcode0x3D, name: "XXX"},
		{opcode: 0x3E, fn: c.opcode0x3E, name: "XXX"},
		{opcode: 0x3F, fn: c.opcode0x3F, name: "XXX"},
		{opcode: 0x40, fn: c.opcode0x40, name: "XXX"},
		{opcode: 0x41, fn: c.opcode0x41, name: "XXX"},
		{opcode: 0x42, fn: c.opcode0x42, name: "XXX"},
		{opcode: 0x43, fn: c.opcode0x43, name: "XXX"},
		{opcode: 0x44, fn: c.opcode0x44, name: "XXX"},
		{opcode: 0x45, fn: c.opcode0x45, name: "XXX"},
		{opcode: 0x46, fn: c.opcode0x46, name: "XXX"},
		{opcode: 0x47, fn: c.opcode0x47, name: "XXX"},
		{opcode: 0x48, fn: c.opcode0x48, name: "XXX"},
		{opcode: 0x49, fn: c.opcode0x49, name: "XXX"},
		{opcode: 0x4A, fn: c.opcode0x4A, name: "XXX"},
		{opcode: 0x4B, fn: c.opcode0x4B, name: "XXX"},
		{opcode: 0x4C, fn: c.opcode0x4C, name: "XXX"},
		{opcode: 0x4D, fn: c.opcode0x4D, name: "XXX"},
		{opcode: 0x4E, fn: c.opcode0x4E, name: "XXX"},
		{opcode: 0x4F, fn: c.opcode0x4F, name: "XXX"},
		{opcode: 0x50, fn: c.opcode0x50, name: "XXX"},
		{opcode: 0x51, fn: c.opcode0x51, name: "XXX"},
		{opcode: 0x52, fn: c.opcode0x52, name: "XXX"},
		{opcode: 0x53, fn: c.opcode0x53, name: "XXX"},
		{opcode: 0x54, fn: c.opcode0x54, name: "XXX"},
		{opcode: 0x55, fn: c.opcode0x55, name: "XXX"},
		{opcode: 0x56, fn: c.opcode0x56, name: "XXX"},
		{opcode: 0x57, fn: c.opcode0x57, name: "XXX"},
		{opcode: 0x58, fn: c.opcode0x58, name: "XXX"},
		{opcode: 0x59, fn: c.opcode0x59, name: "XXX"},
		{opcode: 0x5A, fn: c.opcode0x5A, name: "XXX"},
		{opcode: 0x5B, fn: c.opcode0x5B, name: "XXX"},
		{opcode: 0x5C, fn: c.opcode0x5C, name: "XXX"},
		{opcode: 0x5D, fn: c.opcode0x5D, name: "XXX"},
		{opcode: 0x5E, fn: c.opcode0x5E, name: "XXX"},
		{opcode: 0x5F, fn: c.opcode0x5F, name: "XXX"},
		{opcode: 0x60, fn: c.opcode0x60, name: "XXX"},
		{opcode: 0x61, fn: c.opcode0x61, name: "XXX"},
		{opcode: 0x62, fn: c.opcode0x62, name: "XXX"},
		{opcode: 0x63, fn: c.opcode0x63, name: "XXX"},
		{opcode: 0x64, fn: c.opcode0x64, name: "XXX"},
		{opcode: 0x65, fn: c.opcode0x65, name: "XXX"},
		{opcode: 0x66, fn: c.opcode0x66, name: "XXX"},
		{opcode: 0x67, fn: c.opcode0x67, name: "XXX"},
		{opcode: 0x68, fn: c.opcode0x68, name: "XXX"},
		{opcode: 0x69, fn: c.opcode0x69, name: "XXX"},
		{opcode: 0x6A, fn: c.opcode0x6A, name: "XXX"},
		{opcode: 0x6B, fn: c.opcode0x6B, name: "XXX"},
		{opcode: 0x6C, fn: c.opcode0x6C, name: "XXX"},
		{opcode: 0x6D, fn: c.opcode0x6D, name: "XXX"},
		{opcode: 0x6E, fn: c.opcode0x6E, name: "XXX"},
		{opcode: 0x6F, fn: c.opcode0x6F, name: "XXX"},
		{opcode: 0x70, fn: c.opcode0x70, name: "XXX"},
		{opcode: 0x71, fn: c.opcode0x71, name: "XXX"},
		{opcode: 0x72, fn: c.opcode0x72, name: "XXX"},
		{opcode: 0x73, fn: c.opcode0x73, name: "XXX"},
		{opcode: 0x74, fn: c.opcode0x74, name: "XXX"},
		{opcode: 0x75, fn: c.opcode0x75, name: "XXX"},
		{opcode: 0x76, fn: c.opcode0x76, name: "XXX"},
		{opcode: 0x77, fn: c.opcode0x77, name: "XXX"},
		{opcode: 0x78, fn: c.opcode0x78, name: "XXX"},
		{opcode: 0x79, fn: c.opcode0x79, name: "XXX"},
		{opcode: 0x7A, fn: c.opcode0x7A, name: "XXX"},
		{opcode: 0x7B, fn: c.opcode0x7B, name: "XXX"},
		{opcode: 0x7C, fn: c.opcode0x7C, name: "XXX"},
		{opcode: 0x7D, fn: c.opcode0x7D, name: "XXX"},
		{opcode: 0x7E, fn: c.opcode0x7E, name: "XXX"},
		{opcode: 0x7F, fn: c.opcode0x7F, name: "XXX"},
		{opcode: 0x80, fn: c.opcode0x80, name: "XXX"},
		{opcode: 0x81, fn: c.opcode0x81, name: "XXX"},
		{opcode: 0x82, fn: c.opcode0x82, name: "XXX"},
		{opcode: 0x83, fn: c.opcode0x83, name: "XXX"},
		{opcode: 0x84, fn: c.opcode0x84, name: "XXX"},
		{opcode: 0x85, fn: c.opcode0x85, name: "XXX"},
		{opcode: 0x86, fn: c.opcode0x86, name: "XXX"},
		{opcode: 0x87, fn: c.opcode0x87, name: "XXX"},
		{opcode: 0x88, fn: c.opcode0x88, name: "XXX"},
		{opcode: 0x89, fn: c.opcode0x89, name: "XXX"},
		{opcode: 0x8A, fn: c.opcode0x8A, name: "XXX"},
		{opcode: 0x8B, fn: c.opcode0x8B, name: "XXX"},
		{opcode: 0x8C, fn: c.opcode0x8C, name: "XXX"},
		{opcode: 0x8D, fn: c.opcode0x8D, name: "XXX"},
		{opcode: 0x8E, fn: c.opcode0x8E, name: "XXX"},
		{opcode: 0x8F, fn: c.opcode0x8F, name: "XXX"},
		{opcode: 0x90, fn: c.opcode0x90, name: "XXX"},
		{opcode: 0x91, fn: c.opcode0x91, name: "XXX"},
		{opcode: 0x92, fn: c.opcode0x92, name: "XXX"},
		{opcode: 0x93, fn: c.opcode0x93, name: "XXX"},
		{opcode: 0x94, fn: c.opcode0x94, name: "XXX"},
		{opcode: 0x95, fn: c.opcode0x95, name: "XXX"},
		{opcode: 0x96, fn: c.opcode0x96, name: "XXX"},
		{opcode: 0x97, fn: c.opcode0x97, name: "XXX"},
		{opcode: 0x98, fn: c.opcode0x98, name: "XXX"},
		{opcode: 0x99, fn: c.opcode0x99, name: "XXX"},
		{opcode: 0x9A, fn: c.opcode0x9A, name: "XXX"},
		{opcode: 0x9B, fn: c.opcode0x9B, name: "XXX"},
		{opcode: 0x9C, fn: c.opcode0x9C, name: "XXX"},
		{opcode: 0x9D, fn: c.opcode0x9D, name: "XXX"},
		{opcode: 0x9E, fn: c.opcode0x9E, name: "XXX"},
		{opcode: 0x9F, fn: c.opcode0x9F, name: "XXX"},
		{opcode: 0xA0, fn: c.opcode0xA0, name: "XXX"},
		{opcode: 0xA1, fn: c.opcode0xA1, name: "XXX"},
		{opcode: 0xA2, fn: c.opcode0xA2, name: "XXX"},
		{opcode: 0xA3, fn: c.opcode0xA3, name: "XXX"},
		{opcode: 0xA4, fn: c.opcode0xA4, name: "XXX"},
		{opcode: 0xA5, fn: c.opcode0xA5, name: "XXX"},
		{opcode: 0xA6, fn: c.opcode0xA6, name: "XXX"},
		{opcode: 0xA7, fn: c.opcode0xA7, name: "XXX"},
		{opcode: 0xA8, fn: c.opcode0xA8, name: "XXX"},
		{opcode: 0xA9, fn: c.opcode0xA9, name: "XXX"},
		{opcode: 0xAA, fn: c.opcode0xAA, name: "XXX"},
		{opcode: 0xAB, fn: c.opcode0xAB, name: "XXX"},
		{opcode: 0xAC, fn: c.opcode0xAC, name: "XXX"},
		{opcode: 0xAD, fn: c.opcode0xAD, name: "XXX"},
		{opcode: 0xAE, fn: c.opcode0xAE, name: "XXX"},
		{opcode: 0xAF, fn: c.opcode0xAF, name: "XXX"},
		{opcode: 0xB0, fn: c.opcode0xB0, name: "XXX"},
		{opcode: 0xB1, fn: c.opcode0xB1, name: "XXX"},
		{opcode: 0xB2, fn: c.opcode0xB2, name: "XXX"},
		{opcode: 0xB3, fn: c.opcode0xB3, name: "XXX"},
		{opcode: 0xB4, fn: c.opcode0xB4, name: "XXX"},
		{opcode: 0xB5, fn: c.opcode0xB5, name: "XXX"},
		{opcode: 0xB6, fn: c.opcode0xB6, name: "XXX"},
		{opcode: 0xB7, fn: c.opcode0xB7, name: "XXX"},
		{opcode: 0xB8, fn: c.opcode0xB8, name: "XXX"},
		{opcode: 0xB9, fn: c.opcode0xB9, name: "XXX"},
		{opcode: 0xBA, fn: c.opcode0xBA, name: "XXX"},
		{opcode: 0xBB, fn: c.opcode0xBB, name: "XXX"},
		{opcode: 0xBC, fn: c.opcode0xBC, name: "XXX"},
		{opcode: 0xBD, fn: c.opcode0xBD, name: "XXX"},
		{opcode: 0xBE, fn: c.opcode0xBE, name: "XXX"},
		{opcode: 0xBF, fn: c.opcode0xBF, name: "XXX"},
		{opcode: 0xC0, fn: c.opcode0xC0, name: "XXX"},
		{opcode: 0xC1, fn: c.opcode0xC1, name: "XXX"},
		{opcode: 0xC2, fn: c.opcode0xC2, name: "XXX"},
		{opcode: 0xC3, fn: c.opcode0xC3, name: "XXX"},
		{opcode: 0xC4, fn: c.opcode0xC4, name: "XXX"},
		{opcode: 0xC5, fn: c.opcode0xC5, name: "XXX"},
		{opcode: 0xC6, fn: c.opcode0xC6, name: "XXX"},
		{opcode: 0xC7, fn: c.opcode0xC7, name: "XXX"},
		{opcode: 0xC8, fn: c.opcode0xC8, name: "XXX"},
		{opcode: 0xC9, fn: c.opcode0xC9, name: "XXX"},
		{opcode: 0xCA, fn: c.opcode0xCA, name: "XXX"},
		{opcode: 0xCB, fn: c.opcode0xCB, name: "XXX"},
		{opcode: 0xCC, fn: c.opcode0xCC, name: "XXX"},
		{opcode: 0xCD, fn: c.opcode0xCD, name: "XXX"},
		{opcode: 0xCE, fn: c.opcode0xCE, name: "XXX"},
		{opcode: 0xCF, fn: c.opcode0xCF, name: "XXX"},
		{opcode: 0xD0, fn: c.opcode0xD0, name: "XXX"},
		{opcode: 0xD1, fn: c.opcode0xD1, name: "XXX"},
		{opcode: 0xD2, fn: c.opcode0xD2, name: "XXX"},
		{opcode: 0xD3, fn: c.opcode0xD3, name: "XXX"},
		{opcode: 0xD4, fn: c.opcode0xD4, name: "XXX"},
		{opcode: 0xD5, fn: c.opcode0xD5, name: "XXX"},
		{opcode: 0xD6, fn: c.opcode0xD6, name: "XXX"},
		{opcode: 0xD7, fn: c.opcode0xD7, name: "XXX"},
		{opcode: 0xD8, fn: c.opcode0xD8, name: "XXX"},
		{opcode: 0xD9, fn: c.opcode0xD9, name: "XXX"},
		{opcode: 0xDA, fn: c.opcode0xDA, name: "XXX"},
		{opcode: 0xDB, fn: c.opcode0xDB, name: "XXX"},
		{opcode: 0xDC, fn: c.opcode0xDC, name: "XXX"},
		{opcode: 0xDD, fn: c.opcode0xDD, name: "XXX"},
		{opcode: 0xDE, fn: c.opcode0xDE, name: "XXX"},
		{opcode: 0xDF, fn: c.opcode0xDF, name: "XXX"},
		{opcode: 0xE0, fn: c.opcode0xE0, name: "XXX"},
		{opcode: 0xE1, fn: c.opcode0xE1, name: "XXX"},
		{opcode: 0xE2, fn: c.opcode0xE2, name: "XXX"},
		{opcode: 0xE3, fn: c.opcode0xE3, name: "XXX"},
		{opcode: 0xE4, fn: c.opcode0xE4, name: "XXX"},
		{opcode: 0xE5, fn: c.opcode0xE5, name: "XXX"},
		{opcode: 0xE6, fn: c.opcode0xE6, name: "XXX"},
		{opcode: 0xE7, fn: c.opcode0xE7, name: "XXX"},
		{opcode: 0xE8, fn: c.opcode0xE8, name: "XXX"},
		{opcode: 0xE9, fn: c.opcode0xE9, name: "XXX"},
		{opcode: 0xEA, fn: c.opcode0xEA, name: "XXX"},
		{opcode: 0xEB, fn: c.opcode0xEB, name: "XXX"},
		{opcode: 0xEC, fn: c.opcode0xEC, name: "XXX"},
		{opcode: 0xED, fn: c.opcode0xED, name: "XXX"},
		{opcode: 0xEE, fn: c.opcode0xEE, name: "XXX"},
		{opcode: 0xEF, fn: c.opcode0xEF, name: "XXX"},
		{opcode: 0xF0, fn: c.opcode0xF0, name: "XXX"},
		{opcode: 0xF1, fn: c.opcode0xF1, name: "XXX"},
		{opcode: 0xF2, fn: c.opcode0xF2, name: "XXX"},
		{opcode: 0xF3, fn: c.opcode0xF3, name: "XXX"},
		{opcode: 0xF4, fn: c.opcode0xF4, name: "XXX"},
		{opcode: 0xF5, fn: c.opcode0xF5, name: "XXX"},
		{opcode: 0xF6, fn: c.opcode0xF6, name: "XXX"},
		{opcode: 0xF7, fn: c.opcode0xF7, name: "XXX"},
		{opcode: 0xF8, fn: c.opcode0xF8, name: "XXX"},
		{opcode: 0xF9, fn: c.opcode0xF9, name: "XXX"},
		{opcode: 0xFA, fn: c.opcode0xFA, name: "XXX"},
		{opcode: 0xFB, fn: c.opcode0xFB, name: "XXX"},
		{opcode: 0xFC, fn: c.opcode0xFC, name: "XXX"},
		{opcode: 0xFD, fn: c.opcode0xFD, name: "XXX"},
		{opcode: 0xFE, fn: c.opcode0xFE, name: "XXX"},
		{opcode: 0xFF, fn: c.opcode0xFF, name: "XXX"},
	}
}

// NOP
func (c *CPU) opcode0x00() {
}

// LD BC, nn
func (c *CPU) opcode0x01() {
	c.instruction_LD_r8_d8(c.BC.GetLowPointer(), c.PC.Get())
	c.PC.Increment()
	c.instruction_LD_r8_d8(c.BC.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// LD (BC), A
func (c *CPU) opcode0x02() {
	c.instruction_LD_d8_r8(c.BC.Get(), c.AF.GetHigh())
}

// INC BC
func (c *CPU) opcode0x03() {
	c.BC.Increment()
}

// INC B
func (c *CPU) opcode0x04() {
	c.instruction_INC(c.BC.GetHighPointer())
}

// DEC B
func (c *CPU) opcode0x05() {
	c.instruction_DEC(c.BC.GetHighPointer())
}

// LD B, n
func (c *CPU) opcode0x06() {
	c.instruction_LD_r8_d8(c.BC.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// RLCA
func (c *CPU) opcode0x07() {
	c.instruction_RLC(c.AF.GetHighPointer(), true)
}

// LD (nn), SP
func (c *CPU) opcode0x08() {
	low := c.MemoryRead(c.PC.Get())
	c.PC.Increment()
	high := c.MemoryRead(c.PC.Get())
	c.PC.Increment()
	var addr uint16 = (uint16(high) << 8) + uint16(low)
	c.MemoryWrite(addr, c.SP.GetLow())
	c.MemoryWrite(addr+1, c.SP.GetHigh())
}

// ADD HL, BC
func (c *CPU) opcode0x09() {
	c.instruction_ADD_HL(c.BC.Get())
}

// LD A, (BC)
func (c *CPU) opcode0x0A() {
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), c.BC.Get())
}

// DEC BC
func (c *CPU) opcode0x0B() {
	c.BC.Decrement()
}

// INC C
func (c *CPU) opcode0x0C() {
	c.instruction_INC(c.BC.GetLowPointer())
}

// DEC C
func (c *CPU) opcode0x0D() {
	c.instruction_DEC(c.BC.GetLowPointer())
}

// LD C, n
func (c *CPU) opcode0x0E() {
	c.instruction_LD_r8_d8(c.BC.GetLowPointer(), c.PC.Get())
	c.PC.Increment()
}

// RRCA
func (c *CPU) opcode0x0F() {
	c.instruction_RRC(c.AF.GetHighPointer(), true)
}

// STOP
func (c *CPU) opcode0x10() {
	c.PC.Increment()
}

// LD DE, nn
func (c *CPU) opcode0x11() {
	c.instruction_LD_r8_d8(c.DE.GetLowPointer(), c.PC.Get())
	c.PC.Increment()
	c.instruction_LD_r8_d8(c.DE.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// LD (DE), A
func (c *CPU) opcode0x12() {
	c.instruction_LD_d8_r8(c.DE.Get(), c.AF.GetHigh())
}

// INC DE
func (c *CPU) opcode0x13() {
	c.DE.Increment()
}

// INC D
func (c *CPU) opcode0x14() {
	c.instruction_INC(c.DE.GetHighPointer())
}

// DEC D
func (c *CPU) opcode0x15() {
	c.instruction_DEC(c.DE.GetHighPointer())
}

// LD D, n
func (c *CPU) opcode0x16() {
	c.instruction_LD_r8_d8(c.DE.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// RLA
func (c *CPU) opcode0x17() {
	c.instruction_RL(c.AF.GetHighPointer(), true)
}

// JR n
func (c *CPU) opcode0x18() {
	// calculate by signed int
	c.PC.Set(uint16(int(c.PC.Get()) + 1 + int(int8(c.MemoryRead(c.PC.Get())))))
}

// ADD HL, DE
func (c *CPU) opcode0x19() {
	c.instruction_ADD_HL(c.DE.Get())
}

// LD A, (DE)
func (c *CPU) opcode0x1A() {
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), c.DE.Get())
}

// DEC DE
func (c *CPU) opcode0x1B() {
	c.DE.Decrement()
}

// INC E
func (c *CPU) opcode0x1C() {
	c.instruction_INC(c.DE.GetLowPointer())
}

// DEC E
func (c *CPU) opcode0x1D() {
	c.instruction_DEC(c.DE.GetLowPointer())
}

// LD E, n
func (c *CPU) opcode0x1E() {
	c.instruction_LD_r8_d8(c.DE.GetLowPointer(), c.PC.Get())
	c.PC.Increment()
}

// RRA
func (c *CPU) opcode0x1F() {
	c.instruction_RR(c.AF.GetHighPointer(), true)
}

// JR NZ, n
func (c *CPU) opcode0x20() {
	if c.CheckFlag(CPUFLAG_ZERO) == false {
		c.PC.Set(uint16(int(c.PC.Get()) + 1 + int(int8(c.MemoryRead(c.PC.Get())))))
		c.BranchTaken = true
	} else {
		c.PC.Increment()
	}
}

// LD HL, nn
func (c *CPU) opcode0x21() {
	c.instruction_LD_r8_d8(c.HL.GetLowPointer(), c.PC.Get())
	c.PC.Increment()
	c.instruction_LD_r8_d8(c.HL.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// LD (HLI), A
func (c *CPU) opcode0x22() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.AF.GetHigh())
	c.HL.Increment()
}

// INC HL
func (c *CPU) opcode0x23() {
	c.HL.Increment()
}

// INC H
func (c *CPU) opcode0x24() {
	c.instruction_INC(c.HL.GetHighPointer())
}

// DEC H
func (c *CPU) opcode0x25() {
	c.instruction_DEC(c.HL.GetHighPointer())
}

// LD H, n
func (c *CPU) opcode0x26() {
	c.instruction_LD_r8_d8(c.HL.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// DAA
func (c *CPU) opcode0x27() {
	var a int = int(c.AF.GetHigh())

	if c.CheckFlag(CPUFLAG_SUB) == false {
		if c.CheckFlag(CPUFLAG_HALF) || ((a & 0xF) > 9) {
			a += 0x06
		}

		if c.CheckFlag(CPUFLAG_CARRY) || (a > 0x9F) {
			a += 0x60
		}
	} else {
		if c.CheckFlag(CPUFLAG_HALF) {
			a = (a - 6) & 0xFF
		}

		if c.CheckFlag(CPUFLAG_CARRY) {
			a -= 0x60
		}
	}

	c.ClearFlag(CPUFLAG_HALF)
	c.ClearFlag(CPUFLAG_ZERO)

	if (a & 0x100) == 0x100 {
		c.SetFlag(CPUFLAG_CARRY)
	}

	a &= 0xFF
	c.SetZeroFlag(byte(a))
	c.AF.SetHigh(byte(a))
}

// JR Z, n
func (c *CPU) opcode0x28() {
	if c.CheckFlag(CPUFLAG_ZERO) {
		c.PC.Set(uint16(int(c.PC.Get()) + 1 + int(int8(c.MemoryRead(c.PC.Get())))))
		c.BranchTaken = true
	} else {
		c.PC.Increment()
	}
}

// ADD HL, HL
func (c *CPU) opcode0x29() {
	c.instruction_ADD_HL(c.HL.Get())
}

// LD A, (HLI)
func (c *CPU) opcode0x2A() {
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), c.HL.Get())
	c.HL.Increment()
}

// DEC HL
func (c *CPU) opcode0x2B() {
	c.HL.Decrement()
}

// INC L
func (c *CPU) opcode0x2C() {
	c.instruction_INC(c.HL.GetLowPointer())
}

// DEC L
func (c *CPU) opcode0x2D() {
	c.instruction_DEC(c.HL.GetLowPointer())
}

// LD L, n
func (c *CPU) opcode0x2E() {
	c.instruction_LD_r8_d8(c.HL.GetLowPointer(), c.PC.Get())
	c.PC.Increment()
}

// CPL
func (c *CPU) opcode0x2F() {
	c.AF.SetHigh(^c.AF.GetHigh())
	c.SetFlag(CPUFLAG_HALF)
	c.SetFlag(CPUFLAG_SUB)
}

// JR NC, n
func (c *CPU) opcode0x30() {
	if c.CheckFlag(CPUFLAG_CARRY) == false {
		c.PC.Set(uint16(int(c.PC.Get()) + 1 + int(int8(c.MemoryRead(c.PC.Get())))))
		c.BranchTaken = true
	} else {
		c.PC.Increment()
	}
}

// LD SP, nn
func (c *CPU) opcode0x31() {
	c.SP.SetLow(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
	c.SP.SetHigh(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// LD (HLD), A
func (c *CPU) opcode0x32() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.AF.GetHigh())
	c.HL.Decrement()
}

// INC SP
func (c *CPU) opcode0x33() {
	c.SP.Increment()
}

// INC (HL)
func (c *CPU) opcode0x34() {
	c.instruction_INC_HL()
}

// DEC (HL)
func (c *CPU) opcode0x35() {
	c.instruction_DEC_HL()
}

// LD (HL), n
func (c *CPU) opcode0x36() {
	c.MemoryWrite(c.HL.Get(), c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// SCF
func (c *CPU) opcode0x37() {
	c.SetFlag(CPUFLAG_CARRY)
	c.ClearFlag(CPUFLAG_HALF)
	c.ClearFlag(CPUFLAG_SUB)
}

// JR C, n
func (c *CPU) opcode0x38() {
	if c.CheckFlag(CPUFLAG_CARRY) {
		c.PC.Set(uint16(int(c.PC.Get()) + 1 + int(int8(c.MemoryRead(c.PC.Get())))))
		c.BranchTaken = true
	} else {
		c.PC.Increment()
	}
}

// ADD HL, SP
func (c *CPU) opcode0x39() {
	c.instruction_ADD_HL(c.SP.Get())
}

// LD A, (HLD)
func (c *CPU) opcode0x3A() {
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), c.HL.Get())
	c.HL.Decrement()
}

// DEC SP
func (c *CPU) opcode0x3B() {
	c.SP.Decrement()
}

// INC A
func (c *CPU) opcode0x3C() {
	c.instruction_INC(c.AF.GetHighPointer())
}

// DEC A
func (c *CPU) opcode0x3D() {
	c.instruction_DEC(c.AF.GetHighPointer())
}

// LD A, n
func (c *CPU) opcode0x3E() {
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), c.PC.Get())
	c.PC.Increment()
}

// CCF
func (c *CPU) opcode0x3F() {
	c.FlipFlag(CPUFLAG_CARRY)
	c.ClearFlag(CPUFLAG_HALF)
	c.ClearFlag(CPUFLAG_SUB)
}

// LD B, B
func (c *CPU) opcode0x40() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.BC.GetHigh())
}

// LD B, C
func (c *CPU) opcode0x41() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.BC.GetLow())
}

// LD B, D
func (c *CPU) opcode0x42() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.DE.GetHigh())
}

// LD B, E
func (c *CPU) opcode0x43() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.DE.GetLow())
}

// LD B, H
func (c *CPU) opcode0x44() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.HL.GetHigh())
}

// LD B, L
func (c *CPU) opcode0x45() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.HL.GetLow())
}

// LD B, (HL)
func (c *CPU) opcode0x46() {
	c.instruction_LD_r8_d8(c.BC.GetHighPointer(), c.HL.Get())
}

// LD B, A
func (c *CPU) opcode0x47() {
	c.instruction_LD_r8_r8(c.BC.GetHighPointer(), c.AF.GetHigh())
}

// LD C, B
func (c *CPU) opcode0x48() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.BC.GetHigh())
}

// LD C, C
func (c *CPU) opcode0x49() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.BC.GetLow())
}

// LD C, D
func (c *CPU) opcode0x4A() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.DE.GetHigh())
}

// LD C, E
func (c *CPU) opcode0x4B() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.DE.GetLow())
}

// LD C, H
func (c *CPU) opcode0x4C() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.HL.GetHigh())
}

// LD C, L
func (c *CPU) opcode0x4D() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.HL.GetLow())
}

// LD C, (HL)
func (c *CPU) opcode0x4E() {
	c.instruction_LD_r8_d8(c.BC.GetLowPointer(), c.HL.Get())
}

// LD C, A
func (c *CPU) opcode0x4F() {
	c.instruction_LD_r8_r8(c.BC.GetLowPointer(), c.AF.GetHigh())
}

// LD D, B
func (c *CPU) opcode0x50() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.BC.GetHigh())
}

// LD D, C
func (c *CPU) opcode0x51() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.BC.GetLow())
}

// LD D, D
func (c *CPU) opcode0x52() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.DE.GetHigh())
}

// LD D, E
func (c *CPU) opcode0x53() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.DE.GetLow())
}

// LD D, H
func (c *CPU) opcode0x54() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.HL.GetHigh())
}

// LD D, L
func (c *CPU) opcode0x55() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.HL.GetLow())
}

// LD D, (HL)
func (c *CPU) opcode0x56() {
	c.instruction_LD_r8_d8(c.DE.GetHighPointer(), c.HL.Get())
}

// LD D, A
func (c *CPU) opcode0x57() {
	c.instruction_LD_r8_r8(c.DE.GetHighPointer(), c.AF.GetHigh())
}

// LD E, B
func (c *CPU) opcode0x58() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.BC.GetHigh())
}

// LD E, C
func (c *CPU) opcode0x59() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.BC.GetLow())
}

// LD E, D
func (c *CPU) opcode0x5A() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.DE.GetHigh())
}

// LD E, E
func (c *CPU) opcode0x5B() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.DE.GetLow())
}

// LD E, H
func (c *CPU) opcode0x5C() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.HL.GetHigh())
}

// LD E, L
func (c *CPU) opcode0x5D() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.HL.GetLow())
}

// LD E, (HL)
func (c *CPU) opcode0x5E() {
	c.instruction_LD_r8_d8(c.DE.GetLowPointer(), c.HL.Get())
}

// LD E, A
func (c *CPU) opcode0x5F() {
	c.instruction_LD_r8_r8(c.DE.GetLowPointer(), c.AF.GetHigh())
}

// LD H, B
func (c *CPU) opcode0x60() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.BC.GetHigh())
}

// LD H, C
func (c *CPU) opcode0x61() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.BC.GetLow())
}

// LD H, D
func (c *CPU) opcode0x62() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.DE.GetHigh())
}

// LD, H, E
func (c *CPU) opcode0x63() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.DE.GetLow())
}

// LD H, H
func (c *CPU) opcode0x64() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.HL.GetHigh())
}

// LD H, L
func (c *CPU) opcode0x65() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.HL.GetLow())
}

// LD H, (HL)
func (c *CPU) opcode0x66() {
	c.instruction_LD_r8_d8(c.HL.GetHighPointer(), c.HL.Get())
}

// LD H, A
func (c *CPU) opcode0x67() {
	c.instruction_LD_r8_r8(c.HL.GetHighPointer(), c.AF.GetHigh())
}

// LD L, B
func (c *CPU) opcode0x68() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.BC.GetHigh())
}

// LD L, C
func (c *CPU) opcode0x69() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.BC.GetLow())
}

// LD L, D
func (c *CPU) opcode0x6A() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.DE.GetHigh())
}

// LD L, E
func (c *CPU) opcode0x6B() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.DE.GetLow())
}

// LD L, H
func (c *CPU) opcode0x6C() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.HL.GetHigh())
}

// LD L, L
func (c *CPU) opcode0x6D() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.HL.GetLow())
}

// LD L, (HL)
func (c *CPU) opcode0x6E() {
	c.instruction_LD_r8_d8(c.HL.GetLowPointer(), c.HL.Get())
}

// LD L, A
func (c *CPU) opcode0x6F() {
	c.instruction_LD_r8_r8(c.HL.GetLowPointer(), c.AF.GetHigh())
}

// LD (HL), B
func (c *CPU) opcode0x70() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.BC.GetHigh())
}

// LD (HL), C
func (c *CPU) opcode0x71() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.BC.GetLow())
}

// LD (HL), D
func (c *CPU) opcode0x72() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.DE.GetHigh())
}

// LD (HL), E
func (c *CPU) opcode0x73() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.DE.GetLow())
}

// LD (HL), H
func (c *CPU) opcode0x74() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.HL.GetHigh())
}

// LD (HL), L
func (c *CPU) opcode0x75() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.HL.GetLow())
}

// HALT
func (c *CPU) opcode0x76() {
	if c.IMECycles > 0 {
		// If EI is pending interrupts are triggered before Halt
		c.IMECycles = 0
		c.IME = true
		c.PC.Decrement()
	} else {
		if_reg := c.MemoryRetrieve(0xFF0F)
		ie_reg := c.MemoryRetrieve(0xFFFF)

		c.Halt = true

		if !c.IME && ((if_reg & ie_reg & 0x1F) != 0) {
			c.SkipPCBug = true
		}
	}
}

// LD (HL), A
func (c *CPU) opcode0x77() {
	c.instruction_LD_d8_r8(c.HL.Get(), c.AF.GetHigh())
}

// LD A, B
func (c *CPU) opcode0x78() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.BC.GetHigh())
}

// LD A, C
func (c *CPU) opcode0x79() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.BC.GetLow())
}

// LD A, D
func (c *CPU) opcode0x7A() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.DE.GetHigh())
}

// LD A, E
func (c *CPU) opcode0x7B() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.DE.GetLow())
}

// LD A, H
func (c *CPU) opcode0x7C() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.HL.GetHigh())
}

// LD A, L
func (c *CPU) opcode0x7D() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.HL.GetLow())
}

// LD A, (HL)
func (c *CPU) opcode0x7E() {
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), c.HL.Get())
}

// LD A, A
func (c *CPU) opcode0x7F() {
	c.instruction_LD_r8_r8(c.AF.GetHighPointer(), c.AF.GetHigh())
}

// ADD A, B
func (c *CPU) opcode0x80() {
	c.instruction_ADD(c.BC.GetHigh())
}

// ADD A, C
func (c *CPU) opcode0x81() {
	c.instruction_ADD(c.BC.GetLow())
}

// ADD A, D
func (c *CPU) opcode0x82() {
	c.instruction_ADD(c.DE.GetHigh())
}

// ADD A, E
func (c *CPU) opcode0x83() {
	c.instruction_ADD(c.DE.GetLow())
}

// ADD A, H
func (c *CPU) opcode0x84() {
	c.instruction_ADD(c.HL.GetHigh())
}

// ADD A, L
func (c *CPU) opcode0x85() {
	c.instruction_ADD(c.HL.GetLow())
}

// ADD A, (HL)
func (c *CPU) opcode0x86() {
	c.instruction_ADD(c.MemoryRead(c.HL.Get()))
}

// ADD A, A
func (c *CPU) opcode0x87() {
	c.instruction_ADD(c.AF.GetHigh())
}

// ADC A, B
func (c *CPU) opcode0x88() {
	c.instruction_ADC(c.BC.GetHigh())
}

// ADC A, C
func (c *CPU) opcode0x89() {
	c.instruction_ADC(c.BC.GetLow())
}

// ADC A, D
func (c *CPU) opcode0x8A() {
	c.instruction_ADC(c.DE.GetHigh())
}

// ADC A, E
func (c *CPU) opcode0x8B() {
	c.instruction_ADC(c.DE.GetLow())
}

// ADC A, H
func (c *CPU) opcode0x8C() {
	c.instruction_ADC(c.HL.GetHigh())
}

// ADC A, L
func (c *CPU) opcode0x8D() {
	c.instruction_ADC(c.HL.GetLow())
}

// ADC A, (HL)
func (c *CPU) opcode0x8E() {
	c.instruction_ADC(c.MemoryRead(c.HL.Get()))
}

// ADC A, A
func (c *CPU) opcode0x8F() {
	c.instruction_ADC(c.AF.GetHigh())
}

// SUB B
func (c *CPU) opcode0x90() {
	c.instruction_SUB(c.BC.GetHigh())
}

// SUB C
func (c *CPU) opcode0x91() {
	c.instruction_SUB(c.BC.GetLow())
}

// SUB D
func (c *CPU) opcode0x92() {
	c.instruction_SUB(c.DE.GetHigh())
}

// SUB E
func (c *CPU) opcode0x93() {
	c.instruction_SUB(c.DE.GetLow())
}

// SUB H
func (c *CPU) opcode0x94() {
	c.instruction_SUB(c.HL.GetHigh())
}

// SUB L
func (c *CPU) opcode0x95() {
	c.instruction_SUB(c.HL.GetLow())
}

// SUB (HL)
func (c *CPU) opcode0x96() {
	c.instruction_SUB(c.MemoryRead(c.HL.Get()))
}

// SUB A
func (c *CPU) opcode0x97() {
	c.instruction_SUB(c.AF.GetHigh())
}

// SBC B
func (c *CPU) opcode0x98() {
	c.instruction_SBC(c.BC.GetHigh())
}

// SBC C
func (c *CPU) opcode0x99() {
	c.instruction_SBC(c.BC.GetLow())
}

// SBC D
func (c *CPU) opcode0x9A() {
	c.instruction_SBC(c.DE.GetHigh())
}

// SBC E
func (c *CPU) opcode0x9B() {
	c.instruction_SBC(c.DE.GetLow())
}

// SBC H
func (c *CPU) opcode0x9C() {
	c.instruction_SBC(c.HL.GetHigh())
}

// SBC L
func (c *CPU) opcode0x9D() {
	c.instruction_SBC(c.HL.GetLow())
}

// SBC (HL)
func (c *CPU) opcode0x9E() {
	c.instruction_SBC(c.MemoryRead(c.HL.Get()))
}

// SBC A
func (c *CPU) opcode0x9F() {
	c.instruction_SBC(c.AF.GetHigh())
}

// AND B
func (c *CPU) opcode0xA0() {
	c.instruction_AND(c.BC.GetHigh())
}

// AND C
func (c *CPU) opcode0xA1() {
	c.instruction_AND(c.BC.GetLow())
}

// AND D
func (c *CPU) opcode0xA2() {
	c.instruction_AND(c.DE.GetHigh())
}

// AND E
func (c *CPU) opcode0xA3() {
	c.instruction_AND(c.DE.GetLow())
}

// AND H
func (c *CPU) opcode0xA4() {
	c.instruction_AND(c.HL.GetHigh())
}

// AND L
func (c *CPU) opcode0xA5() {
	c.instruction_AND(c.HL.GetLow())
}

// AND (HL)
func (c *CPU) opcode0xA6() {
	c.instruction_AND(c.MemoryRead(c.HL.Get()))
}

// AND A
func (c *CPU) opcode0xA7() {
	c.instruction_AND(c.AF.GetHigh())
}

// XOR B
func (c *CPU) opcode0xA8() {
	c.instruction_XOR(c.BC.GetHigh())
}

// XOR C
func (c *CPU) opcode0xA9() {
	c.instruction_XOR(c.BC.GetLow())
}

// XOR D
func (c *CPU) opcode0xAA() {
	c.instruction_XOR(c.DE.GetHigh())
}

// XOR E
func (c *CPU) opcode0xAB() {
	c.instruction_XOR(c.DE.GetLow())
}

// XOR H
func (c *CPU) opcode0xAC() {
	c.instruction_XOR(c.HL.GetHigh())
}

// XOR L
func (c *CPU) opcode0xAD() {
	c.instruction_XOR(c.HL.GetLow())
}

// XOR (HL)
func (c *CPU) opcode0xAE() {
	c.instruction_XOR(c.MemoryRead(c.HL.Get()))
}

// XOR A
func (c *CPU) opcode0xAF() {
	c.instruction_XOR(c.AF.GetHigh())
}

// OR B
func (c *CPU) opcode0xB0() {
	c.instruction_OR(c.BC.GetHigh())
}

// OR C
func (c *CPU) opcode0xB1() {
	c.instruction_OR(c.BC.GetLow())
}

// OR D
func (c *CPU) opcode0xB2() {
	c.instruction_OR(c.DE.GetHigh())
}

// OR E
func (c *CPU) opcode0xB3() {
	c.instruction_OR(c.DE.GetLow())
}

// OR H
func (c *CPU) opcode0xB4() {
	c.instruction_OR(c.HL.GetHigh())
}

// OR L
func (c *CPU) opcode0xB5() {
	c.instruction_OR(c.HL.GetLow())
}

// OR (HL)
func (c *CPU) opcode0xB6() {
	c.instruction_OR(c.MemoryRead(c.HL.Get()))
}

// OR A
func (c *CPU) opcode0xB7() {
	c.instruction_OR(c.AF.GetHigh())
}

// CP B
func (c *CPU) opcode0xB8() {
	c.instruction_CP(c.BC.GetHigh())
}

// CP C
func (c *CPU) opcode0xB9() {
	c.instruction_CP(c.BC.GetLow())
}

// CP D
func (c *CPU) opcode0xBA() {
	c.instruction_CP(c.DE.GetHigh())
}

// CP E
func (c *CPU) opcode0xBB() {
	c.instruction_CP(c.DE.GetLow())
}

// CP H
func (c *CPU) opcode0xBC() {
	c.instruction_CP(c.HL.GetHigh())
}

// CP L
func (c *CPU) opcode0xBD() {
	c.instruction_CP(c.HL.GetLow())
}

// CP (HL)
func (c *CPU) opcode0xBE() {
	c.instruction_CP(c.MemoryRead(c.HL.Get()))
}

// CP A
func (c *CPU) opcode0xBF() {
	c.instruction_CP(c.AF.GetHigh())
}

// RET NZ
func (c *CPU) opcode0xC0() {
	if c.CheckFlag(CPUFLAG_ZERO) == false {
		c.instruction_StackPop(c.PC)
		c.BranchTaken = true
	}
}

// POP BC
func (c *CPU) opcode0xC1() {
	c.instruction_StackPop(c.BC)
}

// JP NZ, nn
func (c *CPU) opcode0xC2() {
	if c.CheckFlag(CPUFLAG_ZERO) == false {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// JP nn
func (c *CPU) opcode0xC3() {
	var low byte = c.MemoryRead(c.PC.Get())
	c.PC.Increment()
	var high byte = c.MemoryRead(c.PC.Get())
	c.PC.SetHigh(high)
	c.PC.SetLow(low)
}

// CALL NZ, nn
func (c *CPU) opcode0xC4() {
	if c.CheckFlag(CPUFLAG_ZERO) == false {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		c.instruction_StackPush(c.PC)
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// PUSH BC
func (c *CPU) opcode0xC5() {
	c.instruction_StackPush(c.BC)
}

// ADD A, n
func (c *CPU) opcode0xC6() {
	c.instruction_ADD(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 00H
func (c *CPU) opcode0xC7() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0000)
}

// RET Z
func (c *CPU) opcode0xC8() {
	if c.CheckFlag(CPUFLAG_ZERO) {
		c.instruction_StackPop(c.PC)
		c.BranchTaken = true
	}
}

// RET
func (c *CPU) opcode0xC9() {
	c.instruction_StackPop(c.PC)
}

// JP Z, nn
func (c *CPU) opcode0xCA() {
	if c.CheckFlag(CPUFLAG_ZERO) {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// CB prefixed instruction
func (c *CPU) opcode0xCB() {
}

// CALL Z, nn
func (c *CPU) opcode0xCC() {
	if c.CheckFlag(CPUFLAG_ZERO) {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		c.instruction_StackPush(c.PC)
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// CALL nn
func (c *CPU) opcode0xCD() {
	var low byte = c.MemoryRead(c.PC.Get())
	c.PC.Increment()
	var high byte = c.MemoryRead(c.PC.Get())
	c.PC.Increment()
	c.instruction_StackPush(c.PC)
	c.PC.SetHigh(high)
	c.PC.SetLow(low)
}

// ADC A, n
func (c *CPU) opcode0xCE() {
	c.instruction_ADC(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 08H
func (c *CPU) opcode0xCF() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0008)
}

// RET NC
func (c *CPU) opcode0xD0() {
	if c.CheckFlag(CPUFLAG_CARRY) == false {
		c.instruction_StackPop(c.PC)
		c.BranchTaken = true
	}
}

// POP DE
func (c *CPU) opcode0xD1() {
	c.instruction_StackPop(c.DE)
}

// JP NC, nn
func (c *CPU) opcode0xD2() {
	if c.CheckFlag(CPUFLAG_CARRY) == false {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// Invalid Opcode
func (c *CPU) opcode0xD3() {
	c.instruction_InvalidOpcode()
}

// CALL NC, nn
func (c *CPU) opcode0xD4() {
	if c.CheckFlag(CPUFLAG_CARRY) == false {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		c.instruction_StackPush(c.PC)
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// PUSH DE
func (c *CPU) opcode0xD5() {
	c.instruction_StackPush(c.DE)
}

// SUB n
func (c *CPU) opcode0xD6() {
	c.instruction_SUB(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 10H
func (c *CPU) opcode0xD7() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0010)
}

// RET C
func (c *CPU) opcode0xD8() {
	if c.CheckFlag(CPUFLAG_CARRY) {
		c.instruction_StackPop(c.PC)
		c.BranchTaken = true
	}
}

// RETI
func (c *CPU) opcode0xD9() {
	c.instruction_StackPop(c.PC)
	c.IME = true
}

// JP C, nn
func (c *CPU) opcode0xDA() {
	if c.CheckFlag(CPUFLAG_CARRY) {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// Invalid Opcode
func (c *CPU) opcode0xDB() {
	c.instruction_InvalidOpcode()
}

// CALL C, nn
func (c *CPU) opcode0xDC() {
	if c.CheckFlag(CPUFLAG_CARRY) {
		var low byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		var high byte = c.MemoryRead(c.PC.Get())
		c.PC.Increment()
		c.instruction_StackPush(c.PC)
		c.PC.SetHigh(high)
		c.PC.SetLow(low)
		c.BranchTaken = true
	} else {
		c.PC.Increment()
		c.PC.Increment()
	}
}

// Invalid Opcode
func (c *CPU) opcode0xDD() {
	c.instruction_InvalidOpcode()
}

// SBC n
func (c *CPU) opcode0xDE() {
	c.instruction_SBC(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 18H
func (c *CPU) opcode0xDF() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0018)
}

// LD (0xFF00+n), A
func (c *CPU) opcode0xE0() {
	var addr uint16 = 0xFF00 + uint16(c.MemoryRead(c.PC.Get()))
	c.instruction_LD_d8_r8(addr, c.AF.GetHigh())
	c.PC.Increment()
}

// POP HL
func (c *CPU) opcode0xE1() {
	c.instruction_StackPop(c.HL)
}

// LD (0xFF00+C), A
func (c *CPU) opcode0xE2() {
	var addr uint16 = 0xFF00 + uint16(c.BC.GetLow())
	c.instruction_LD_d8_r8(addr, c.AF.GetHigh())
	// c.PC.Increment()
}

// Invalid Opcode
func (c *CPU) opcode0xE3() {
	c.instruction_InvalidOpcode()
}

// Invalid Opcode
func (c *CPU) opcode0xE4() {
	c.instruction_InvalidOpcode()
}

// PUSH HL
func (c *CPU) opcode0xE5() {
	c.instruction_StackPush(c.HL)
}

// AND n
func (c *CPU) opcode0xE6() {
	c.instruction_AND(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 20H
func (c *CPU) opcode0xE7() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0020)
}

// ADD SP, n
func (c *CPU) opcode0xE8() {
	c.instruction_ADD_SP(int8(c.MemoryRead(c.PC.Get())))
	c.PC.Increment()
}

// JP (HL)
func (c *CPU) opcode0xE9() {
	c.PC.Set(c.HL.Get())
}

// LD (nn), A
func (c *CPU) opcode0xEA() {
	tmp := NewCPURegister()
	tmp.SetLow(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
	tmp.SetHigh(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
	c.instruction_LD_d8_r8(tmp.Get(), c.AF.GetHigh())
}

// Invalid Opcode
func (c *CPU) opcode0xEB() {
	c.instruction_InvalidOpcode()
}

// Invalid Opcode
func (c *CPU) opcode0xEC() {
	c.instruction_InvalidOpcode()
}

// Invalid Opcode
func (c *CPU) opcode0xED() {
	c.instruction_InvalidOpcode()
}

// XOR n
func (c *CPU) opcode0xEE() {
	c.instruction_XOR(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 28H
func (c *CPU) opcode0xEF() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0028)
}

// LD A, (0xFF00+n)
func (c *CPU) opcode0xF0() {
	var addr uint16 = 0xFF00 + uint16(c.MemoryRead(c.PC.Get()))
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), addr)
	c.PC.Increment()
}

// POP AF
func (c *CPU) opcode0xF1() {
	c.instruction_StackPop(c.AF)
	c.AF.SetLow(c.AF.GetLow() & 0xF0)
}

// LD A, (C)
func (c *CPU) opcode0xF2() {
	var addr uint16 = 0xFF00 + uint16(c.BC.GetLow())
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), addr)
}

// DI
func (c *CPU) opcode0xF3() {
	c.IME = false
	c.IMECycles = 0
}

// Invalid Opcode
func (c *CPU) opcode0xF4() {
	c.instruction_InvalidOpcode()
}

// PUSH AF
func (c *CPU) opcode0xF5() {
	c.instruction_StackPush(c.AF)
}

// OR n
func (c *CPU) opcode0xF6() {
	c.instruction_OR(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 30H
func (c *CPU) opcode0xF7() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0030)
}

// LD HL, SP+n
func (c *CPU) opcode0xF8() {
	var n int8 = int8(c.MemoryRead(c.PC.Get()))
	var result uint16 = uint16(int16(c.SP.Get()) + int16(n))
	c.ClearAllFlags()

	if ((c.SP.Get() ^ uint16(n) ^ result) & 0x100) == 0x100 {
		c.SetFlag(CPUFLAG_CARRY)
	}
	if ((c.SP.Get() ^ uint16(n) ^ result) & 0x10) == 0x10 {
		c.SetFlag(CPUFLAG_HALF)
	}
	c.HL.Set(result)
	c.PC.Increment()
}

// LD SP, HL
func (c *CPU) opcode0xF9() {
	c.SP.Set(c.HL.Get())
}

// LD A, (nn)
func (c *CPU) opcode0xFA() {
	tmp := NewCPURegister()
	tmp.SetLow(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
	tmp.SetHigh(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
	c.instruction_LD_r8_d8(c.AF.GetHighPointer(), tmp.Get())
}

// EI
func (c *CPU) opcode0xFB() {
	var ei_cycles int = int(OPCodeMachineCycles[0xFB]) * c.adjustedCycles(4)
	c.IMECycles = ei_cycles + 1
}

// Invalid Opcode
func (c *CPU) opcode0xFC() {
	c.instruction_InvalidOpcode()
}

// Invalid Opcode
func (c *CPU) opcode0xFD() {
	c.instruction_InvalidOpcode()
}

// CP n
func (c *CPU) opcode0xFE() {
	c.instruction_CP(c.MemoryRead(c.PC.Get()))
	c.PC.Increment()
}

// RST 38H
func (c *CPU) opcode0xFF() {
	c.instruction_StackPush(c.PC)
	c.PC.Set(0x0038)
}
