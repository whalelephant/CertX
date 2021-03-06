// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credentials/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// this line is used by starport scaffolding # 3
type QueryGetCredentialRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetCredentialRequest) Reset()         { *m = QueryGetCredentialRequest{} }
func (m *QueryGetCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCredentialRequest) ProtoMessage()    {}
func (*QueryGetCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28518a19b54a2e65, []int{0}
}
func (m *QueryGetCredentialRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCredentialRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCredentialRequest.Merge(m, src)
}
func (m *QueryGetCredentialRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCredentialRequest proto.InternalMessageInfo

func (m *QueryGetCredentialRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetCredentialResponse struct {
	Credential *Credential `protobuf:"bytes,1,opt,name=Credential,proto3" json:"Credential,omitempty"`
}

func (m *QueryGetCredentialResponse) Reset()         { *m = QueryGetCredentialResponse{} }
func (m *QueryGetCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCredentialResponse) ProtoMessage()    {}
func (*QueryGetCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28518a19b54a2e65, []int{1}
}
func (m *QueryGetCredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCredentialResponse.Merge(m, src)
}
func (m *QueryGetCredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCredentialResponse proto.InternalMessageInfo

func (m *QueryGetCredentialResponse) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type QueryAllCredentialRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCredentialRequest) Reset()         { *m = QueryAllCredentialRequest{} }
func (m *QueryAllCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCredentialRequest) ProtoMessage()    {}
func (*QueryAllCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28518a19b54a2e65, []int{2}
}
func (m *QueryAllCredentialRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCredentialRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCredentialRequest.Merge(m, src)
}
func (m *QueryAllCredentialRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCredentialRequest proto.InternalMessageInfo

func (m *QueryAllCredentialRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCredentialResponse struct {
	Credential []*Credential       `protobuf:"bytes,1,rep,name=Credential,proto3" json:"Credential,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCredentialResponse) Reset()         { *m = QueryAllCredentialResponse{} }
func (m *QueryAllCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCredentialResponse) ProtoMessage()    {}
func (*QueryAllCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28518a19b54a2e65, []int{3}
}
func (m *QueryAllCredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCredentialResponse.Merge(m, src)
}
func (m *QueryAllCredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCredentialResponse proto.InternalMessageInfo

func (m *QueryAllCredentialResponse) GetCredential() []*Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *QueryAllCredentialResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetCredentialRequest)(nil), "whalelephant.certx.credentials.QueryGetCredentialRequest")
	proto.RegisterType((*QueryGetCredentialResponse)(nil), "whalelephant.certx.credentials.QueryGetCredentialResponse")
	proto.RegisterType((*QueryAllCredentialRequest)(nil), "whalelephant.certx.credentials.QueryAllCredentialRequest")
	proto.RegisterType((*QueryAllCredentialResponse)(nil), "whalelephant.certx.credentials.QueryAllCredentialResponse")
}

func init() { proto.RegisterFile("credentials/query.proto", fileDescriptor_28518a19b54a2e65) }

var fileDescriptor_28518a19b54a2e65 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4e, 0xdb, 0x30,
	0x1c, 0xc6, 0xeb, 0x74, 0xdb, 0xc1, 0xd3, 0x76, 0xf0, 0x65, 0x5b, 0x54, 0x45, 0x53, 0x0e, 0xdb,
	0xd4, 0x4d, 0xb6, 0xda, 0x1d, 0x2a, 0xb8, 0x15, 0x24, 0x2a, 0x71, 0xa2, 0x3d, 0x21, 0x6e, 0x4e,
	0x62, 0x25, 0x96, 0xd2, 0x38, 0x8d, 0x5d, 0xa0, 0x42, 0x5c, 0xe0, 0x05, 0x90, 0x78, 0x12, 0xae,
	0x88, 0x07, 0xe0, 0x58, 0x89, 0x0b, 0x47, 0xd4, 0xf2, 0x20, 0xa8, 0x49, 0x44, 0x5c, 0x9a, 0xd2,
	0xaa, 0x97, 0x28, 0x8a, 0xff, 0xdf, 0xe7, 0xdf, 0xf7, 0xd9, 0x81, 0xdf, 0xdc, 0x84, 0x79, 0x2c,
	0x52, 0x9c, 0x86, 0x92, 0x0c, 0x86, 0x2c, 0x19, 0xe1, 0x38, 0x11, 0x4a, 0x20, 0xeb, 0x24, 0xa0,
	0x21, 0x0b, 0x59, 0x1c, 0xd0, 0x48, 0x61, 0x97, 0x25, 0xea, 0x14, 0x6b, 0xb3, 0x66, 0xcd, 0x17,
	0xc2, 0x0f, 0x19, 0xa1, 0x31, 0x27, 0x34, 0x8a, 0x84, 0xa2, 0x8a, 0x8b, 0x48, 0x66, 0x6a, 0xb3,
	0xee, 0x0a, 0xd9, 0x17, 0x92, 0x38, 0x54, 0xb2, 0xcc, 0x96, 0x1c, 0x37, 0x1c, 0xa6, 0x68, 0x83,
	0xc4, 0xd4, 0xe7, 0x51, 0x3a, 0x9c, 0xcf, 0xd6, 0x74, 0x84, 0xe2, 0x3d, 0x5b, 0xb5, 0xff, 0xc2,
	0x1f, 0xdd, 0x99, 0xbe, 0xc3, 0xd4, 0xee, 0xeb, 0x5a, 0x8f, 0x0d, 0x86, 0x4c, 0x2a, 0xf4, 0x15,
	0x1a, 0xdc, 0xfb, 0x0e, 0x7e, 0x82, 0x3f, 0x1f, 0x7a, 0x06, 0xf7, 0xec, 0x00, 0x9a, 0x65, 0xc3,
	0x32, 0x16, 0x91, 0x64, 0x68, 0x1f, 0xc2, 0xe2, 0x6b, 0xaa, 0xfa, 0xdc, 0xac, 0xe3, 0xf7, 0x73,
	0x62, 0xcd, 0x47, 0x53, 0xdb, 0x6e, 0x8e, 0xd5, 0x0e, 0xc3, 0x45, 0xac, 0x3d, 0x08, 0x8b, 0x94,
	0xf9, 0x46, 0xbf, 0x70, 0x56, 0x09, 0x9e, 0x55, 0x82, 0xb3, 0xa6, 0xf3, 0x4a, 0xf0, 0x01, 0xf5,
	0x59, 0xae, 0xed, 0x69, 0x4a, 0xfb, 0x06, 0xe4, 0x79, 0xde, 0xec, 0xb2, 0x24, 0x4f, 0x75, 0xf3,
	0x3c, 0xa8, 0x33, 0x87, 0x6c, 0xa4, 0xc8, 0xbf, 0x57, 0x22, 0x67, 0x20, 0x3a, 0x73, 0xf3, 0xb2,
	0x0a, 0x3f, 0xa6, 0xcc, 0xe8, 0x0e, 0xe8, 0x7c, 0x68, 0x6b, 0x15, 0xd9, 0xd2, 0x63, 0x36, 0xb7,
	0x37, 0x91, 0x66, 0x6c, 0x76, 0xeb, 0xe2, 0xe1, 0xf9, 0xda, 0x68, 0x20, 0x42, 0x74, 0x0f, 0x92,
	0x7a, 0x90, 0xf2, 0x9b, 0x47, 0xce, 0xb8, 0x77, 0x8e, 0x6e, 0x01, 0xfc, 0x52, 0xf8, 0xb5, 0xc3,
	0x75, 0x13, 0x94, 0xdd, 0x88, 0x35, 0x13, 0x94, 0x1e, 0xb3, 0xdd, 0x4c, 0x13, 0xfc, 0x43, 0xf5,
	0xf5, 0x13, 0xec, 0x74, 0xef, 0x27, 0x16, 0x18, 0x4f, 0x2c, 0xf0, 0x34, 0xb1, 0xc0, 0xd5, 0xd4,
	0xaa, 0x8c, 0xa7, 0x56, 0xe5, 0x71, 0x6a, 0x55, 0x8e, 0x5a, 0x3e, 0x57, 0xc1, 0xd0, 0xc1, 0xae,
	0xe8, 0x2f, 0xfa, 0x1d, 0xe6, 0xcf, 0x79, 0x5f, 0x35, 0x8a, 0x99, 0x74, 0x3e, 0xa5, 0xff, 0xe3,
	0xff, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x96, 0x0c, 0x30, 0x32, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a credential by id.
	Credential(ctx context.Context, in *QueryGetCredentialRequest, opts ...grpc.CallOption) (*QueryGetCredentialResponse, error)
	// Queries a list of credential items.
	CredentialAll(ctx context.Context, in *QueryAllCredentialRequest, opts ...grpc.CallOption) (*QueryAllCredentialResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Credential(ctx context.Context, in *QueryGetCredentialRequest, opts ...grpc.CallOption) (*QueryGetCredentialResponse, error) {
	out := new(QueryGetCredentialResponse)
	err := c.cc.Invoke(ctx, "/whalelephant.certx.credentials.Query/Credential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CredentialAll(ctx context.Context, in *QueryAllCredentialRequest, opts ...grpc.CallOption) (*QueryAllCredentialResponse, error) {
	out := new(QueryAllCredentialResponse)
	err := c.cc.Invoke(ctx, "/whalelephant.certx.credentials.Query/CredentialAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a credential by id.
	Credential(context.Context, *QueryGetCredentialRequest) (*QueryGetCredentialResponse, error)
	// Queries a list of credential items.
	CredentialAll(context.Context, *QueryAllCredentialRequest) (*QueryAllCredentialResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Credential(ctx context.Context, req *QueryGetCredentialRequest) (*QueryGetCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Credential not implemented")
}
func (*UnimplementedQueryServer) CredentialAll(ctx context.Context, req *QueryAllCredentialRequest) (*QueryAllCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Credential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Credential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whalelephant.certx.credentials.Query/Credential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Credential(ctx, req.(*QueryGetCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CredentialAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CredentialAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whalelephant.certx.credentials.Query/CredentialAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CredentialAll(ctx, req.(*QueryAllCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "whalelephant.certx.credentials.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Credential",
			Handler:    _Query_Credential_Handler,
		},
		{
			MethodName: "CredentialAll",
			Handler:    _Query_CredentialAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credentials/query.proto",
}

func (m *QueryGetCredentialRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCredentialRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCredentialRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCredentialResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Credential != nil {
		{
			size, err := m.Credential.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCredentialRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCredentialRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCredentialRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCredentialResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Credential) > 0 {
		for iNdEx := len(m.Credential) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Credential[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCredentialRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetCredentialResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Credential != nil {
		l = m.Credential.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCredentialRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCredentialResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Credential) > 0 {
		for _, e := range m.Credential {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCredentialRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCredentialRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCredentialRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCredentialResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Credential == nil {
				m.Credential = &Credential{}
			}
			if err := m.Credential.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCredentialRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCredentialRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCredentialRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCredentialResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credential = append(m.Credential, &Credential{})
			if err := m.Credential[len(m.Credential)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
