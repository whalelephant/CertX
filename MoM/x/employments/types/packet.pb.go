// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: employments/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EmploymentsPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*EmploymentsPacketData_NoData
	//	*EmploymentsPacketData_ECredentialPacket
	Packet isEmploymentsPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *EmploymentsPacketData) Reset()         { *m = EmploymentsPacketData{} }
func (m *EmploymentsPacketData) String() string { return proto.CompactTextString(m) }
func (*EmploymentsPacketData) ProtoMessage()    {}
func (*EmploymentsPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11ab9ba307251bc, []int{0}
}
func (m *EmploymentsPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmploymentsPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmploymentsPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmploymentsPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmploymentsPacketData.Merge(m, src)
}
func (m *EmploymentsPacketData) XXX_Size() int {
	return m.Size()
}
func (m *EmploymentsPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_EmploymentsPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_EmploymentsPacketData proto.InternalMessageInfo

type isEmploymentsPacketData_Packet interface {
	isEmploymentsPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type EmploymentsPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type EmploymentsPacketData_ECredentialPacket struct {
	ECredentialPacket *ECredentialPacketData `protobuf:"bytes,2,opt,name=eCredentialPacket,proto3,oneof" json:"eCredentialPacket,omitempty"`
}

func (*EmploymentsPacketData_NoData) isEmploymentsPacketData_Packet()            {}
func (*EmploymentsPacketData_ECredentialPacket) isEmploymentsPacketData_Packet() {}

func (m *EmploymentsPacketData) GetPacket() isEmploymentsPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *EmploymentsPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*EmploymentsPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *EmploymentsPacketData) GetECredentialPacket() *ECredentialPacketData {
	if x, ok := m.GetPacket().(*EmploymentsPacketData_ECredentialPacket); ok {
		return x.ECredentialPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EmploymentsPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EmploymentsPacketData_NoData)(nil),
		(*EmploymentsPacketData_ECredentialPacket)(nil),
	}
}

type ECredentialsPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*ECredentialsPacketData_NoData
	//	*ECredentialsPacketData_ECredentialPacket
	Packet isECredentialsPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *ECredentialsPacketData) Reset()         { *m = ECredentialsPacketData{} }
func (m *ECredentialsPacketData) String() string { return proto.CompactTextString(m) }
func (*ECredentialsPacketData) ProtoMessage()    {}
func (*ECredentialsPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11ab9ba307251bc, []int{1}
}
func (m *ECredentialsPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ECredentialsPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ECredentialsPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ECredentialsPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ECredentialsPacketData.Merge(m, src)
}
func (m *ECredentialsPacketData) XXX_Size() int {
	return m.Size()
}
func (m *ECredentialsPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_ECredentialsPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_ECredentialsPacketData proto.InternalMessageInfo

type isECredentialsPacketData_Packet interface {
	isECredentialsPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ECredentialsPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type ECredentialsPacketData_ECredentialPacket struct {
	ECredentialPacket *ECredentialPacketData `protobuf:"bytes,2,opt,name=eCredentialPacket,proto3,oneof" json:"eCredentialPacket,omitempty"`
}

func (*ECredentialsPacketData_NoData) isECredentialsPacketData_Packet()            {}
func (*ECredentialsPacketData_ECredentialPacket) isECredentialsPacketData_Packet() {}

func (m *ECredentialsPacketData) GetPacket() isECredentialsPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *ECredentialsPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*ECredentialsPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *ECredentialsPacketData) GetECredentialPacket() *ECredentialPacketData {
	if x, ok := m.GetPacket().(*ECredentialsPacketData_ECredentialPacket); ok {
		return x.ECredentialPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ECredentialsPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ECredentialsPacketData_NoData)(nil),
		(*ECredentialsPacketData_ECredentialPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11ab9ba307251bc, []int{2}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// this line is used by starport scaffolding # ibc/packet/proto/message
// ECredentialPacketData defines a struct for the packet payload
type ECredentialPacketData struct {
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Claim   string `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *ECredentialPacketData) Reset()         { *m = ECredentialPacketData{} }
func (m *ECredentialPacketData) String() string { return proto.CompactTextString(m) }
func (*ECredentialPacketData) ProtoMessage()    {}
func (*ECredentialPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11ab9ba307251bc, []int{3}
}
func (m *ECredentialPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ECredentialPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ECredentialPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ECredentialPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ECredentialPacketData.Merge(m, src)
}
func (m *ECredentialPacketData) XXX_Size() int {
	return m.Size()
}
func (m *ECredentialPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_ECredentialPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_ECredentialPacketData proto.InternalMessageInfo

func (m *ECredentialPacketData) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ECredentialPacketData) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

// ECredentialPacketAck defines a struct for the packet acknowledgment
type ECredentialPacketAck struct {
}

func (m *ECredentialPacketAck) Reset()         { *m = ECredentialPacketAck{} }
func (m *ECredentialPacketAck) String() string { return proto.CompactTextString(m) }
func (*ECredentialPacketAck) ProtoMessage()    {}
func (*ECredentialPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a11ab9ba307251bc, []int{4}
}
func (m *ECredentialPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ECredentialPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ECredentialPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ECredentialPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ECredentialPacketAck.Merge(m, src)
}
func (m *ECredentialPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *ECredentialPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ECredentialPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_ECredentialPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmploymentsPacketData)(nil), "whalelephant.mom.employments.EmploymentsPacketData")
	proto.RegisterType((*ECredentialsPacketData)(nil), "whalelephant.mom.employments.ECredentialsPacketData")
	proto.RegisterType((*NoData)(nil), "whalelephant.mom.employments.NoData")
	proto.RegisterType((*ECredentialPacketData)(nil), "whalelephant.mom.employments.ECredentialPacketData")
	proto.RegisterType((*ECredentialPacketAck)(nil), "whalelephant.mom.employments.ECredentialPacketAck")
}

func init() { proto.RegisterFile("employments/packet.proto", fileDescriptor_a11ab9ba307251bc) }

var fileDescriptor_a11ab9ba307251bc = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcd, 0x2d, 0xc8,
	0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0xcf, 0x48, 0xcc, 0x49, 0xcd, 0x49, 0x2d, 0xc8, 0x48,
	0xcc, 0x2b, 0xd1, 0xcb, 0xcd, 0xcf, 0xd5, 0x43, 0x52, 0xaa, 0x74, 0x8c, 0x91, 0x4b, 0xd4, 0x15,
	0xc1, 0x0f, 0x00, 0xeb, 0x74, 0x49, 0x2c, 0x49, 0x14, 0xb2, 0xe3, 0x62, 0xcb, 0xcb, 0x07, 0xb1,
	0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x54, 0xf4, 0xf0, 0x19, 0xa4, 0xe7, 0x07, 0x56, 0xeb,
	0xc1, 0x10, 0x04, 0xd5, 0x25, 0x94, 0xcc, 0x25, 0x98, 0xea, 0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57,
	0x92, 0x99, 0x98, 0x03, 0x31, 0x58, 0x82, 0x09, 0x6c, 0x94, 0x31, 0x7e, 0xa3, 0x5c, 0xd1, 0xb5,
	0x41, 0x4d, 0xc6, 0x34, 0xcf, 0x89, 0x83, 0x8b, 0x0d, 0xe2, 0x59, 0xa5, 0xe3, 0x8c, 0x5c, 0x62,
	0x48, 0x1a, 0x87, 0xb0, 0x4f, 0x38, 0xb8, 0xd8, 0x20, 0x4e, 0x50, 0x72, 0xe7, 0x12, 0xc5, 0x6a,
	0x82, 0x90, 0x04, 0x17, 0x7b, 0x71, 0x69, 0x52, 0x56, 0x6a, 0x72, 0x09, 0xd8, 0x4b, 0x9c, 0x41,
	0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x72, 0x4e, 0x62, 0x66, 0x2e, 0xd8, 0x7d, 0x9c, 0x41, 0x10,
	0x8e, 0x92, 0x18, 0x97, 0x08, 0x86, 0x41, 0x8e, 0xc9, 0xd9, 0x4e, 0xfe, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x8f, 0xec, 0x45, 0xfd, 0xe4, 0xd4, 0xa2, 0x92, 0x08, 0x7d, 0xdf, 0x7c, 0x5f, 0xfd,
	0x0a, 0x7d, 0xe4, 0x34, 0x57, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0x73, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x86, 0x9a, 0x5c, 0x8f, 0x02, 0x00, 0x00,
}

func (m *EmploymentsPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmploymentsPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmploymentsPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *EmploymentsPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmploymentsPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *EmploymentsPacketData_ECredentialPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmploymentsPacketData_ECredentialPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ECredentialPacket != nil {
		{
			size, err := m.ECredentialPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ECredentialsPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ECredentialsPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ECredentialsPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ECredentialsPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ECredentialsPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ECredentialsPacketData_ECredentialPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ECredentialsPacketData_ECredentialPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ECredentialPacket != nil {
		{
			size, err := m.ECredentialPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ECredentialPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ECredentialPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ECredentialPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Claim) > 0 {
		i -= len(m.Claim)
		copy(dAtA[i:], m.Claim)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Claim)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ECredentialPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ECredentialPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ECredentialPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EmploymentsPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *EmploymentsPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *EmploymentsPacketData_ECredentialPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ECredentialPacket != nil {
		l = m.ECredentialPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *ECredentialsPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *ECredentialsPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *ECredentialsPacketData_ECredentialPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ECredentialPacket != nil {
		l = m.ECredentialPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ECredentialPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Claim)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *ECredentialPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EmploymentsPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmploymentsPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmploymentsPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &EmploymentsPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ECredentialPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ECredentialPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &EmploymentsPacketData_ECredentialPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ECredentialsPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ECredentialsPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ECredentialsPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &ECredentialsPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ECredentialPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ECredentialPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &ECredentialsPacketData_ECredentialPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ECredentialPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ECredentialPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ECredentialPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claim = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ECredentialPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ECredentialPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ECredentialPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
