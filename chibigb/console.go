package chibigb

import (
	"image/color"
)

type Console struct {
	CPU    *CPU
	Memory *Memory
	PPU    *PPU
	APU    *APU

	Cartridge  *Cartridge
	Controller *Controller
}

func NewConsole() *Console {
	console := &Console{}
	console.Cartridge = NewCartridge(console)

	console.CPU = NewCPU(console)
	console.PPU = NewPPU(console)
	console.APU = NewAPU(console)
	console.Memory = NewMemory(console)
	console.Controller = NewController(console)

	console.CPU.Reset()
	console.PPU.Reset()
	console.APU.Reset()
	console.Memory.Reset()
	console.Controller.Reset()

	return console
}

func (console *Console) RunToVBlank() {
	var vblank bool = false
	var totalClocks uint64 = 0
	var clockCycles byte

	for vblank == false {
		clockCycles = console.CPU.Step(75)
		console.CPU.UpdateTimers(clockCycles)
		console.CPU.UpdateSerial(clockCycles)
		vblank = console.PPU.Step(&clockCycles)
		console.APU.Step(clockCycles)
		console.Controller.Step(clockCycles)
		totalClocks += uint64(clockCycles)

		if totalClocks > 702240 {
			vblank = true
		}
	}
}

func (console *Console) LoadFromBuffer(data []byte) {
	console.Cartridge.LoadFromBuffer(data)

	// loads the first 32KB only (bank 0 and 1)
	for i := 0; i < 0x8000; i++ {
		console.Memory.Load(uint16(i), console.Cartridge.ROM[i])
	}

	console.Memory.SetMemoryBankController()
}

func (console *Console) SetPixels(pixels []byte) {
	for i := 0; i < GAMEBOY_WIDTH*GAMEBOY_HEIGHT; i++ {
		var color color.RGBA = console.GetLCDColor(i)
		// R, G, B, A (4)
		pixels[(i * 4)] = color.R
		pixels[(i*4)+1] = color.G
		pixels[(i*4)+2] = color.B
		pixels[(i*4)+3] = color.A
	}
}

func (console *Console) SetSoundBuffer(buffer []byte) {
	buffer = console.APU.ReadSoundBuffer(buffer)
}

func (console *Console) GetLCDColor(index int) color.RGBA {
	var color color.RGBA
	color.A = 0xFF

	switch console.PPU.FrameBuffer[index] {
	case 0:
		color.R = 0xC6
		color.G = 0xDE
		color.B = 0x8C
	case 1:
		color.R = 0x84
		color.G = 0xA5
		color.B = 0x63
	case 2:
		color.R = 0x39
		color.G = 0x61
		color.B = 0x39
	case 3:
		color.R = 0x08
		color.G = 0x18
		color.B = 0x10
	}

	return color
}

func (console *Console) SetButtonState(button int, pressed bool) {
	// set key in constroller
	if pressed {
		console.Controller.KeyPressed(GameBoyKeys(button))
	} else {
		console.Controller.KeyReleased(GameBoyKeys(button))
	}
}

func CheckBit(value byte, bit int) bool {
	return (value & (0x01 << bit)) != 0
}

func SetBit(value byte, bit int) byte {
	return (value | (0x01 << bit))
}

func ClearBit(value byte, bit int) byte {
	return (value & ^(0x01 << bit))
}
