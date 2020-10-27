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
		return types.NewOwner(addr, []string{}), false
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &owner)
	return owner, true
}

func (k Keeper) SetOwner(ctx sdk.Context, addr sdk.AccAddress, ids []string) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOwnerKey(addr)

	var owner types.Owner
	owner.Address = addr
	owner.IDs = ids

	value := k.cdc.MustMarshalBinaryLengthPrefixed(owner)
	store.Set(key, value)
}
