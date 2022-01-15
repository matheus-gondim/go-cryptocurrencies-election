// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/cryptocurrency.proto

package pb

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

type Cryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol  string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Like    int64  `protobuf:"varint,4,opt,name=like,proto3" json:"like,omitempty"`
	Dislike int64  `protobuf:"varint,5,opt,name=dislike,proto3" json:"dislike,omitempty"`
}

func (x *Cryptocurrency) Reset() {
	*x = Cryptocurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cryptocurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cryptocurrency) ProtoMessage() {}

func (x *Cryptocurrency) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cryptocurrency.ProtoReflect.Descriptor instead.
func (*Cryptocurrency) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{0}
}

func (x *Cryptocurrency) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cryptocurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cryptocurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Cryptocurrency) GetLike() int64 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *Cryptocurrency) GetDislike() int64 {
	if x != nil {
		return x.Dislike
	}
	return 0
}

type CreateCryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *CreateCryptocurrency) Reset() {
	*x = CreateCryptocurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCryptocurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCryptocurrency) ProtoMessage() {}

func (x *CreateCryptocurrency) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCryptocurrency.ProtoReflect.Descriptor instead.
func (*CreateCryptocurrency) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCryptocurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCryptocurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type GetCryptocurrencyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCryptocurrencyParams) Reset() {
	*x = GetCryptocurrencyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCryptocurrencyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptocurrencyParams) ProtoMessage() {}

func (x *GetCryptocurrencyParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCryptocurrencyParams.ProtoReflect.Descriptor instead.
func (*GetCryptocurrencyParams) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{2}
}

type CryptocurrencyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cryptocurrencies []*Cryptocurrency `protobuf:"bytes,1,rep,name=cryptocurrencies,proto3" json:"cryptocurrencies,omitempty"`
}

func (x *CryptocurrencyList) Reset() {
	*x = CryptocurrencyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencyList) ProtoMessage() {}

func (x *CryptocurrencyList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencyList.ProtoReflect.Descriptor instead.
func (*CryptocurrencyList) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{3}
}

func (x *CryptocurrencyList) GetCryptocurrencies() []*Cryptocurrency {
	if x != nil {
		return x.Cryptocurrencies
	}
	return nil
}

type CryptocurrencyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CryptocurrencyId) Reset() {
	*x = CryptocurrencyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencyId) ProtoMessage() {}

func (x *CryptocurrencyId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencyId.ProtoReflect.Descriptor instead.
func (*CryptocurrencyId) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{4}
}

func (x *CryptocurrencyId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CryptocurrencyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CryptocurrencyMessage) Reset() {
	*x = CryptocurrencyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencyMessage) ProtoMessage() {}

func (x *CryptocurrencyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencyMessage.ProtoReflect.Descriptor instead.
func (*CryptocurrencyMessage) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{5}
}

func (x *CryptocurrencyMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateCryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *UpdateCryptocurrency) Reset() {
	*x = UpdateCryptocurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cryptocurrency_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptocurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptocurrency) ProtoMessage() {}

func (x *UpdateCryptocurrency) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cryptocurrency_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptocurrency.ProtoReflect.Descriptor instead.
func (*UpdateCryptocurrency) Descriptor() ([]byte, []int) {
	return file_proto_cryptocurrency_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCryptocurrency) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCryptocurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCryptocurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

var File_proto_cryptocurrency_proto protoreflect.FileDescriptor

var file_proto_cryptocurrency_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x19, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x51, 0x0a, 0x12, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a,
	0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31,
	0x0a, 0x15, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x52, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x32, 0xa0, 0x03, 0x0a, 0x16, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x12, 0x15, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x45, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x18, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x37, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x11, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x32, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cryptocurrency_proto_rawDescOnce sync.Once
	file_proto_cryptocurrency_proto_rawDescData = file_proto_cryptocurrency_proto_rawDesc
)

func file_proto_cryptocurrency_proto_rawDescGZIP() []byte {
	file_proto_cryptocurrency_proto_rawDescOnce.Do(func() {
		file_proto_cryptocurrency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cryptocurrency_proto_rawDescData)
	})
	return file_proto_cryptocurrency_proto_rawDescData
}

var file_proto_cryptocurrency_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_cryptocurrency_proto_goTypes = []interface{}{
	(*Cryptocurrency)(nil),          // 0: Cryptocurrency
	(*CreateCryptocurrency)(nil),    // 1: CreateCryptocurrency
	(*GetCryptocurrencyParams)(nil), // 2: GetCryptocurrencyParams
	(*CryptocurrencyList)(nil),      // 3: CryptocurrencyList
	(*CryptocurrencyId)(nil),        // 4: CryptocurrencyId
	(*CryptocurrencyMessage)(nil),   // 5: CryptocurrencyMessage
	(*UpdateCryptocurrency)(nil),    // 6: UpdateCryptocurrency
}
var file_proto_cryptocurrency_proto_depIdxs = []int32{
	0, // 0: CryptocurrencyList.cryptocurrencies:type_name -> Cryptocurrency
	1, // 1: CryptocurrencyElection.CreateNew:input_type -> CreateCryptocurrency
	2, // 2: CryptocurrencyElection.FindCryptocurrencies:input_type -> GetCryptocurrencyParams
	4, // 3: CryptocurrencyElection.FindById:input_type -> CryptocurrencyId
	4, // 4: CryptocurrencyElection.DeleteById:input_type -> CryptocurrencyId
	6, // 5: CryptocurrencyElection.UpdateById:input_type -> UpdateCryptocurrency
	4, // 6: CryptocurrencyElection.UpvoteById:input_type -> CryptocurrencyId
	4, // 7: CryptocurrencyElection.DownvoteById:input_type -> CryptocurrencyId
	0, // 8: CryptocurrencyElection.CreateNew:output_type -> Cryptocurrency
	3, // 9: CryptocurrencyElection.FindCryptocurrencies:output_type -> CryptocurrencyList
	0, // 10: CryptocurrencyElection.FindById:output_type -> Cryptocurrency
	5, // 11: CryptocurrencyElection.DeleteById:output_type -> CryptocurrencyMessage
	5, // 12: CryptocurrencyElection.UpdateById:output_type -> CryptocurrencyMessage
	0, // 13: CryptocurrencyElection.UpvoteById:output_type -> Cryptocurrency
	0, // 14: CryptocurrencyElection.DownvoteById:output_type -> Cryptocurrency
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cryptocurrency_proto_init() }
func file_proto_cryptocurrency_proto_init() {
	if File_proto_cryptocurrency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cryptocurrency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cryptocurrency); i {
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
		file_proto_cryptocurrency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCryptocurrency); i {
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
		file_proto_cryptocurrency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCryptocurrencyParams); i {
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
		file_proto_cryptocurrency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencyList); i {
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
		file_proto_cryptocurrency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencyId); i {
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
		file_proto_cryptocurrency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencyMessage); i {
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
		file_proto_cryptocurrency_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCryptocurrency); i {
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
			RawDescriptor: file_proto_cryptocurrency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cryptocurrency_proto_goTypes,
		DependencyIndexes: file_proto_cryptocurrency_proto_depIdxs,
		MessageInfos:      file_proto_cryptocurrency_proto_msgTypes,
	}.Build()
	File_proto_cryptocurrency_proto = out.File
	file_proto_cryptocurrency_proto_rawDesc = nil
	file_proto_cryptocurrency_proto_goTypes = nil
	file_proto_cryptocurrency_proto_depIdxs = nil
}
