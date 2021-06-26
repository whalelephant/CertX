package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendProof{}

func NewMsgSendProof(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	subject string,
	verifier string,
	issuer string,
	claim string,
	signature string,
) *MsgSendProof {
	return &MsgSendProof{
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

func (msg *MsgSendProof) Route() string {
	return RouterKey
}

func (msg *MsgSendProof) Type() string {
	return "SendProof"
}

func (msg *MsgSendProof) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
