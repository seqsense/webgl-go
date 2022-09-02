//go:build tiny
// +build tiny

package webgl

import (
	"reflect"
	"unsafe"
)

func float32SliceAsByteSlice(floats []float32) []byte {
	l := 4 * len(floats)
	n := uintptr(unsafe.Pointer(&l))

	up := unsafe.Pointer(&(floats[0]))
	pi := (*[1]byte)(up)
	buf := (*pi)[:]
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	sh.Len = n
	sh.Cap = n

	return buf
}
