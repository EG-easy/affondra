package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteItem{}

// MsgDeleteItem defines delete item message
type MsgDeleteItem struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgDeleteItem is a constructor function for MsgDeleteItem
func NewMsgDeleteItem(id string, creator sdk.AccAddress) MsgDeleteItem {
	return MsgDeleteItem{
		ID:      id,
		Creator: creator,
	}
}

// Route Implements Msg
func (msg MsgDeleteItem) Route() string {
	return RouterKey
}

// Type Implements Msg
func (msg MsgDeleteItem) Type() string {
	return "DeleteItem"
}

// GetSigners Implements Msg
func (msg MsgDeleteItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes Implements Msg
func (msg MsgDeleteItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Implements Msg
func (msg MsgDeleteItem) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
