package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Item struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Denom string `json:"denom" yaml:"denom"`
    NftId string `json:"nftId" yaml:"nftId"`
    Price int32 `json:"price" yaml:"price"`
    Affiliate int32 `json:"affiliate" yaml:"affiliate"`
    InSale bool `json:"inSale" yaml:"inSale"`
}