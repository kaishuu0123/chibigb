package chibigb

import "log"

const (
	// must be power of two
	SAMPLES_PER_SECONDS = 44100
	CLOCKS_PER_SECONDS  = 4194304
	CLOCKS_PER_SAMPLE   = CLOCKS_PER_SECONDS / SAMPLES_PER_SECONDS
)

type APU struct {
	console *Console

	buffer      *APURingBuffer
	leftSample  uint32
	rightSample uint32
	numSamples  uint32

	lastLeft           float64
	lastRight          float64
	lastCorrectedLeft  float64
	lastCorrectedRight float64

	allSoundsOn bool

	sweepTimeCounter    int
	envelopeTimeCounter int
	lengthTimeCounter   int

	waves [4]APUWave

	// cart chip sounds. never used by any game?
	vInToLeftSpeaker  bool
	vInToRightSpeaker bool

	leftSpeakerVolume  byte // left=S02 in docs
	rightSpeakerVolume byte // right=S1 in docs
}

func NewAPU(console *Console) *APU {
	apu := &APU{
		console: console,
		buffer:  NewAPURingBuffer(),
	}

	apu.waves[0].waveType = APU_WAVE_FORM_SQUARE
	apu.waves[1].waveType = APU_WAVE_FORM_SQUARE
	apu.waves[2].waveType = APU_WAVE_FORM_WAVE
	apu.waves[3].waveType = APU_WAVE_FORM_NOISE

	apu.waves[3].polyFeedbackRegister = 0x01

	return apu
}

func (apu *APU) Reset() {
	// NOTHING DONE
}

func (apu *APU) Step(cycleCount byte) {
	for i := 0; i < int(cycleCount); i++ {
		apu.runCycle()
	}
}

func (apu *APU) runCycle() {
	apu.lengthTimeCounter++
	if apu.lengthTimeCounter >= 16384 {
		apu.runLengthCycle()
		apu.lengthTimeCounter = 0
	}

	apu.envelopeTimeCounter++
	if apu.envelopeTimeCounter >= 65536 {
		apu.runEnvelopeCycle()
		apu.envelopeTimeCounter = 0
	}

	if !apu.buffer.full() {
		apu.generateSample()
	}

	apu.sweepTimeCounter++
	if apu.sweepTimeCounter >= 32768 {
		apu.waves[0].runSweepCycle()
		apu.sweepTimeCounter = 0
	}
}

func (apu *APU) Read(addr uint16) byte {
	switch addr {
	case 0xFF10:
		return apu.waves[0].readSweepRegister()
	case 0xFF11:
		return apu.waves[0].readLengthDutyRegister()
	case 0xFF12:
		return apu.waves[0].readEnvelopeRegister()
	case 0xFF13:
		return apu.waves[0].readFreqLowRegister()
	case 0xFF14:
		return apu.waves[0].readFreqHighRegister()

	case 0xFF15:
		// unmapped bytes
		return 0xFF

	case 0xFF16:
		return apu.waves[1].readLengthDutyRegister()
	case 0xFF17:
		return apu.waves[1].readEnvelopeRegister()
	case 0xFF18:
		return apu.waves[1].readFreqLowRegister()
	case 0xFF19:
		return apu.waves[1].readFreqHighRegister()

	case 0xFF1A:
		if apu.waves[2].enabled {
			return (0x01 << 7) | 0x7F
		} else {
			return 0x7F
		}

	case 0xFF1B:
		return apu.waves[2].readLengthDataRegister()
	case 0xFF1C:
		return apu.waves[2].readWaveOutLevelRegister()
	case 0xFF1D:
		return apu.waves[2].readFreqLowRegister()
	case 0xFF1E:
		return apu.waves[2].readFreqHighRegister()

	case 0xFF1F:
		// unmapped bytes
		return 0xFF

	case 0xFF20:
		return apu.waves[3].readLengthDataRegister()
	case 0xFF21:
		return apu.waves[3].readEnvelopeRegister()
	case 0xFF22:
		return apu.waves[3].readPolyCounterRegister()
	case 0xFF23:
		return apu.waves[3].readFreqHighRegister()

	case 0xFF24:
		return apu.readVolumeRegister()
	case 0xFF25:
		return apu.readSpeakerSelectRegister()
	case 0xFF26:
		return apu.readEnabledRegister()
	}

	switch {
	case addr >= 0xFF27 && addr < 0xFF30:
		return 0xFF
	case addr >= 0xFF30 && addr < 0xFF40:
		return apu.waves[2].readWavePatternValue(addr)
	}

	log.Fatalf("Unknown address %X\n", addr)
	return 0x00
}

func (apu *APU) Write(addr uint16, value byte) {
	switch addr {
	case 0xFF10:
		apu.waves[0].writeSweepRegister(value)
	case 0xFF11:
		apu.waves[0].writeLengthDutyRegister(value)
	case 0xFF12:
		apu.waves[0].writeEnvelopeRegister(value)
	case 0xFF13:
		apu.waves[0].writeFreqLowRegister(value)
	case 0xFF14:
		apu.waves[0].writeFreqHighRegister(value)

	case 0xFF15:
		// NOTHING DONE

	case 0xFF16:
		apu.waves[1].writeLengthDutyRegister(value)
	case 0xFF17:
		apu.waves[1].writeEnvelopeRegister(value)
	case 0xFF18:
		apu.waves[1].writeFreqLowRegister(value)
	case 0xFF19:
		apu.waves[1].writeFreqHighRegister(value)

	case 0xFF1A:
		apu.waves[2].writeEnabledRegister(value)
	case 0xFF1B:
		apu.waves[2].writeLengthDataRegister(value)
	case 0xFF1C:
		apu.waves[2].writeWaveOutLevelRegister(value)
	case 0xFF1D:
		apu.waves[2].writeFreqLowRegister(value)
	case 0xFF1E:
		apu.waves[2].writeFreqHighRegister(value)

	case 0xFF1F:
		// NOTHING DONE

	case 0xFF20:
		apu.waves[3].writeLengthDataRegister(value)
	case 0xFF21:
		apu.waves[3].writeEnvelopeRegister(value)
	case 0xFF22:
		apu.waves[3].writePolyCounterRegister(value)
	case 0xFF23:
		apu.waves[3].writeFreqHighRegister(value)

	case 0xFF24:
		apu.writeVolumeRegister(value)
	case 0xFF25:
		apu.writeSpeakerSelectRegister(value)
	case 0xFF26:
		apu.writeEnabledRegister(value)
	}

	switch {
	case addr >= 0xFF27 && addr < 0xFF30:
		// NOTHING DONE
	case addr >= 0xFF30 && addr < 0xFF40:
		apu.waves[2].writeWavePatternValue(addr, value)
	}
}

func (apu *APU) ReadSoundBuffer(toFill []byte) []byte {
	if int(apu.buffer.size()) < len(toFill) {
	}
	for int(apu.buffer.size()) < len(toFill) {
		apu.generateSample()
	}
	return apu.buffer.Read(toFill)
}

func (apu *APU) generateSample() {
	apu.runFreqCycle()

	leftSam, rightSam := uint32(0), uint32(0)
	if apu.allSoundsOn {

		// IMPORTANT TODO:
		// probably have to reintroduce wave bias fix
		// from float-land. should probably center at
		// 7.5 and let the end dc blocker get it around zero?

		left0, right0 := apu.waves[0].GetSample()
		left1, right1 := apu.waves[1].GetSample()
		left2, right2 := apu.waves[2].GetSample()
		left3, right3 := apu.waves[3].GetSample()
		leftSam += uint32(left0 + left1 + left2 + left3)
		rightSam += uint32(right0 + right1 + right2 + right3)
		leftSam *= uint32(apu.leftSpeakerVolume + 1)
		rightSam *= uint32(apu.rightSpeakerVolume + 1)
		// will need to div by 4*8*15
	}
	apu.leftSample += leftSam
	apu.rightSample += rightSam
	apu.numSamples++

	if apu.numSamples >= CLOCKS_PER_SAMPLE {
		left := float64(apu.leftSample) / float64(apu.numSamples)
		right := float64(apu.rightSample) / float64(apu.numSamples)
		left /= 4 * 8 * 15
		right /= 4 * 8 * 15

		// dc blocker to center waveform
		correctedLeft := left - apu.lastLeft + 0.995*apu.lastCorrectedLeft
		apu.lastCorrectedLeft = correctedLeft
		apu.lastLeft = left
		left = correctedLeft

		correctedRight := right - apu.lastRight + 0.995*apu.lastCorrectedRight
		apu.lastCorrectedRight = correctedRight
		apu.lastRight = right
		right = correctedRight

		iSampleL, iSampleR := int16(left*32767.0), int16(right*32767.0)
		apu.buffer.Write([]byte{
			byte(iSampleL & 0xff),
			byte(iSampleL >> 8),
			byte(iSampleR & 0xff),
			byte(iSampleR >> 8),
		})

		apu.leftSample = 0
		apu.rightSample = 0
		apu.numSamples = 0
	}
}

func (apu *APU) runFreqCycle() {
	apu.waves[0].runFreqCycle()
	apu.waves[1].runFreqCycle()
	apu.waves[2].runFreqCycle()
	apu.waves[3].runFreqCycle()
}

func (apu *APU) runLengthCycle() {
	apu.waves[0].runLengthCycle()
	apu.waves[1].runLengthCycle()
	apu.waves[2].runLengthCycle()
	apu.waves[3].runLengthCycle()
}

func (apu *APU) runEnvelopeCycle() {
	apu.waves[0].runEnvelopeCycle()
	apu.waves[1].runEnvelopeCycle()
	apu.waves[2].runEnvelopeCycle()
	apu.waves[3].runEnvelopeCycle()
}

func (apu *APU) readVolumeRegister() byte {
	value := apu.rightSpeakerVolume<<4 | apu.leftSpeakerVolume
	if apu.vInToLeftSpeaker {
		value |= 0x80
	}
	if apu.vInToRightSpeaker {
		value |= 0x08
	}
	return value
}

func (apu *APU) writeVolumeRegister(value byte) {
	apu.vInToLeftSpeaker = value&0x80 != 0
	apu.vInToRightSpeaker = value&0x08 != 0
	apu.rightSpeakerVolume = (value >> 4) & 0x07
	apu.leftSpeakerVolume = value & 0x07
}

func (apu *APU) readSpeakerSelectRegister() byte {
	var result byte

	if apu.waves[0].rightSpeakerEnabled {
		result |= 1
	}
	if apu.waves[1].rightSpeakerEnabled {
		result |= 1 << 1
	}
	if apu.waves[2].rightSpeakerEnabled {
		result |= 1 << 2
	}
	if apu.waves[3].rightSpeakerEnabled {
		result |= 1 << 3
	}
	if apu.waves[0].leftSpeakerEnabled {
		result |= 1 << 4
	}
	if apu.waves[1].leftSpeakerEnabled {
		result |= 1 << 5
	}
	if apu.waves[2].leftSpeakerEnabled {
		result |= 1 << 6
	}
	if apu.waves[3].leftSpeakerEnabled {
		result |= 1 << 7
	}

	return result
}

func (apu *APU) writeSpeakerSelectRegister(value byte) {
	apu.waves[0].rightSpeakerEnabled = (value & 0x01) != 0
	apu.waves[1].rightSpeakerEnabled = (value & 0x02) != 0
	apu.waves[2].rightSpeakerEnabled = (value & 0x04) != 0
	apu.waves[3].rightSpeakerEnabled = (value & 0x08) != 0
	apu.waves[0].leftSpeakerEnabled = (value & 0x10) != 0
	apu.waves[1].leftSpeakerEnabled = (value & 0x20) != 0
	apu.waves[2].leftSpeakerEnabled = (value & 0x40) != 0
	apu.waves[3].leftSpeakerEnabled = (value & 0x80) != 0
}

func (apu *APU) readEnabledRegister() byte {
	var result byte

	if apu.waves[0].enabled {
		result |= 0x01
	}
	if apu.waves[1].enabled {
		result |= 0x01 << 1
	}
	if apu.waves[2].enabled {
		result |= 0x01 << 2
	}
	if apu.waves[3].enabled {
		result |= 0x01 << 3
	}

	result |= 0x01 << 4
	result |= 0x01 << 5
	result |= 0x01 << 6
	if apu.allSoundsOn {
		result |= 0x01 << 7
	}

	return result
}

func (apu *APU) writeEnabledRegister(value byte) {
	apu.allSoundsOn = (value & 0x80) != 0
}
