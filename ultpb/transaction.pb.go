// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package ultpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Every operation is related to a specific account and
// each transaction could contain one or more operations
type OpType int32

const (
	OpType_CREATE_ACCOUNT OpType = 0
	OpType_PAYMENT        OpType = 1
	OpType_PATH_PAYMENT   OpType = 2
	OpType_TRUST          OpType = 3
)

var OpType_name = map[int32]string{
	0: "CREATE_ACCOUNT",
	1: "PAYMENT",
	2: "PATH_PAYMENT",
	3: "TRUST",
}
var OpType_value = map[string]int32{
	"CREATE_ACCOUNT": 0,
	"PAYMENT":        1,
	"PATH_PAYMENT":   2,
	"TRUST":          3,
}

func (x OpType) String() string {
	return proto.EnumName(OpType_name, int32(x))
}
func (OpType) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// Operation
type Op struct {
	// operation type of this transaction
	OpType OpType `protobuf:"varint,1,opt,name=OpType,enum=ultpb.OpType" json:"OpType,omitempty"`
	// defacto operation
	//
	// Types that are valid to be assigned to Op:
	//	*Op_CreateAccount
	//	*Op_Payment
	//	*Op_PathPayment
	//	*Op_Trust
	Op isOp_Op `protobuf_oneof:"Op"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type isOp_Op interface {
	isOp_Op()
}

type Op_CreateAccount struct {
	CreateAccount *CreateAccountOp `protobuf:"bytes,2,opt,name=CreateAccount,oneof"`
}
type Op_Payment struct {
	Payment *PaymentOp `protobuf:"bytes,3,opt,name=Payment,oneof"`
}
type Op_PathPayment struct {
	PathPayment *PathPaymentOp `protobuf:"bytes,4,opt,name=PathPayment,oneof"`
}
type Op_Trust struct {
	Trust *TrustOp `protobuf:"bytes,5,opt,name=Trust,oneof"`
}

func (*Op_CreateAccount) isOp_Op() {}
func (*Op_Payment) isOp_Op()       {}
func (*Op_PathPayment) isOp_Op()   {}
func (*Op_Trust) isOp_Op()         {}

func (m *Op) GetOp() isOp_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *Op) GetOpType() OpType {
	if m != nil {
		return m.OpType
	}
	return OpType_CREATE_ACCOUNT
}

func (m *Op) GetCreateAccount() *CreateAccountOp {
	if x, ok := m.GetOp().(*Op_CreateAccount); ok {
		return x.CreateAccount
	}
	return nil
}

func (m *Op) GetPayment() *PaymentOp {
	if x, ok := m.GetOp().(*Op_Payment); ok {
		return x.Payment
	}
	return nil
}

func (m *Op) GetPathPayment() *PathPaymentOp {
	if x, ok := m.GetOp().(*Op_PathPayment); ok {
		return x.PathPayment
	}
	return nil
}

func (m *Op) GetTrust() *TrustOp {
	if x, ok := m.GetOp().(*Op_Trust); ok {
		return x.Trust
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Op) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Op_OneofMarshaler, _Op_OneofUnmarshaler, _Op_OneofSizer, []interface{}{
		(*Op_CreateAccount)(nil),
		(*Op_Payment)(nil),
		(*Op_PathPayment)(nil),
		(*Op_Trust)(nil),
	}
}

func _Op_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Op)
	// Op
	switch x := m.Op.(type) {
	case *Op_CreateAccount:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateAccount); err != nil {
			return err
		}
	case *Op_Payment:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Payment); err != nil {
			return err
		}
	case *Op_PathPayment:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PathPayment); err != nil {
			return err
		}
	case *Op_Trust:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Trust); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Op.Op has unexpected type %T", x)
	}
	return nil
}

func _Op_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Op)
	switch tag {
	case 2: // Op.CreateAccount
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateAccountOp)
		err := b.DecodeMessage(msg)
		m.Op = &Op_CreateAccount{msg}
		return true, err
	case 3: // Op.Payment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PaymentOp)
		err := b.DecodeMessage(msg)
		m.Op = &Op_Payment{msg}
		return true, err
	case 4: // Op.PathPayment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PathPaymentOp)
		err := b.DecodeMessage(msg)
		m.Op = &Op_PathPayment{msg}
		return true, err
	case 5: // Op.Trust
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TrustOp)
		err := b.DecodeMessage(msg)
		m.Op = &Op_Trust{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Op_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Op)
	// Op
	switch x := m.Op.(type) {
	case *Op_CreateAccount:
		s := proto.Size(x.CreateAccount)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Op_Payment:
		s := proto.Size(x.Payment)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Op_PathPayment:
		s := proto.Size(x.PathPayment)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Op_Trust:
		s := proto.Size(x.Trust)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Transaction
type Tx struct {
	// the source account for sending asset to the new account
	AccountID string `protobuf:"bytes,1,opt,name=AccountID" json:"AccountID,omitempty"`
	// fee to pay
	Fee int64 `protobuf:"varint,2,opt,name=Fee" json:"Fee,omitempty"`
	// extra note about the transaction
	Note string `protobuf:"bytes,3,opt,name=Note" json:"Note,omitempty"`
	// unique transaction sequence number
	SeqNum uint64 `protobuf:"varint,4,opt,name=SeqNum" json:"SeqNum,omitempty"`
	// list of operations
	OpList []*Op `protobuf:"bytes,5,rep,name=OpList" json:"OpList,omitempty"`
}

func (m *Tx) Reset()                    { *m = Tx{} }
func (m *Tx) String() string            { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()               {}
func (*Tx) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *Tx) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Tx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Tx) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Tx) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *Tx) GetOpList() []*Op {
	if m != nil {
		return m.OpList
	}
	return nil
}

// Transaction set for consensus
type TxSet struct {
	// hash of previous closed ledger header
	PrevLedgerHash string `protobuf:"bytes,1,opt,name=PrevLedgerHash" json:"PrevLedgerHash,omitempty"`
	// list of transaction hashes
	TxList []*Tx `protobuf:"bytes,2,rep,name=TxList" json:"TxList,omitempty"`
}

func (m *TxSet) Reset()                    { *m = TxSet{} }
func (m *TxSet) String() string            { return proto.CompactTextString(m) }
func (*TxSet) ProtoMessage()               {}
func (*TxSet) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *TxSet) GetPrevLedgerHash() string {
	if m != nil {
		return m.PrevLedgerHash
	}
	return ""
}

func (m *TxSet) GetTxList() []*Tx {
	if m != nil {
		return m.TxList
	}
	return nil
}

// Create a new account by sending asset to it
type CreateAccountOp struct {
	// destination account
	AccountID string `protobuf:"bytes,1,opt,name=AccountID" json:"AccountID,omitempty"`
	// initial balance in ULUs
	Balance int64 `protobuf:"varint,2,opt,name=Balance" json:"Balance,omitempty"`
}

func (m *CreateAccountOp) Reset()                    { *m = CreateAccountOp{} }
func (m *CreateAccountOp) String() string            { return proto.CompactTextString(m) }
func (*CreateAccountOp) ProtoMessage()               {}
func (*CreateAccountOp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *CreateAccountOp) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CreateAccountOp) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

// Peer to peer payment
type PaymentOp struct {
	// destination account of the payment
	AccountID string `protobuf:"bytes,1,opt,name=AccountID" json:"AccountID,omitempty"`
	// asset type of payment
	Asset *Asset `protobuf:"bytes,2,opt,name=Asset" json:"Asset,omitempty"`
	// amount of payment in specified asset type
	Amount int64 `protobuf:"varint,3,opt,name=Amount" json:"Amount,omitempty"`
}

func (m *PaymentOp) Reset()                    { *m = PaymentOp{} }
func (m *PaymentOp) String() string            { return proto.CompactTextString(m) }
func (*PaymentOp) ProtoMessage()               {}
func (*PaymentOp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *PaymentOp) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *PaymentOp) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *PaymentOp) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// Path payment
type PathPaymentOp struct {
	// source asset to send
	SrcAsset *Asset `protobuf:"bytes,1,opt,name=SrcAsset" json:"SrcAsset,omitempty"`
	// amount of source asset to send
	SrcAmount int64 `protobuf:"varint,2,opt,name=SrcAmount" json:"SrcAmount,omitempty"`
	// destination account of the path payment
	AccountID string `protobuf:"bytes,3,opt,name=AccountID" json:"AccountID,omitempty"`
	// destination asset type
	DstAsset *Asset `protobuf:"bytes,4,opt,name=DstAsset" json:"DstAsset,omitempty"`
	// payment asset path
	Path []*Asset `protobuf:"bytes,5,rep,name=Path" json:"Path,omitempty"`
}

func (m *PathPaymentOp) Reset()                    { *m = PathPaymentOp{} }
func (m *PathPaymentOp) String() string            { return proto.CompactTextString(m) }
func (*PathPaymentOp) ProtoMessage()               {}
func (*PathPaymentOp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *PathPaymentOp) GetSrcAsset() *Asset {
	if m != nil {
		return m.SrcAsset
	}
	return nil
}

func (m *PathPaymentOp) GetSrcAmount() int64 {
	if m != nil {
		return m.SrcAmount
	}
	return 0
}

func (m *PathPaymentOp) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *PathPaymentOp) GetDstAsset() *Asset {
	if m != nil {
		return m.DstAsset
	}
	return nil
}

func (m *PathPaymentOp) GetPath() []*Asset {
	if m != nil {
		return m.Path
	}
	return nil
}

// Trust management
type TrustOp struct {
	// asset for trust
	Asset *Asset `protobuf:"bytes,2,opt,name=Asset" json:"Asset,omitempty"`
	// limit of trust
	Limit int64 `protobuf:"varint,3,opt,name=Limit" json:"Limit,omitempty"`
}

func (m *TrustOp) Reset()                    { *m = TrustOp{} }
func (m *TrustOp) String() string            { return proto.CompactTextString(m) }
func (*TrustOp) ProtoMessage()               {}
func (*TrustOp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *TrustOp) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *TrustOp) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Trust authorization management
type AllowTrustOp struct {
	// accountID of trustor
	AccountID string `protobuf:"bytes,1,opt,name=AccountID" json:"AccountID,omitempty"`
	// asset of trust
	Asset *Asset `protobuf:"bytes,2,opt,name=Asset" json:"Asset,omitempty"`
	// authorization flag
	Authorized int32 `protobuf:"varint,3,opt,name=Authorized" json:"Authorized,omitempty"`
}

func (m *AllowTrustOp) Reset()                    { *m = AllowTrustOp{} }
func (m *AllowTrustOp) String() string            { return proto.CompactTextString(m) }
func (*AllowTrustOp) ProtoMessage()               {}
func (*AllowTrustOp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *AllowTrustOp) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *AllowTrustOp) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *AllowTrustOp) GetAuthorized() int32 {
	if m != nil {
		return m.Authorized
	}
	return 0
}

// Offer management
type OfferOp struct {
	// asset for selling
	SellingAsset *Asset `protobuf:"bytes,1,opt,name=SellingAsset" json:"SellingAsset,omitempty"`
	// asset for buying
	BuyingAsset *Asset `protobuf:"bytes,2,opt,name=BuyingAsset" json:"BuyingAsset,omitempty"`
	// amount of asset for selling
	Amount int64 `protobuf:"varint,3,opt,name=Amount" json:"Amount,omitempty"`
	// price in fractional format
	Price *Price `protobuf:"bytes,4,opt,name=Price" json:"Price,omitempty"`
	// ID of this offer
	OfferID string `protobuf:"bytes,5,opt,name=OfferID" json:"OfferID,omitempty"`
	// whether the offer is passive
	Passive int32 `protobuf:"varint,6,opt,name=Passive" json:"Passive,omitempty"`
}

func (m *OfferOp) Reset()                    { *m = OfferOp{} }
func (m *OfferOp) String() string            { return proto.CompactTextString(m) }
func (*OfferOp) ProtoMessage()               {}
func (*OfferOp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *OfferOp) GetSellingAsset() *Asset {
	if m != nil {
		return m.SellingAsset
	}
	return nil
}

func (m *OfferOp) GetBuyingAsset() *Asset {
	if m != nil {
		return m.BuyingAsset
	}
	return nil
}

func (m *OfferOp) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OfferOp) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *OfferOp) GetOfferID() string {
	if m != nil {
		return m.OfferID
	}
	return ""
}

func (m *OfferOp) GetPassive() int32 {
	if m != nil {
		return m.Passive
	}
	return 0
}

func init() {
	proto.RegisterType((*Op)(nil), "ultpb.Op")
	proto.RegisterType((*Tx)(nil), "ultpb.Tx")
	proto.RegisterType((*TxSet)(nil), "ultpb.TxSet")
	proto.RegisterType((*CreateAccountOp)(nil), "ultpb.CreateAccountOp")
	proto.RegisterType((*PaymentOp)(nil), "ultpb.PaymentOp")
	proto.RegisterType((*PathPaymentOp)(nil), "ultpb.PathPaymentOp")
	proto.RegisterType((*TrustOp)(nil), "ultpb.TrustOp")
	proto.RegisterType((*AllowTrustOp)(nil), "ultpb.AllowTrustOp")
	proto.RegisterType((*OfferOp)(nil), "ultpb.OfferOp")
	proto.RegisterEnum("ultpb.OpType", OpType_name, OpType_value)
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6b, 0xdb, 0x4c,
	0x10, 0x8d, 0x2c, 0xcb, 0x8e, 0x46, 0x8e, 0x3f, 0x7f, 0x4b, 0x08, 0x22, 0x94, 0x92, 0x0a, 0x5a,
	0x4c, 0x29, 0xa6, 0xa4, 0x97, 0x9e, 0x0a, 0x8a, 0x93, 0x90, 0x40, 0x6a, 0x89, 0xf5, 0xe6, 0xd0,
	0x53, 0x50, 0x94, 0x4d, 0x22, 0x90, 0x25, 0x55, 0x5a, 0xa5, 0x4e, 0x6f, 0xa5, 0xbf, 0xac, 0xbf,
	0xa3, 0x7f, 0xa6, 0xec, 0x68, 0xa5, 0x44, 0xc6, 0x21, 0x85, 0xde, 0x34, 0xef, 0xcd, 0x9b, 0x7d,
	0x33, 0x3b, 0x5a, 0xf8, 0x5f, 0xe4, 0x41, 0x52, 0x04, 0xa1, 0x88, 0xd2, 0x64, 0x92, 0xe5, 0xa9,
	0x48, 0x89, 0x51, 0xc6, 0x22, 0xbb, 0xdc, 0xb5, 0x82, 0xa2, 0xe0, 0xa2, 0xc2, 0x76, 0xad, 0xf4,
	0xfa, 0x9a, 0xe7, 0x55, 0xe0, 0xfc, 0xe8, 0x40, 0xc7, 0xcb, 0xc8, 0x6b, 0xe8, 0x79, 0x19, 0xbb,
	0xcf, 0xb8, 0xad, 0xed, 0x69, 0xe3, 0xe1, 0xfe, 0xd6, 0x04, 0x85, 0x93, 0x0a, 0xa4, 0x8a, 0x24,
	0x9f, 0x60, 0x6b, 0x9a, 0xf3, 0x40, 0x70, 0x37, 0x0c, 0xd3, 0x32, 0x11, 0x76, 0x67, 0x4f, 0x1b,
	0x5b, 0xfb, 0x3b, 0x2a, 0xbb, 0xc5, 0x79, 0xd9, 0xc9, 0x06, 0x6d, 0xa7, 0x93, 0x77, 0xd0, 0xf7,
	0x83, 0xfb, 0x05, 0x4f, 0x84, 0xad, 0xa3, 0x72, 0xa4, 0x94, 0x0a, 0x45, 0x4d, 0x9d, 0x42, 0x3e,
	0x82, 0xe5, 0x07, 0xe2, 0xb6, 0x56, 0x74, 0x51, 0xb1, 0xdd, 0x28, 0x1a, 0x06, 0x55, 0x8f, 0x53,
	0xc9, 0x1b, 0x30, 0x58, 0x5e, 0x16, 0xc2, 0x36, 0x50, 0x33, 0x54, 0x1a, 0xc4, 0x30, 0xbb, 0xa2,
	0x0f, 0xba, 0xb2, 0x79, 0xe7, 0xa7, 0x06, 0x1d, 0xb6, 0x24, 0x2f, 0xc0, 0x54, 0x3e, 0x4f, 0x0f,
	0x71, 0x0c, 0x26, 0x7d, 0x00, 0xc8, 0x08, 0xf4, 0x63, 0xce, 0xb1, 0x61, 0x9d, 0xca, 0x4f, 0x42,
	0xa0, 0x3b, 0x4b, 0x05, 0xc7, 0x4e, 0x4c, 0x8a, 0xdf, 0x64, 0x07, 0x7a, 0x73, 0xfe, 0x75, 0x56,
	0x2e, 0xd0, 0x6d, 0x97, 0xaa, 0x88, 0xbc, 0x92, 0xf3, 0x3d, 0x8b, 0xd0, 0x91, 0x3e, 0xb6, 0xf6,
	0xcd, 0x66, 0xbe, 0x54, 0x11, 0x0e, 0x05, 0x83, 0x2d, 0xe7, 0x5c, 0x9a, 0x1f, 0xfa, 0x39, 0xbf,
	0x3b, 0xe3, 0x57, 0x37, 0x3c, 0x3f, 0x09, 0x8a, 0x5b, 0x65, 0x66, 0x05, 0x95, 0x35, 0xd9, 0x12,
	0x6b, 0x76, 0x5a, 0x35, 0xd9, 0x92, 0x2a, 0xc2, 0x39, 0x85, 0xff, 0x56, 0xee, 0xe4, 0x99, 0x2e,
	0x6d, 0xe8, 0x1f, 0x04, 0x71, 0x90, 0x84, 0x75, 0xa7, 0x75, 0xe8, 0x70, 0x30, 0x9b, 0x71, 0x3f,
	0x53, 0xc4, 0x01, 0xc3, 0x95, 0xfb, 0xa6, 0xb6, 0x63, 0xa0, 0x7c, 0x21, 0x46, 0x2b, 0x4a, 0x0e,
	0xca, 0x5d, 0xe0, 0x0a, 0xe9, 0x78, 0x8e, 0x8a, 0x9c, 0x5f, 0x1a, 0x6c, 0xb5, 0xae, 0x96, 0x8c,
	0x61, 0x73, 0x9e, 0x87, 0x55, 0x41, 0x6d, 0x4d, 0xc1, 0x86, 0x95, 0xae, 0xe4, 0xf7, 0xa2, 0xd9,
	0x4c, 0x9d, 0x3e, 0x00, 0x6d, 0xcf, 0xfa, 0xaa, 0xe7, 0x31, 0x6c, 0x1e, 0x16, 0xa2, 0x3a, 0xa5,
	0xbb, 0xee, 0x94, 0x9a, 0x25, 0x7b, 0xd0, 0x95, 0x06, 0xd5, 0x45, 0xb6, 0xb3, 0x90, 0x71, 0xa6,
	0xd0, 0x57, 0x9b, 0xf6, 0x57, 0xa3, 0xd8, 0x06, 0xe3, 0x2c, 0x5a, 0x44, 0xf5, 0x24, 0xaa, 0xc0,
	0xc9, 0x60, 0xe0, 0xc6, 0x71, 0xfa, 0xad, 0xae, 0xf4, 0xef, 0x23, 0x7f, 0x09, 0xe0, 0x96, 0xe2,
	0x36, 0xcd, 0xa3, 0xef, 0xfc, 0x0a, 0x0f, 0x33, 0xe8, 0x23, 0xc4, 0xf9, 0xad, 0x41, 0xdf, 0x93,
	0x4f, 0x83, 0x97, 0x91, 0xf7, 0x30, 0x98, 0xf3, 0x38, 0x8e, 0x92, 0x9b, 0xa7, 0x07, 0xdf, 0xca,
	0x20, 0x13, 0xb0, 0x0e, 0xca, 0xfb, 0x46, 0xb0, 0xce, 0xc7, 0xe3, 0x84, 0xa7, 0x16, 0x40, 0x76,
	0xe2, 0xe7, 0x51, 0xc8, 0x57, 0x6e, 0x01, 0x31, 0x5a, 0x51, 0x72, 0x4b, 0xd1, 0xe8, 0xe9, 0x21,
	0xfe, 0xe0, 0x26, 0xad, 0x43, 0xc9, 0xf8, 0x41, 0x51, 0x44, 0x77, 0xdc, 0xee, 0x61, 0x83, 0x75,
	0xf8, 0xf6, 0xb8, 0x7e, 0xe1, 0x08, 0x81, 0xe1, 0x94, 0x1e, 0xb9, 0xec, 0xe8, 0xc2, 0x9d, 0x4e,
	0xbd, 0xf3, 0x19, 0x1b, 0x6d, 0x10, 0x0b, 0xfa, 0xbe, 0xfb, 0xe5, 0xf3, 0xd1, 0x8c, 0x8d, 0x34,
	0x32, 0x82, 0x81, 0xef, 0xb2, 0x93, 0x8b, 0x1a, 0xe9, 0x10, 0x13, 0x0c, 0x46, 0xcf, 0xe7, 0x6c,
	0xa4, 0x5f, 0xf6, 0xf0, 0xdd, 0xfc, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0x73, 0x0a, 0x3f, 0x96,
	0x6d, 0x05, 0x00, 0x00,
}
