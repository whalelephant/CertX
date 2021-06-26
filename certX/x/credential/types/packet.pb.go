// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credential/packet.proto

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

type CredentialPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*CredentialPacketData_NoData
	//	*CredentialPacketData_ProofPacket
	Packet isCredentialPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *CredentialPacketData) Reset()         { *m = CredentialPacketData{} }
func (m *CredentialPacketData) String() string { return proto.CompactTextString(m) }
func (*CredentialPacketData) ProtoMessage()    {}
func (*CredentialPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfeb20e43b12fb, []int{0}
}
func (m *CredentialPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialPacketData.Merge(m, src)
}
func (m *CredentialPacketData) XXX_Size() int {
	return m.Size()
}
func (m *CredentialPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialPacketData proto.InternalMessageInfo

type isCredentialPacketData_Packet interface {
	isCredentialPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CredentialPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type CredentialPacketData_ProofPacket struct {
	ProofPacket *ProofPacketData `protobuf:"bytes,2,opt,name=proofPacket,proto3,oneof" json:"proofPacket,omitempty"`
}

func (*CredentialPacketData_NoData) isCredentialPacketData_Packet()      {}
func (*CredentialPacketData_ProofPacket) isCredentialPacketData_Packet() {}

func (m *CredentialPacketData) GetPacket() isCredentialPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *CredentialPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*CredentialPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *CredentialPacketData) GetProofPacket() *ProofPacketData {
	if x, ok := m.GetPacket().(*CredentialPacketData_ProofPacket); ok {
		return x.ProofPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CredentialPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CredentialPacketData_NoData)(nil),
		(*CredentialPacketData_ProofPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfeb20e43b12fb, []int{1}
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
// ProofPacketData defines a struct for the packet payload
type ProofPacketData struct {
	Subject   string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Verifier  string `protobuf:"bytes,2,opt,name=verifier,proto3" json:"verifier,omitempty"`
	Issuer    string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Claim     string `protobuf:"bytes,4,opt,name=claim,proto3" json:"claim,omitempty"`
	Signature string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ProofPacketData) Reset()         { *m = ProofPacketData{} }
func (m *ProofPacketData) String() string { return proto.CompactTextString(m) }
func (*ProofPacketData) ProtoMessage()    {}
func (*ProofPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfeb20e43b12fb, []int{2}
}
func (m *ProofPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofPacketData.Merge(m, src)
}
func (m *ProofPacketData) XXX_Size() int {
	return m.Size()
}
func (m *ProofPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_ProofPacketData proto.InternalMessageInfo

func (m *ProofPacketData) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ProofPacketData) GetVerifier() string {
	if m != nil {
		return m.Verifier
	}
	return ""
}

func (m *ProofPacketData) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *ProofPacketData) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *ProofPacketData) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

// ProofPacketAck defines a struct for the packet acknowledgment
type ProofPacketAck struct {
}

func (m *ProofPacketAck) Reset()         { *m = ProofPacketAck{} }
func (m *ProofPacketAck) String() string { return proto.CompactTextString(m) }
func (*ProofPacketAck) ProtoMessage()    {}
func (*ProofPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfeb20e43b12fb, []int{3}
}
func (m *ProofPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofPacketAck.Merge(m, src)
}
func (m *ProofPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *ProofPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_ProofPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CredentialPacketData)(nil), "whalelephant.certx.credential.CredentialPacketData")
	proto.RegisterType((*NoData)(nil), "whalelephant.certx.credential.NoData")
	proto.RegisterType((*ProofPacketData)(nil), "whalelephant.certx.credential.ProofPacketData")
	proto.RegisterType((*ProofPacketAck)(nil), "whalelephant.certx.credential.ProofPacketAck")
}

func init() { proto.RegisterFile("credential/packet.proto", fileDescriptor_67bfeb20e43b12fb) }

var fileDescriptor_67bfeb20e43b12fb = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x6a, 0x63, 0x3b, 0x05, 0x95, 0xa5, 0x68, 0x10, 0x0d, 0x12, 0x10, 0x3c, 0x6d,
	0x40, 0xc1, 0xab, 0x58, 0x3d, 0x78, 0x92, 0x92, 0x93, 0x78, 0xdb, 0xae, 0xd3, 0x76, 0x6d, 0x9a,
	0x0d, 0x9b, 0x8d, 0xd6, 0xb7, 0x10, 0x7c, 0x16, 0xdf, 0xc1, 0x63, 0x8f, 0x1e, 0xa5, 0x7d, 0x11,
	0xe9, 0x26, 0x36, 0xc1, 0x83, 0x5e, 0x96, 0xfd, 0x67, 0xe6, 0xff, 0x76, 0x76, 0x06, 0xf6, 0x84,
	0xc6, 0x07, 0x4c, 0x8c, 0xe4, 0x71, 0x98, 0x72, 0x31, 0x46, 0xc3, 0x52, 0xad, 0x8c, 0xa2, 0x87,
	0xcf, 0x23, 0x1e, 0x63, 0x8c, 0xe9, 0x88, 0x27, 0x86, 0x09, 0xd4, 0x66, 0xca, 0xaa, 0xda, 0xe0,
	0x9d, 0x40, 0xe7, 0x6a, 0x25, 0x7b, 0xd6, 0x79, 0xcd, 0x0d, 0xa7, 0x17, 0xe0, 0x26, 0x6a, 0x79,
	0xf3, 0xc8, 0x11, 0x39, 0x69, 0x9f, 0x1e, 0xb3, 0x3f, 0x41, 0xec, 0xd6, 0x16, 0xdf, 0x38, 0x51,
	0x69, 0xa3, 0x11, 0xb4, 0x53, 0xad, 0xd4, 0xa0, 0x60, 0x7a, 0x6b, 0x96, 0xc2, 0xfe, 0xa1, 0xf4,
	0x2a, 0x47, 0x89, 0xab, 0x43, 0xba, 0x4d, 0x70, 0x8b, 0xcf, 0x05, 0x4d, 0x70, 0x8b, 0x17, 0x83,
	0x37, 0x02, 0xdb, 0xbf, 0x6c, 0xd4, 0x83, 0xcd, 0x2c, 0xef, 0x3f, 0xa2, 0x30, 0xb6, 0xfb, 0x56,
	0xf4, 0x23, 0xe9, 0x3e, 0x34, 0x9f, 0x50, 0xcb, 0x81, 0x44, 0x6d, 0x5b, 0x6a, 0x45, 0x2b, 0x4d,
	0x77, 0xc1, 0x95, 0x59, 0x96, 0xa3, 0xf6, 0xd6, 0x6d, 0xa6, 0x54, 0xb4, 0x03, 0x0d, 0x11, 0x73,
	0x39, 0xf1, 0x36, 0x6c, 0xb8, 0x10, 0xf4, 0x00, 0x5a, 0x99, 0x1c, 0x26, 0xdc, 0xe4, 0x1a, 0xbd,
	0x86, 0xcd, 0x54, 0x81, 0x60, 0x07, 0xb6, 0x6a, 0x4d, 0x5d, 0x8a, 0x71, 0xb7, 0xf7, 0x31, 0xf7,
	0xc9, 0x6c, 0xee, 0x93, 0xaf, 0xb9, 0x4f, 0x5e, 0x17, 0xbe, 0x33, 0x5b, 0xf8, 0xce, 0xe7, 0xc2,
	0x77, 0xee, 0xcf, 0x87, 0xd2, 0x8c, 0xf2, 0x3e, 0x13, 0x6a, 0x12, 0xd6, 0xc7, 0x13, 0x2e, 0xc7,
	0x73, 0x57, 0x9e, 0xd3, 0xb0, 0xb6, 0x61, 0xf3, 0x92, 0x62, 0xd6, 0x77, 0xed, 0x86, 0xcf, 0xbe,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x8d, 0xf3, 0x2d, 0xfc, 0x01, 0x00, 0x00,
}

func (m *CredentialPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *CredentialPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *CredentialPacketData_ProofPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialPacketData_ProofPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProofPacket != nil {
		{
			size, err := m.ProofPacket.MarshalToSizedBuffer(dAtA[:i])
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

func (m *ProofPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Claim) > 0 {
		i -= len(m.Claim)
		copy(dAtA[i:], m.Claim)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Claim)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Verifier) > 0 {
		i -= len(m.Verifier)
		copy(dAtA[i:], m.Verifier)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Verifier)))
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

func (m *ProofPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *CredentialPacketData) Size() (n int) {
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

func (m *CredentialPacketData_NoData) Size() (n int) {
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
func (m *CredentialPacketData_ProofPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProofPacket != nil {
		l = m.ProofPacket.Size()
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

func (m *ProofPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Verifier)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Claim)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *ProofPacketAck) Size() (n int) {
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
func (m *CredentialPacketData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CredentialPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.Packet = &CredentialPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofPacket", wireType)
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
			v := &ProofPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &CredentialPacketData_ProofPacket{v}
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
func (m *ProofPacketData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProofPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Verifier", wireType)
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
			m.Verifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
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
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			m.Signature = string(dAtA[iNdEx:postIndex])
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
func (m *ProofPacketAck) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProofPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
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