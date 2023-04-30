// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: src/proto/api.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId uint32  `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	DueDate  string  `protobuf:"bytes,2,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	Value    float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddInvoiceRequest) Reset() {
	*x = AddInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInvoiceRequest) ProtoMessage() {}

func (x *AddInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInvoiceRequest.ProtoReflect.Descriptor instead.
func (*AddInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_api_proto_rawDescGZIP(), []int{0}
}

func (x *AddInvoiceRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *AddInvoiceRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

func (x *AddInvoiceRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base      *Base  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	FirtsName string `protobuf:"bytes,2,opt,name=firtsName,proto3" json:"firtsName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_src_proto_api_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *User) GetFirtsName() string {
	if x != nil {
		return x.FirtsName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *Base                  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ClientId uint32                 `protobuf:"varint,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Value    float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	DueDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_src_proto_api_proto_rawDescGZIP(), []int{2}
}

func (x *Invoice) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Invoice) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Invoice) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Invoice) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Version   uint32                 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_src_proto_api_proto_rawDescGZIP(), []int{3}
}

func (x *Base) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Base) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Base) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Base) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Base) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Base) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type BasicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId int32  `protobuf:"varint,1,opt,name=referenceId,proto3" json:"referenceId,omitempty"`
	Success     bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BasicResponse) Reset() {
	*x = BasicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicResponse) ProtoMessage() {}

func (x *BasicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicResponse.ProtoReflect.Descriptor instead.
func (*BasicResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_api_proto_rawDescGZIP(), []int{4}
}

func (x *BasicResponse) GetReferenceId() int32 {
	if x != nil {
		return x.ReferenceId
	}
	return 0
}

func (x *BasicResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BasicResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_src_proto_api_proto protoreflect.FileDescriptor

var file_src_proto_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x74, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x74,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x0d,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x62, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_api_proto_rawDescOnce sync.Once
	file_src_proto_api_proto_rawDescData = file_src_proto_api_proto_rawDesc
)

func file_src_proto_api_proto_rawDescGZIP() []byte {
	file_src_proto_api_proto_rawDescOnce.Do(func() {
		file_src_proto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_api_proto_rawDescData)
	})
	return file_src_proto_api_proto_rawDescData
}

var file_src_proto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_src_proto_api_proto_goTypes = []interface{}{
	(*AddInvoiceRequest)(nil),     // 0: project.proto.AddInvoiceRequest
	(*User)(nil),                  // 1: project.proto.User
	(*Invoice)(nil),               // 2: project.proto.Invoice
	(*Base)(nil),                  // 3: project.proto.Base
	(*BasicResponse)(nil),         // 4: project.proto.BasicResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_src_proto_api_proto_depIdxs = []int32{
	3, // 0: project.proto.User.base:type_name -> project.proto.Base
	3, // 1: project.proto.Invoice.base:type_name -> project.proto.Base
	5, // 2: project.proto.Invoice.dueDate:type_name -> google.protobuf.Timestamp
	5, // 3: project.proto.Base.deleted_at:type_name -> google.protobuf.Timestamp
	5, // 4: project.proto.Base.created_at:type_name -> google.protobuf.Timestamp
	5, // 5: project.proto.Base.updated_at:type_name -> google.protobuf.Timestamp
	0, // 6: project.proto.FinancialService.AddInvoice:input_type -> project.proto.AddInvoiceRequest
	4, // 7: project.proto.FinancialService.AddInvoice:output_type -> project.proto.BasicResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_src_proto_api_proto_init() }
func file_src_proto_api_proto_init() {
	if File_src_proto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInvoiceRequest); i {
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
		file_src_proto_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_src_proto_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_src_proto_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_src_proto_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicResponse); i {
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
			RawDescriptor: file_src_proto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_api_proto_goTypes,
		DependencyIndexes: file_src_proto_api_proto_depIdxs,
		MessageInfos:      file_src_proto_api_proto_msgTypes,
	}.Build()
	File_src_proto_api_proto = out.File
	file_src_proto_api_proto_rawDesc = nil
	file_src_proto_api_proto_goTypes = nil
	file_src_proto_api_proto_depIdxs = nil
}
