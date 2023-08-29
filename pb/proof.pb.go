// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/proof.proto

package pb

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

type Proof struct {
	// Start index of the leaves that match the queried namespace.ID.
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// End index (non-inclusive) of the leaves that match the queried
	// namespace.ID.
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	// Nodes hold the tree nodes necessary for the Merkle range proof.
	Nodes [][]byte `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// LeafHash contains the namespace.ID if NMT does not have it and
	// it should be proven. LeafHash is necessary to prove the Absence Proof.
	// This field will be empty in case of Inclusion Proof.
	LeafHash []byte `protobuf:"bytes,4,opt,name=leafHash,proto3" json:"leafHash,omitempty"`
	// The isMaxNamespaceIgnored flag influences the calculation of the
	// namespace ID range for intermediate nodes in the tree.
	IsMaxNamespaceIgnored bool `protobuf:"varint,5,opt,name=isMaxNamespaceIgnored,proto3" json:"isMaxNamespaceIgnored,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e2daa763cd7daf3, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Proof) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Proof) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Proof) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *Proof) GetIsMaxNamespaceIgnored() bool {
	if m != nil {
		return m.IsMaxNamespaceIgnored
	}
	return false
}

func init() {
	proto.RegisterType((*Proof)(nil), "proof.pb.Proof")
}

func init() { proto.RegisterFile("pb/proof.proto", fileDescriptor_2e2daa763cd7daf3) }

var fileDescriptor_2e2daa763cd7daf3 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0x2f,
	0x28, 0xca, 0xcf, 0x4f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x72, 0x92, 0x94,
	0xa6, 0x33, 0x72, 0xb1, 0x06, 0x80, 0x38, 0x42, 0x22, 0x5c, 0xac, 0xc5, 0x25, 0x89, 0x45, 0x25,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x6a, 0x5e, 0x8a,
	0x04, 0x13, 0x58, 0x0c, 0xc4, 0x04, 0xa9, 0xcb, 0xcb, 0x4f, 0x49, 0x2d, 0x96, 0x60, 0x56, 0x60,
	0xd6, 0xe0, 0x09, 0x82, 0x70, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0x52, 0x13, 0xd3, 0x3c, 0x12, 0x8b,
	0x33, 0x24, 0x58, 0x14, 0x18, 0x35, 0x78, 0x82, 0xe0, 0x7c, 0x21, 0x13, 0x2e, 0xd1, 0xcc, 0x62,
	0xdf, 0xc4, 0x0a, 0xbf, 0xc4, 0xdc, 0xd4, 0xe2, 0x82, 0xc4, 0xe4, 0x54, 0xcf, 0xf4, 0xbc, 0xfc,
	0xa2, 0xd4, 0x14, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0xec, 0x92, 0x4e, 0xe6, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17,
	0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x9a, 0x93, 0x5a, 0x5c, 0x92, 0x99, 0x98, 0x5f, 0x94, 0xae,
	0x9f, 0x97, 0x5b, 0xa2, 0x5f, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x64, 0x62, 0x02, 0x2f, 0xf5, 0x00, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsMaxNamespaceIgnored {
		i--
		if m.IsMaxNamespaceIgnored {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.LeafHash) > 0 {
		i -= len(m.LeafHash)
		copy(dAtA[i:], m.LeafHash)
		i = encodeVarintProof(dAtA, i, uint64(len(m.LeafHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Nodes[iNdEx])
			copy(dAtA[i:], m.Nodes[iNdEx])
			i = encodeVarintProof(dAtA, i, uint64(len(m.Nodes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.End != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x10
	}
	if m.Start != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovProof(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovProof(uint64(m.End))
	}
	if len(m.Nodes) > 0 {
		for _, b := range m.Nodes {
			l = len(b)
			n += 1 + l + sovProof(uint64(l))
		}
	}
	l = len(m.LeafHash)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	if m.IsMaxNamespaceIgnored {
		n += 2
	}
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, make([]byte, postIndex-iNdEx))
			copy(m.Nodes[len(m.Nodes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeafHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeafHash = append(m.LeafHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LeafHash == nil {
				m.LeafHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsMaxNamespaceIgnored", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
			m.IsMaxNamespaceIgnored = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)