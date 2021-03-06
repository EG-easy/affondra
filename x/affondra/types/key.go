package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
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

// ItemPrefix
const (
	ItemPrefix = "item-"
)

// KeyPrefix for query
var (
	CollectionsKeyPrefix = []byte{0x00} // key for Item Collection
	OwnersKeyPrefix      = []byte{0x01}
)

// GetCollectionKey gets the key of a collection
func GetCollectionKey(denom string) []byte {
	h := tmhash.New()
	_, err := h.Write([]byte(denom))
	if err != nil {
		panic(err)
	}
	bs := h.Sum(nil)

	return append(CollectionsKeyPrefix, bs...)
}

// SplitOwnerKey returns address as a key
func SplitOwnerKey(key []byte) (sdk.AccAddress, []byte) {
	if len(key) != 53 {
		panic(fmt.Sprintf("unexpected key length %d", len(key)))
	}

	address := key[1 : sdk.AddrLen+1]
	IDsBz := key[sdk.AddrLen+1:]

	return sdk.AccAddress(address), IDsBz
}

// GetOwnerKey returns owner address as a key
func GetOwnerKey(address sdk.AccAddress) []byte {
	fmt.Printf("OwnersKeyPrefix: %b\n", OwnersKeyPrefix)
	return append(OwnersKeyPrefix, address.Bytes()...)
}
