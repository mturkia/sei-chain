// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/contract.proto

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

type ContractInfo struct {
	CodeId                 uint64   `protobuf:"varint,1,opt,name=codeId,proto3" json:"codeId,omitempty"`
	ContractAddr           string   `protobuf:"bytes,2,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
	NeedHook               bool     `protobuf:"varint,3,opt,name=NeedHook,proto3" json:"NeedHook,omitempty"`
	NeedOrderMatching      bool     `protobuf:"varint,4,opt,name=NeedOrderMatching,proto3" json:"NeedOrderMatching,omitempty"`
	DependentContractAddrs []string `protobuf:"bytes,5,rep,name=dependentContractAddrs,proto3" json:"dependentContractAddrs,omitempty"`
	NumIncomingPaths       int64    `protobuf:"varint,6,opt,name=numIncomingPaths,proto3" json:"numIncomingPaths,omitempty"`
}

func (m *ContractInfo) Reset()         { *m = ContractInfo{} }
func (m *ContractInfo) String() string { return proto.CompactTextString(m) }
func (*ContractInfo) ProtoMessage()    {}
func (*ContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee35557664974a8a, []int{0}
}
func (m *ContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractInfo.Merge(m, src)
}
func (m *ContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractInfo proto.InternalMessageInfo

func (m *ContractInfo) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *ContractInfo) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *ContractInfo) GetNeedHook() bool {
	if m != nil {
		return m.NeedHook
	}
	return false
}

func (m *ContractInfo) GetNeedOrderMatching() bool {
	if m != nil {
		return m.NeedOrderMatching
	}
	return false
}

func (m *ContractInfo) GetDependentContractAddrs() []string {
	if m != nil {
		return m.DependentContractAddrs
	}
	return nil
}

func (m *ContractInfo) GetNumIncomingPaths() int64 {
	if m != nil {
		return m.NumIncomingPaths
	}
	return 0
}

func init() {
	proto.RegisterType((*ContractInfo)(nil), "seiprotocol.seichain.dex.ContractInfo")
}

func init() { proto.RegisterFile("dex/contract.proto", fileDescriptor_ee35557664974a8a) }

var fileDescriptor_ee35557664974a8a = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3d, 0x4b, 0xc3, 0x40,
	0x18, 0xc7, 0x7b, 0xa6, 0x86, 0xf6, 0xe8, 0xa0, 0x37, 0x94, 0xc3, 0xe1, 0x08, 0x9d, 0x82, 0xd8,
	0x64, 0x10, 0xdc, 0xb5, 0x83, 0x66, 0xf0, 0x85, 0x8c, 0x6e, 0xe9, 0xdd, 0x63, 0x72, 0x68, 0xee,
	0x09, 0xb9, 0x2b, 0xc4, 0x6f, 0xe1, 0xc7, 0x72, 0xec, 0xe8, 0x28, 0xc9, 0x97, 0x70, 0x94, 0x86,
	0x56, 0x94, 0xe2, 0xf6, 0x7f, 0xe3, 0x79, 0xe0, 0x47, 0x99, 0x82, 0x26, 0x96, 0x68, 0x5c, 0x9d,
	0x49, 0x17, 0x55, 0x35, 0x3a, 0x64, 0xdc, 0x82, 0xee, 0x95, 0xc4, 0x97, 0xc8, 0x82, 0x96, 0x45,
	0xa6, 0x4d, 0xa4, 0xa0, 0x99, 0x7d, 0x11, 0x3a, 0x59, 0x6c, 0xc7, 0x89, 0x79, 0x42, 0x36, 0xa5,
	0xbe, 0x44, 0x05, 0x89, 0xe2, 0x24, 0x20, 0xe1, 0x30, 0xdd, 0x3a, 0x36, 0xa3, 0x93, 0xdd, 0xd1,
	0x4b, 0xa5, 0x6a, 0x7e, 0x10, 0x90, 0x70, 0x9c, 0xfe, 0xc9, 0xd8, 0x09, 0x1d, 0xdd, 0x01, 0xa8,
	0x1b, 0xc4, 0x67, 0xee, 0x05, 0x24, 0x1c, 0xa5, 0x3f, 0x9e, 0x9d, 0xd1, 0xe3, 0x8d, 0xbe, 0xaf,
	0x15, 0xd4, 0xb7, 0x99, 0x93, 0x85, 0x36, 0x39, 0x1f, 0xf6, 0xa3, 0xfd, 0x82, 0x5d, 0xd0, 0xa9,
	0x82, 0x0a, 0x8c, 0x02, 0xe3, 0x16, 0xbf, 0x5e, 0x58, 0x7e, 0x18, 0x78, 0xe1, 0x38, 0xfd, 0xa7,
	0x65, 0xa7, 0xf4, 0xc8, 0xac, 0xca, 0xc4, 0x48, 0x2c, 0xb5, 0xc9, 0x1f, 0x32, 0x57, 0x58, 0xee,
	0x07, 0x24, 0xf4, 0xd2, 0xbd, 0xfc, 0xea, 0xfa, 0xbd, 0x15, 0x64, 0xdd, 0x0a, 0xf2, 0xd9, 0x0a,
	0xf2, 0xd6, 0x89, 0xc1, 0xba, 0x13, 0x83, 0x8f, 0x4e, 0x0c, 0x1e, 0xe7, 0xb9, 0x76, 0xc5, 0x6a,
	0x19, 0x49, 0x2c, 0x63, 0x0b, 0x7a, 0xbe, 0x43, 0xd7, 0x9b, 0x9e, 0x5d, 0xdc, 0xc4, 0x1b, 0xcc,
	0xee, 0xb5, 0x02, 0xbb, 0xf4, 0xfb, 0xfe, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x41, 0xbf,
	0x86, 0x7a, 0x01, 0x00, 0x00,
}

func (m *ContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumIncomingPaths != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.NumIncomingPaths))
		i--
		dAtA[i] = 0x30
	}
	if len(m.DependentContractAddrs) > 0 {
		for iNdEx := len(m.DependentContractAddrs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DependentContractAddrs[iNdEx])
			copy(dAtA[i:], m.DependentContractAddrs[iNdEx])
			i = encodeVarintContract(dAtA, i, uint64(len(m.DependentContractAddrs[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.NeedOrderMatching {
		i--
		if m.NeedOrderMatching {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.NeedHook {
		i--
		if m.NeedHook {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractAddr) > 0 {
		i -= len(m.ContractAddr)
		copy(dAtA[i:], m.ContractAddr)
		i = encodeVarintContract(dAtA, i, uint64(len(m.ContractAddr)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovContract(uint64(m.CodeId))
	}
	l = len(m.ContractAddr)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.NeedHook {
		n += 2
	}
	if m.NeedOrderMatching {
		n += 2
	}
	if len(m.DependentContractAddrs) > 0 {
		for _, s := range m.DependentContractAddrs {
			l = len(s)
			n += 1 + l + sovContract(uint64(l))
		}
	}
	if m.NumIncomingPaths != 0 {
		n += 1 + sovContract(uint64(m.NumIncomingPaths))
	}
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: ContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedHook", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
			m.NeedHook = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedOrderMatching", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
			m.NeedOrderMatching = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DependentContractAddrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DependentContractAddrs = append(m.DependentContractAddrs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumIncomingPaths", wireType)
			}
			m.NumIncomingPaths = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumIncomingPaths |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)
