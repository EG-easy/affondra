package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidAffiliatePrice = sdkerrors.Register(ModuleName, 1, "[affondra] Affiliate should be less than Price")
	ErrAlreadyOnSale         = sdkerrors.Register(ModuleName, 2, "[affondra] The NFT is already on Sale")
	ErrUnknownID             = sdkerrors.Register(ModuleName, 3, "[affondra] Unknown Item collection")
	ErrUnknownItem           = sdkerrors.Register(ModuleName, 4, "[affondra] Unknown Item collection")
	ErrItemAlreadyExists     = sdkerrors.Register(ModuleName, 5, "[affondra] Item Already exists")
	ErrUnknownCollection     = sdkerrors.Register(ModuleName, 6, "[affondra] Unknown Item collection")
)
