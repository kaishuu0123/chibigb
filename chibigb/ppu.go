package chibigb

const GAMEBOY_WIDTH = 160
const GAMEBOY_HEIGHT = 144

type GameBoyColorPixelFormat int

const (
	GBC_PIXEL_RGB565 GameBoyColorPixelFormat = 0
	GBC_PIXEL_RGB555                         = 1
	GBC_PIXEL_BGR565                         = 2
	GBC_PIXEL_BGR555                         = 3
)

type PPU struct {
	console *Console

	SpriteXCacheBuffer [GAMEBOY_WIDTH * GAMEBOY_HEIGHT]int
	ColorCacheBuffer   [GAMEBOY_WIDTH * GAMEBOY_HEIGHT]byte
	FrameBuffer        [GAMEBOY_WIDTH * GAMEBOY_HEIGHT]byte
	ColorFrameBuffer   [GAMEBOY_WIDTH * GAMEBOY_HEIGHT]uint16

	statusMode              int
	statusModeCounter       int
	statusModeCounterAUX    int
	statusModeLYCounter     int
	screenEnableDelayCycles int
	statusVBlankLine        int
	pixelCounter            int
	tileCycleCounter        int
	screenEnabled           bool

	GameBoyColorSpritePalettes     [8][4][2]uint16
	GameBoyColorBackgroundPalettes [8][4][2]uint16

	scanLineTransfered bool

	windowLine int
	hideFrames int

	irq48Signal             byte
	gameBoyColorPixelFormat GameBoyColorPixelFormat
}

func NewPPU(console *Console) *PPU {
	ppu := &PPU{
		console:                 console,
		screenEnabled:           true,
		gameBoyColorPixelFormat: GBC_PIXEL_RGB565,
	}

	return ppu
}

func (p *PPU) Reset() {
	p.statusMode = 1
	p.statusModeCounter = 0
	p.statusModeCounterAUX = 0
	p.statusModeLYCounter = 144
	p.screenEnableDelayCycles = 0
	p.statusVBlankLine = 0
	p.windowLine = 0
	p.pixelCounter = 0
	p.tileCycleCounter = 0
	p.screenEnabled = true
	p.scanLineTransfered = false
	p.hideFrames = 0
	p.irq48Signal = 0
}

func (p *PPU) Step(clockCycles *byte) bool {
	var vblank bool = false
	p.statusModeCounter += int(*clockCycles)

	if p.screenEnabled {
		switch p.statusMode {
		// During H-BLANK
		case 0:
			if p.statusModeCounter >= 204 {
				p.statusModeCounter -= 204
				p.statusMode = 2

				p.statusModeLYCounter++
				p.console.Memory.Load(0xFF44, byte(p.statusModeLYCounter))
				p.CompareLYToLYC()

				if p.statusModeLYCounter == 144 {
					p.statusMode = 1
					p.statusVBlankLine = 0
					p.statusModeCounterAUX = p.statusModeCounter

					p.console.CPU.RequestInterrupt(INTERRUPT_VBLANK)

					p.irq48Signal &= 0x09
					var stat byte = p.console.Memory.Retrieve(0xFF41)
					if CheckBit(stat, 4) {
						if CheckBit(p.irq48Signal, 0) == false && CheckBit(p.irq48Signal, 3) == false {
							p.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
						}
						p.irq48Signal = SetBit(p.irq48Signal, 1)
					}
					p.irq48Signal &= 0x0E

					if p.hideFrames > 0 {
						p.hideFrames--
					} else {
						vblank = true
					}
					p.windowLine = 0
				} else {
					p.irq48Signal &= 0x09
					var stat byte = p.console.Memory.Retrieve(0xFF41)
					if CheckBit(stat, 5) {
						if p.irq48Signal == 0 {
							p.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
						}
						p.irq48Signal = SetBit(p.irq48Signal, 2)
					}
					p.irq48Signal &= 0x0E
				}

				p.UpdateStatRegister()
			}
		// During V-BLANK
		case 1:
			p.statusModeCounterAUX += int(*clockCycles)
			if p.statusModeCounterAUX >= 456 {
				p.statusModeCounterAUX -= 456
				p.statusVBlankLine++

				if p.statusVBlankLine <= 9 {
					p.statusModeLYCounter++
					p.console.Memory.Load(0xFF44, byte(p.statusModeLYCounter))
					p.CompareLYToLYC()
				}
			}

			if p.statusModeCounter >= 4104 && p.statusModeCounterAUX >= 4 && p.statusModeLYCounter == 153 {
				p.statusModeLYCounter = 0
				p.console.Memory.Load(0xFF44, byte(p.statusModeLYCounter))
				p.CompareLYToLYC()
			}

			if p.statusModeCounter >= 4560 {
				p.statusModeCounter -= 4560
				p.statusMode = 2
				p.UpdateStatRegister()
				p.irq48Signal &= 0x07
				p.irq48Signal &= 0x0A

				var stat byte = p.console.Memory.Retrieve(0xFF41)
				if CheckBit(stat, 5) {
					if p.irq48Signal == 0 {
						p.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
					}
					p.irq48Signal = SetBit(p.irq48Signal, 2)
				}
				p.irq48Signal &= 0x0D
			}
		// During searching OAM RAM
		case 2:
			if p.statusModeCounter >= 80 {
				p.statusModeCounter -= 80
				p.statusMode = 3
				p.scanLineTransfered = false
				p.irq48Signal &= 0x08
				p.UpdateStatRegister()
			}
		// During transfering data to LCD driver
		case 3:
			if p.pixelCounter < 160 {
				p.tileCycleCounter += int(*clockCycles)
				var lcdc byte = p.console.Memory.Retrieve(0xFF40)

				if p.screenEnabled && CheckBit(lcdc, 7) {
					for p.tileCycleCounter >= 3 {
						p.RenderBG(p.statusModeLYCounter, p.pixelCounter)

						p.pixelCounter += 4
						p.tileCycleCounter -= 3
						if p.pixelCounter >= 160 {
							break
						}
					}
				}
			}

			if p.statusModeCounter >= 160 && p.scanLineTransfered == false {
				p.ScanLine(p.statusModeLYCounter)
				p.scanLineTransfered = true
			}

			if p.statusModeCounter >= 172 {
				p.pixelCounter = 0
				p.statusModeCounter -= 172
				p.statusMode = 0
				p.tileCycleCounter = 0
				p.UpdateStatRegister()

				p.irq48Signal &= 0x08
				var stat byte = p.console.Memory.Retrieve(0xFF41)
				if CheckBit(stat, 3) {
					if CheckBit(p.irq48Signal, 3) == false {
						p.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
					}
					p.irq48Signal = SetBit(p.irq48Signal, 0)
				}
			}
		}
	} else {
		// Screen disabled
		if p.screenEnableDelayCycles > 0 {
			p.screenEnableDelayCycles -= int(*clockCycles)

			if p.screenEnableDelayCycles <= 0 {
				p.screenEnableDelayCycles = 0
				p.screenEnabled = true
				p.hideFrames = 3
				p.statusMode = 0
				p.statusModeCounter = 0
				p.statusModeCounterAUX = 0
				p.statusModeLYCounter = 0
				p.windowLine = 0
				p.statusVBlankLine = 0
				p.pixelCounter = 0
				p.tileCycleCounter = 0
				p.console.Memory.Load(0xFF44, byte(p.statusModeLYCounter))
				p.irq48Signal = 0

				var stat byte = p.console.Memory.Retrieve(0xFF41)
				if CheckBit(stat, 5) {
					p.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
					p.irq48Signal = SetBit(p.irq48Signal, 2)
				}

				p.CompareLYToLYC()
			}
		} else if p.statusModeCounter >= 70224 {
			p.statusModeCounter -= 70224
			vblank = true
		}
	}

	return vblank
}

func (p *PPU) CompareLYToLYC() {
	if p.screenEnabled {
		var lyc byte = p.console.Memory.Retrieve(0xFF45)
		var stat byte = p.console.Memory.Retrieve(0xFF41)

		if lyc == byte(p.statusModeLYCounter) {
			stat = SetBit(stat, 2)
			if CheckBit(stat, 6) {
				if p.irq48Signal == 0 {
					p.console.CPU.RequestInterrupt(INTERRUPT_LCDSTAT)
				}
				p.irq48Signal = SetBit(p.irq48Signal, 3)
			}
		} else {
			stat = ClearBit(stat, 2)
			p.irq48Signal = ClearBit(p.irq48Signal, 3)
		}

		p.console.Memory.Load(0xFF41, stat)
	}
}

func (p *PPU) UpdateStatRegister() {
	// Updates the STAT register with current mode
	var stat byte = p.console.Memory.Retrieve(0xFF41)
	p.console.Memory.Load(0xFF41, (stat&0xFC)|(byte(p.statusMode)&0x03))
}

func (p *PPU) RenderBG(line int, pixel int) {
	var lcdc byte = p.console.Memory.Retrieve(0xFF40)
	var lineWidth int = (line * GAMEBOY_WIDTH)

	if CheckBit(lcdc, 0) {
		var pixelsToRender int = 4
		var offsetXInit int = pixel & 0x07
		var offsetXEnd int = offsetXInit + pixelsToRender
		var screenTile int = pixel >> 3
		var tileStartAddr int
		if CheckBit(lcdc, 4) {
			tileStartAddr = 0x8000
		} else {
			tileStartAddr = 0x8800
		}
		var mapStartAddr int
		if CheckBit(lcdc, 3) {
			mapStartAddr = 0x9C00
		} else {
			mapStartAddr = 0x9800
		}
		var scrollX byte = p.console.Memory.Retrieve(0xFF43)
		var scrollY byte = p.console.Memory.Retrieve(0xFF42)
		var lineScrolled byte = byte(line) + scrollY
		var lineScrolled32 int = (int(lineScrolled) >> 3) << 5
		var tilePixelY int = int(lineScrolled) & 0x07
		var tilePixelY2 int = int(tilePixelY) << 1
		var palette byte = p.console.Memory.Retrieve(0xFF47)

		for offsetX := offsetXInit; offsetX < offsetXEnd; offsetX++ {
			var screenPixelX int = (screenTile << 3) + offsetX
			var mapPixelX byte = byte(screenPixelX) + scrollX
			var mapTileX int = int(mapPixelX) >> 3
			var mapTileOffsetX int = int(mapPixelX) & 0x07
			var mapTileAddr uint16 = uint16(mapStartAddr + lineScrolled32 + mapTileX)
			var mapTile int = 0

			if tileStartAddr == 0x8800 {
				mapTile = int(int8(p.console.Memory.Retrieve(mapTileAddr)))
				mapTile += 128
			} else {
				mapTile = int(p.console.Memory.Retrieve(mapTileAddr))
			}

			var mapTile16 int = mapTile << 4
			var byte1 byte = 0
			var byte2 byte = 0
			var finalPixelY2 int

			finalPixelY2 = tilePixelY2

			var tileAddress int = tileStartAddr + mapTile16 + finalPixelY2

			byte1 = p.console.Memory.Retrieve(uint16(tileAddress))
			byte2 = p.console.Memory.Retrieve(uint16(tileAddress + 1))

			var pixelXInTile int = mapTileOffsetX
			var pixelXInTileBit int = 0x01 << (7 - pixelXInTile)
			var pixelData int
			if (byte1 & byte(pixelXInTileBit)) != 0 {
				pixelData = 1
			} else {
				pixelData = 0
			}
			if (byte2 & byte(pixelXInTileBit)) != 0 {
				pixelData |= 2
			} else {
				pixelData |= 0
			}

			var index int = lineWidth + screenPixelX
			p.ColorCacheBuffer[index] = byte(pixelData & 0x03)

			var color byte = (palette >> (pixelData << 1)) & 0x03
			p.ColorFrameBuffer[index] = uint16(color)
			p.FrameBuffer[index] = color
		}
	} else {
		for x := 0; x < 4; x++ {
			var position int = lineWidth + pixel + x
			p.FrameBuffer[position] = 0
			p.ColorCacheBuffer[position] = 0
		}
	}
}

func (p *PPU) ScanLine(line int) {
	var lcdc byte = p.console.Memory.Retrieve(0xFF40)

	if p.screenEnabled && CheckBit(lcdc, 7) {
		p.RenderWindow(line)
		p.RenderSprites(line)
	} else {
		var lineWidth int = (line * GAMEBOY_WIDTH)
		for x := 0; x < GAMEBOY_WIDTH; x++ {
			p.FrameBuffer[lineWidth+x] = 0x0
		}
	}
}

func (p *PPU) RenderWindow(line int) {
	if p.windowLine > 143 {
		return
	}

	var lcdc byte = p.console.Memory.Retrieve(0xFF40)
	if CheckBit(lcdc, 5) == false {
		return
	}

	var wx int = int(p.console.Memory.Retrieve(0xFF4B)) - 7
	if wx > 159 {
		return
	}

	var wy byte = p.console.Memory.Retrieve(0xFF4A)
	if (wy > 143) || (wy > byte(line)) {
		return
	}

	var tiles int
	if CheckBit(lcdc, 4) {
		tiles = 0x8000
	} else {
		tiles = 0x8800
	}
	var maps int
	if CheckBit(lcdc, 6) {
		maps = 0x9C00
	} else {
		maps = 0x9800
	}
	var lineAdjusted int = p.windowLine
	var y32 int = (lineAdjusted >> 3) << 5
	var pixelY = lineAdjusted & 0x07
	var pixelY2 = pixelY << 1
	var lineWidth int = line * GAMEBOY_WIDTH
	var palette byte = p.console.Memory.Retrieve(0xFF47)

	for x := 0; x < 32; x++ {
		var tile int = 0
		if tiles == 0x8800 {
			tile = int(int8(p.console.Memory.Retrieve(uint16(maps + y32 + x))))
			tile += 128
		} else {
			tile = int(p.console.Memory.Retrieve(uint16(maps + y32 + x)))
		}

		var mapOffsetX int = x << 3
		var tile16 = tile << 4
		var byte1 byte = 0
		var byte2 byte = 0
		var finalPixelY2 int
		finalPixelY2 = pixelY2
		var tileAddress int = tiles + tile16 + finalPixelY2

		byte1 = p.console.Memory.Retrieve(uint16(tileAddress))
		byte2 = p.console.Memory.Retrieve(uint16(tileAddress + 1))

		for pixelX := 0; pixelX < 8; pixelX++ {
			var bufferX int = (mapOffsetX + pixelX + wx)

			if bufferX < 0 || bufferX >= GAMEBOY_WIDTH {
				continue
			}

			var pixelXPos int = pixelX
			var pixel int
			if (byte1 & (0x01 << (7 - pixelXPos))) != 0 {
				pixel = 1
			} else {
				pixel = 0
			}
			if (byte2 & (0x01 << (7 - pixelXPos))) != 0 {
				pixel |= 2
			} else {
				pixel |= 0
			}

			var position int = lineWidth + bufferX
			p.ColorCacheBuffer[position] = byte(pixel & 0x03)
			var color byte = (palette >> (pixel << 1)) & 0x03
			p.ColorFrameBuffer[position] = uint16(color)
			p.FrameBuffer[position] = color
		}
	}

	p.windowLine++
}

func (p *PPU) RenderSprites(line int) {
	var lcdc = p.console.Memory.Retrieve(0xFF40)
	if CheckBit(lcdc, 1) == false {
		return
	}

	var spriteHeight int
	if CheckBit(lcdc, 2) {
		spriteHeight = 16
	} else {
		spriteHeight = 8
	}
	var lineWidth int = line * GAMEBOY_WIDTH
	var visibleSprites [40]bool
	var spriteLimit int = 0

	for sprite := 0; sprite < 40; sprite++ {
		var sprite4 int = sprite << 2
		var spriteY int = int(p.console.Memory.Retrieve(uint16(0xFE00+sprite4))) - 16
		if spriteY > line || spriteY+spriteHeight <= line {
			visibleSprites[sprite] = false
			continue
		}

		spriteLimit++
		visibleSprites[sprite] = spriteLimit <= 10
	}

	for sprite := 39; sprite >= 0; sprite-- {
		if visibleSprites[sprite] == false {
			continue
		}

		var sprite4 int = sprite << 2
		var spriteX int = int(p.console.Memory.Retrieve(uint16(0xFE00+sprite4+1))) - 8

		if spriteX < -7 || spriteX >= GAMEBOY_WIDTH {
			continue
		}

		var spriteY int = int(p.console.Memory.Retrieve(uint16(0xFE00+sprite4))) - 16
		var spriteTile16 int
		if spriteHeight == 16 {
			spriteTile16 = (int(p.console.Memory.Retrieve(uint16(0xFE00+sprite4+2))) & 0xFE) << 4
		} else {
			spriteTile16 = (int(p.console.Memory.Retrieve(uint16(0xFE00+sprite4+2))) & 0xFF) << 4
		}
		var spriteFlags byte = p.console.Memory.Retrieve(uint16(0xFE00 + sprite4 + 3))
		var spritePalette int
		if CheckBit(spriteFlags, 4) {
			spritePalette = 1
		} else {
			spritePalette = 0
		}
		var palette byte
		if spritePalette != 0 {
			palette = p.console.Memory.Retrieve(0xFF49)
		} else {
			palette = p.console.Memory.Retrieve(0xFF48)
		}
		var xFlip bool = CheckBit(spriteFlags, 5)
		var yFlip bool = CheckBit(spriteFlags, 6)
		var aboveBG bool = !CheckBit(spriteFlags, 7)
		var tiles int = 0x8000
		var pixelY int
		if yFlip {
			if spriteHeight == 16 {
				pixelY = 15 - (line - spriteY)
			} else {
				pixelY = 7 - (line - spriteY)
			}
		} else {
			pixelY = line - spriteY
		}
		var byte1 byte = 0
		var byte2 byte = 0
		var pixelY2 int = 0
		var offset int = 0

		if spriteHeight == 16 && pixelY >= 8 {
			pixelY2 = (pixelY - 8) << 1
			offset = 16
		} else {
			pixelY2 = pixelY << 1
		}

		var tileAddress int = tiles + spriteTile16 + pixelY2 + offset

		byte1 = p.console.Memory.Retrieve(uint16(tileAddress))
		byte2 = p.console.Memory.Retrieve(uint16(tileAddress + 1))

		for pixelX := 0; pixelX < 8; pixelX++ {
			var shift int
			if xFlip {
				shift = pixelX
			} else {
				shift = 7 - pixelX
			}
			var pixel int
			if (byte1 & (0x01 << shift)) != 0 {
				pixel = 1
			} else {
				pixel = 0
			}
			if (byte2 & (0x01 << shift)) != 0 {
				pixel |= 2
			} else {
				pixel |= 0
			}

			if pixel == 0 {
				continue
			}

			var bufferX int = spriteX + pixelX

			if bufferX < 0 || bufferX >= GAMEBOY_WIDTH {
				continue
			}

			var position int = lineWidth + bufferX
			var colorCache byte = p.ColorCacheBuffer[position]

			var spriteXCache = p.SpriteXCacheBuffer[position]
			if CheckBit(colorCache, 3) && (spriteXCache < spriteX) {
				continue
			}

			if aboveBG == false && (colorCache&0x03) != 0 {
				continue
			}

			p.ColorCacheBuffer[position] = SetBit(colorCache, 3)
			p.SpriteXCacheBuffer[position] = spriteX

			var color byte = (palette >> (pixel << 1)) & 0x03
			p.ColorFrameBuffer[position] = uint16(color)
			p.FrameBuffer[position] = color
		}
	}
}

func (p *PPU) ResetWindowLine() {
	wy := p.console.Memory.Retrieve(0xFF4A)

	if p.windowLine == 0 && p.statusModeLYCounter < 144 && byte(p.statusModeLYCounter) > wy {
		p.windowLine = 144
	}
}

func (p *PPU) EnableScreen() {
	if p.screenEnabled == false {
		p.screenEnableDelayCycles = 244
	}
}

func (p *PPU) DisableScreen() {
	p.screenEnabled = false
	p.console.Memory.Load(0xFF44, 0x00)
	stat := p.console.Memory.Retrieve(0xFF41)
	stat &= 0x7C
	p.console.Memory.Load(0xFF41, stat)
	p.statusMode = 0
	p.statusModeCounter = 0
	p.statusModeCounterAUX = 0
	p.statusModeLYCounter = 0
	p.irq48Signal = 0
}

func (p *PPU) GetIRQ48Signal() byte {
	return p.irq48Signal
}

func (p *PPU) SetIRQ48Signal(signal byte) {
	p.irq48Signal = signal
}

func (p *PPU) GetCurrentStatusMode() int {
	return p.statusMode
}
