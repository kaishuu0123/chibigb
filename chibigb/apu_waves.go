package chibigb

type EnvelopeDirection bool

const (
	ENVELOPE_UP   EnvelopeDirection = true
	ENVELOPE_DOWN                   = false
)

type SweepDirection bool

const (
	SWEEP_UP   SweepDirection = false
	SWEEP_DOWN                = true
)

type APUWaveType int

const (
	APU_WAVE_FORM_SQUARE APUWaveType = 0
	APU_WAVE_FORM_WAVE               = 1
	APU_WAVE_FORM_NOISE              = 2
)

type APUWave struct {
	waveType APUWaveType

	enabled             bool
	leftSpeakerEnabled  bool
	rightSpeakerEnabled bool

	envelopeDirection  EnvelopeDirection
	envelopeStartValue byte
	envelopeSweepValue byte
	currentEnvelope    byte
	envelopeCounter    byte

	timer        uint32
	freqDivider  uint32
	freqRegister uint16

	sweepCounter   byte
	sweepDirection SweepDirection
	sweepTime      byte
	sweepShift     byte

	lengthData    uint16
	currentLength uint16

	waveDuty                byte
	waveDutySequenceCounter byte

	waveOutLevel      byte // waves[2] only
	wavePatternRAM    [16]byte
	wavePatternCursor byte

	polyFeedbackRegister uint16 // sound[3] only
	polyDivisorShift     byte
	polyDivisorBase      byte
	poly7BitMode         bool
	polySample           byte

	playsContinuously bool
	restartRequested  bool
}

func (a *APUWave) runFreqCycle() {
	a.timer++

	if a.timer >= a.freqDivider {
		a.timer = 0
		switch a.waveType {
		case APU_WAVE_FORM_SQUARE:
			a.waveDutySequenceCounter = (a.waveDutySequenceCounter + 1) & 0x07
		case APU_WAVE_FORM_WAVE:
			a.wavePatternCursor = (a.wavePatternCursor + 1) & 0x1F
		case APU_WAVE_FORM_NOISE:
			a.updatePolyCounter()
		}
	}
}

func (a *APUWave) runLengthCycle() {
	if a.currentLength > 0 && !a.playsContinuously {
		a.currentLength--
		if a.currentLength == 0 {
			a.enabled = false
		}
	}
	if a.restartRequested {
		a.enabled = true
		a.restartRequested = false
		if a.lengthData == 0 {
			if a.waveType == APU_WAVE_FORM_WAVE {
				a.lengthData = 256
			} else {
				a.lengthData = 64
			}
		}
		a.currentLength = a.lengthData
		a.currentEnvelope = a.envelopeStartValue
		a.sweepCounter = 0
		a.wavePatternCursor = 0
		a.polyFeedbackRegister = 0xFFFF
	}
}

func (a *APUWave) runSweepCycle() {
	if a.sweepTime != 0 {
		if a.sweepCounter < a.sweepTime {
			a.sweepCounter++
		} else {
			a.sweepCounter = 0
			var nextFreq uint16
			if a.sweepDirection == SWEEP_UP {
				nextFreq = a.freqRegister + (a.freqRegister >> uint16(a.sweepShift))
			} else {
				nextFreq = a.freqRegister - (a.freqRegister >> uint16(a.sweepShift))
			}
			if nextFreq > 2047 {
				a.enabled = false
			} else {
				a.freqRegister = nextFreq
				a.updateFreq()
			}
		}
	}
}

func (a *APUWave) runEnvelopeCycle() {
	// more complicated, see GBSOUND
	if a.envelopeSweepValue != 0 {
		if a.envelopeCounter < a.envelopeSweepValue {
			a.envelopeCounter++
		} else {
			a.envelopeCounter = 0
			if a.envelopeDirection == ENVELOPE_UP {
				if a.currentEnvelope < 0x0f {
					a.currentEnvelope++
				}
			} else {
				if a.currentEnvelope > 0x00 {
					a.currentEnvelope--
				}
			}
		}
	}
}

var dutyCycleTable = [4][8]byte{
	{0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 1, 1, 1},
	{0, 1, 1, 1, 1, 1, 1, 0},
}

func (a *APUWave) inDutyCycle() bool {
	sel := a.waveDuty
	counter := a.waveDutySequenceCounter
	return dutyCycleTable[sel][counter] == 1
}

func (a *APUWave) GetSample() (byte, byte) {
	sample := byte(0)
	if a.enabled {
		switch a.waveType {
		case APU_WAVE_FORM_SQUARE:
			vol := a.currentEnvelope
			if a.inDutyCycle() {
				sample = vol
			} else {
				sample = 0
			}
		case APU_WAVE_FORM_WAVE:
			if a.waveOutLevel > 0 {
				sampleByte := a.wavePatternRAM[a.wavePatternCursor/2]
				if a.wavePatternCursor&1 == 0 {
					sample = sampleByte >> 4
				} else {
					sample = sampleByte & 0x0f
				}
			}
		case APU_WAVE_FORM_NOISE:
			if a.freqDivider > 0 {
				vol := a.currentEnvelope
				sample = vol * a.polySample
			}
		}
	}

	left, right := byte(0), byte(0)
	if a.leftSpeakerEnabled {
		left = sample
	}
	if a.rightSpeakerEnabled {
		right = sample
	}

	return left, right
}

func (a *APUWave) updatePolyCounter() {
	newHigh := (a.polyFeedbackRegister & 0x01) ^ ((a.polyFeedbackRegister >> 1) & 0x01)

	a.polyFeedbackRegister >>= 1
	a.polyFeedbackRegister &^= 1 << 14
	a.polyFeedbackRegister |= newHigh << 14

	if a.poly7BitMode {
		a.polyFeedbackRegister &^= 1 << 6
		a.polyFeedbackRegister |= newHigh << 6
	}
	if a.polyFeedbackRegister&0x01 == 0 {
		a.polySample = 1
	} else {
		a.polySample = 0
	}
}

func (a *APUWave) updateFreq() {
	switch a.waveType {
	case APU_WAVE_FORM_SQUARE:
		a.freqDivider = 4 * (2048 - uint32(a.freqRegister)) // 32 mul for freq, div by 8 for duty seq
	case APU_WAVE_FORM_WAVE:
		a.freqDivider = 2 * (2048 - uint32(a.freqRegister))
	case APU_WAVE_FORM_NOISE:
		divider := uint32(8)
		if a.polyDivisorBase > 0 {
			if a.polyDivisorShift < 14 {
				divider = uint32(a.polyDivisorBase) << uint32(a.polyDivisorShift+4)
			} else {
				divider = 0 // invalid shift value - disable audio
			}
		}
		a.freqDivider = divider
	default:
		panic("unexpected sound type")
	}
}

func (a *APUWave) readSweepRegister() byte {
	val := a.sweepTime << 4
	val |= a.sweepShift
	if a.sweepDirection == SWEEP_DOWN {
		val |= 0x08
	}

	return val | 0x80
}

func (a *APUWave) writeSweepRegister(value byte) {
	a.sweepTime = (value >> 4) & 0x07
	a.sweepShift = value & 0x07
	if (value & 0x08) != 0 {
		a.sweepDirection = SWEEP_DOWN
	} else {
		a.sweepDirection = SWEEP_UP
	}
}

func (a *APUWave) readLengthDutyRegister() byte {
	return (a.waveDuty << 6) | 0x3f
}

func (a *APUWave) writeLengthDutyRegister(value byte) {
	a.lengthData = 64 - uint16(value&0x3f)
	a.waveDuty = value >> 6
}

func (a *APUWave) readEnvelopeRegister() byte {
	value := a.envelopeStartValue<<4 | a.envelopeSweepValue
	if a.envelopeDirection == ENVELOPE_UP {
		value |= 0x08
	}

	return value
}

func (a *APUWave) writeEnvelopeRegister(value byte) {
	a.envelopeStartValue = value >> 4
	if a.envelopeStartValue == 0 {
		a.enabled = false
	}
	if (value & 0x08) != 0 {
		a.envelopeDirection = ENVELOPE_UP
	} else {
		a.envelopeDirection = ENVELOPE_DOWN
	}
	a.envelopeSweepValue = value & 0x07
}

func (a *APUWave) readWaveOutLevelRegister() byte {
	return (a.waveOutLevel << 5) | 0x9F
}

func (a *APUWave) writeWaveOutLevelRegister(value byte) {
	a.waveOutLevel = (value >> 5) & 0x03
}

func (a *APUWave) readLengthDataRegister() byte {
	switch a.waveType {
	case APU_WAVE_FORM_WAVE:
		return byte(256 - a.lengthData)
	case APU_WAVE_FORM_NOISE:
		return byte(64 - a.lengthData)
	}

	panic("readLengthDataRegister: unexpected sound type.")
}

func (a *APUWave) writeLengthDataRegister(value byte) {
	switch a.waveType {
	case APU_WAVE_FORM_WAVE:
		a.lengthData = 256 - uint16(value)
	case APU_WAVE_FORM_NOISE:
		a.lengthData = 16 - uint16(value)
	default:
		panic("writeLengthDataRegister: unexpected sound type.")
	}
}

func (a *APUWave) readFreqLowRegister() byte {
	return 0xFF
}

func (a *APUWave) writeFreqLowRegister(value byte) {
	a.freqRegister &^= 0x00FF
	a.freqRegister |= uint16(value)
	a.updateFreq()
}

func (a *APUWave) readFreqHighRegister() byte {
	value := byte(0xFF)
	if a.playsContinuously {
		value &^= 0x40
	}
	return value
}

func (a *APUWave) writeFreqHighRegister(value byte) {
	if (value & 0x80) != 0 {
		a.restartRequested = true
	}
	a.playsContinuously = (value & 0x40) == 0
	a.freqRegister &^= 0xFF00
	a.freqRegister |= uint16(value&0x07) << 8
	a.updateFreq()
}

func (a *APUWave) readWavePatternValue(addr uint16) byte {
	return a.wavePatternRAM[addr-0xFF30]
}

func (a *APUWave) writeWavePatternValue(addr uint16, value byte) {
	a.wavePatternRAM[addr-0xFF30] = value
}

func (a *APUWave) readPolyCounterRegister() byte {
	val := byte(0)

	if a.poly7BitMode {
		val |= 0x08
	}
	val |= a.polyDivisorShift << 4
	val |= a.polyDivisorBase

	return val
}

func (a *APUWave) writePolyCounterRegister(val byte) {
	a.poly7BitMode = val&0x08 != 0
	a.polyDivisorShift = val >> 4
	a.polyDivisorBase = val & 0x07
}

func (a *APUWave) writeEnabledRegister(value byte) {
	a.enabled = (value & 0x80) != 0
}
