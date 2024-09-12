package libkaspawallet

import (
	"github.com/coinexcom/kaspad/domain/dagconfig"
	"github.com/pkg/errors"
)

// Sign signs the transaction with the given private keys
func Sign(params *dagconfig.Params, mnemonics []string, serializedPSTx []byte, ecdsa bool) ([]byte, error) {
	return nil, errors.New("fail")
}
