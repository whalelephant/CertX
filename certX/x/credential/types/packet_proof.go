package types

// ValidateBasic is used for validating the packet
func (p ProofPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ProofPacketData) GetBytes() ([]byte, error) {
	var modulePacket CredentialPacketData

	modulePacket.Packet = &CredentialPacketData_ProofPacket{&p}

	return modulePacket.Marshal()
}
