// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eCredentials/query.proto

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
type QueryGetECredentialRecordRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetECredentialRecordRequest) Reset()         { *m = QueryGetECredentialRecordRequest{} }
func (m *QueryGetECredentialRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetECredentialRecordRequest) ProtoMessage()    {}
func (*QueryGetECredentialRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72133c1e63390b40, []int{0}
}
func (m *QueryGetECredentialRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetECredentialRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetECredentialRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetECredentialRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetECredentialRecordRequest.Merge(m, src)
}
func (m *QueryGetECredentialRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetECredentialRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetECredentialRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetECredentialRecordRequest proto.InternalMessageInfo

func (m *QueryGetECredentialRecordRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetECredentialRecordResponse struct {
	ECredentialRecord *ECredentialRecord `protobuf:"bytes,1,opt,name=ECredentialRecord,proto3" json:"ECredentialRecord,omitempty"`
}

func (m *QueryGetECredentialRecordResponse) Reset()         { *m = QueryGetECredentialRecordResponse{} }
func (m *QueryGetECredentialRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetECredentialRecordResponse) ProtoMessage()    {}
func (*QueryGetECredentialRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72133c1e63390b40, []int{1}
}
func (m *QueryGetECredentialRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetECredentialRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetECredentialRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetECredentialRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetECredentialRecordResponse.Merge(m, src)
}
func (m *QueryGetECredentialRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetECredentialRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetECredentialRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetECredentialRecordResponse proto.InternalMessageInfo

func (m *QueryGetECredentialRecordResponse) GetECredentialRecord() *ECredentialRecord {
	if m != nil {
		return m.ECredentialRecord
	}
	return nil
}

type QueryAllECredentialRecordRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllECredentialRecordRequest) Reset()         { *m = QueryAllECredentialRecordRequest{} }
func (m *QueryAllECredentialRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllECredentialRecordRequest) ProtoMessage()    {}
func (*QueryAllECredentialRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72133c1e63390b40, []int{2}
}
func (m *QueryAllECredentialRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllECredentialRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllECredentialRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllECredentialRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllECredentialRecordRequest.Merge(m, src)
}
func (m *QueryAllECredentialRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllECredentialRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllECredentialRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllECredentialRecordRequest proto.InternalMessageInfo

func (m *QueryAllECredentialRecordRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllECredentialRecordResponse struct {
	ECredentialRecord []*ECredentialRecord `protobuf:"bytes,1,rep,name=ECredentialRecord,proto3" json:"ECredentialRecord,omitempty"`
	Pagination        *query.PageResponse  `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllECredentialRecordResponse) Reset()         { *m = QueryAllECredentialRecordResponse{} }
func (m *QueryAllECredentialRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllECredentialRecordResponse) ProtoMessage()    {}
func (*QueryAllECredentialRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72133c1e63390b40, []int{3}
}
func (m *QueryAllECredentialRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllECredentialRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllECredentialRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllECredentialRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllECredentialRecordResponse.Merge(m, src)
}
func (m *QueryAllECredentialRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllECredentialRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllECredentialRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllECredentialRecordResponse proto.InternalMessageInfo

func (m *QueryAllECredentialRecordResponse) GetECredentialRecord() []*ECredentialRecord {
	if m != nil {
		return m.ECredentialRecord
	}
	return nil
}

func (m *QueryAllECredentialRecordResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetECredentialRecordRequest)(nil), "whalelephant.certx.eCredentials.QueryGetECredentialRecordRequest")
	proto.RegisterType((*QueryGetECredentialRecordResponse)(nil), "whalelephant.certx.eCredentials.QueryGetECredentialRecordResponse")
	proto.RegisterType((*QueryAllECredentialRecordRequest)(nil), "whalelephant.certx.eCredentials.QueryAllECredentialRecordRequest")
	proto.RegisterType((*QueryAllECredentialRecordResponse)(nil), "whalelephant.certx.eCredentials.QueryAllECredentialRecordResponse")
}

func init() { proto.RegisterFile("eCredentials/query.proto", fileDescriptor_72133c1e63390b40) }

var fileDescriptor_72133c1e63390b40 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xa9, 0x7a, 0x18, 0x41, 0x70, 0xf0, 0x50, 0x8a, 0xc4, 0x1a, 0x44, 0xc5, 0xc3,
	0x0c, 0x8d, 0x82, 0x3f, 0x2e, 0xd2, 0x8a, 0xf6, 0xaa, 0x39, 0x89, 0x27, 0x27, 0xc9, 0x23, 0x1d,
	0x99, 0x66, 0xd2, 0xcc, 0x54, 0x2d, 0xe2, 0x45, 0xf0, 0x2e, 0xf8, 0x17, 0x79, 0xf3, 0x24, 0x05,
	0x2f, 0x8a, 0x17, 0x69, 0xf7, 0x0f, 0x59, 0x9a, 0x64, 0xd9, 0xb4, 0xe9, 0x26, 0x5d, 0x76, 0x2f,
	0x81, 0x90, 0xf7, 0x3e, 0xf3, 0xfd, 0xbc, 0x3c, 0x06, 0x77, 0xe0, 0x59, 0x0a, 0x21, 0xc4, 0x46,
	0x70, 0xa9, 0xd9, 0x74, 0x06, 0xe9, 0x9c, 0x26, 0xa9, 0x32, 0x8a, 0xdc, 0xf8, 0x30, 0xe6, 0x12,
	0x24, 0x24, 0x63, 0x1e, 0x1b, 0x1a, 0x40, 0x6a, 0x3e, 0xd2, 0x72, 0x71, 0xf7, 0x7a, 0xa4, 0x54,
	0x24, 0x81, 0xf1, 0x44, 0x30, 0x1e, 0xc7, 0xca, 0x70, 0x23, 0x54, 0xac, 0xf3, 0xf6, 0xee, 0xbd,
	0x40, 0xe9, 0x89, 0xd2, 0xcc, 0xe7, 0x1a, 0x72, 0x2e, 0x7b, 0xdf, 0xf7, 0xc1, 0xf0, 0x3e, 0x4b,
	0x78, 0x24, 0xe2, 0xac, 0xb8, 0xa8, 0xbd, 0xb5, 0x11, 0xa2, 0xf4, 0xe2, 0x41, 0xa0, 0xd2, 0x30,
	0xaf, 0x72, 0x5c, 0xdc, 0x7b, 0xb5, 0xe6, 0x8c, 0xc0, 0x3c, 0xdf, 0x2e, 0xf1, 0x60, 0x3a, 0x03,
	0x6d, 0xc8, 0x15, 0x6c, 0x89, 0xb0, 0x83, 0x7a, 0xe8, 0xee, 0x05, 0xcf, 0x12, 0xa1, 0xf3, 0x15,
	0xe1, 0x9b, 0x35, 0x4d, 0x3a, 0x51, 0xb1, 0x06, 0xf2, 0x16, 0x5f, 0xad, 0x7c, 0xcc, 0x20, 0x97,
	0x5d, 0x97, 0x36, 0x8c, 0x81, 0x56, 0xb1, 0x55, 0x98, 0xf3, 0xae, 0xc8, 0x3e, 0x90, 0xf2, 0xc4,
	0xec, 0x2f, 0x30, 0x3e, 0x9e, 0x4c, 0x71, 0xfc, 0x6d, 0x9a, 0x8f, 0x91, 0xae, 0xc7, 0x48, 0xf3,
	0xdf, 0x53, 0x8c, 0x91, 0xbe, 0xe4, 0x11, 0x14, 0xbd, 0x5e, 0xa9, 0xd3, 0xf9, 0x75, 0xe4, 0xbc,
	0xfb, 0xb0, 0x7a, 0xe7, 0xf6, 0xb9, 0x39, 0x93, 0xd1, 0x86, 0x8f, 0x95, 0xf9, 0xdc, 0x69, 0xf4,
	0xc9, 0xe3, 0x95, 0x85, 0xdc, 0x1f, 0x6d, 0x7c, 0x31, 0x13, 0x22, 0xff, 0xd0, 0x8e, 0xd4, 0x64,
	0xd0, 0x98, 0xb7, 0x69, 0x6f, 0xba, 0xc3, 0xb3, 0x20, 0xf2, 0xc8, 0xce, 0xd3, 0x2f, 0xbf, 0x0f,
	0xbe, 0x5b, 0x8f, 0xc9, 0x43, 0x56, 0x66, 0xb1, 0x8c, 0xc5, 0xea, 0x37, 0x9c, 0x7d, 0x12, 0xe1,
	0x67, 0xf2, 0x17, 0xe1, 0x6b, 0x15, 0xfc, 0x40, 0xca, 0x7d, 0x05, 0x6b, 0x96, 0x6b, 0x5f, 0xc1,
	0xba, 0x95, 0x71, 0x9e, 0x64, 0x82, 0x0f, 0x88, 0x7b, 0x7a, 0xc1, 0xa1, 0xf7, 0x73, 0x69, 0xa3,
	0xc5, 0xd2, 0x46, 0xff, 0x97, 0x36, 0xfa, 0xb6, 0xb2, 0x5b, 0x8b, 0x95, 0xdd, 0xfa, 0xb3, 0xb2,
	0x5b, 0x6f, 0x1e, 0x45, 0xc2, 0x8c, 0x67, 0x3e, 0x0d, 0xd4, 0xa4, 0xca, 0x7d, 0x5d, 0x3c, 0xb7,
	0xf8, 0x66, 0x9e, 0x80, 0xf6, 0x2f, 0x65, 0xf7, 0xc2, 0xfd, 0xc3, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x51, 0xab, 0x48, 0x36, 0xc4, 0x04, 0x00, 0x00,
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
	// Queries a eCredentialRecord by id.
	ECredentialRecord(ctx context.Context, in *QueryGetECredentialRecordRequest, opts ...grpc.CallOption) (*QueryGetECredentialRecordResponse, error)
	// Queries a list of eCredentialRecord items.
	ECredentialRecordAll(ctx context.Context, in *QueryAllECredentialRecordRequest, opts ...grpc.CallOption) (*QueryAllECredentialRecordResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ECredentialRecord(ctx context.Context, in *QueryGetECredentialRecordRequest, opts ...grpc.CallOption) (*QueryGetECredentialRecordResponse, error) {
	out := new(QueryGetECredentialRecordResponse)
	err := c.cc.Invoke(ctx, "/whalelephant.certx.eCredentials.Query/ECredentialRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ECredentialRecordAll(ctx context.Context, in *QueryAllECredentialRecordRequest, opts ...grpc.CallOption) (*QueryAllECredentialRecordResponse, error) {
	out := new(QueryAllECredentialRecordResponse)
	err := c.cc.Invoke(ctx, "/whalelephant.certx.eCredentials.Query/ECredentialRecordAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a eCredentialRecord by id.
	ECredentialRecord(context.Context, *QueryGetECredentialRecordRequest) (*QueryGetECredentialRecordResponse, error)
	// Queries a list of eCredentialRecord items.
	ECredentialRecordAll(context.Context, *QueryAllECredentialRecordRequest) (*QueryAllECredentialRecordResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ECredentialRecord(ctx context.Context, req *QueryGetECredentialRecordRequest) (*QueryGetECredentialRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ECredentialRecord not implemented")
}
func (*UnimplementedQueryServer) ECredentialRecordAll(ctx context.Context, req *QueryAllECredentialRecordRequest) (*QueryAllECredentialRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ECredentialRecordAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ECredentialRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetECredentialRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ECredentialRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whalelephant.certx.eCredentials.Query/ECredentialRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ECredentialRecord(ctx, req.(*QueryGetECredentialRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ECredentialRecordAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllECredentialRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ECredentialRecordAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whalelephant.certx.eCredentials.Query/ECredentialRecordAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ECredentialRecordAll(ctx, req.(*QueryAllECredentialRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "whalelephant.certx.eCredentials.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ECredentialRecord",
			Handler:    _Query_ECredentialRecord_Handler,
		},
		{
			MethodName: "ECredentialRecordAll",
			Handler:    _Query_ECredentialRecordAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eCredentials/query.proto",
}

func (m *QueryGetECredentialRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetECredentialRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetECredentialRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryGetECredentialRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetECredentialRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetECredentialRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ECredentialRecord != nil {
		{
			size, err := m.ECredentialRecord.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllECredentialRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllECredentialRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllECredentialRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllECredentialRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllECredentialRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllECredentialRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.ECredentialRecord) > 0 {
		for iNdEx := len(m.ECredentialRecord) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ECredentialRecord[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetECredentialRecordRequest) Size() (n int) {
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

func (m *QueryGetECredentialRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ECredentialRecord != nil {
		l = m.ECredentialRecord.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllECredentialRecordRequest) Size() (n int) {
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

func (m *QueryAllECredentialRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ECredentialRecord) > 0 {
		for _, e := range m.ECredentialRecord {
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
func (m *QueryGetECredentialRecordRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetECredentialRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetECredentialRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryGetECredentialRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetECredentialRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetECredentialRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ECredentialRecord", wireType)
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
			if m.ECredentialRecord == nil {
				m.ECredentialRecord = &ECredentialRecord{}
			}
			if err := m.ECredentialRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllECredentialRecordRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllECredentialRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllECredentialRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllECredentialRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllECredentialRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllECredentialRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ECredentialRecord", wireType)
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
			m.ECredentialRecord = append(m.ECredentialRecord, &ECredentialRecord{})
			if err := m.ECredentialRecord[len(m.ECredentialRecord)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
