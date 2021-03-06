package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRecord{}

func NewMsgCreateRecord(creator string, subject string, role string, since string) *MsgCreateRecord {
	return &MsgCreateRecord{
		Creator: creator,
		Subject: subject,
		Role:    role,
		Since:   since,
	}
}

func (msg *MsgCreateRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateRecord) Type() string {
	return "CreateRecord"
}

func (msg *MsgCreateRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRecord{}

func NewMsgUpdateRecord(creator string, id uint64, subject string, role string, since string) *MsgUpdateRecord {
	return &MsgUpdateRecord{
		Id:      id,
		Creator: creator,
		Subject: subject,
		Role:    role,
		Since:   since,
	}
}

func (msg *MsgUpdateRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRecord) Type() string {
	return "UpdateRecord"
}

func (msg *MsgUpdateRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRecord{}

func NewMsgDeleteRecord(creator string, id uint64) *MsgDeleteRecord {
	return &MsgDeleteRecord{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRecord) Type() string {
	return "DeleteRecord"
}

func (msg *MsgDeleteRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
