// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credentials/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgSendVerifiableCredential struct {
	Sender           string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	Subject          string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Verifier         string `protobuf:"bytes,6,opt,name=verifier,proto3" json:"verifier,omitempty"`
	Issuer           string `protobuf:"bytes,7,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Claim            string `protobuf:"bytes,8,opt,name=claim,proto3" json:"claim,omitempty"`
	Signature        string `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MsgSendVerifiableCredential) Reset()         { *m = MsgSendVerifiableCredential{} }
func (m *MsgSendVerifiableCredential) String() string { return proto.CompactTextString(m) }
func (*MsgSendVerifiableCredential) ProtoMessage()    {}
func (*MsgSendVerifiableCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deef4b20050de6e, []int{0}
}
func (m *MsgSendVerifiableCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendVerifiableCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendVerifiableCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendVerifiableCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendVerifiableCredential.Merge(m, src)
}
func (m *MsgSendVerifiableCredential) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendVerifiableCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendVerifiableCredential.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendVerifiableCredential proto.InternalMessageInfo

func (m *MsgSendVerifiableCredential) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendVerifiableCredential) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetVerifier() string {
	if m != nil {
		return m.Verifier
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *MsgSendVerifiableCredential) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type MsgSendVerifiableCredentialResponse struct {
}

func (m *MsgSendVerifiableCredentialResponse) Reset()         { *m = MsgSendVerifiableCredentialResponse{} }
func (m *MsgSendVerifiableCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendVerifiableCredentialResponse) ProtoMessage()    {}
func (*MsgSendVerifiableCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7deef4b20050de6e, []int{1}
}
func (m *MsgSendVerifiableCredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendVerifiableCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendVerifiableCredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendVerifiableCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendVerifiableCredentialResponse.Merge(m, src)
}
func (m *MsgSendVerifiableCredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendVerifiableCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendVerifiableCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendVerifiableCredentialResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendVerifiableCredential)(nil), "whalelephant.certx.credentials.MsgSendVerifiableCredential")
	proto.RegisterType((*MsgSendVerifiableCredentialResponse)(nil), "whalelephant.certx.credentials.MsgSendVerifiableCredentialResponse")
}

func init() { proto.RegisterFile("credentials/tx.proto", fileDescriptor_7deef4b20050de6e) }

var fileDescriptor_7deef4b20050de6e = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xeb, 0xfe, 0xaf, 0x27, 0x64, 0x55, 0xc8, 0x2a, 0xc8, 0xaa, 0x8a, 0x90, 0x2a, 0x86,
	0x44, 0x82, 0x81, 0x81, 0x8d, 0xb2, 0x30, 0x74, 0xa0, 0x20, 0x84, 0xd8, 0x9c, 0xf4, 0x48, 0x8c,
	0x12, 0x27, 0xb2, 0x1d, 0x28, 0x6f, 0xc1, 0xc2, 0xc6, 0xc8, 0xc3, 0x30, 0x76, 0x64, 0x44, 0xed,
	0x8b, 0xa0, 0x3a, 0x6d, 0x29, 0x42, 0x64, 0x61, 0x89, 0xf2, 0x7d, 0x77, 0xf6, 0xf7, 0xd3, 0xf9,
	0x70, 0xdb, 0x57, 0x30, 0x06, 0x69, 0x04, 0x8f, 0xb4, 0x6b, 0x26, 0x4e, 0xaa, 0x12, 0x93, 0x10,
	0xf6, 0x18, 0xf2, 0x08, 0x22, 0x48, 0x43, 0x2e, 0x8d, 0xe3, 0x83, 0x32, 0x13, 0x67, 0xa3, 0xb1,
	0xf7, 0x52, 0xc6, 0x3b, 0x43, 0x1d, 0x5c, 0x82, 0x1c, 0x5f, 0x83, 0x12, 0x77, 0x82, 0x7b, 0x11,
	0x0c, 0xd6, 0x0d, 0x64, 0x1b, 0xd7, 0x35, 0xc8, 0x31, 0x28, 0x8a, 0xba, 0xa8, 0xdf, 0x1a, 0x2d,
	0x15, 0x21, 0xb8, 0x9a, 0x26, 0xca, 0xd0, 0xb2, 0x75, 0xed, 0x3f, 0xd9, 0xc5, 0x2d, 0x3f, 0xe4,
	0x52, 0x42, 0x74, 0x7e, 0x46, 0x2b, 0xb6, 0xf0, 0x6d, 0x90, 0x03, 0xbc, 0x65, 0x44, 0x0c, 0x49,
	0x66, 0xae, 0x44, 0x0c, 0xda, 0xf0, 0x38, 0xa5, 0xd5, 0x2e, 0xea, 0x57, 0x47, 0xbf, 0x7c, 0x42,
	0x71, 0x43, 0x67, 0xde, 0x3d, 0xf8, 0x86, 0xd6, 0xec, 0x3d, 0x2b, 0x49, 0x3a, 0xb8, 0xf9, 0x60,
	0x39, 0x41, 0xd1, 0xba, 0x2d, 0xad, 0xf5, 0x82, 0x55, 0x68, 0x9d, 0x81, 0xa2, 0x8d, 0x9c, 0x35,
	0x57, 0xa4, 0x8d, 0x6b, 0x7e, 0xc4, 0x45, 0x4c, 0x9b, 0xd6, 0xce, 0xc5, 0x82, 0x56, 0x8b, 0x40,
	0x72, 0x93, 0x29, 0xa0, 0xad, 0x9c, 0x76, 0x6d, 0xf4, 0xf6, 0xf1, 0x5e, 0xc1, 0x58, 0x46, 0xa0,
	0xd3, 0x44, 0x6a, 0x38, 0x7c, 0x43, 0xb8, 0x32, 0xd4, 0x01, 0x79, 0x45, 0x98, 0xfe, 0x39, 0xc3,
	0x13, 0xa7, 0xf8, 0x11, 0x9c, 0x82, 0xa4, 0xce, 0xe0, 0x1f, 0x87, 0x57, 0x98, 0xa7, 0x17, 0xef,
	0x33, 0x86, 0xa6, 0x33, 0x86, 0x3e, 0x67, 0x0c, 0x3d, 0xcf, 0x59, 0x69, 0x3a, 0x67, 0xa5, 0x8f,
	0x39, 0x2b, 0xdd, 0x1e, 0x07, 0xc2, 0x84, 0x99, 0xe7, 0xf8, 0x49, 0xec, 0x6e, 0x06, 0xb9, 0x8b,
	0xa0, 0x9b, 0xe5, 0x77, 0xe2, 0xfe, 0xd8, 0xad, 0xa7, 0x14, 0xb4, 0x57, 0xb7, 0xfb, 0x75, 0xf4,
	0x15, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xc0, 0xb2, 0x8b, 0x77, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	SendVerifiableCredential(ctx context.Context, in *MsgSendVerifiableCredential, opts ...grpc.CallOption) (*MsgSendVerifiableCredentialResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendVerifiableCredential(ctx context.Context, in *MsgSendVerifiableCredential, opts ...grpc.CallOption) (*MsgSendVerifiableCredentialResponse, error) {
	out := new(MsgSendVerifiableCredentialResponse)
	err := c.cc.Invoke(ctx, "/whalelephant.certx.credentials.Msg/SendVerifiableCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	SendVerifiableCredential(context.Context, *MsgSendVerifiableCredential) (*MsgSendVerifiableCredentialResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendVerifiableCredential(ctx context.Context, req *MsgSendVerifiableCredential) (*MsgSendVerifiableCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifiableCredential not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendVerifiableCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendVerifiableCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendVerifiableCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whalelephant.certx.credentials.Msg/SendVerifiableCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendVerifiableCredential(ctx, req.(*MsgSendVerifiableCredential))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "whalelephant.certx.credentials.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVerifiableCredential",
			Handler:    _Msg_SendVerifiableCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credentials/tx.proto",
}

func (m *MsgSendVerifiableCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendVerifiableCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendVerifiableCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Claim) > 0 {
		i -= len(m.Claim)
		copy(dAtA[i:], m.Claim)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Claim)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Verifier) > 0 {
		i -= len(m.Verifier)
		copy(dAtA[i:], m.Verifier)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Verifier)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendVerifiableCredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendVerifiableCredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendVerifiableCredentialResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendVerifiableCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Verifier)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Claim)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendVerifiableCredentialResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendVerifiableCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendVerifiableCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendVerifiableCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Verifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claim = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSendVerifiableCredentialResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendVerifiableCredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendVerifiableCredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
