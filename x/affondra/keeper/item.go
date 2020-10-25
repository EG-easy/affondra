package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// CreateItem creates a item
func (k Keeper) CreateItem(ctx sdk.Context, item types.Item) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ItemPrefix + item.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(item)
	store.Set(key, value)
}

// GetItem returns the item information
func (k Keeper) GetItem(ctx sdk.Context, key string) (types.Item, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.Item
	byteKey := []byte(types.ItemPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// SetItem sets a item
func (k Keeper) SetItem(ctx sdk.Context, item types.Item) {
	itemKey := item.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(item)
	key := []byte(types.ItemPrefix + itemKey)
	store.Set(key, bz)
}

// DeleteItem deletes a item
func (k Keeper) DeleteItem(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ItemPrefix + key))
}

//
// Functions used by querier
//

func listItem(ctx sdk.Context, k Keeper) ([]byte, error) {
	var itemList []types.Item
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ItemPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var item types.Item
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &item)
		itemList = append(itemList, item)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, itemList)
	return res, nil
}

func getItem(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	fmt.Printf("key from getItem: %s", key)
	item, err := k.GetItem(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, item)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetItemOwner(ctx sdk.Context, key string) sdk.AccAddress {
	item, err := k.GetItem(ctx, key)
	if err != nil {
		return nil
	}
	return item.Creator
}

// Check if the key exists in the store
func (k Keeper) ItemExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ItemPrefix + key))
}
