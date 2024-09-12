package bip32

import (
	"github.com/pkg/errors"
)

const hardenedIndexStart = 0x80000000

// NewMaster returns a new extended private key based on the given seed and version
func NewMaster(seed []byte, version [4]byte) (*ExtendedKey, error) {
	return nil, errors.New("fail secp256k1")
}

func isHardened(i uint32) bool {
	return i >= hardenedIndexStart
}

// Child return the i'th derived child of extKey.
func (extKey *ExtendedKey) Child(i uint32) (*ExtendedKey, error) {
	return nil, errors.New("fail")
}
