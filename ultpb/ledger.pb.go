// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger.proto

package ultpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LedgerHeader struct {
	// version of pb message, each time we update the pb message definition, the
	// version number will incremented. The specific version number to use is
	// hard coded and changed each time we do some changes.
	Version uint32 `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	// hash of the previous ledger
	PrevLedgerHash string `protobuf:"bytes,2,opt,name=PrevLedgerHash" json:"PrevLedgerHash,omitempty"`
	// hash of the transaction list in this ledger
	TxListHash string `protobuf:"bytes,3,opt,name=TxListHash" json:"TxListHash,omitempty"`
	// maximun number of transactions allowed in a transaction list
	MaxTxListSize uint32 `protobuf:"varint,4,opt,name=MaxTxListSize" json:"MaxTxListSize,omitempty"`
	// sequence number of this ledger
	SeqNum uint64 `protobuf:"varint,5,opt,name=SeqNum" json:"SeqNum,omitempty"`
	// total number of tokens in existence
	TotalTokens uint64 `protobuf:"varint,6,opt,name=TotalTokens" json:"TotalTokens,omitempty"`
	// base fee per operation
	BaseFee uint64 `protobuf:"varint,7,opt,name=BaseFee" json:"BaseFee,omitempty"`
	// base reserver for an account
	BaseReserve uint64 `protobuf:"varint,8,opt,name=BaseReserve" json:"BaseReserve,omitempty"`
	// close time
	CloseTime int64 `protobuf:"varint,9,opt,name=CloseTime" json:"CloseTime,omitempty"`
}

func (m *LedgerHeader) Reset()                    { *m = LedgerHeader{} }
func (m *LedgerHeader) String() string            { return proto.CompactTextString(m) }
func (*LedgerHeader) ProtoMessage()               {}
func (*LedgerHeader) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *LedgerHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *LedgerHeader) GetPrevLedgerHash() string {
	if m != nil {
		return m.PrevLedgerHash
	}
	return ""
}

func (m *LedgerHeader) GetTxListHash() string {
	if m != nil {
		return m.TxListHash
	}
	return ""
}

func (m *LedgerHeader) GetMaxTxListSize() uint32 {
	if m != nil {
		return m.MaxTxListSize
	}
	return 0
}

func (m *LedgerHeader) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *LedgerHeader) GetTotalTokens() uint64 {
	if m != nil {
		return m.TotalTokens
	}
	return 0
}

func (m *LedgerHeader) GetBaseFee() uint64 {
	if m != nil {
		return m.BaseFee
	}
	return 0
}

func (m *LedgerHeader) GetBaseReserve() uint64 {
	if m != nil {
		return m.BaseReserve
	}
	return 0
}

func (m *LedgerHeader) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func init() {
	proto.RegisterType((*LedgerHeader)(nil), "ultpb.LedgerHeader")
}

func init() { proto.RegisterFile("ledger.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x70, 0xd2, 0x3f, 0x5b, 0x77, 0x6c, 0x3d, 0xcc, 0x41, 0xe6, 0x20, 0x12, 0x44, 0x64,
	0x4f, 0x5e, 0x7c, 0x03, 0x05, 0xf1, 0x50, 0x45, 0xd2, 0xc5, 0x7b, 0x4a, 0x3f, 0x74, 0x71, 0xdb,
	0xd4, 0x24, 0x2d, 0xc5, 0x47, 0xf2, 0x29, 0x65, 0xd3, 0x2e, 0xae, 0xbd, 0xe5, 0xfb, 0xcd, 0x37,
	0x84, 0x84, 0xc6, 0x35, 0x16, 0xef, 0xf0, 0xb7, 0x6b, 0xef, 0xa2, 0xe3, 0xe1, 0xa6, 0x8e, 0xeb,
	0xf9, 0xd5, 0x4f, 0x8f, 0xc6, 0xd3, 0xe4, 0x4f, 0xb0, 0x0b, 0x78, 0x16, 0x1a, 0xbd, 0xc1, 0x87,
	0xca, 0xad, 0x44, 0x69, 0x55, 0x4c, 0x4c, 0x1b, 0xf9, 0x86, 0xce, 0x5e, 0x3d, 0xb6, 0x87, 0xb6,
	0x0d, 0x1f, 0xd2, 0xd3, 0xaa, 0xc8, 0xcd, 0x91, 0xf2, 0x25, 0x51, 0xb9, 0x9b, 0x56, 0x21, 0xa6,
	0x4e, 0x3f, 0x75, 0x3a, 0xc2, 0xd7, 0x34, 0x79, 0xb6, 0xbb, 0x3d, 0xcc, 0xaa, 0x6f, 0xc8, 0x20,
	0xdd, 0xf3, 0x1f, 0xf9, 0x9c, 0xb2, 0x19, 0xbe, 0x5e, 0x36, 0x4b, 0x19, 0x6a, 0x55, 0x0c, 0xcc,
	0x21, 0xb1, 0xa6, 0xd3, 0xd2, 0x45, 0x5b, 0x97, 0xee, 0x13, 0xab, 0x20, 0x59, 0x1a, 0x76, 0xa9,
	0x79, 0xc1, 0xbd, 0x0d, 0x78, 0x04, 0x64, 0x94, 0xa6, 0x6d, 0x6c, 0x76, 0x9b, 0xa3, 0x41, 0x80,
	0xdf, 0x42, 0x4e, 0xf6, 0xbb, 0x1d, 0xe2, 0x0b, 0xca, 0x1f, 0x6a, 0x17, 0x50, 0x56, 0x4b, 0x48,
	0xae, 0x55, 0xd1, 0x37, 0x7f, 0x30, 0xcf, 0xd2, 0xd7, 0xdd, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0xf8, 0x5b, 0x0a, 0x4a, 0x01, 0x00, 0x00,
}
