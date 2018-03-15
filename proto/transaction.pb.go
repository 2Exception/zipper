// Code generated by protoc-gen-go.
// source: transaction.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TransactionType int32

const (
	TransactionType_Atomic          TransactionType = 0
	TransactionType_AcrossChain     TransactionType = 1
	TransactionType_Merged          TransactionType = 2
	TransactionType_Backfront       TransactionType = 3
	TransactionType_Distribut       TransactionType = 4
	TransactionType_Issue           TransactionType = 5
	TransactionType_IssueUpdate     TransactionType = 6
	TransactionType_JSContractInit  TransactionType = 7
	TransactionType_LuaContractInit TransactionType = 8
	TransactionType_ContractInvoke  TransactionType = 9
	TransactionType_ContractQuery   TransactionType = 10
	TransactionType_Security        TransactionType = 11
)

var TransactionType_name = map[int32]string{
	0:  "Atomic",
	1:  "AcrossChain",
	2:  "Merged",
	3:  "Backfront",
	4:  "Distribut",
	5:  "Issue",
	6:  "IssueUpdate",
	7:  "JSContractInit",
	8:  "LuaContractInit",
	9:  "ContractInvoke",
	10: "ContractQuery",
	11: "Security",
}
var TransactionType_value = map[string]int32{
	"Atomic":          0,
	"AcrossChain":     1,
	"Merged":          2,
	"Backfront":       3,
	"Distribut":       4,
	"Issue":           5,
	"IssueUpdate":     6,
	"JSContractInit":  7,
	"LuaContractInit": 8,
	"ContractInvoke":  9,
	"ContractQuery":   10,
	"Security":        11,
}

func (x TransactionType) String() string {
	return proto1.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type ContractSpec struct {
	Addr   []byte   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Code   []byte   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Params []string `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
}

func (m *ContractSpec) Reset()                    { *m = ContractSpec{} }
func (m *ContractSpec) String() string            { return proto1.CompactTextString(m) }
func (*ContractSpec) ProtoMessage()               {}
func (*ContractSpec) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ContractSpec) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *ContractSpec) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ContractSpec) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type TxHeader struct {
	FromChain  []byte          `protobuf:"bytes,1,opt,name=fromChain,proto3" json:"fromChain,omitempty"`
	ToChain    []byte          `protobuf:"bytes,2,opt,name=toChain,proto3" json:"toChain,omitempty"`
	Type       TransactionType `protobuf:"varint,3,opt,name=type,enum=proto.TransactionType" json:"type,omitempty"`
	Nonce      uint32          `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
	Sender     string          `protobuf:"bytes,5,opt,name=sender" json:"sender,omitempty"`
	Recipient  string          `protobuf:"bytes,6,opt,name=recipient" json:"recipient,omitempty"`
	AssetID    uint32          `protobuf:"varint,7,opt,name=assetID" json:"assetID,omitempty"`
	Amount     int64           `protobuf:"varint,8,opt,name=amount" json:"amount,omitempty"`
	Fee        int64           `protobuf:"varint,9,opt,name=fee" json:"fee,omitempty"`
	Signature  []byte          `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
	CreateTime uint32          `protobuf:"varint,11,opt,name=createTime" json:"createTime,omitempty"`
}

func (m *TxHeader) Reset()                    { *m = TxHeader{} }
func (m *TxHeader) String() string            { return proto1.CompactTextString(m) }
func (*TxHeader) ProtoMessage()               {}
func (*TxHeader) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TxHeader) GetFromChain() []byte {
	if m != nil {
		return m.FromChain
	}
	return nil
}

func (m *TxHeader) GetToChain() []byte {
	if m != nil {
		return m.ToChain
	}
	return nil
}

func (m *TxHeader) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_Atomic
}

func (m *TxHeader) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TxHeader) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *TxHeader) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *TxHeader) GetAssetID() uint32 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *TxHeader) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxHeader) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TxHeader) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TxHeader) GetCreateTime() uint32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type Transaction struct {
	Header       *TxHeader     `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Payload      []byte        `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Meta         []byte        `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	ContractSpec *ContractSpec `protobuf:"bytes,4,opt,name=contractSpec" json:"contractSpec,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto1.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Transaction) GetHeader() *TxHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Transaction) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Transaction) GetContractSpec() *ContractSpec {
	if m != nil {
		return m.ContractSpec
	}
	return nil
}

func init() {
	proto1.RegisterType((*ContractSpec)(nil), "proto.ContractSpec")
	proto1.RegisterType((*TxHeader)(nil), "proto.TxHeader")
	proto1.RegisterType((*Transaction)(nil), "proto.Transaction")
	proto1.RegisterEnum("proto.TransactionType", TransactionType_name, TransactionType_value)
}

func init() { proto1.RegisterFile("transaction.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xc1, 0x72, 0xd3, 0x30,
	0x10, 0xc5, 0x71, 0xe2, 0xc4, 0xeb, 0xa4, 0x51, 0x55, 0xa6, 0xa3, 0x03, 0xc3, 0x78, 0x72, 0x21,
	0xd3, 0x43, 0x0f, 0xe1, 0xc0, 0xb9, 0xb4, 0x07, 0xc2, 0x00, 0x33, 0x38, 0xe1, 0x03, 0x54, 0x79,
	0x53, 0x34, 0xc5, 0x92, 0x47, 0x92, 0x19, 0xfc, 0x31, 0xfc, 0x09, 0x5f, 0xc1, 0x17, 0x31, 0x92,
	0x6d, 0x92, 0xf6, 0xe4, 0x7d, 0x6f, 0x77, 0x9e, 0xde, 0xbe, 0x35, 0x9c, 0x3b, 0xc3, 0x95, 0xe5,
	0xc2, 0x49, 0xad, 0xae, 0x6b, 0xa3, 0x9d, 0xa6, 0x93, 0xf0, 0x59, 0x7d, 0x81, 0xf9, 0xad, 0x56,
	0xce, 0x70, 0xe1, 0x76, 0x35, 0x0a, 0x4a, 0x61, 0xcc, 0xcb, 0xd2, 0xb0, 0x28, 0x8f, 0xd6, 0xf3,
	0x22, 0xd4, 0x9e, 0x13, 0xba, 0x44, 0x36, 0xea, 0x38, 0x5f, 0xd3, 0x4b, 0x48, 0x6a, 0x6e, 0x78,
	0x65, 0x59, 0x9c, 0xc7, 0xeb, 0xb4, 0xe8, 0xd1, 0xea, 0xcf, 0x08, 0x66, 0xfb, 0x5f, 0x1f, 0x90,
	0x97, 0x68, 0xe8, 0x2b, 0x48, 0x0f, 0x46, 0x57, 0xb7, 0xdf, 0xb9, 0x54, 0xbd, 0xe2, 0x91, 0xa0,
	0x0c, 0xa6, 0x4e, 0x77, 0xbd, 0x4e, 0x79, 0x80, 0xf4, 0x0a, 0xc6, 0xae, 0xad, 0x91, 0xc5, 0x79,
	0xb4, 0x3e, 0xdb, 0x5c, 0x76, 0x8e, 0xaf, 0xf7, 0xc7, 0x1d, 0xf6, 0x6d, 0x8d, 0x45, 0x98, 0xa1,
	0x2f, 0x61, 0xa2, 0xb4, 0x12, 0xc8, 0xc6, 0x79, 0xb4, 0x5e, 0x14, 0x1d, 0xf0, 0xf6, 0x2c, 0xaa,
	0x12, 0x0d, 0x9b, 0xe4, 0x91, 0xb7, 0xd7, 0x21, 0xef, 0xc8, 0xa0, 0x90, 0xb5, 0x44, 0xe5, 0x58,
	0x12, 0x5a, 0x47, 0xc2, 0x3b, 0xe2, 0xd6, 0xa2, 0xdb, 0xde, 0xb1, 0x69, 0x50, 0x1b, 0xa0, 0xd7,
	0xe3, 0x95, 0x6e, 0x94, 0x63, 0xb3, 0x3c, 0x5a, 0xc7, 0x45, 0x8f, 0x28, 0x81, 0xf8, 0x80, 0xc8,
	0xd2, 0x40, 0xfa, 0xd2, 0xbf, 0x60, 0xe5, 0x83, 0xe2, 0xae, 0x31, 0xc8, 0xa0, 0xdb, 0xf9, 0x3f,
	0x41, 0x5f, 0x03, 0x08, 0x83, 0xdc, 0xe1, 0x5e, 0x56, 0xc8, 0xb2, 0xf0, 0xc8, 0x09, 0xb3, 0xfa,
	0x1d, 0x41, 0x76, 0xb2, 0x27, 0x7d, 0x03, 0x49, 0x97, 0x65, 0x88, 0x2f, 0xdb, 0x2c, 0x87, 0x2c,
	0xfa, 0x88, 0x8b, 0xbe, 0xed, 0xad, 0xd7, 0xbc, 0xfd, 0xa1, 0x79, 0x39, 0x84, 0xd9, 0x43, 0x7f,
	0xbd, 0x0a, 0x1d, 0x0f, 0x61, 0xce, 0x8b, 0x50, 0xd3, 0x77, 0x30, 0x17, 0x27, 0x57, 0x0f, 0xd9,
	0x65, 0x9b, 0x8b, 0x5e, 0xfc, 0xf4, 0x87, 0x28, 0x9e, 0x0c, 0x5e, 0xfd, 0x8d, 0x60, 0xf9, 0xec,
	0x0e, 0x14, 0x20, 0xb9, 0x71, 0xba, 0x92, 0x82, 0xbc, 0xa0, 0x4b, 0xc8, 0x6e, 0x84, 0xd1, 0xd6,
	0x86, 0x43, 0x92, 0xc8, 0x37, 0x3f, 0xa3, 0x79, 0xc0, 0x92, 0x8c, 0xe8, 0x02, 0xd2, 0xf7, 0x5c,
	0x3c, 0x1e, 0x8c, 0x56, 0x8e, 0xc4, 0x1e, 0xde, 0x49, 0xeb, 0x8c, 0xbc, 0x6f, 0x1c, 0x19, 0xd3,
	0x14, 0x26, 0x5b, 0x6b, 0x1b, 0x24, 0x13, 0xaf, 0x12, 0xca, 0x6f, 0x75, 0xc9, 0x1d, 0x92, 0x84,
	0x52, 0x38, 0xfb, 0xb8, 0x1b, 0x6c, 0x6d, 0x95, 0x74, 0x64, 0x4a, 0x2f, 0x60, 0xf9, 0xa9, 0xe1,
	0x4f, 0xc8, 0x99, 0x1f, 0x3c, 0x32, 0x3f, 0xf5, 0x23, 0x92, 0x94, 0x9e, 0xc3, 0x62, 0xe0, 0xbe,
	0x36, 0x68, 0x5a, 0x02, 0x74, 0x0e, 0xb3, 0x1d, 0x8a, 0xc6, 0x48, 0xd7, 0x92, 0xec, 0x3e, 0x09,
	0x6b, 0xbf, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xab, 0x38, 0x9b, 0x84, 0x26, 0x03, 0x00, 0x00,
}
