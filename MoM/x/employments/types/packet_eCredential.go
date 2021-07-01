package types

// ValidateBasic is used for validating the packet
func (p ECredentialPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ECredentialPacketData) GetBytes() ([]byte, error) {
	var modulePacket EmploymentsPacketData

	modulePacket.Packet = &EmploymentsPacketData_ECredentialPacket{&p}

	return modulePacket.Marshal()
}
