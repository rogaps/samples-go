// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: trading/proto/trading/trading.proto

package trading

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
	return file_trading_proto_trading_trading_proto_enumTypes[0].Descriptor()
}

func (TradingSide) Type() protoreflect.EnumType {
	return &file_trading_proto_trading_trading_proto_enumTypes[0]
}

func (x TradingSide) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradingSide.Descriptor instead.
func (TradingSide) EnumDescriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{0}
}

type TransactionStatus int32

const (
	TransactionStatus_PENDING TransactionStatus = 0
	TransactionStatus_FAILED  TransactionStatus = 1
	TransactionStatus_SUCCESS TransactionStatus = 2
)

// Enum value maps for TransactionStatus.
var (
	TransactionStatus_name = map[int32]string{
		0: "PENDING",
		1: "FAILED",
		2: "SUCCESS",
	}
	TransactionStatus_value = map[string]int32{
		"PENDING": 0,
		"FAILED":  1,
		"SUCCESS": 2,
	}
)

func (x TransactionStatus) Enum() *TransactionStatus {
	p := new(TransactionStatus)
	*p = x
	return p
}

func (x TransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_trading_proto_trading_trading_proto_enumTypes[1].Descriptor()
}

func (TransactionStatus) Type() protoreflect.EnumType {
	return &file_trading_proto_trading_trading_proto_enumTypes[1]
}

func (x TransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionStatus.Descriptor instead.
func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{1}
}

type CreateTradingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAssetType   string      `protobuf:"bytes,1,opt,name=base_asset_type,json=baseAssetType,proto3" json:"base_asset_type,omitempty"`
	QuoteAssetType  string      `protobuf:"bytes,2,opt,name=quote_asset_type,json=quoteAssetType,proto3" json:"quote_asset_type,omitempty"`
	Side            TradingSide `protobuf:"varint,3,opt,name=side,proto3,enum=trading.TradingSide" json:"side,omitempty"`
	Amount          string      `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountAssetType string      `protobuf:"bytes,5,opt,name=amount_asset_type,json=amountAssetType,proto3" json:"amount_asset_type,omitempty"`
}

func (x *CreateTradingRequest) Reset() {
	*x = CreateTradingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_trading_trading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradingRequest) ProtoMessage() {}

func (x *CreateTradingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_trading_trading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradingRequest.ProtoReflect.Descriptor instead.
func (*CreateTradingRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTradingRequest) GetBaseAssetType() string {
	if x != nil {
		return x.BaseAssetType
	}
	return ""
}

func (x *CreateTradingRequest) GetQuoteAssetType() string {
	if x != nil {
		return x.QuoteAssetType
	}
	return ""
}

func (x *CreateTradingRequest) GetSide() TradingSide {
	if x != nil {
		return x.Side
	}
	return TradingSide_BUY
}

func (x *CreateTradingRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CreateTradingRequest) GetAmountAssetType() string {
	if x != nil {
		return x.AmountAssetType
	}
	return ""
}

type CreateTradingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId   string      `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	BaseAssetType   string      `protobuf:"bytes,2,opt,name=base_asset_type,json=baseAssetType,proto3" json:"base_asset_type,omitempty"`
	QuoteAssetType  string      `protobuf:"bytes,3,opt,name=quote_asset_type,json=quoteAssetType,proto3" json:"quote_asset_type,omitempty"`
	Side            TradingSide `protobuf:"varint,4,opt,name=side,proto3,enum=trading.TradingSide" json:"side,omitempty"`
	Amount          string      `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountAssetType string      `protobuf:"bytes,6,opt,name=amount_asset_type,json=amountAssetType,proto3" json:"amount_asset_type,omitempty"`
	Total           string      `protobuf:"bytes,7,opt,name=total,proto3" json:"total,omitempty"`
	TotalAssetType  string      `protobuf:"bytes,8,opt,name=total_asset_type,json=totalAssetType,proto3" json:"total_asset_type,omitempty"`
}

func (x *CreateTradingResponse) Reset() {
	*x = CreateTradingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_trading_trading_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradingResponse) ProtoMessage() {}

func (x *CreateTradingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_trading_trading_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradingResponse.ProtoReflect.Descriptor instead.
func (*CreateTradingResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTradingResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *CreateTradingResponse) GetBaseAssetType() string {
	if x != nil {
		return x.BaseAssetType
	}
	return ""
}

func (x *CreateTradingResponse) GetQuoteAssetType() string {
	if x != nil {
		return x.QuoteAssetType
	}
	return ""
}

func (x *CreateTradingResponse) GetSide() TradingSide {
	if x != nil {
		return x.Side
	}
	return TradingSide_BUY
}

func (x *CreateTradingResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CreateTradingResponse) GetAmountAssetType() string {
	if x != nil {
		return x.AmountAssetType
	}
	return ""
}

func (x *CreateTradingResponse) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *CreateTradingResponse) GetTotalAssetType() string {
	if x != nil {
		return x.TotalAssetType
	}
	return ""
}

type CommitTradingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *CommitTradingRequest) Reset() {
	*x = CommitTradingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_trading_trading_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitTradingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitTradingRequest) ProtoMessage() {}

func (x *CommitTradingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_trading_trading_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitTradingRequest.ProtoReflect.Descriptor instead.
func (*CommitTradingRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{2}
}

func (x *CommitTradingRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type CommitTradingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string            `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status        TransactionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=trading.TransactionStatus" json:"status,omitempty"`
}

func (x *CommitTradingResponse) Reset() {
	*x = CommitTradingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_trading_trading_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitTradingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitTradingResponse) ProtoMessage() {}

func (x *CommitTradingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_trading_trading_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitTradingResponse.ProtoReflect.Descriptor instead.
func (*CommitTradingResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{3}
}

func (x *CommitTradingResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *CommitTradingResponse) GetStatus() TransactionStatus {
	if x != nil {
		return x.Status
	}
	return TransactionStatus_PENDING
}

type RegisterActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskToken string `protobuf:"bytes,2,opt,name=task_token,json=taskToken,proto3" json:"task_token,omitempty"`
}

func (x *RegisterActivityRequest) Reset() {
	*x = RegisterActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_trading_trading_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterActivityRequest) ProtoMessage() {}

func (x *RegisterActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_trading_trading_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterActivityRequest.ProtoReflect.Descriptor instead.
func (*RegisterActivityRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterActivityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterActivityRequest) GetTaskToken() string {
	if x != nil {
		return x.TaskToken
	}
	return ""
}

type RegisterActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterActivityResponse) Reset() {
	*x = RegisterActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_trading_trading_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterActivityResponse) ProtoMessage() {}

func (x *RegisterActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_trading_trading_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterActivityResponse.ProtoReflect.Descriptor instead.
func (*RegisterActivityResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_trading_trading_proto_rawDescGZIP(), []int{5}
}

var File_trading_proto_trading_trading_proto protoreflect.FileDescriptor

var file_trading_proto_trading_trading_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0xd6,
	0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x69, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73,
	0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x69,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04,
	0x73, 0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28,
	0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a, 0x17, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2a, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x64, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c,
	0x4c, 0x10, 0x01, 0x2a, 0x39, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x32, 0x8f,
	0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x72,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_trading_proto_trading_trading_proto_rawDescOnce sync.Once
	file_trading_proto_trading_trading_proto_rawDescData = file_trading_proto_trading_trading_proto_rawDesc
)

func file_trading_proto_trading_trading_proto_rawDescGZIP() []byte {
	file_trading_proto_trading_trading_proto_rawDescOnce.Do(func() {
		file_trading_proto_trading_trading_proto_rawDescData = protoimpl.X.CompressGZIP(file_trading_proto_trading_trading_proto_rawDescData)
	})
	return file_trading_proto_trading_trading_proto_rawDescData
}

var file_trading_proto_trading_trading_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_trading_proto_trading_trading_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_trading_proto_trading_trading_proto_goTypes = []interface{}{
	(TradingSide)(0),                 // 0: trading.TradingSide
	(TransactionStatus)(0),           // 1: trading.TransactionStatus
	(*CreateTradingRequest)(nil),     // 2: trading.CreateTradingRequest
	(*CreateTradingResponse)(nil),    // 3: trading.CreateTradingResponse
	(*CommitTradingRequest)(nil),     // 4: trading.CommitTradingRequest
	(*CommitTradingResponse)(nil),    // 5: trading.CommitTradingResponse
	(*RegisterActivityRequest)(nil),  // 6: trading.RegisterActivityRequest
	(*RegisterActivityResponse)(nil), // 7: trading.RegisterActivityResponse
}
var file_trading_proto_trading_trading_proto_depIdxs = []int32{
	0, // 0: trading.CreateTradingRequest.side:type_name -> trading.TradingSide
	0, // 1: trading.CreateTradingResponse.side:type_name -> trading.TradingSide
	1, // 2: trading.CommitTradingResponse.status:type_name -> trading.TransactionStatus
	2, // 3: trading.TradingService.CreateTrading:input_type -> trading.CreateTradingRequest
	4, // 4: trading.TradingService.CommitTrading:input_type -> trading.CommitTradingRequest
	6, // 5: trading.TradingService.RegisterActivity:input_type -> trading.RegisterActivityRequest
	3, // 6: trading.TradingService.CreateTrading:output_type -> trading.CreateTradingResponse
	5, // 7: trading.TradingService.CommitTrading:output_type -> trading.CommitTradingResponse
	7, // 8: trading.TradingService.RegisterActivity:output_type -> trading.RegisterActivityResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trading_proto_trading_trading_proto_init() }
func file_trading_proto_trading_trading_proto_init() {
	if File_trading_proto_trading_trading_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trading_proto_trading_trading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradingRequest); i {
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
		file_trading_proto_trading_trading_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradingResponse); i {
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
		file_trading_proto_trading_trading_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitTradingRequest); i {
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
		file_trading_proto_trading_trading_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitTradingResponse); i {
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
		file_trading_proto_trading_trading_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterActivityRequest); i {
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
		file_trading_proto_trading_trading_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterActivityResponse); i {
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
			RawDescriptor: file_trading_proto_trading_trading_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trading_proto_trading_trading_proto_goTypes,
		DependencyIndexes: file_trading_proto_trading_trading_proto_depIdxs,
		EnumInfos:         file_trading_proto_trading_trading_proto_enumTypes,
		MessageInfos:      file_trading_proto_trading_trading_proto_msgTypes,
	}.Build()
	File_trading_proto_trading_trading_proto = out.File
	file_trading_proto_trading_trading_proto_rawDesc = nil
	file_trading_proto_trading_trading_proto_goTypes = nil
	file_trading_proto_trading_trading_proto_depIdxs = nil
}
