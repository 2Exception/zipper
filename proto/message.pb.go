// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProtoID int32

const (
	ProtoID_NA              ProtoID = 0
	ProtoID_SyncWorker      ProtoID = 1
	ProtoID_ConsensusWorker ProtoID = 2
)

var ProtoID_name = map[int32]string{
	0: "NA",
	1: "SyncWorker",
	2: "ConsensusWorker",
}
var ProtoID_value = map[string]int32{
	"NA":              0,
	"SyncWorker":      1,
	"ConsensusWorker": 2,
}

func (x ProtoID) String() string {
	return proto1.EnumName(ProtoID_name, int32(x))
}
func (ProtoID) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type MsgType int32

const (
	MsgType_BC_GetBlocksMsg     MsgType = 0
	MsgType_BC_GetInvMsg        MsgType = 1
	MsgType_BC_GetDataMsg       MsgType = 2
	MsgType_BC_OnBlockMsg       MsgType = 3
	MsgType_BC_OnTransactionMsg MsgType = 4
	MsgType_BC_OnStatusMSg      MsgType = 5
	MsgType_BC_OnConsensusMSg   MsgType = 6
)

var MsgType_name = map[int32]string{
	0: "BC_GetBlocksMsg",
	1: "BC_GetInvMsg",
	2: "BC_GetDataMsg",
	3: "BC_OnBlockMsg",
	4: "BC_OnTransactionMsg",
	5: "BC_OnStatusMSg",
	6: "BC_OnConsensusMSg",
}
var MsgType_value = map[string]int32{
	"BC_GetBlocksMsg":     0,
	"BC_GetInvMsg":        1,
	"BC_GetDataMsg":       2,
	"BC_OnBlockMsg":       3,
	"BC_OnTransactionMsg": 4,
	"BC_OnStatusMSg":      5,
	"BC_OnConsensusMSg":   6,
}

func (x MsgType) String() string {
	return proto1.EnumName(MsgType_name, int32(x))
}
func (MsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type InvType int32

const (
	InvType_block       InvType = 0
	InvType_transaction InvType = 1
)

var InvType_name = map[int32]string{
	0: "block",
	1: "transaction",
}
var InvType_value = map[string]int32{
	"block":       0,
	"transaction": 1,
}

func (x InvType) String() string {
	return proto1.EnumName(InvType_name, int32(x))
}
func (InvType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type GetBlocksMsg struct {
	Version       uint32   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LocatorHashes []string `protobuf:"bytes,2,rep,name=locatorHashes" json:"locatorHashes,omitempty"`
	HashStop      string   `protobuf:"bytes,3,opt,name=hashStop" json:"hashStop,omitempty"`
}

func (m *GetBlocksMsg) Reset()                    { *m = GetBlocksMsg{} }
func (m *GetBlocksMsg) String() string            { return proto1.CompactTextString(m) }
func (*GetBlocksMsg) ProtoMessage()               {}
func (*GetBlocksMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetBlocksMsg) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetBlocksMsg) GetLocatorHashes() []string {
	if m != nil {
		return m.LocatorHashes
	}
	return nil
}

func (m *GetBlocksMsg) GetHashStop() string {
	if m != nil {
		return m.HashStop
	}
	return ""
}

type StatusMsg struct {
	Version     uint32 `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	StartHeight uint32 `protobuf:"varint,2,opt,name=StartHeight" json:"StartHeight,omitempty"`
}

func (m *StatusMsg) Reset()                    { *m = StatusMsg{} }
func (m *StatusMsg) String() string            { return proto1.CompactTextString(m) }
func (*StatusMsg) ProtoMessage()               {}
func (*StatusMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *StatusMsg) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StatusMsg) GetStartHeight() uint32 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

type GetInvMsg struct {
	Type  InvType  `protobuf:"varint,1,opt,name=type,enum=proto.InvType" json:"type,omitempty"`
	Hashs []string `protobuf:"bytes,2,rep,name=hashs" json:"hashs,omitempty"`
}

func (m *GetInvMsg) Reset()                    { *m = GetInvMsg{} }
func (m *GetInvMsg) String() string            { return proto1.CompactTextString(m) }
func (*GetInvMsg) ProtoMessage()               {}
func (*GetInvMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetInvMsg) GetType() InvType {
	if m != nil {
		return m.Type
	}
	return InvType_block
}

func (m *GetInvMsg) GetHashs() []string {
	if m != nil {
		return m.Hashs
	}
	return nil
}

type OnBlockMsg struct {
	Block *Block `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *OnBlockMsg) Reset()                    { *m = OnBlockMsg{} }
func (m *OnBlockMsg) String() string            { return proto1.CompactTextString(m) }
func (*OnBlockMsg) ProtoMessage()               {}
func (*OnBlockMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *OnBlockMsg) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type OnTransactionMsg struct {
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *OnTransactionMsg) Reset()                    { *m = OnTransactionMsg{} }
func (m *OnTransactionMsg) String() string            { return proto1.CompactTextString(m) }
func (*OnTransactionMsg) ProtoMessage()               {}
func (*OnTransactionMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *OnTransactionMsg) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type GetDataMsg struct {
	InvList []*GetInvMsg `protobuf:"bytes,1,rep,name=InvList" json:"InvList,omitempty"`
}

func (m *GetDataMsg) Reset()                    { *m = GetDataMsg{} }
func (m *GetDataMsg) String() string            { return proto1.CompactTextString(m) }
func (*GetDataMsg) ProtoMessage()               {}
func (*GetDataMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetDataMsg) GetInvList() []*GetInvMsg {
	if m != nil {
		return m.InvList
	}
	return nil
}

func init() {
	proto1.RegisterType((*GetBlocksMsg)(nil), "proto.GetBlocksMsg")
	proto1.RegisterType((*StatusMsg)(nil), "proto.StatusMsg")
	proto1.RegisterType((*GetInvMsg)(nil), "proto.GetInvMsg")
	proto1.RegisterType((*OnBlockMsg)(nil), "proto.OnBlockMsg")
	proto1.RegisterType((*OnTransactionMsg)(nil), "proto.OnTransactionMsg")
	proto1.RegisterType((*GetDataMsg)(nil), "proto.GetDataMsg")
	proto1.RegisterEnum("proto.ProtoID", ProtoID_name, ProtoID_value)
	proto1.RegisterEnum("proto.MsgType", MsgType_name, MsgType_value)
	proto1.RegisterEnum("proto.InvType", InvType_name, InvType_value)
}

func init() { proto1.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0xb1, 0x09, 0x50, 0x8f, 0x81, 0x2c, 0x93, 0x56, 0xb5, 0x38, 0x59, 0x56, 0x2b, 0x21,
	0x0e, 0x51, 0x45, 0xab, 0xaa, 0xd7, 0x26, 0xa9, 0x00, 0xa9, 0x69, 0x2a, 0x3b, 0x6a, 0x8f, 0xd5,
	0x06, 0xad, 0x0c, 0x4d, 0xba, 0x8b, 0xbc, 0x13, 0x24, 0xbe, 0x4a, 0x3f, 0x6d, 0xb4, 0x7f, 0x00,
	0xfb, 0xb4, 0x9a, 0xdf, 0x9b, 0x79, 0x3b, 0xfb, 0x16, 0x06, 0xff, 0x84, 0xd6, 0xbc, 0x14, 0x97,
	0xdb, 0x4a, 0x91, 0xc2, 0x8e, 0x3d, 0xc6, 0xf1, 0xc3, 0x93, 0x5a, 0x3d, 0x3a, 0x36, 0x1e, 0x51,
	0xc5, 0xa5, 0xe6, 0x2b, 0xda, 0x28, 0xe9, 0x50, 0xf6, 0x17, 0xfa, 0x73, 0x41, 0x57, 0xa6, 0x49,
	0xdf, 0xea, 0x12, 0x13, 0xe8, 0xed, 0x44, 0xa5, 0x37, 0x4a, 0x26, 0x41, 0x1a, 0x4c, 0x06, 0xf9,
	0xa1, 0xc4, 0x77, 0x30, 0x78, 0x52, 0x2b, 0x4e, 0xaa, 0x5a, 0x70, 0xbd, 0x16, 0x3a, 0x09, 0xd3,
	0xf6, 0x24, 0xca, 0x9b, 0x10, 0xc7, 0xf0, 0x6a, 0xcd, 0xf5, 0xba, 0x20, 0xb5, 0x4d, 0xda, 0x69,
	0x30, 0x89, 0xf2, 0x63, 0x9d, 0xcd, 0x21, 0x2a, 0x88, 0xd3, 0xf3, 0xe1, 0xa2, 0x5f, 0xcd, 0x8b,
	0x7c, 0x89, 0x29, 0xc4, 0x05, 0xf1, 0x8a, 0x16, 0x62, 0x53, 0xae, 0x29, 0x09, 0xad, 0x5a, 0x47,
	0xd9, 0x37, 0x88, 0xe6, 0x82, 0x96, 0x72, 0x67, 0x8c, 0x32, 0x38, 0xa3, 0xfd, 0x56, 0x58, 0x97,
	0xe1, 0x6c, 0xe8, 0xde, 0x75, 0xb9, 0x94, 0xbb, 0xfb, 0xfd, 0x56, 0xe4, 0x56, 0xc3, 0xd7, 0xd0,
	0x31, 0x5b, 0x1c, 0x76, 0x76, 0x45, 0xf6, 0x01, 0xe0, 0x4e, 0xda, 0xa7, 0x3b, 0x9f, 0x8e, 0xcd,
	0xca, 0x1a, 0xc5, 0xb3, 0xbe, 0x37, 0xb2, 0x7a, 0xee, 0xa4, 0x6c, 0x01, 0xec, 0x4e, 0xde, 0x9f,
	0x42, 0x34, 0x73, 0x9f, 0x20, 0xae, 0xc5, 0xea, 0xa7, 0xd1, 0x4f, 0xd7, 0x7a, 0xf3, 0x7a, 0x5b,
	0xf6, 0x05, 0x60, 0x2e, 0xe8, 0x86, 0x13, 0x37, 0x1e, 0x53, 0xe8, 0x2d, 0xe5, 0xee, 0xfb, 0x46,
	0x53, 0x12, 0xa4, 0xed, 0x49, 0x3c, 0x63, 0x7e, 0xfe, 0xf8, 0xcc, 0xfc, 0xd0, 0x30, 0xfd, 0x0c,
	0xbd, 0x9f, 0x46, 0x5b, 0xde, 0x60, 0x17, 0xc2, 0x1f, 0x5f, 0x59, 0x0b, 0x87, 0x00, 0xc5, 0x5e,
	0xae, 0x7e, 0xab, 0xea, 0x51, 0x54, 0x2c, 0xc0, 0x0b, 0x38, 0xbf, 0x56, 0x52, 0x0b, 0xa9, 0x9f,
	0xb5, 0x87, 0xe1, 0xf4, 0x7f, 0x00, 0xbd, 0x5b, 0x5d, 0x9a, 0x54, 0x4c, 0xc3, 0xd5, 0xf5, 0x9f,
	0xfa, 0xc7, 0xb3, 0x16, 0x32, 0xe8, 0x3b, 0xe8, 0x6e, 0x64, 0x01, 0x8e, 0x60, 0xe0, 0x88, 0xdf,
	0x93, 0x85, 0x1e, 0x9d, 0x62, 0x63, 0x6d, 0x7c, 0x0b, 0x17, 0x16, 0x35, 0x73, 0x61, 0x67, 0x88,
	0x30, 0xb4, 0x82, 0xff, 0xf4, 0xa2, 0x64, 0x1d, 0x7c, 0x03, 0x23, 0xcb, 0x8e, 0xfb, 0x19, 0xdc,
	0x9d, 0xbe, 0xb7, 0x01, 0xd8, 0xdd, 0x22, 0xff, 0x0f, 0xac, 0x85, 0xe7, 0x8d, 0x68, 0x59, 0xf0,
	0xd0, 0xb5, 0xa9, 0x7c, 0x7c, 0x09, 0x00, 0x00, 0xff, 0xff, 0x05, 0xde, 0x57, 0x38, 0xec, 0x02,
	0x00, 0x00,
}
