package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateItem{}

// MsgCreateItem defines create item message
type MsgCreateItem struct {
	ID          string
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	Denom       string         `json:"denom" yaml:"denom"`
	NftId       string         `json:"nftId" yaml:"nftId"`
	Price       sdk.Coin       `json:"price" yaml:"price"`
	Affiliate   sdk.Coin       `json:"affiliate" yaml:"affiliate"`
	Description string         `json:"description" yaml:"description"`
	InSale      bool           `json:"inSale" yaml:"inSale"`
}

// NewMsgCreateItem is a constructor function for MsgCreateItem
func NewMsgCreateItem(creator sdk.AccAddress, denom string, nftId string, price sdk.Coin, affiliate sdk.Coin, description string, inSale bool) MsgCreateItem {
	return MsgCreateItem{
		ID:          uuid.New().String(),
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
func (msg MsgCreateItem) Route() string {
	return RouterKey
}

// Type Implements Msg
func (msg MsgCreateItem) Type() string {
	return "CreateItem"
}

// GetSigners Implements Msg
func (msg MsgCreateItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes Implements Msg
func (msg MsgCreateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Implements Msg
func (msg MsgCreateItem) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
