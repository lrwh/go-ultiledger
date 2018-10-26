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
)

var OpType_name = map[int32]string{
	0: "CREATE_ACCOUNT",
	1: "PAYMENT",
	2: "PATH_PAYMENT",
}
var OpType_value = map[string]int32{
	"CREATE_ACCOUNT": 0,
	"PAYMENT":        1,
	"PATH_PAYMENT":   2,
}

func (x OpType) String() string {
	return proto.EnumName(OpType_name, int32(x))
}
func (OpType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

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
	Op isOp_Op `protobuf_oneof:"Op"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

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

func (*Op_CreateAccount) isOp_Op() {}
func (*Op_Payment) isOp_Op()       {}
func (*Op_PathPayment) isOp_Op()   {}

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

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Op) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Op_OneofMarshaler, _Op_OneofUnmarshaler, _Op_OneofSizer, []interface{}{
		(*Op_CreateAccount)(nil),
		(*Op_Payment)(nil),
		(*Op_PathPayment)(nil),
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
	Fee uint64 `protobuf:"varint,2,opt,name=Fee" json:"Fee,omitempty"`
	// extra note about the transaction
	Note string `protobuf:"bytes,3,opt,name=Note" json:"Note,omitempty"`
	// unique transaction sequence number
	SequenceNumber uint64 `protobuf:"varint,4,opt,name=SequenceNumber" json:"SequenceNumber,omitempty"`
	// list of operations
	OpList []*Op `protobuf:"bytes,5,rep,name=OpList" json:"OpList,omitempty"`
}

func (m *Tx) Reset()                    { *m = Tx{} }
func (m *Tx) String() string            { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()               {}
func (*Tx) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Tx) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Tx) GetFee() uint64 {
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

func (m *Tx) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
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
func (*TxSet) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

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
	Balance uint64 `protobuf:"varint,2,opt,name=Balance" json:"Balance,omitempty"`
}

func (m *CreateAccountOp) Reset()                    { *m = CreateAccountOp{} }
func (m *CreateAccountOp) String() string            { return proto.CompactTextString(m) }
func (*CreateAccountOp) ProtoMessage()               {}
func (*CreateAccountOp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *CreateAccountOp) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CreateAccountOp) GetBalance() uint64 {
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
func (*PaymentOp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

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
	SrcAmount uint64 `protobuf:"varint,2,opt,name=SrcAmount" json:"SrcAmount,omitempty"`
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
func (*PathPaymentOp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *PathPaymentOp) GetSrcAsset() *Asset {
	if m != nil {
		return m.SrcAsset
	}
	return nil
}

func (m *PathPaymentOp) GetSrcAmount() uint64 {
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
	Limit uint64 `protobuf:"varint,3,opt,name=Limit" json:"Limit,omitempty"`
}

func (m *TrustOp) Reset()                    { *m = TrustOp{} }
func (m *TrustOp) String() string            { return proto.CompactTextString(m) }
func (*TrustOp) ProtoMessage()               {}
func (*TrustOp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *TrustOp) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *TrustOp) GetLimit() uint64 {
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
	Authorized uint32 `protobuf:"varint,3,opt,name=Authorized" json:"Authorized,omitempty"`
}

func (m *AllowTrustOp) Reset()                    { *m = AllowTrustOp{} }
func (m *AllowTrustOp) String() string            { return proto.CompactTextString(m) }
func (*AllowTrustOp) ProtoMessage()               {}
func (*AllowTrustOp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

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

func (m *AllowTrustOp) GetAuthorized() uint32 {
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
	Amount uint64 `protobuf:"varint,3,opt,name=Amount" json:"Amount,omitempty"`
	// price in fractional format
	Price *Price `protobuf:"bytes,4,opt,name=Price" json:"Price,omitempty"`
	// ID of this offer
	OfferID string `protobuf:"bytes,5,opt,name=OfferID" json:"OfferID,omitempty"`
	// whether the offer is passive
	Passive uint32 `protobuf:"varint,6,opt,name=Passive" json:"Passive,omitempty"`
}

func (m *OfferOp) Reset()                    { *m = OfferOp{} }
func (m *OfferOp) String() string            { return proto.CompactTextString(m) }
func (*OfferOp) ProtoMessage()               {}
func (*OfferOp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

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

func (m *OfferOp) GetAmount() uint64 {
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

func (m *OfferOp) GetPassive() uint32 {
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

func init() { proto.RegisterFile("transaction.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xde, 0xb4, 0x49, 0xbb, 0x3d, 0x69, 0xfb, 0xeb, 0x6f, 0x58, 0x96, 0xb0, 0x88, 0xd4, 0x80,
	0x52, 0x44, 0x8a, 0xd4, 0x1b, 0x41, 0x10, 0xd2, 0x3f, 0xd2, 0x42, 0x6d, 0xc2, 0x34, 0x5e, 0x78,
	0xb5, 0xa4, 0xd9, 0xd9, 0x6d, 0x20, 0x4d, 0x62, 0x32, 0x59, 0x5b, 0x5f, 0xc3, 0x27, 0xf2, 0x31,
	0xc4, 0x97, 0x91, 0x99, 0x4c, 0xb2, 0x4d, 0xa8, 0xac, 0xe0, 0xdd, 0x9c, 0xef, 0x3b, 0xe7, 0xcc,
	0x77, 0xbe, 0x39, 0x0c, 0xfc, 0x4f, 0x63, 0x27, 0x48, 0x1c, 0x97, 0x7a, 0x61, 0x30, 0x8c, 0xe2,
	0x90, 0x86, 0x48, 0x49, 0x7d, 0x1a, 0x6d, 0xae, 0x54, 0x27, 0x49, 0x08, 0xcd, 0xb0, 0x2b, 0x35,
	0xbc, 0xbd, 0x25, 0x71, 0x16, 0xe8, 0x3f, 0x25, 0xa8, 0x99, 0x11, 0x7a, 0x0e, 0x0d, 0x33, 0xb2,
	0x0f, 0x11, 0xd1, 0xa4, 0xbe, 0x34, 0xe8, 0x8e, 0x3a, 0x43, 0x5e, 0x38, 0xcc, 0x40, 0x2c, 0x48,
	0xf4, 0x1e, 0x3a, 0x93, 0x98, 0x38, 0x94, 0x18, 0xae, 0x1b, 0xa6, 0x01, 0xd5, 0x6a, 0x7d, 0x69,
	0xa0, 0x8e, 0x2e, 0x45, 0x76, 0x89, 0x33, 0xa3, 0xf9, 0x19, 0x2e, 0xa7, 0xa3, 0x57, 0xd0, 0xb4,
	0x9c, 0xc3, 0x8e, 0x04, 0x54, 0xab, 0xf3, 0xca, 0x9e, 0xa8, 0x14, 0x28, 0xaf, 0xc9, 0x53, 0xd0,
	0x5b, 0x50, 0x2d, 0x87, 0x6e, 0xf3, 0x0a, 0x99, 0x57, 0x5c, 0x14, 0x15, 0x05, 0xc3, 0xab, 0x8e,
	0x53, 0xc7, 0x32, 0x1b, 0x4a, 0xff, 0x2e, 0x41, 0xcd, 0xde, 0xa3, 0x27, 0xd0, 0x12, 0xf7, 0x2f,
	0xa6, 0x7c, 0xbc, 0x16, 0x7e, 0x00, 0x50, 0x0f, 0xea, 0x1f, 0x08, 0xe1, 0x83, 0xc8, 0x98, 0x1d,
	0x11, 0x02, 0x79, 0x15, 0x52, 0xc2, 0x15, 0xb6, 0x30, 0x3f, 0xa3, 0x17, 0xd0, 0x5d, 0x93, 0x2f,
	0x29, 0x09, 0x5c, 0xb2, 0x4a, 0x77, 0x1b, 0x12, 0x73, 0x35, 0x32, 0xae, 0xa0, 0xe8, 0x19, 0xf3,
	0x71, 0xe9, 0x25, 0x54, 0x53, 0xfa, 0xf5, 0x81, 0x3a, 0x6a, 0x15, 0x3e, 0x62, 0x41, 0xe8, 0x18,
	0x14, 0x7b, 0xbf, 0x26, 0x94, 0xf5, 0xb4, 0x62, 0x72, 0xbf, 0x24, 0x37, 0x77, 0x24, 0x9e, 0x3b,
	0xc9, 0x56, 0x88, 0xab, 0xa0, 0xac, 0xa7, 0xbd, 0xe7, 0x3d, 0x6b, 0xa5, 0x9e, 0xf6, 0x1e, 0x0b,
	0x42, 0x5f, 0xc0, 0x7f, 0x15, 0xef, 0x1f, 0x99, 0x5a, 0x83, 0xe6, 0xd8, 0xf1, 0x9d, 0xc0, 0xcd,
	0x27, 0xcf, 0x43, 0x9d, 0x40, 0xab, 0xb0, 0xf5, 0x91, 0x26, 0x3a, 0x28, 0x06, 0xdb, 0x2b, 0xb1,
	0x05, 0x6d, 0xa1, 0x8b, 0x63, 0x38, 0xa3, 0xd0, 0x25, 0x34, 0x8c, 0x1d, 0x5f, 0x15, 0x66, 0x67,
	0x1d, 0x8b, 0x48, 0xff, 0x21, 0x41, 0xa7, 0xf4, 0x84, 0x68, 0x00, 0xe7, 0xeb, 0xd8, 0xcd, 0x1a,
	0x4a, 0x27, 0x1a, 0x16, 0x2c, 0x53, 0xc5, 0xce, 0xbb, 0x62, 0x03, 0x65, 0xfc, 0x00, 0x94, 0x35,
	0xd7, 0xab, 0x9a, 0x07, 0x70, 0x3e, 0x4d, 0x68, 0x76, 0x8b, 0x7c, 0xea, 0x96, 0x9c, 0x45, 0x7d,
	0x90, 0x99, 0x40, 0xf1, 0x90, 0xe5, 0x2c, 0xce, 0xe8, 0x13, 0x68, 0xda, 0x71, 0x9a, 0x30, 0xf1,
	0x7f, 0x63, 0xc5, 0x05, 0x28, 0x4b, 0x6f, 0xe7, 0x65, 0x4e, 0xc8, 0x38, 0x0b, 0xf4, 0x08, 0xda,
	0x86, 0xef, 0x87, 0x5f, 0xf3, 0x4e, 0xff, 0x6e, 0xf9, 0x53, 0x00, 0x23, 0xa5, 0xdb, 0x30, 0xf6,
	0xbe, 0x91, 0x1b, 0x7e, 0x59, 0x07, 0x1f, 0x21, 0xfa, 0x2f, 0x09, 0x9a, 0x26, 0xfb, 0x02, 0xcc,
	0x08, 0xbd, 0x86, 0xf6, 0x9a, 0xf8, 0xbe, 0x17, 0xdc, 0xfd, 0xd9, 0xf8, 0x52, 0x06, 0x1a, 0x82,
	0x3a, 0x4e, 0x0f, 0x45, 0xc1, 0x29, 0x1d, 0xc7, 0x09, 0x95, 0x05, 0x90, 0xf3, 0x05, 0x60, 0x93,
	0x58, 0xb1, 0xe7, 0x92, 0xca, 0x2b, 0x70, 0x0c, 0x67, 0x14, 0xdb, 0x52, 0x2e, 0x74, 0x31, 0xd5,
	0x14, 0xee, 0x44, 0x1e, 0x32, 0xc6, 0x72, 0x92, 0xc4, 0xbb, 0x27, 0x5a, 0x83, 0x0f, 0x98, 0x87,
	0x2f, 0xdf, 0xe5, 0x3f, 0x19, 0x42, 0xd0, 0x9d, 0xe0, 0x99, 0x61, 0xcf, 0xae, 0x8d, 0xc9, 0xc4,
	0xfc, 0xb4, 0xb2, 0x7b, 0x67, 0x48, 0x85, 0xa6, 0x65, 0x7c, 0xfe, 0x38, 0x5b, 0xd9, 0x3d, 0x09,
	0xf5, 0xa0, 0x6d, 0x19, 0xf6, 0xfc, 0x3a, 0x47, 0x6a, 0x9b, 0x06, 0xff, 0x14, 0xdf, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xae, 0xe9, 0x6b, 0x89, 0x4a, 0x05, 0x00, 0x00,
}
