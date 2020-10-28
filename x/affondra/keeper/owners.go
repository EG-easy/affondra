package keeper

import (
	"github.com/EG-easy/affondra/x/affondra/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetOwner(ctx sdk.Context, addr sdk.AccAddress) (owner types.Owner, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOwnerKey(addr)
	b := store.Get(key)
	if b == nil {
		return types.NewOwner(addr, []types.Item{}), false
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &owner)
	return owner, true
}

func (k Keeper) SetOwner(ctx sdk.Context, addr sdk.AccAddress, owner types.Owner) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOwnerKey(addr)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(owner)
	store.Set(key, value)
}
