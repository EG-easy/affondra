package keeper

import (
	"github.com/EG-easy/affondra/x/affondra/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCollection sets the entire collection of a single denom
func (k Keeper) SetCollection(ctx sdk.Context, denom string, collection types.Collection) {
	store := ctx.KVStore(k.storeKey)
	collectionKey := types.GetCollectionKey(denom)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(collection)
	store.Set(collectionKey, bz)
}

// GetCollection returns a collection of Items
func (k Keeper) GetCollection(ctx sdk.Context, denom string) (collection types.Collection, found bool) {
	store := ctx.KVStore(k.storeKey)
	collectionKey := types.GetCollectionKey(denom)
	bz := store.Get(collectionKey)
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &collection)
	return collection, true
}
