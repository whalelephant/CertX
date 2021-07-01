package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/whalelephant/certX/MoM/x/employments/types"
)

func (k msgServer) SendECredential(goCtx context.Context, msg *types.MsgSendECredential) (*types.MsgSendECredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if msg.claim exits for sender

    // get record

    // create string from record 

    // sign recordStr

    // encrypt with ekey  (sign and record)
    

	// Construct the packet
	var packet types.ECredentialPacketData

	packet.Subject = msg.Subject
	packet.Claim = msg.Claim

	// Transmit the packet
	err := k.TransmitECredentialPacket(
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

	return &types.MsgSendECredentialResponse{}, nil
}
