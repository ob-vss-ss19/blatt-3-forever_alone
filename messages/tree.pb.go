// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tree.proto

package messages

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HelloWorld struct {
}

func (m *HelloWorld) Reset()      { *m = HelloWorld{} }
func (*HelloWorld) ProtoMessage() {}
func (*HelloWorld) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{0}
}
func (m *HelloWorld) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloWorld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloWorld.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloWorld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorld.Merge(m, src)
}
func (m *HelloWorld) XXX_Size() int {
	return m.Size()
}
func (m *HelloWorld) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorld.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorld proto.InternalMessageInfo

type Create struct {
	MaxSize int32 `protobuf:"varint,2,opt,name=maxSize,proto3" json:"maxSize,omitempty"`
}

func (m *Create) Reset()      { *m = Create{} }
func (*Create) ProtoMessage() {}
func (*Create) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{1}
}
func (m *Create) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Create) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Create.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Create) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Create.Merge(m, src)
}
func (m *Create) XXX_Size() int {
	return m.Size()
}
func (m *Create) XXX_DiscardUnknown() {
	xxx_messageInfo_Create.DiscardUnknown(m)
}

var xxx_messageInfo_Create proto.InternalMessageInfo

func (m *Create) GetMaxSize() int32 {
	if m != nil {
		return m.MaxSize
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloWorld)(nil), "messages.HelloWorld")
	proto.RegisterType((*Create)(nil), "messages.Create")
}

func init() { proto.RegisterFile("tree.proto", fileDescriptor_cb3889276909882a) }

var fileDescriptor_cb3889276909882a = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x29, 0x4a, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d,
	0x56, 0xe2, 0xe1, 0xe2, 0xf2, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0x52,
	0xe2, 0x62, 0x73, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x15, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0xac, 0x08,
	0xce, 0xac, 0x4a, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x9d, 0x4c, 0x2e, 0x3c,
	0x94, 0x63, 0xb8, 0xf1, 0x50, 0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b,
	0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x2d, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0xa0, 0x5b, 0x9f, 0x86, 0x00, 0x00, 0x00,
}

func (this *HelloWorld) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloWorld)
	if !ok {
		that2, ok := that.(HelloWorld)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *Create) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Create)
	if !ok {
		that2, ok := that.(Create)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxSize != that1.MaxSize {
		return false
	}
	return true
}
func (this *HelloWorld) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&messages.HelloWorld{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Create) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&messages.Create{")
	s = append(s, "MaxSize: "+fmt.Sprintf("%#v", this.MaxSize)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTree(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HelloWorld) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloWorld) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Create) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Create) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTree(dAtA, i, uint64(m.MaxSize))
	}
	return i, nil
}

func encodeVarintTree(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HelloWorld) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Create) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxSize != 0 {
		n += 1 + sovTree(uint64(m.MaxSize))
	}
	return n
}

func sovTree(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTree(x uint64) (n int) {
	return sovTree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HelloWorld) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloWorld{`,
		`}`,
	}, "")
	return s
}
func (this *Create) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Create{`,
		`MaxSize:` + fmt.Sprintf("%v", this.MaxSize) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTree(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HelloWorld) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: HelloWorld: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloWorld: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTree
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTree
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
func (m *Create) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Create: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Create: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSize", wireType)
			}
			m.MaxSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTree
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTree
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
func skipTree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTree
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
				return 0, ErrInvalidLengthTree
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTree
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTree
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTree(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTree
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTree = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTree   = fmt.Errorf("proto: integer overflow")
)
