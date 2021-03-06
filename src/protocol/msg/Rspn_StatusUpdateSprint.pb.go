// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Rspn_StatusUpdateSprint.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Rspn_StatusUpdateSprint struct {
	Player           []*PlayerInfo `protobuf:"bytes,1,rep,name=Player" json:"Player,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Rspn_StatusUpdateSprint) Reset()                    { *m = Rspn_StatusUpdateSprint{} }
func (m *Rspn_StatusUpdateSprint) String() string            { return proto.CompactTextString(m) }
func (*Rspn_StatusUpdateSprint) ProtoMessage()               {}
func (*Rspn_StatusUpdateSprint) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Rspn_StatusUpdateSprint) GetPlayer() []*PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

func init() {
	proto.RegisterType((*Rspn_StatusUpdateSprint)(nil), "msg.Rspn_StatusUpdateSprint")
}

func init() { proto.RegisterFile("Rspn_StatusUpdateSprint.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0d, 0x2a, 0x2e, 0xc8,
	0x8b, 0x0f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x0e, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x0d, 0x2e, 0x28,
	0xca, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x12,
	0x08, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x87, 0x08, 0x2b, 0x59, 0x71, 0x89,
	0xe3, 0xd0, 0x27, 0x24, 0xcf, 0xc5, 0x06, 0x51, 0x2e, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xc4,
	0xaf, 0x97, 0x5b, 0x9c, 0xae, 0x87, 0x30, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x9f, 0x6d,
	0xd0, 0x72, 0x00, 0x00, 0x00,
}
