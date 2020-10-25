package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "affondra"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	PollPrefix = "poll-"
)

const (
	VotePrefix = "vote-"
)

const (
	ItemPrefix = "item-"
)

var (
	OwnersKeyPrefix = []byte{0x00}
)

func SplitOwnerKey(key []byte) (sdk.AccAddress, []byte) {
	if len(key) != 53 {
		panic(fmt.Sprintf("unexpected key length %d", len(key)))
	}

	address := key[1 : sdk.AddrLen+1]
	denomHashBz := key[sdk.AddrLen+1:]

	return sdk.AccAddress(address), denomHashBz
}

func GetOwnerKey(address sdk.AccAddress) []byte {
	return append(OwnersKeyPrefix, address.Bytes()...)
}
