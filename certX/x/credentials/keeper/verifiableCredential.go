package keeper

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/whalelephant/certX/certX/x/credentials/types"
)

// TransmitVerifiableCredentialPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitVerifiableCredentialPacket(
	ctx sdk.Context,
	packetData types.VerifiableCredentialPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvVerifiableCredentialPacket processes packet reception
func (k Keeper) OnRecvVerifiableCredentialPacket(ctx sdk.Context, packet channeltypes.Packet, data types.VerifiableCredentialPacketData) (packetAck types.VerifiableCredentialPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	fmt.Println("Got VerifiableCredentialPacketData")

	fmt.Println("checking sign")
	decodedSig, err := base64.StdEncoding.DecodeString(data.GetSignature())
	if err != nil {
		fmt.Println("sig not decodable: ", err)
		return packetAck, err
	}

    // 33 secp256k1 pubkey len; 64 sign len
	pubkey := decodedSig[0:33]
	sig := decodedSig[33:]
	msg := data.GetSubject() + data.GetVerifier() + data.GetIssuer() + data.GetClaim()

	var key secp256k1.PubKey
	key.Key = pubkey
	addr := key.Address().String()

	// Issuer did:method:identifier (identifier is address in Bech32 format)
	didComponents := strings.Split(data.GetIssuer(), ":")
	fmt.Println("didC[2]: ", didComponents[2])

	if !key.VerifySignature([]byte(msg), sig) {
		fmt.Println("sig not verified: ")
		return packetAck, err
	} else if addr != didComponents[2] {
		fmt.Println("signer not issuer")
		return packetAck, err
	}

	var recvProof types.Credential
    // Demo purpose only, Creator should be the Issuer
	recvProof.Creator = packet.GetSourcePort() + packet.GetDestChannel()
	recvProof.Issuer = data.GetIssuer()
	recvProof.Subject = data.GetSubject()
	recvProof.Verifier = data.GetVerifier()
	recvProof.Claim = data.GetClaim()
	recvProof.Signature = data.GetSignature()
	// can return ID as packetAck
	_ = k.AppendCredential(ctx, recvProof)
	return packetAck, nil
}

// OnAcknowledgementVerifiableCredentialPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementVerifiableCredentialPacket(ctx sdk.Context, packet channeltypes.Packet, data types.VerifiableCredentialPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.VerifiableCredentialPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutVerifiableCredentialPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutVerifiableCredentialPacket(ctx sdk.Context, packet channeltypes.Packet, data types.VerifiableCredentialPacketData) error {

	// TODO: packet timeout logic

	return nil
}
