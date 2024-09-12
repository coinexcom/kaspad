package model

import "github.com/coinexcom/kaspad/domain/consensus/model/externalapi"

// VirtualBlockHash is a marker hash for the virtual block
var VirtualBlockHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
})

// VirtualGenesisBlockHash is a marker hash for the virtual genesis block
var VirtualGenesisBlockHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe,
	0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe,
	0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe,
	0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe, 0xfe,
})
