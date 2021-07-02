package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendECredential{}, "employments/SendECredential", nil)

	cdc.RegisterConcrete(&MsgCreateRecord{}, "employments/CreateRecord", nil)
	cdc.RegisterConcrete(&MsgUpdateRecord{}, "employments/UpdateRecord", nil)
	cdc.RegisterConcrete(&MsgDeleteRecord{}, "employments/DeleteRecord", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendECredential{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRecord{},
		&MsgUpdateRecord{},
		&MsgDeleteRecord{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
   	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
