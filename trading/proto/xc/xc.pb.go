// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: trading/proto/xc/xc.proto

package xc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TradingSide int32

const (
	TradingSide_BUY  TradingSide = 0
	TradingSide_SELL TradingSide = 1
)

// Enum value maps for TradingSide.
var (
	TradingSide_name = map[int32]string{
		0: "BUY",
		1: "SELL",
	}
	TradingSide_value = map[string]int32{
		"BUY":  0,
		"SELL": 1,
	}
)

func (x TradingSide) Enum() *TradingSide {
	p := new(TradingSide)
	*p = x
	return p
}

func (x TradingSide) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradingSide) Descriptor() protoreflect.EnumDescriptor {
	return file_trading_proto_xc_xc_proto_enumTypes[0].Descriptor()
}

func (TradingSide) Type() protoreflect.EnumType {
	return &file_trading_proto_xc_xc_proto_enumTypes[0]
}

func (x TradingSide) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradingSide.Descriptor instead.
func (TradingSide) EnumDescriptor() ([]byte, []int) {
	return file_trading_proto_xc_xc_proto_rawDescGZIP(), []int{0}
}

type GetExchangeAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAssetType   string      `protobuf:"bytes,1,opt,name=base_asset_type,json=baseAssetType,proto3" json:"base_asset_type,omitempty"`
	QuoteAssetType  string      `protobuf:"bytes,2,opt,name=quote_asset_type,json=quoteAssetType,proto3" json:"quote_asset_type,omitempty"`
	Side            TradingSide `protobuf:"varint,3,opt,name=side,proto3,enum=xc.TradingSide" json:"side,omitempty"`
	Amount          string      `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountAssetType string      `protobuf:"bytes,5,opt,name=amount_asset_type,json=amountAssetType,proto3" json:"amount_asset_type,omitempty"`
}

func (x *GetExchangeAmountRequest) Reset() {
	*x = GetExchangeAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_xc_xc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeAmountRequest) ProtoMessage() {}

func (x *GetExchangeAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_xc_xc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeAmountRequest.ProtoReflect.Descriptor instead.
func (*GetExchangeAmountRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_xc_xc_proto_rawDescGZIP(), []int{0}
}

func (x *GetExchangeAmountRequest) GetBaseAssetType() string {
	if x != nil {
		return x.BaseAssetType
	}
	return ""
}

func (x *GetExchangeAmountRequest) GetQuoteAssetType() string {
	if x != nil {
		return x.QuoteAssetType
	}
	return ""
}

func (x *GetExchangeAmountRequest) GetSide() TradingSide {
	if x != nil {
		return x.Side
	}
	return TradingSide_BUY
}

func (x *GetExchangeAmountRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *GetExchangeAmountRequest) GetAmountAssetType() string {
	if x != nil {
		return x.AmountAssetType
	}
	return ""
}

type GetExchangeAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAssetType  string      `protobuf:"bytes,1,opt,name=base_asset_type,json=baseAssetType,proto3" json:"base_asset_type,omitempty"`
	QuoteAssetType string      `protobuf:"bytes,2,opt,name=quote_asset_type,json=quoteAssetType,proto3" json:"quote_asset_type,omitempty"`
	Side           TradingSide `protobuf:"varint,3,opt,name=side,proto3,enum=xc.TradingSide" json:"side,omitempty"`
	BaseAmount     string      `protobuf:"bytes,4,opt,name=base_amount,json=baseAmount,proto3" json:"base_amount,omitempty"`
	QuoteAmount    string      `protobuf:"bytes,5,opt,name=quote_amount,json=quoteAmount,proto3" json:"quote_amount,omitempty"`
}

func (x *GetExchangeAmountResponse) Reset() {
	*x = GetExchangeAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_xc_xc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeAmountResponse) ProtoMessage() {}

func (x *GetExchangeAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_xc_xc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeAmountResponse.ProtoReflect.Descriptor instead.
func (*GetExchangeAmountResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_xc_xc_proto_rawDescGZIP(), []int{1}
}

func (x *GetExchangeAmountResponse) GetBaseAssetType() string {
	if x != nil {
		return x.BaseAssetType
	}
	return ""
}

func (x *GetExchangeAmountResponse) GetQuoteAssetType() string {
	if x != nil {
		return x.QuoteAssetType
	}
	return ""
}

func (x *GetExchangeAmountResponse) GetSide() TradingSide {
	if x != nil {
		return x.Side
	}
	return TradingSide_BUY
}

func (x *GetExchangeAmountResponse) GetBaseAmount() string {
	if x != nil {
		return x.BaseAmount
	}
	return ""
}

func (x *GetExchangeAmountResponse) GetQuoteAmount() string {
	if x != nil {
		return x.QuoteAmount
	}
	return ""
}

type PlaceOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAssetType  string      `protobuf:"bytes,1,opt,name=base_asset_type,json=baseAssetType,proto3" json:"base_asset_type,omitempty"`
	QuoteAssetType string      `protobuf:"bytes,2,opt,name=quote_asset_type,json=quoteAssetType,proto3" json:"quote_asset_type,omitempty"`
	Side           TradingSide `protobuf:"varint,3,opt,name=side,proto3,enum=xc.TradingSide" json:"side,omitempty"`
	BaseAmount     string      `protobuf:"bytes,4,opt,name=base_amount,json=baseAmount,proto3" json:"base_amount,omitempty"`
	QuoteAmount    string      `protobuf:"bytes,5,opt,name=quote_amount,json=quoteAmount,proto3" json:"quote_amount,omitempty"`
}

func (x *PlaceOrderRequest) Reset() {
	*x = PlaceOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_xc_xc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderRequest) ProtoMessage() {}

func (x *PlaceOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_xc_xc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderRequest.ProtoReflect.Descriptor instead.
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_xc_xc_proto_rawDescGZIP(), []int{2}
}

func (x *PlaceOrderRequest) GetBaseAssetType() string {
	if x != nil {
		return x.BaseAssetType
	}
	return ""
}

func (x *PlaceOrderRequest) GetQuoteAssetType() string {
	if x != nil {
		return x.QuoteAssetType
	}
	return ""
}

func (x *PlaceOrderRequest) GetSide() TradingSide {
	if x != nil {
		return x.Side
	}
	return TradingSide_BUY
}

func (x *PlaceOrderRequest) GetBaseAmount() string {
	if x != nil {
		return x.BaseAmount
	}
	return ""
}

func (x *PlaceOrderRequest) GetQuoteAmount() string {
	if x != nil {
		return x.QuoteAmount
	}
	return ""
}

type PlaceOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *PlaceOrderResponse) Reset() {
	*x = PlaceOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_xc_xc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResponse) ProtoMessage() {}

func (x *PlaceOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_xc_xc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResponse.ProtoReflect.Descriptor instead.
func (*PlaceOrderResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_xc_xc_proto_rawDescGZIP(), []int{3}
}

func (x *PlaceOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_trading_proto_xc_xc_proto protoreflect.FileDescriptor

var file_trading_proto_xc_xc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x78, 0x63, 0x2f, 0x78, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x78, 0x63, 0x22,
	0xd5, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x78,
	0x63, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73,
	0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x78, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xce, 0x01, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x78, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x2f, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x2a, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x64,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4c, 0x4c, 0x10, 0x01, 0x32, 0x9e, 0x01, 0x0a, 0x09, 0x58, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x78, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x78, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x78, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x78, 0x63,
	0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_trading_proto_xc_xc_proto_rawDescOnce sync.Once
	file_trading_proto_xc_xc_proto_rawDescData = file_trading_proto_xc_xc_proto_rawDesc
)

func file_trading_proto_xc_xc_proto_rawDescGZIP() []byte {
	file_trading_proto_xc_xc_proto_rawDescOnce.Do(func() {
		file_trading_proto_xc_xc_proto_rawDescData = protoimpl.X.CompressGZIP(file_trading_proto_xc_xc_proto_rawDescData)
	})
	return file_trading_proto_xc_xc_proto_rawDescData
}

var file_trading_proto_xc_xc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trading_proto_xc_xc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_trading_proto_xc_xc_proto_goTypes = []interface{}{
	(TradingSide)(0),                  // 0: xc.TradingSide
	(*GetExchangeAmountRequest)(nil),  // 1: xc.GetExchangeAmountRequest
	(*GetExchangeAmountResponse)(nil), // 2: xc.GetExchangeAmountResponse
	(*PlaceOrderRequest)(nil),         // 3: xc.PlaceOrderRequest
	(*PlaceOrderResponse)(nil),        // 4: xc.PlaceOrderResponse
}
var file_trading_proto_xc_xc_proto_depIdxs = []int32{
	0, // 0: xc.GetExchangeAmountRequest.side:type_name -> xc.TradingSide
	0, // 1: xc.GetExchangeAmountResponse.side:type_name -> xc.TradingSide
	0, // 2: xc.PlaceOrderRequest.side:type_name -> xc.TradingSide
	1, // 3: xc.XcService.GetExchangeAmount:input_type -> xc.GetExchangeAmountRequest
	3, // 4: xc.XcService.PlaceOrder:input_type -> xc.PlaceOrderRequest
	2, // 5: xc.XcService.GetExchangeAmount:output_type -> xc.GetExchangeAmountResponse
	4, // 6: xc.XcService.PlaceOrder:output_type -> xc.PlaceOrderResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trading_proto_xc_xc_proto_init() }
func file_trading_proto_xc_xc_proto_init() {
	if File_trading_proto_xc_xc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trading_proto_xc_xc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeAmountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trading_proto_xc_xc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeAmountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trading_proto_xc_xc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trading_proto_xc_xc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trading_proto_xc_xc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trading_proto_xc_xc_proto_goTypes,
		DependencyIndexes: file_trading_proto_xc_xc_proto_depIdxs,
		EnumInfos:         file_trading_proto_xc_xc_proto_enumTypes,
		MessageInfos:      file_trading_proto_xc_xc_proto_msgTypes,
	}.Build()
	File_trading_proto_xc_xc_proto = out.File
	file_trading_proto_xc_xc_proto_rawDesc = nil
	file_trading_proto_xc_xc_proto_goTypes = nil
	file_trading_proto_xc_xc_proto_depIdxs = nil
}
