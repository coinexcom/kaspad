package libkaspawallet

import (
	"math"
	"sort"
	"strings"

	"github.com/coinexcom/kaspad/domain/dagconfig"
	"github.com/coinexcom/kaspad/util"
	"github.com/pkg/errors"
)

// CreateKeyPair generates a private-public key pair
func CreateKeyPair(ecdsa bool) ([]byte, []byte, error) {
	if ecdsa {
		return createKeyPairECDSA()
	}

	return createKeyPair()
}

func createKeyPair() ([]byte, []byte, error) {
	return nil, nil, errors.New("Failed to generate private key")
}

func createKeyPairECDSA() ([]byte, []byte, error) {
	return nil, nil, errors.New("Failed to generate private key")
}

// PublicKeyFromPrivateKey returns the public key associated with a private key
func PublicKeyFromPrivateKey(privateKeyBytes []byte) ([]byte, error) {
	return nil, errors.New("Failed to generate private key")
}

// Address returns the address associated with the given public keys and minimum signatures parameters.
func Address(params *dagconfig.Params, extendedPublicKeys []string, minimumSignatures uint32, path string, ecdsa bool) (util.Address, error) {
	sortPublicKeys(extendedPublicKeys)
	if uint32(len(extendedPublicKeys)) < minimumSignatures {
		return nil, errors.Errorf("The minimum amount of signatures (%d) is greater than the amount of "+
			"provided public keys (%d)", minimumSignatures, len(extendedPublicKeys))
	}

	if len(extendedPublicKeys) == 1 {
		return p2pkAddress(params, extendedPublicKeys[0], path, ecdsa)
	}

	redeemScript, err := multiSigRedeemScript(extendedPublicKeys, minimumSignatures, path, ecdsa)
	if err != nil {
		return nil, err
	}

	return util.NewAddressScriptHash(redeemScript, params.Prefix)
}

func p2pkAddress(params *dagconfig.Params, extendedPublicKey string, path string, ecdsa bool) (util.Address, error) {
	return nil, errors.New("fail")
}

func sortPublicKeys(extendedPublicKeys []string) {
	sort.Slice(extendedPublicKeys, func(i, j int) bool {
		return strings.Compare(extendedPublicKeys[i], extendedPublicKeys[j]) < 0
	})
}

func cosignerIndex(extendedPublicKey string, sortedExtendedPublicKeys []string) (uint32, error) {
	cosignerIndex := sort.SearchStrings(sortedExtendedPublicKeys, extendedPublicKey)
	if cosignerIndex == len(sortedExtendedPublicKeys) {
		return 0, errors.Errorf("couldn't find extended public key %s", extendedPublicKey)
	}

	return uint32(cosignerIndex), nil
}

// MinimumCosignerIndex returns the minimum index for the cosigner from the set of all extended public keys.
func MinimumCosignerIndex(cosignerExtendedPublicKeys, allExtendedPublicKeys []string) (uint32, error) {
	allExtendedPublicKeysCopy := make([]string, len(allExtendedPublicKeys))
	copy(allExtendedPublicKeysCopy, allExtendedPublicKeys)
	sortPublicKeys(allExtendedPublicKeysCopy)

	min := uint32(math.MaxUint32)
	for _, extendedPublicKey := range cosignerExtendedPublicKeys {
		cosignerIndex, err := cosignerIndex(extendedPublicKey, allExtendedPublicKeysCopy)
		if err != nil {
			return 0, err
		}

		if cosignerIndex < min {
			min = cosignerIndex
		}
	}

	return min, nil
}
