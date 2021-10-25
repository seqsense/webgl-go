package webgl

type BufferData interface {
	Bytes() []byte
}

type Float32ArrayBuffer []float32

func (b Float32ArrayBuffer) Bytes() []byte {
	return float32SliceAsByteSlice([]float32(b))
}

type ByteArrayBuffer []byte

func (b ByteArrayBuffer) Bytes() []byte {
	return b
}

func (b ByteArrayBuffer) UInt32Slice() []uint32 {
	return byteSliceAsUInt32Slice(b)
}

type Uint16ArrayBuffer []uint16

func (b Uint16ArrayBuffer) Bytes() []byte {
	return uint16SliceAsByteSlice([]uint16(b))
}
