package bip32

import (
	"github.com/pkg/errors"
)

const (
	versionSerializationLen     = 4
	depthSerializationLen       = 1
	fingerprintSerializationLen = 4
	childNumberSerializationLen = 4
	chainCodeSerializationLen   = 32
	keySerializationLen         = 33
	checkSumLen                 = 4
)

const extendedKeySerializationLen = versionSerializationLen +
	depthSerializationLen +
	fingerprintSerializationLen +
	childNumberSerializationLen +
	chainCodeSerializationLen +
	keySerializationLen +
	checkSumLen

// DeserializeExtendedKey deserialized the given base58 string and returns an extended key
func DeserializeExtendedKey(extKeyString string) (*ExtendedKey, error) {
	return nil, errors.New("fail")
}
