package types

// ValidateBasic is used for validating the packet
func (p VerifiableCredentialPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p VerifiableCredentialPacketData) GetBytes() ([]byte, error) {
	var modulePacket CredentialsPacketData

	modulePacket.Packet = &CredentialsPacketData_VerifiableCredentialPacket{&p}

	return modulePacket.Marshal()
}
