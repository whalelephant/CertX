package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendECredential{}

func NewMsgSendECredential(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	subject string,
	claim string,
) *MsgSendECredential {
	return &MsgSendECredential{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Subject:          subject,
		Claim:            claim,
	}
}

func (msg *MsgSendECredential) Route() string {
	return RouterKey
}

func (msg *MsgSendECredential) Type() string {
	return "SendECredential"
}

func (msg *MsgSendECredential) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendECredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendECredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
