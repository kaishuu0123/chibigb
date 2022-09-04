package chibigb

type GameBoyKeys int

const (
	KEY_A      GameBoyKeys = 4
	KEY_B                  = 5
	KEY_START              = 7
	KEY_SELECT             = 6
	KEY_RIGHT              = 0
	KEY_LEFT               = 1
	KEY_UP                 = 2
	KEY_DOWN               = 3
)

type Controller struct {
	console     *Console
	joypadState byte
	p1          byte
	inputCycles int
}

func NewController(console *Console) *Controller {
	return &Controller{
		console: console,
	}
}

func (c *Controller) Reset() {
	c.joypadState = 0xFF
	c.p1 = 0xFF
	c.inputCycles = 0
}

func (c *Controller) KeyPressed(key GameBoyKeys) {
	c.joypadState = ClearBit(c.joypadState, int(key))
}

func (c *Controller) KeyReleased(key GameBoyKeys) {
	c.joypadState = SetBit(c.joypadState, int(key))
}

func (c *Controller) Update() {
	var current byte = c.p1 & 0xF0

	switch current & 0x30 {
	case 0x10:
		var topJoypad byte = (c.joypadState >> 4) & 0x0F
		current |= topJoypad
	case 0x20:
		var bottomJoypad byte = c.joypadState & 0x0F
		current |= bottomJoypad
	case 0x30:
		current |= 0x0F
	}

	if (c.p1 & ^current & 0x0F) != 0 {
		c.console.CPU.RequestInterrupt(INTERRUPT_JOYPAD)
	}

	c.p1 = current
}

func (c *Controller) Step(clockCycles byte) {
	c.inputCycles += int(clockCycles)

	// Joypad Poll Speed (64 Hz)
	if c.inputCycles >= 65536 {
		c.inputCycles -= 65536
		c.Update()
	}
}

func (c *Controller) Read() byte {
	return c.p1
}

func (c *Controller) Write(value byte) {
	c.p1 = (c.p1 & 0xCF) | (value & 0x30)
	c.Update()
}
