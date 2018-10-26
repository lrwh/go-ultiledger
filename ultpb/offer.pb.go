// Code generated by protoc-gen-go. DO NOT EDIT.
// source: offer.proto

package ultpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Price struct {
	Numerator   uint64 `protobuf:"varint,1,opt,name=Numerator" json:"Numerator,omitempty"`
	Denominator uint64 `protobuf:"varint,2,opt,name=Denominator" json:"Denominator,omitempty"`
}

func (m *Price) Reset()                    { *m = Price{} }
func (m *Price) String() string            { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()               {}
func (*Price) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Price) GetNumerator() uint64 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *Price) GetDenominator() uint64 {
	if m != nil {
		return m.Denominator
	}
	return 0
}

type Offer struct {
	// seller of the offer
	AccountID string `protobuf:"bytes,1,opt,name=AccountID" json:"AccountID,omitempty"`
	// ID of the offer
	OfferID string `protobuf:"bytes,2,opt,name=OfferID" json:"OfferID,omitempty"`
	// asset for selling
	SellingAsset *Asset `protobuf:"bytes,3,opt,name=SellingAsset" json:"SellingAsset,omitempty"`
	// asset for buying
	BuyingAsset *Asset `protobuf:"bytes,4,opt,name=BuyingAsset" json:"BuyingAsset,omitempty"`
	// amount of asset for selling
	Amount uint64 `protobuf:"varint,5,opt,name=Amount" json:"Amount,omitempty"`
	// price of the offer
	Price *Price `protobuf:"bytes,6,opt,name=Price" json:"Price,omitempty"`
	// passive offer?
	Passive uint32 `protobuf:"varint,7,opt,name=Passive" json:"Passive,omitempty"`
}

func (m *Offer) Reset()                    { *m = Offer{} }
func (m *Offer) String() string            { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()               {}
func (*Offer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Offer) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Offer) GetOfferID() string {
	if m != nil {
		return m.OfferID
	}
	return ""
}

func (m *Offer) GetSellingAsset() *Asset {
	if m != nil {
		return m.SellingAsset
	}
	return nil
}

func (m *Offer) GetBuyingAsset() *Asset {
	if m != nil {
		return m.BuyingAsset
	}
	return nil
}

func (m *Offer) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Offer) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Offer) GetPassive() uint32 {
	if m != nil {
		return m.Passive
	}
	return 0
}

func init() {
	proto.RegisterType((*Price)(nil), "ultpb.Price")
	proto.RegisterType((*Offer)(nil), "ultpb.Offer")
}

func init() { proto.RegisterFile("offer.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x89, 0x6e, 0x5b, 0x3a, 0x59, 0x2f, 0x73, 0x90, 0x20, 0x1e, 0x4a, 0x4f, 0x3d, 0x15,
	0xd1, 0x4f, 0x50, 0x29, 0xc8, 0x5e, 0x74, 0x89, 0x9f, 0xa0, 0x5b, 0xb2, 0x12, 0x68, 0x93, 0x25,
	0x7f, 0x04, 0x3f, 0xb9, 0x57, 0xc9, 0xd8, 0x75, 0x77, 0xc1, 0xe3, 0xfc, 0xe6, 0xcd, 0xcb, 0xcb,
	0x03, 0x6e, 0xf7, 0x7b, 0xe5, 0xda, 0x83, 0xb3, 0xc1, 0x62, 0x16, 0xa7, 0x70, 0xd8, 0xdd, 0xf1,
	0xc1, 0x7b, 0x15, 0x7e, 0x59, 0xfd, 0x02, 0xd9, 0xd6, 0xe9, 0x51, 0xe1, 0x3d, 0x94, 0xaf, 0x71,
	0x56, 0x6e, 0x08, 0xd6, 0x09, 0x56, 0xb1, 0x66, 0x25, 0x4f, 0x00, 0x2b, 0xe0, 0xbd, 0x32, 0x76,
	0xd6, 0x86, 0xf6, 0x57, 0xb4, 0x3f, 0x47, 0xf5, 0x37, 0x83, 0xec, 0x2d, 0x3d, 0x96, 0x9c, 0xba,
	0x71, 0xb4, 0xd1, 0x84, 0x4d, 0x4f, 0x4e, 0xa5, 0x3c, 0x01, 0x14, 0x50, 0x90, 0x6c, 0xd3, 0x93,
	0x4b, 0x29, 0x8f, 0x23, 0x3e, 0xc0, 0xfa, 0x5d, 0x4d, 0x93, 0x36, 0x1f, 0x5d, 0x0a, 0x28, 0xae,
	0x2b, 0xd6, 0xf0, 0xc7, 0x75, 0x4b, 0xa9, 0x5b, 0x62, 0xf2, 0x42, 0x81, 0x2d, 0xf0, 0xe7, 0xf8,
	0xf5, 0x77, 0xb0, 0xfa, 0xe7, 0xe0, 0x5c, 0x80, 0xb7, 0x90, 0x77, 0x73, 0xca, 0x21, 0x32, 0xfa,
	0xc0, 0x32, 0x61, 0xbd, 0x94, 0x20, 0xf2, 0x0b, 0x07, 0x62, 0x72, 0xe9, 0x47, 0x40, 0xb1, 0x1d,
	0xbc, 0xd7, 0x9f, 0x4a, 0x14, 0x15, 0x6b, 0x6e, 0xe4, 0x71, 0xdc, 0xe5, 0xd4, 0xe4, 0xd3, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x11, 0xc1, 0x75, 0x6c, 0x01, 0x00, 0x00,
}
