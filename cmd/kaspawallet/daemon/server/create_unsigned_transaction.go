package server

import (
	"context"
	"github.com/coinexcom/kaspad/cmd/kaspawallet/daemon/pb"
	"github.com/coinexcom/kaspad/cmd/kaspawallet/libkaspawallet"
	"github.com/coinexcom/kaspad/domain/consensus/utils/constants"
	"github.com/pkg/errors"
)

// TODO: Implement a better fee estimation mechanism
const feePerInput = 10000

// The minimal change amount to target in order to avoid large storage mass (see KIP9 for more details).
// By having at least 0.2KAS in the change output we make sure that every transaction with send value >= 0.2KAS
// should succeed (at most 50K storage mass for each output, thus overall lower than standard mass upper bound which is 100K gram)
const minChangeTarget = constants.SompiPerKaspa / 5

func (s *server) CreateUnsignedTransactions(_ context.Context, request *pb.CreateUnsignedTransactionsRequest) (
	*pb.CreateUnsignedTransactionsResponse, error,
) {

	return &pb.CreateUnsignedTransactionsResponse{}, nil
}

func (s *server) createUnsignedTransactions(address string, amount uint64, isSendAll bool, fromAddressesString []string, useExistingChangeAddress bool) ([][]byte, error) {
	return nil, nil
}

func (s *server) selectUTXOs(spendAmount uint64, isSendAll bool, feePerInput uint64, fromAddresses []*walletAddress) (
	selectedUTXOs []*libkaspawallet.UTXO, totalReceived uint64, changeSompi uint64, err error) {

	selectedUTXOs = []*libkaspawallet.UTXO{}
	totalValue := uint64(0)

	dagInfo, err := s.rpcClient.GetBlockDAGInfo()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, utxo := range s.utxosSortedByAmount {
		if (fromAddresses != nil && !walletAddressesContain(fromAddresses, utxo.address)) ||
			!s.isUTXOSpendable(utxo, dagInfo.VirtualDAAScore) {
			continue
		}

		if broadcastTime, ok := s.usedOutpoints[*utxo.Outpoint]; ok {
			if s.usedOutpointHasExpired(broadcastTime) {
				delete(s.usedOutpoints, *utxo.Outpoint)
			} else {
				continue
			}
		}

		selectedUTXOs = append(selectedUTXOs, &libkaspawallet.UTXO{
			Outpoint:       utxo.Outpoint,
			UTXOEntry:      utxo.UTXOEntry,
			DerivationPath: s.walletAddressPath(utxo.address),
		})

		totalValue += utxo.UTXOEntry.Amount()

		fee := feePerInput * uint64(len(selectedUTXOs))
		totalSpend := spendAmount + fee
		// Two break cases (if not send all):
		// 		1. totalValue == totalSpend, so there's no change needed -> number of outputs = 1, so a single input is sufficient
		// 		2. totalValue > totalSpend, so there will be change and 2 outputs, therefor in order to not struggle with --
		//		   2.1 go-nodes dust patch we try and find at least 2 inputs (even though the next one is not necessary in terms of spend value)
		// 		   2.2 KIP9 we try and make sure that the change amount is not too small
		if !isSendAll && (totalValue == totalSpend || (totalValue >= totalSpend+minChangeTarget && len(selectedUTXOs) > 1)) {
			break
		}
	}

	fee := feePerInput * uint64(len(selectedUTXOs))
	var totalSpend uint64
	if isSendAll {
		totalSpend = totalValue
		totalReceived = totalValue - fee
	} else {
		totalSpend = spendAmount + fee
		totalReceived = spendAmount
	}
	if totalValue < totalSpend {
		return nil, 0, 0, errors.Errorf("Insufficient funds for send: %f required, while only %f available",
			float64(totalSpend)/constants.SompiPerKaspa, float64(totalValue)/constants.SompiPerKaspa)
	}

	return selectedUTXOs, totalReceived, totalValue - totalSpend, nil
}

func walletAddressesContain(addresses []*walletAddress, contain *walletAddress) bool {
	for _, address := range addresses {
		if *address == *contain {
			return true
		}
	}

	return false
}
