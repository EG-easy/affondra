package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/voter/x/voter/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreatePoll creates a poll
func (k Keeper) CreatePoll(ctx sdk.Context, poll types.Poll) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PollPrefix + poll.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(poll)
	store.Set(key, value)
}

// GetPoll returns the poll information
func (k Keeper) GetPoll(ctx sdk.Context, key string) (types.Poll, error) {
	store := ctx.KVStore(k.storeKey)
	var poll types.Poll
	byteKey := []byte(types.PollPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &poll)
	if err != nil {
		return poll, err
	}
	return poll, nil
}

// SetPoll sets a poll
func (k Keeper) SetPoll(ctx sdk.Context, poll types.Poll) {
	pollKey := poll.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(poll)
	key := []byte(types.PollPrefix + pollKey)
	store.Set(key, bz)
}

// DeletePoll deletes a poll
func (k Keeper) DeletePoll(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.PollPrefix + key))
}

//
// Functions used by querier
//

func listPoll(ctx sdk.Context, k Keeper) ([]byte, error) {
	var pollList []types.Poll
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PollPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var poll types.Poll
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &poll)
		pollList = append(pollList, poll)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, pollList)
	return res, nil
}

func getPoll(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	poll, err := k.GetPoll(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, poll)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetPollOwner(ctx sdk.Context, key string) sdk.AccAddress {
	poll, err := k.GetPoll(ctx, key)
	if err != nil {
		return nil
	}
	return poll.Creator
}


// Check if the key exists in the store
func (k Keeper) PollExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PollPrefix + key))
}
