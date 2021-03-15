// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pooltoy/v1beta1/user.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type User struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Creator     string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	UserAccount string `protobuf:"bytes,3,opt,name=user_account,json=userAccount,proto3" json:"user_account,omitempty" yaml:"userAccount"`
	IsAdmin     bool   `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty" yaml:"isAdmin"`
	Name        string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty" yaml:"email"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbd8b8e3bf177559, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *User) GetUserAccount() string {
	if m != nil {
		return m.UserAccount
	}
	return ""
}

func (m *User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "pooltoy.v1beta1.User")
}

func init() { proto.RegisterFile("pooltoy/v1beta1/user.proto", fileDescriptor_bbd8b8e3bf177559) }

var fileDescriptor_bbd8b8e3bf177559 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x9b, 0xd8, 0xdf, 0x69, 0xb5, 0x32, 0x88, 0xa4, 0x05, 0x93, 0x32, 0x82, 0x74, 0xa1,
	0x0d, 0xa5, 0x2b, 0xdd, 0xb5, 0x7b, 0x37, 0x01, 0x37, 0x6e, 0xca, 0x24, 0x19, 0xd3, 0x81, 0x24,
	0xb7, 0x24, 0x13, 0x31, 0xaf, 0xe0, 0xca, 0xc7, 0x72, 0xd9, 0xa5, 0xab, 0x20, 0xed, 0x1b, 0xe4,
	0x09, 0x24, 0x33, 0xa9, 0x74, 0x77, 0xcf, 0xf9, 0xbe, 0xcb, 0x25, 0x19, 0x34, 0xde, 0x02, 0x84,
	0x02, 0x72, 0xfb, 0x7d, 0xee, 0x32, 0x41, 0xe7, 0x76, 0x96, 0xb2, 0x64, 0xb6, 0x4d, 0x40, 0x00,
	0x1e, 0xd6, 0x6c, 0x56, 0xb3, 0xf1, 0xc8, 0x83, 0x34, 0x82, 0x74, 0x2d, 0xb1, 0xad, 0x82, 0x72,
	0xc7, 0x57, 0x01, 0x04, 0xa0, 0xfa, 0x6a, 0xaa, 0xdb, 0x51, 0x00, 0x10, 0x84, 0xcc, 0x96, 0xc9,
	0xcd, 0xde, 0x6c, 0x1a, 0xe7, 0x0a, 0x91, 0x4f, 0x1d, 0x35, 0x5f, 0x52, 0x96, 0xe0, 0x1b, 0xa4,
	0x73, 0xdf, 0xd0, 0x26, 0xda, 0xb4, 0xb7, 0x3a, 0x2f, 0x0b, 0xab, 0x97, 0xd3, 0x28, 0x7c, 0x22,
	0xdc, 0x27, 0x8e, 0xce, 0x7d, 0x7c, 0x8f, 0x3a, 0x5e, 0xc2, 0xa8, 0x80, 0xc4, 0xd0, 0xa5, 0x83,
	0xcb, 0xc2, 0xba, 0x50, 0x4e, 0x0d, 0x88, 0x73, 0x54, 0xf0, 0x23, 0x1a, 0x54, 0x1f, 0xb0, 0xa6,
	0x9e, 0x07, 0x59, 0x2c, 0x8c, 0x33, 0xb9, 0x72, 0x5d, 0x16, 0x16, 0x56, 0x2b, 0x15, 0x5d, 0x2a,
	0x48, 0x9c, 0xfe, 0x49, 0xc2, 0x0f, 0xa8, 0xcb, 0xd3, 0x35, 0xf5, 0x23, 0x1e, 0x1b, 0xcd, 0x89,
	0x36, 0xed, 0x9e, 0x5e, 0xe2, 0xe9, 0xb2, 0x02, 0xc4, 0xe9, 0xd4, 0x13, 0xbe, 0x45, 0xcd, 0x98,
	0x46, 0xcc, 0x68, 0xc9, 0x0b, 0xc3, 0xb2, 0xb0, 0xfa, 0x4a, 0xad, 0x5a, 0xe2, 0x48, 0x88, 0xef,
	0x50, 0x8b, 0x45, 0x94, 0x87, 0x46, 0x5b, 0x5a, 0x97, 0x65, 0x61, 0x0d, 0x94, 0x25, 0x6b, 0xe2,
	0x28, 0xbc, 0x7a, 0xfe, 0xde, 0x9b, 0xda, 0x6e, 0x6f, 0x6a, 0xbf, 0x7b, 0x53, 0xfb, 0x3a, 0x98,
	0x8d, 0xdd, 0xc1, 0x6c, 0xfc, 0x1c, 0xcc, 0xc6, 0xeb, 0x22, 0xe0, 0x62, 0x93, 0xb9, 0x33, 0x0f,
	0x22, 0x9b, 0xc7, 0x82, 0x25, 0xde, 0x86, 0xf2, 0xd8, 0x65, 0x49, 0xc8, 0x63, 0xfb, 0xf8, 0x76,
	0x1f, 0xff, 0x93, 0xc8, 0xb7, 0x2c, 0x75, 0xdb, 0xf2, 0x17, 0x2f, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf7, 0x4b, 0x0b, 0x64, 0xdd, 0x01, 0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if m.IsAdmin {
		i--
		if m.IsAdmin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.UserAccount) > 0 {
		i -= len(m.UserAccount)
		copy(dAtA[i:], m.UserAccount)
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserAccount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UserAccount)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.IsAdmin {
		n += 2
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsAdmin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsAdmin = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
