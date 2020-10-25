package keeper

import (
	"github.com/EG-easy/affondra/x/affondra/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress) (owner types.Owner) {
	var itemCollections []types.ItemCollection
	k.IterateItemCollections(ctx, types.GetOwnerKey(address),
		func(_ sdk.AccAddress, itemCollection types.ItemCollection) (stop bool) {
			itemCollections = append(itemCollections, itemCollection)
			return false
		},
	)

	return types.NewOwner(address, itemCollections...)
}

//func (k Keeper) IterateOwners(ctx sdk.Context, handler func(owner types.Owner) (stop bool)) {
//	store := ctx.KVStore(k.storeKey)
//	iterator := sdk.KVStorePrefixIterator(store, types.OwnersKeyPrefix)
//	defer iterator.Close()
//	for ; iterator.Valid(); iterator.Next() {
//		var owner types.Owner
//
//		address, _ = types.SplitOwnerKey(iterator.Key())
//		owner = key.GetOwner(ctx, address)
//		if handler(owner) {
//			break
//		}
//	}
//}

func (k Keeper) IterateItemCollections(ctx sdk.Context, prefix []byte, handler func(owner sdk.AccAddress, itemCollection types.ItemCollection) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var itemCollection types.ItemCollection
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &itemCollection)

		owner, _ := types.SplitOwnerKey(iterator.Key())
		if handler(owner, itemCollection) {
			break
		}

	}
}
