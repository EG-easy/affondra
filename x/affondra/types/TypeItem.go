package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Item struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
	NftId     string         `json:"nftId" yaml:"nftId"`
	Price     sdk.Coin       `json:"price" yaml:"price"`
	Affiliate sdk.Coin       `json:"affiliate" yaml:"affiliate"`
	InSale    bool           `json:"inSale" yaml:"inSale"`
}
