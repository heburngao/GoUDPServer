// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Rspn_HeartBeating.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Rspn_HeartBeating struct {
	Status           *int32        `protobuf:"varint,1,req,name=Status,def=1" json:"Status,omitempty"`
	PlayerList       []*PlayerInfo `protobuf:"bytes,2,rep,name=PlayerList" json:"PlayerList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Rspn_HeartBeating) Reset()                    { *m = Rspn_HeartBeating{} }
func (m *Rspn_HeartBeating) String() string            { return proto.CompactTextString(m) }
func (*Rspn_HeartBeating) ProtoMessage()               {}
func (*Rspn_HeartBeating) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

const Default_Rspn_HeartBeating_Status int32 = 1

func (m *Rspn_HeartBeating) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_Rspn_HeartBeating_Status
}

func (m *Rspn_HeartBeating) GetPlayerList() []*PlayerInfo {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

func init() {
	proto.RegisterType((*Rspn_HeartBeating)(nil), "msg.Rspn_HeartBeating")
}

func init() { proto.RegisterFile("Rspn_HeartBeating.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0f, 0x2a, 0x2e, 0xc8,
	0x8b, 0xf7, 0x48, 0x4d, 0x2c, 0x2a, 0x71, 0x4a, 0x4d, 0x2c, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x12, 0x08, 0xc8, 0x49, 0xac, 0x4c, 0x2d,
	0xf2, 0xcc, 0x4b, 0xcb, 0x87, 0x08, 0x2b, 0x79, 0x73, 0x09, 0x62, 0xe8, 0x10, 0x12, 0xe4, 0x62,
	0x0b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd2, 0x60, 0xb5, 0x62, 0x34, 0x14,
	0x52, 0xe6, 0xe2, 0x82, 0xe8, 0xf5, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36,
	0xe2, 0xd7, 0xcb, 0x2d, 0x4e, 0xd7, 0x43, 0x18, 0x09, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x7b,
	0xea, 0xe2, 0x7d, 0x00, 0x00, 0x00,
}