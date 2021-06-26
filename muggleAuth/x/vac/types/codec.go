package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendVerifiableCredential{}, "vac/SendVerifiableCredential", nil)

	cdc.RegisterConcrete(&MsgCreateCredential{}, "vac/CreateCredential", nil)
	cdc.RegisterConcrete(&MsgUpdateCredential{}, "vac/UpdateCredential", nil)
	cdc.RegisterConcrete(&MsgDeleteCredential{}, "vac/DeleteCredential", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendVerifiableCredential{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCredential{},
		&MsgUpdateCredential{},
		&MsgDeleteCredential{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
