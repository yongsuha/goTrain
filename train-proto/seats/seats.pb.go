// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: seats.proto

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

type AddSeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId     int64  `protobuf:"varint,1,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatNumber  string `protobuf:"bytes,2,opt,name=SeatNumber,proto3" json:"SeatNumber,omitempty"`
	SeatType    string `protobuf:"bytes,3,opt,name=SeatType,proto3" json:"SeatType,omitempty"`
	IsAvailable int64  `protobuf:"varint,4,opt,name=IsAvailable,proto3" json:"IsAvailable,omitempty"` // 1. 可选 2. 不可选
}

func (x *AddSeatReq) Reset() {
	*x = AddSeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSeatReq) ProtoMessage() {}

func (x *AddSeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSeatReq.ProtoReflect.Descriptor instead.
func (*AddSeatReq) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{0}
}

func (x *AddSeatReq) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *AddSeatReq) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *AddSeatReq) GetSeatType() string {
	if x != nil {
		return x.SeatType
	}
	return ""
}

func (x *AddSeatReq) GetIsAvailable() int64 {
	if x != nil {
		return x.IsAvailable
	}
	return 0
}

type AddSeatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *AddSeatResp) Reset() {
	*x = AddSeatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSeatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSeatResp) ProtoMessage() {}

func (x *AddSeatResp) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSeatResp.ProtoReflect.Descriptor instead.
func (*AddSeatResp) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{1}
}

func (x *AddSeatResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelSeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId int64 `protobuf:"varint,1,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
}

func (x *DelSeatReq) Reset() {
	*x = DelSeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelSeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelSeatReq) ProtoMessage() {}

func (x *DelSeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelSeatReq.ProtoReflect.Descriptor instead.
func (*DelSeatReq) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{2}
}

func (x *DelSeatReq) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

type UpdateSeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId     int64  `protobuf:"varint,1,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatNumber  string `protobuf:"bytes,2,opt,name=SeatNumber,proto3" json:"SeatNumber,omitempty"`
	IsAvailable int64  `protobuf:"varint,3,opt,name=IsAvailable,proto3" json:"IsAvailable,omitempty"`
}

func (x *UpdateSeatReq) Reset() {
	*x = UpdateSeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeatReq) ProtoMessage() {}

func (x *UpdateSeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeatReq.ProtoReflect.Descriptor instead.
func (*UpdateSeatReq) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSeatReq) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *UpdateSeatReq) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *UpdateSeatReq) GetIsAvailable() int64 {
	if x != nil {
		return x.IsAvailable
	}
	return 0
}

type UpdateSeatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId    int64  `protobuf:"varint,1,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatNumber string `protobuf:"bytes,2,opt,name=SeatNumber,proto3" json:"SeatNumber,omitempty"`
}

func (x *UpdateSeatResp) Reset() {
	*x = UpdateSeatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSeatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeatResp) ProtoMessage() {}

func (x *UpdateSeatResp) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeatResp.ProtoReflect.Descriptor instead.
func (*UpdateSeatResp) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSeatResp) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *UpdateSeatResp) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

type GetSeatDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId    int64  `protobuf:"varint,1,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatNumber string `protobuf:"bytes,2,opt,name=SeatNumber,proto3" json:"SeatNumber,omitempty"`
}

func (x *GetSeatDetailReq) Reset() {
	*x = GetSeatDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeatDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeatDetailReq) ProtoMessage() {}

func (x *GetSeatDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeatDetailReq.ProtoReflect.Descriptor instead.
func (*GetSeatDetailReq) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{5}
}

func (x *GetSeatDetailReq) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *GetSeatDetailReq) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

type GetSeatDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId     int64  `protobuf:"varint,1,opt,name=TrainId,proto3" json:"TrainId,omitempty"`
	SeatNumber  string `protobuf:"bytes,2,opt,name=SeatNumber,proto3" json:"SeatNumber,omitempty"`
	SeatType    string `protobuf:"bytes,3,opt,name=SeatType,proto3" json:"SeatType,omitempty"`
	IsAvailable int64  `protobuf:"varint,4,opt,name=IsAvailable,proto3" json:"IsAvailable,omitempty"`
	SeatId      int64  `protobuf:"varint,5,opt,name=SeatId,proto3" json:"SeatId,omitempty"`
	CreateTime  string `protobuf:"bytes,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime  string `protobuf:"bytes,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *GetSeatDetailResp) Reset() {
	*x = GetSeatDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeatDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeatDetailResp) ProtoMessage() {}

func (x *GetSeatDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_seats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeatDetailResp.ProtoReflect.Descriptor instead.
func (*GetSeatDetailResp) Descriptor() ([]byte, []int) {
	return file_seats_proto_rawDescGZIP(), []int{6}
}

func (x *GetSeatDetailResp) GetTrainId() int64 {
	if x != nil {
		return x.TrainId
	}
	return 0
}

func (x *GetSeatDetailResp) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *GetSeatDetailResp) GetSeatType() string {
	if x != nil {
		return x.SeatType
	}
	return ""
}

func (x *GetSeatDetailResp) GetIsAvailable() int64 {
	if x != nil {
		return x.IsAvailable
	}
	return 0
}

func (x *GetSeatDetailResp) GetSeatId() int64 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *GetSeatDetailResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *GetSeatDetailResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

var File_seats_proto protoreflect.FileDescriptor

var file_seats_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x26, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x4c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0xe3, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x53, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49,
	0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xca, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53, 0x65, 0x61, 0x74,
	0x12, 0x0b, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e,
	0x41, 0x64, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x44,
	0x65, 0x6c, 0x53, 0x65, 0x61, 0x74, 0x12, 0x0b, 0x2e, 0x44, 0x65, 0x6c, 0x53, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seats_proto_rawDescOnce sync.Once
	file_seats_proto_rawDescData = file_seats_proto_rawDesc
)

func file_seats_proto_rawDescGZIP() []byte {
	file_seats_proto_rawDescOnce.Do(func() {
		file_seats_proto_rawDescData = protoimpl.X.CompressGZIP(file_seats_proto_rawDescData)
	})
	return file_seats_proto_rawDescData
}

var file_seats_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_seats_proto_goTypes = []interface{}{
	(*AddSeatReq)(nil),        // 0: AddSeatReq
	(*AddSeatResp)(nil),       // 1: AddSeatResp
	(*DelSeatReq)(nil),        // 2: DelSeatReq
	(*UpdateSeatReq)(nil),     // 3: UpdateSeatReq
	(*UpdateSeatResp)(nil),    // 4: UpdateSeatResp
	(*GetSeatDetailReq)(nil),  // 5: GetSeatDetailReq
	(*GetSeatDetailResp)(nil), // 6: GetSeatDetailResp
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_seats_proto_depIdxs = []int32{
	0, // 0: SeatService.AddSeat:input_type -> AddSeatReq
	2, // 1: SeatService.DelSeat:input_type -> DelSeatReq
	3, // 2: SeatService.UpdateSeat:input_type -> UpdateSeatReq
	5, // 3: SeatService.GetSeatDetail:input_type -> GetSeatDetailReq
	1, // 4: SeatService.AddSeat:output_type -> AddSeatResp
	7, // 5: SeatService.DelSeat:output_type -> google.protobuf.Empty
	4, // 6: SeatService.UpdateSeat:output_type -> UpdateSeatResp
	6, // 7: SeatService.GetSeatDetail:output_type -> GetSeatDetailResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_seats_proto_init() }
func file_seats_proto_init() {
	if File_seats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSeatReq); i {
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
		file_seats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSeatResp); i {
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
		file_seats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelSeatReq); i {
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
		file_seats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSeatReq); i {
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
		file_seats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSeatResp); i {
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
		file_seats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeatDetailReq); i {
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
		file_seats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeatDetailResp); i {
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
			RawDescriptor: file_seats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seats_proto_goTypes,
		DependencyIndexes: file_seats_proto_depIdxs,
		MessageInfos:      file_seats_proto_msgTypes,
	}.Build()
	File_seats_proto = out.File
	file_seats_proto_rawDesc = nil
	file_seats_proto_goTypes = nil
	file_seats_proto_depIdxs = nil
}
