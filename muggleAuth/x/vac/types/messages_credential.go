package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCredential{}

func NewMsgCreateCredential(creator string, issuer string, holder string, claim int32) *MsgCreateCredential {
	return &MsgCreateCredential{
		Creator: creator,
		Issuer:  issuer,
		Holder:  holder,
		Claim:   claim,
	}
}

func (msg *MsgCreateCredential) Route() string {
	return RouterKey
}

func (msg *MsgCreateCredential) Type() string {
	return "CreateCredential"
}

func (msg *MsgCreateCredential) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCredential{}

func NewMsgUpdateCredential(creator string, id uint64, issuer string, holder string, claim int32) *MsgUpdateCredential {
	return &MsgUpdateCredential{
		Id:      id,
		Creator: creator,
		Issuer:  issuer,
		Holder:  holder,
		Claim:   claim,
	}
}

func (msg *MsgUpdateCredential) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCredential) Type() string {
	return "UpdateCredential"
}

func (msg *MsgUpdateCredential) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateCredential{}

func NewMsgDeleteCredential(creator string, id uint64) *MsgDeleteCredential {
	return &MsgDeleteCredential{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteCredential) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCredential) Type() string {
	return "DeleteCredential"
}

func (msg *MsgDeleteCredential) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
