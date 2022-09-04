package chibigb

const APU_RING_BUFFER_SIZE int = 16 * 512 * 4

type APURingBuffer struct {
	writeIndex uint
	readIndex  uint
	buf        [APU_RING_BUFFER_SIZE]byte
}

func NewAPURingBuffer() *APURingBuffer {
	return &APURingBuffer{}
}

func (a *APURingBuffer) Read(preSizedBuf []byte) []byte {
	readCount := 0
	for i := range preSizedBuf {
		if a.size() == 0 {
			break
		}
		preSizedBuf[i] = a.buf[a.mask(a.readIndex)]
		a.readIndex++
		readCount++
	}

	return preSizedBuf[:readCount]
}

func (a *APURingBuffer) Write(bytes []byte) (writeCount int) {
	for _, b := range bytes {
		if a.full() {
			return writeCount
		}
		a.buf[a.mask(a.writeIndex)] = b
		a.writeIndex++
		writeCount++
	}

	return writeCount
}

func (a *APURingBuffer) mask(i uint) uint {
	return i & (uint(len(a.buf)) - 1)
}

func (a *APURingBuffer) size() uint {
	return a.writeIndex - a.readIndex
}

func (a *APURingBuffer) full() bool {
	return a.size() == uint(len(a.buf))
}
