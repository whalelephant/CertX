package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendVerifiableCredential{}

func NewMsgSendVerifiableCredential(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	subject string,
	verifier string,
	issuer string,
	claim string,
	signature string,
) *MsgSendVerifiableCredential {
	return &MsgSendVerifiableCredential{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Subject:          subject,
		Verifier:         verifier,
		Issuer:           issuer,
		Claim:            claim,
		Signature:        signature,
	}
}

func (msg *MsgSendVerifiableCredential) Route() string {
	return RouterKey
}

func (msg *MsgSendVerifiableCredential) Type() string {
	return "SendVerifiableCredential"
}

func (msg *MsgSendVerifiableCredential) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendVerifiableCredential) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendVerifiableCredential) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
