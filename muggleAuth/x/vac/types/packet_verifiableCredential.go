package types

import (
   	sdk "github.com/cosmos/cosmos-sdk/types"

)

// ValidateBasic is used for validating the packet
func (p VerifiableCredentialPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p VerifiableCredentialPacketData) GetBytes() ([]byte) {
    return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}
