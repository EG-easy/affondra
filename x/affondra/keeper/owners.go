package keeper

import (
	"fmt"

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

	fmt.Printf("itemCollections from GetOwner: %v\n", itemCollections)
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

	fmt.Print("===start IterateItemCollections===\n")
	store := ctx.KVStore(k.storeKey)

	fmt.Printf("store: %v\n", store)
	fmt.Printf("store: %s\n", fmt.Sprintf("%s", store))

	iterator := sdk.KVStorePrefixIterator(store, prefix)

	fmt.Printf("prefix: %b\n", prefix)

	fmt.Printf("iterator: %v\n", iterator)
	fmt.Printf("iterator valid: %v\n", iterator.Valid())
	fmt.Printf("iterator value: %v\n", iterator.Value())
	fmt.Printf("iterator error: %v\n", iterator.Error())
	fmt.Printf("iterator key: %v\n", iterator.Key())

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var itemCollection types.ItemCollection
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &itemCollection)

		fmt.Printf("iterator key: %b\n", iterator.Key())

		owner, _ := types.SplitOwnerKey(iterator.Key())

		fmt.Printf("owner: %b\n", owner)

		if handler(owner, itemCollection) {
			break
		}
	}
}
