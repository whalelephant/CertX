package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/whalelephant/certX/x/credentials/types"
)

func (k msgServer) SendVerifiableCredential(goCtx context.Context, msg *types.MsgSendVerifiableCredential) (*types.MsgSendVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.VerifiableCredentialPacketData

	packet.Subject = msg.Subject
	packet.Verifier = msg.Verifier
	packet.Issuer = msg.Issuer
	packet.Claim = msg.Claim
	packet.Signature = msg.Signature

	// Transmit the packet
	err := k.TransmitVerifiableCredentialPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendVerifiableCredentialResponse{}, nil
}
