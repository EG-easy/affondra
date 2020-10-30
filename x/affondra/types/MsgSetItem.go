package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetItem{}

// MsgSetItem defines set item message
type MsgSetItem struct {
	ID          string         `json:"id" yaml:"id"`
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	Denom       string         `json:"denom" yaml:"denom"`
	NftId       string         `json:"nftId" yaml:"nftId"`
	Price       sdk.Coin       `json:"price" yaml:"price"`
	Affiliate   sdk.Coin       `json:"affiliate" yaml:"affiliate"`
	Description string         `json:"description" yaml:"description"`
	InSale      bool           `json:"inSale" yaml:"inSale"`
}

// NewMsgSetItem is a constructor function for MsgSetItem
func NewMsgSetItem(creator sdk.AccAddress, id string, denom string, nftId string, price sdk.Coin, affiliate sdk.Coin, description string, inSale bool) MsgSetItem {
	return MsgSetItem{
		ID:          id,
		Creator:     creator,
		Denom:       denom,
		NftId:       nftId,
		Price:       price,
		Affiliate:   affiliate,
		Description: description,
		InSale:      inSale,
	}
}

// Route Implements Msg
func (msg MsgSetItem) Route() string {
	return RouterKey
}

// Type Implements Msg
func (msg MsgSetItem) Type() string {
	return "SetItem"
}

// GetSigners Implements Msg
func (msg MsgSetItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes Implements Msg
func (msg MsgSetItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Implements Msg
func (msg MsgSetItem) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
