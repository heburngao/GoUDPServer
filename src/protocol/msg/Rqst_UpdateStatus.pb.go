// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Rqst_UpdateStatus.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Rqst_UpdateStatus struct {
	Info             *StatusInfo `protobuf:"bytes,1,req,name=Info" json:"Info,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Rqst_UpdateStatus) Reset()                    { *m = Rqst_UpdateStatus{} }
func (m *Rqst_UpdateStatus) String() string            { return proto.CompactTextString(m) }
func (*Rqst_UpdateStatus) ProtoMessage()               {}
func (*Rqst_UpdateStatus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Rqst_UpdateStatus) GetInfo() *StatusInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Rqst_UpdateStatus)(nil), "msg.Rqst_UpdateStatus")
}

func init() { proto.RegisterFile("Rqst_UpdateStatus.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0f, 0x2a, 0x2c, 0x2e,
	0x89, 0x0f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x0d, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x12, 0x08, 0xc8, 0x49, 0xac, 0x4c, 0x2d,
	0xf2, 0xcc, 0x4b, 0xcb, 0x87, 0x08, 0x2b, 0x19, 0x71, 0x09, 0x62, 0xe8, 0x10, 0x92, 0xe5, 0x62,
	0x01, 0x29, 0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x36, 0xe2, 0xd7, 0xcb, 0x2d, 0x4e, 0xd7, 0x83,
	0x48, 0x81, 0x84, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0xef, 0xac, 0xfd, 0x64, 0x00, 0x00,
	0x00,
}