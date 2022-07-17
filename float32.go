package xboct

import (
	"unsafe"
)

const (
	bias             = 127
	fullMantissaMask = 1 << 23
	oneBitMask       = 0b1
)

func FractionalPartLength(number float32) uint16 {
	// Returns the IEEE 754 binary representation.
	bits := *(*uint32)(unsafe.Pointer(&number))
	if bits == 0 {
		return 0
	}

	exponent := int(bits<<1>>24) - bias
	if exponent > 23 { // mantissa length
		return 0
	}

	cleanedRowMantissa := bits << 9 >> 9
	var restoredMantissa uint32
	if exponent < 0 {
		restoredMantissa = (cleanedRowMantissa | fullMantissaMask) >> -exponent
	} else {
		restoredMantissa = cleanedRowMantissa << uint32(exponent)
	}

	var i, length uint16 = 1, 0
	var bit uint32
	for ; i < 23; i++ {
		bit = (restoredMantissa >> (23 - i)) & oneBitMask
		if bit == 1 {
			length = i
		}
	}

	return length
}
