// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vac/query.proto

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
	return fileDescriptor_65bf4cde522b592c, []int{0}
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
	return fileDescriptor_65bf4cde522b592c, []int{1}
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
	return fileDescriptor_65bf4cde522b592c, []int{2}
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
	return fileDescriptor_65bf4cde522b592c, []int{3}
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
	proto.RegisterType((*QueryGetCredentialRequest)(nil), "whalelephant.muggleAuth.vac.QueryGetCredentialRequest")
	proto.RegisterType((*QueryGetCredentialResponse)(nil), "whalelephant.muggleAuth.vac.QueryGetCredentialResponse")
	proto.RegisterType((*QueryAllCredentialRequest)(nil), "whalelephant.muggleAuth.vac.QueryAllCredentialRequest")
	proto.RegisterType((*QueryAllCredentialResponse)(nil), "whalelephant.muggleAuth.vac.QueryAllCredentialResponse")
}

func init() { proto.RegisterFile("vac/query.proto", fileDescriptor_65bf4cde522b592c) }

var fileDescriptor_65bf4cde522b592c = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcb, 0x8e, 0xd3, 0x30,
	0x18, 0x85, 0xeb, 0x70, 0x59, 0x18, 0x01, 0x92, 0xc5, 0x02, 0x02, 0x8a, 0x50, 0x16, 0x94, 0x9b,
	0x6c, 0xb5, 0x54, 0xb0, 0x2e, 0x48, 0x74, 0x87, 0xa0, 0x2b, 0xc4, 0xce, 0x49, 0x7e, 0x25, 0x96,
	0xdc, 0x38, 0x8d, 0x9d, 0x42, 0x85, 0xd8, 0xf0, 0x04, 0x48, 0xbc, 0x07, 0x12, 0xe2, 0x25, 0x58,
	0x56, 0x62, 0xc3, 0x72, 0xd4, 0xce, 0x0b, 0xcc, 0x1b, 0x8c, 0x72, 0xd1, 0xc4, 0x9d, 0x69, 0x3b,
	0x9d, 0xd9, 0x3a, 0xff, 0x39, 0xff, 0x77, 0x8e, 0x1d, 0x7c, 0x7b, 0xc6, 0x43, 0x36, 0x2d, 0x20,
	0x9f, 0xd3, 0x2c, 0x57, 0x46, 0x91, 0xfb, 0x9f, 0x13, 0x2e, 0x41, 0x42, 0x96, 0xf0, 0xd4, 0xd0,
	0x49, 0x11, 0xc7, 0x12, 0x86, 0x85, 0x49, 0xe8, 0x8c, 0x87, 0xee, 0x83, 0x58, 0xa9, 0x58, 0x02,
	0xe3, 0x99, 0x60, 0x3c, 0x4d, 0x95, 0xe1, 0x46, 0xa8, 0x54, 0xd7, 0x52, 0xf7, 0x69, 0xa8, 0xf4,
	0x44, 0x69, 0x16, 0x70, 0x0d, 0xb5, 0x27, 0x9b, 0xf5, 0x02, 0x30, 0xbc, 0xc7, 0x32, 0x1e, 0x8b,
	0xb4, 0x1a, 0x6e, 0x66, 0xef, 0x94, 0x7b, 0xc3, 0x1c, 0x22, 0x48, 0x8d, 0xe0, 0xb2, 0x3e, 0xf5,
	0x9f, 0xe1, 0x7b, 0x1f, 0x4a, 0xdd, 0x08, 0xcc, 0x9b, 0x93, 0x6f, 0x63, 0x98, 0x16, 0xa0, 0x0d,
	0xb9, 0x85, 0x1d, 0x11, 0xdd, 0x45, 0x0f, 0xd1, 0xe3, 0xab, 0x63, 0x47, 0x44, 0x3e, 0x60, 0x77,
	0xd3, 0xb0, 0xce, 0x54, 0xaa, 0x81, 0x8c, 0x30, 0x6e, 0x4f, 0x2b, 0xd5, 0x8d, 0x7e, 0x97, 0xee,
	0x08, 0x47, 0x2d, 0x13, 0x4b, 0xea, 0x87, 0x0d, 0xd3, 0x50, 0xca, 0xb3, 0x4c, 0x6f, 0x31, 0x6e,
	0xa3, 0x35, 0x5b, 0x1e, 0xd1, 0xba, 0x07, 0x5a, 0xf6, 0x40, 0xeb, 0x6e, 0x9b, 0x1e, 0xe8, 0x7b,
	0x1e, 0x43, 0xa3, 0x1d, 0x5b, 0x4a, 0xff, 0x17, 0x6a, 0xc2, 0x9c, 0xda, 0xb2, 0x25, 0xcc, 0x95,
	0x4b, 0x86, 0x29, 0x8d, 0x2c, 0x5e, 0xa7, 0x69, 0xe5, 0x3c, 0xde, 0x9a, 0xc2, 0x06, 0xee, 0x1f,
	0x39, 0xf8, 0x5a, 0x05, 0x4c, 0xfe, 0x20, 0x1b, 0x8e, 0xbc, 0xdc, 0x89, 0xb5, 0xf5, 0x76, 0xdd,
	0x57, 0x17, 0xd6, 0xd5, 0x54, 0xfe, 0xe0, 0xfb, 0xbf, 0xc3, 0x9f, 0x0e, 0x25, 0xcf, 0x99, 0x6d,
	0xc0, 0x5a, 0x03, 0xb6, 0xfe, 0xd4, 0xd8, 0x57, 0x11, 0x7d, 0x23, 0xbf, 0x11, 0xbe, 0xd9, 0x9a,
	0x0d, 0xe5, 0x5e, 0xe0, 0x9b, 0x9e, 0xc0, 0x3e, 0xe0, 0x1b, 0x2f, 0xd5, 0x67, 0x15, 0xf8, 0x13,
	0xd2, 0xdd, 0x13, 0xfc, 0xf5, 0xbb, 0xbf, 0x4b, 0x0f, 0x2d, 0x96, 0x1e, 0x3a, 0x58, 0x7a, 0xe8,
	0xc7, 0xca, 0xeb, 0x2c, 0x56, 0x5e, 0xe7, 0xff, 0xca, 0xeb, 0x7c, 0x1a, 0xc4, 0xc2, 0x24, 0x45,
	0x40, 0x43, 0x35, 0x59, 0x37, 0x0b, 0x21, 0x37, 0x1f, 0x6d, 0xcb, 0x2f, 0x95, 0xa9, 0x99, 0x67,
	0xa0, 0x83, 0xeb, 0xd5, 0x4f, 0xf7, 0xe2, 0x38, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x07, 0x2d, 0xde,
	0x04, 0x04, 0x00, 0x00,
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
	// this line is used by starport scaffolding # 2
	Credential(ctx context.Context, in *QueryGetCredentialRequest, opts ...grpc.CallOption) (*QueryGetCredentialResponse, error)
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
	err := c.cc.Invoke(ctx, "/whalelephant.muggleAuth.vac.Query/Credential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CredentialAll(ctx context.Context, in *QueryAllCredentialRequest, opts ...grpc.CallOption) (*QueryAllCredentialResponse, error) {
	out := new(QueryAllCredentialResponse)
	err := c.cc.Invoke(ctx, "/whalelephant.muggleAuth.vac.Query/CredentialAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Credential(context.Context, *QueryGetCredentialRequest) (*QueryGetCredentialResponse, error)
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
		FullMethod: "/whalelephant.muggleAuth.vac.Query/Credential",
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
		FullMethod: "/whalelephant.muggleAuth.vac.Query/CredentialAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CredentialAll(ctx, req.(*QueryAllCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "whalelephant.muggleAuth.vac.Query",
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
	Metadata: "vac/query.proto",
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
