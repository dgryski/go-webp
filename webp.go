// Package webp is a CGO wrapper around libwebp
package webp

/*
#cgo LDFLAGS: -lwebp -L/opt/local/lib/

#include <stdlib.h>
#include <webp/encode.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

func WebPEncodeRGB(rgb []byte, width, height, stride int, quality float64) []byte {
	var coutput *C.uint8_t
	outptr := (**C.uint8_t)(unsafe.Pointer(&coutput))
	length := C.WebPEncodeRGB((*C.uint8_t)(unsafe.Pointer(&rgb[0])), C.int(width), C.int(height), C.int(stride), C.float(quality), outptr)

	var output []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&output)))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(coutput))
	return output
}

func WebPEncodeRGBA(rgb []byte, width, height, stride int, quality float64) []byte {
	var coutput *C.uint8_t
	outptr := (**C.uint8_t)(unsafe.Pointer(&coutput))
	length := C.WebPEncodeRGBA((*C.uint8_t)(unsafe.Pointer(&rgb[0])), C.int(width), C.int(height), C.int(stride), C.float(quality), outptr)

	var output []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&output)))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(coutput))
	return output
}

func WebPEncodeLosslessRGB(rgb []byte, width, height, stride int) []byte {
	var coutput *C.uint8_t
	outptr := (**C.uint8_t)(unsafe.Pointer(&coutput))
	length := C.WebPEncodeLosslessRGB((*C.uint8_t)(unsafe.Pointer(&rgb[0])), C.int(width), C.int(height), C.int(stride), outptr)

	var output []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&output)))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(coutput))
	return output
}

func WebPEncodeLosslessRGBA(rgb []byte, width, height, stride int) []byte {
	var coutput *C.uint8_t
	outptr := (**C.uint8_t)(unsafe.Pointer(&coutput))
	length := C.WebPEncodeLosslessRGBA((*C.uint8_t)(unsafe.Pointer(&rgb[0])), C.int(width), C.int(height), C.int(stride), outptr)

	var output []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&output)))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(coutput))
	return output
}

func FreeWebP(img []byte) {
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&img)))
	C.free(unsafe.Pointer(sliceHeader.Data))
}
