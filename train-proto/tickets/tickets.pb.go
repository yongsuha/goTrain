// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: tickets.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddTicketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TrainId      int64  `protobuf:"varint,2,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatId       int64  `protobuf:"varint,3,opt,name=SeatId,proto3" json:"SeatId,omitempty"`
	Price        int64  `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
	PurchaseTime string `protobuf:"bytes,5,opt,name=PurchaseTime,proto3" json:"PurchaseTime,omitempty"`
}

func (x *AddTicketReq) Reset() {
	*x = AddTicketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTicketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTicketReq) ProtoMessage() {}

func (x *AddTicketReq) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTicketReq.ProtoReflect.Descriptor instead.
func (*AddTicketReq) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{0}
}

func (x *AddTicketReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddTicketReq) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *AddTicketReq) GetSeatId() int64 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *AddTicketReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddTicketReq) GetPurchaseTime() string {
	if x != nil {
		return x.PurchaseTime
	}
	return ""
}

type AddTicketResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId int64 `protobuf:"varint,1,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
}

func (x *AddTicketResp) Reset() {
	*x = AddTicketResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTicketResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTicketResp) ProtoMessage() {}

func (x *AddTicketResp) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTicketResp.ProtoReflect.Descriptor instead.
func (*AddTicketResp) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{1}
}

func (x *AddTicketResp) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type DelTicketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId int64 `protobuf:"varint,1,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
}

func (x *DelTicketReq) Reset() {
	*x = DelTicketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelTicketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelTicketReq) ProtoMessage() {}

func (x *DelTicketReq) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelTicketReq.ProtoReflect.Descriptor instead.
func (*DelTicketReq) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{2}
}

func (x *DelTicketReq) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type GetTicketDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId int64 `protobuf:"varint,1,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
	UserId   int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetTicketDetailReq) Reset() {
	*x = GetTicketDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketDetailReq) ProtoMessage() {}

func (x *GetTicketDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketDetailReq.ProtoReflect.Descriptor instead.
func (*GetTicketDetailReq) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{3}
}

func (x *GetTicketDetailReq) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *GetTicketDetailReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetTicketDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketDetail []*TicketInfo `protobuf:"bytes,1,rep,name=TicketDetail,proto3" json:"TicketDetail,omitempty"`
}

func (x *GetTicketDetailResp) Reset() {
	*x = GetTicketDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketDetailResp) ProtoMessage() {}

func (x *GetTicketDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketDetailResp.ProtoReflect.Descriptor instead.
func (*GetTicketDetailResp) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{4}
}

func (x *GetTicketDetailResp) GetTicketDetail() []*TicketInfo {
	if x != nil {
		return x.TicketDetail
	}
	return nil
}

type TicketInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId     int64  `protobuf:"varint,1,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
	UserId       int64  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TrainId      int64  `protobuf:"varint,3,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatId       int64  `protobuf:"varint,4,opt,name=SeatId,proto3" json:"SeatId,omitempty"`
	Price        int64  `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
	PurchaseTime string `protobuf:"bytes,6,opt,name=PurchaseTime,proto3" json:"PurchaseTime,omitempty"`
	CreateTime   string `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime   string `protobuf:"bytes,8,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *TicketInfo) Reset() {
	*x = TicketInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketInfo) ProtoMessage() {}

func (x *TicketInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketInfo.ProtoReflect.Descriptor instead.
func (*TicketInfo) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{5}
}

func (x *TicketInfo) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *TicketInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TicketInfo) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *TicketInfo) GetSeatId() int64 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *TicketInfo) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TicketInfo) GetPurchaseTime() string {
	if x != nil {
		return x.PurchaseTime
	}
	return ""
}

func (x *TicketInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *TicketInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type UpdateTicketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId     int64  `protobuf:"varint,1,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
	UserId       int64  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TrainId      int64  `protobuf:"varint,3,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatId       int64  `protobuf:"varint,4,opt,name=SeatId,proto3" json:"SeatId,omitempty"`
	Price        int64  `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
	PurchaseTime string `protobuf:"bytes,6,opt,name=PurchaseTime,proto3" json:"PurchaseTime,omitempty"`
}

func (x *UpdateTicketReq) Reset() {
	*x = UpdateTicketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketReq) ProtoMessage() {}

func (x *UpdateTicketReq) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketReq.ProtoReflect.Descriptor instead.
func (*UpdateTicketReq) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTicketReq) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *UpdateTicketReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateTicketReq) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *UpdateTicketReq) GetSeatId() int64 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *UpdateTicketReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateTicketReq) GetPurchaseTime() string {
	if x != nil {
		return x.PurchaseTime
	}
	return ""
}

type UpdateTicketResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId     int64  `protobuf:"varint,1,opt,name=TicketId,proto3" json:"TicketId,omitempty"`
	UserId       int64  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TrainId      int64  `protobuf:"varint,3,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatId       int64  `protobuf:"varint,4,opt,name=SeatId,proto3" json:"SeatId,omitempty"`
	Price        int64  `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
	PurchaseTime string `protobuf:"bytes,6,opt,name=PurchaseTime,proto3" json:"PurchaseTime,omitempty"`
	CreateTime   string `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime   string `protobuf:"bytes,8,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *UpdateTicketResp) Reset() {
	*x = UpdateTicketResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketResp) ProtoMessage() {}

func (x *UpdateTicketResp) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketResp.ProtoReflect.Descriptor instead.
func (*UpdateTicketResp) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTicketResp) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *UpdateTicketResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateTicketResp) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *UpdateTicketResp) GetSeatId() int64 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *UpdateTicketResp) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateTicketResp) GetPurchaseTime() string {
	if x != nil {
		return x.PurchaseTime
	}
	return ""
}

func (x *UpdateTicketResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *UpdateTicketResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

var File_tickets_proto protoreflect.FileDescriptor

var file_tickets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0c, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xec, 0x01, 0x0a,
	0x0a, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xf2, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0xe2, 0x01, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x32, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickets_proto_rawDescOnce sync.Once
	file_tickets_proto_rawDescData = file_tickets_proto_rawDesc
)

func file_tickets_proto_rawDescGZIP() []byte {
	file_tickets_proto_rawDescOnce.Do(func() {
		file_tickets_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickets_proto_rawDescData)
	})
	return file_tickets_proto_rawDescData
}

var file_tickets_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tickets_proto_goTypes = []interface{}{
	(*AddTicketReq)(nil),        // 0: AddTicketReq
	(*AddTicketResp)(nil),       // 1: AddTicketResp
	(*DelTicketReq)(nil),        // 2: DelTicketReq
	(*GetTicketDetailReq)(nil),  // 3: GetTicketDetailReq
	(*GetTicketDetailResp)(nil), // 4: GetTicketDetailResp
	(*TicketInfo)(nil),          // 5: TicketInfo
	(*UpdateTicketReq)(nil),     // 6: UpdateTicketReq
	(*UpdateTicketResp)(nil),    // 7: UpdateTicketResp
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_tickets_proto_depIdxs = []int32{
	5, // 0: GetTicketDetailResp.TicketDetail:type_name -> TicketInfo
	0, // 1: TicketService.AddTicket:input_type -> AddTicketReq
	2, // 2: TicketService.DelTicket:input_type -> DelTicketReq
	3, // 3: TicketService.GetTicketDetail:input_type -> GetTicketDetailReq
	6, // 4: TicketService.UpdateTicket:input_type -> UpdateTicketReq
	1, // 5: TicketService.AddTicket:output_type -> AddTicketResp
	8, // 6: TicketService.DelTicket:output_type -> google.protobuf.Empty
	4, // 7: TicketService.GetTicketDetail:output_type -> GetTicketDetailResp
	7, // 8: TicketService.UpdateTicket:output_type -> UpdateTicketResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tickets_proto_init() }
func file_tickets_proto_init() {
	if File_tickets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tickets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTicketReq); i {
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
		file_tickets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTicketResp); i {
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
		file_tickets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelTicketReq); i {
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
		file_tickets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketDetailReq); i {
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
		file_tickets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketDetailResp); i {
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
		file_tickets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketInfo); i {
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
		file_tickets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketReq); i {
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
		file_tickets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketResp); i {
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
			RawDescriptor: file_tickets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tickets_proto_goTypes,
		DependencyIndexes: file_tickets_proto_depIdxs,
		MessageInfos:      file_tickets_proto_msgTypes,
	}.Build()
	File_tickets_proto = out.File
	file_tickets_proto_rawDesc = nil
	file_tickets_proto_goTypes = nil
	file_tickets_proto_depIdxs = nil
}
