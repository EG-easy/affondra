package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidAffiliatePrice = sdkerrors.Register(ModuleName, 1, "[affondra] Affiliate should be less than Price")
	ErrAlreadyOnSale         = sdkerrors.Register(ModuleName, 2, "[affondra] The NFT is already on Sale")
	ErrUnknownID             = sdkerrors.Register(ModuleName, 3, "[affondra] Unknown Item collection")
)
