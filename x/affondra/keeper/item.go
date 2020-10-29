package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// CreateItem creates a item
func (k Keeper) CreateItem(ctx sdk.Context, item types.Item) (err error) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ItemPrefix + item.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(item)
	store.Set(key, value)

	// register collection info
	collection, found := k.GetCollection(ctx, item.GetDenom())
	if found {
		collection, err = collection.AddItem(item)
		if err != nil {
			return err
		}
	} else {
		collection = types.NewCollection(item.GetDenom(), types.NewItems(item))
	}
	k.SetCollection(ctx, item.GetDenom(), collection)

	//register owner info
	owner, found := k.GetOwner(ctx, item.GetOwner())
	if found {
		owner, err = owner.AddItem(item)
		if err != nil {
			return err
		}
	} else {
		owner = types.NewOwner(item.GetOwner(), types.NewItems(item))
	}
	k.SetOwner(ctx, item.GetOwner(), owner)

	return nil
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

// BuyItem create transactions
func (k Keeper) BuyItem(ctx sdk.Context, key string, receiver sdk.AccAddress) (err error) {
	item, err := k.GetItem(ctx, key)
	if err != nil {
		return err
	}
	item.SetReceiver(receiver)
	k.SetItem(ctx, item)
	return nil
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

// GetItemOwner returns creator of the item
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

// IterateItems iterates over collections and performs a function
func (k Keeper) IterateItems(ctx sdk.Context, handler func(item types.Item) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ItemPrefix))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var item types.Item
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &item)
		if handler(item) {
			break
		}
	}
}
