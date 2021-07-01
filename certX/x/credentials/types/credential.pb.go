// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credentials/credential.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Credential struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id        uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Issuer    string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Verifier  string `protobuf:"bytes,4,opt,name=verifier,proto3" json:"verifier,omitempty"`
	Subject   string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Signature string `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Claim     string `protobuf:"bytes,7,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d919d03bd0f37e3, []int{0}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Credential) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Credential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Credential) GetVerifier() string {
	if m != nil {
		return m.Verifier
	}
	return ""
}

func (m *Credential) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Credential) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Credential) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func init() {
	proto.RegisterType((*Credential)(nil), "whalelephant.certx.credentials.Credential")
}

func init() { proto.RegisterFile("credentials/credential.proto", fileDescriptor_8d919d03bd0f37e3) }

var fileDescriptor_8d919d03bd0f37e3 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe3, 0xd0, 0xa6, 0xd4, 0x03, 0x83, 0x55, 0x21, 0xab, 0xaa, 0xac, 0x8a, 0xa9, 0x53,
	0x3c, 0x30, 0xb0, 0xc3, 0x09, 0xe8, 0x84, 0xd8, 0x1c, 0xe7, 0x91, 0x18, 0xa5, 0x71, 0x64, 0x3b,
	0x50, 0x6e, 0xc1, 0x71, 0x38, 0x02, 0x63, 0x47, 0x46, 0x94, 0x5c, 0x04, 0xc5, 0x09, 0x4d, 0x16,
	0xeb, 0x7d, 0xfe, 0xff, 0xf7, 0x86, 0x0f, 0x6f, 0xa4, 0x81, 0x14, 0x4a, 0xa7, 0x44, 0x61, 0xf9,
	0x38, 0xc7, 0x95, 0xd1, 0x4e, 0x13, 0xf6, 0x9e, 0x8b, 0x02, 0x0a, 0xa8, 0x72, 0x51, 0xba, 0x58,
	0x82, 0x71, 0xc7, 0x78, 0xb2, 0xb0, 0x5e, 0x65, 0x3a, 0xd3, 0xbe, 0xca, 0xbb, 0xa9, 0xdf, 0xba,
	0xf9, 0x42, 0x18, 0x3f, 0x9c, 0x5b, 0x84, 0xe2, 0x85, 0x34, 0x20, 0x9c, 0x36, 0x14, 0x6d, 0xd1,
	0x6e, 0xb9, 0xff, 0x47, 0x72, 0x85, 0x43, 0x95, 0xd2, 0x70, 0x8b, 0x76, 0xb3, 0x7d, 0xa8, 0x52,
	0x72, 0x8d, 0x23, 0x65, 0x6d, 0x0d, 0x86, 0x5e, 0xf8, 0xe2, 0x40, 0x64, 0x8d, 0x2f, 0xdf, 0xc0,
	0xa8, 0x17, 0x05, 0x86, 0xce, 0x7c, 0x72, 0xe6, 0xee, 0xba, 0xad, 0x93, 0x57, 0x90, 0x8e, 0xce,
	0xfb, 0xeb, 0x03, 0x92, 0x0d, 0x5e, 0x5a, 0x95, 0x95, 0xc2, 0xd5, 0x06, 0x68, 0xe4, 0xb3, 0xf1,
	0x83, 0xac, 0xf0, 0x5c, 0x16, 0x42, 0x1d, 0xe8, 0xc2, 0x27, 0x3d, 0xdc, 0x3f, 0x7e, 0x37, 0x0c,
	0x9d, 0x1a, 0x86, 0x7e, 0x1b, 0x86, 0x3e, 0x5b, 0x16, 0x9c, 0x5a, 0x16, 0xfc, 0xb4, 0x2c, 0x78,
	0xbe, 0xcb, 0x94, 0xcb, 0xeb, 0x24, 0x96, 0xfa, 0xc0, 0xa7, 0x56, 0x78, 0x67, 0xe5, 0x69, 0x78,
	0x8f, 0x7c, 0xaa, 0xd3, 0x7d, 0x54, 0x60, 0x93, 0xc8, 0x4b, 0xb9, 0xfd, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0xac, 0x55, 0xf3, 0x6a, 0x01, 0x00, 0x00,
}

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Claim) > 0 {
		i -= len(m.Claim)
		copy(dAtA[i:], m.Claim)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Claim)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Verifier) > 0 {
		i -= len(m.Verifier)
		copy(dAtA[i:], m.Verifier)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Verifier)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintCredential(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredential(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredential(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovCredential(uint64(m.Id))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.Verifier)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.Claim)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	return n
}

func sovCredential(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredential(x uint64) (n int) {
	return sovCredential(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredential
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
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Verifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claim = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredential
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
func skipCredential(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredential
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
					return 0, ErrIntOverflowCredential
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
					return 0, ErrIntOverflowCredential
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
				return 0, ErrInvalidLengthCredential
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredential
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredential
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredential        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredential          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredential = fmt.Errorf("proto: unexpected end of group")
)
