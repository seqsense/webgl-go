package webgl

import (
	"reflect"
	"unsafe"
)

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
	n := 2 * len(b)

	up := unsafe.Pointer(&(b[0]))
	pi := (*[1]byte)(up)
	buf := (*pi)[:]
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	sh.Len = n
	sh.Cap = n

	return buf
}
