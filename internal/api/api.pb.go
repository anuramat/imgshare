// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/api.proto

package api

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID      string `protobuf:"bytes,1,opt,name=FileID,proto3" json:"FileID,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Upvotes     int64  `protobuf:"varint,3,opt,name=Upvotes,proto3" json:"Upvotes,omitempty"`
	Downvotes   int64  `protobuf:"varint,4,opt,name=Downvotes,proto3" json:"Downvotes,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetFileID() string {
	if x != nil {
		return x.FileID
	}
	return ""
}

func (x *Image) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Image) GetUpvotes() int64 {
	if x != nil {
		return x.Upvotes
	}
	return 0
}

func (x *Image) GetDownvotes() int64 {
	if x != nil {
		return x.Downvotes
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

type ImageAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Image  *Image `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *ImageAuthRequest) Reset() {
	*x = ImageAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageAuthRequest) ProtoMessage() {}

func (x *ImageAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageAuthRequest.ProtoReflect.Descriptor instead.
func (*ImageAuthRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ImageAuthRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ImageAuthRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//homework
	Image []*Image `protobuf:"bytes,1,rep,name=Image,proto3" json:"Image,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Images) GetImage() []*Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//homework
	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *Page) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Page) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GalleryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	UserID int64 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GalleryRequest) Reset() {
	*x = GalleryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryRequest) ProtoMessage() {}

func (x *GalleryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryRequest.ProtoReflect.Descriptor instead.
func (*GalleryRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *GalleryRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GalleryRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GalleryImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Total  int32  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Image  *Image `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *GalleryImage) Reset() {
	*x = GalleryImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryImage) ProtoMessage() {}

func (x *GalleryImage) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryImage.ProtoReflect.Descriptor instead.
func (*GalleryImage) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *GalleryImage) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GalleryImage) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GalleryImage) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x79, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4c, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x22, 0x34, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x40, 0x0a, 0x0e, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x0c, 0x47, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x32, 0xc0, 0x03, 0x0a, 0x08, 0x49,
	0x6d, 0x67, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a,
	0x0b, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x32, 0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x6d, 0x61, 0x74, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2d, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_api_proto_goTypes = []interface{}{
	(*Image)(nil),            // 0: api.Image
	(*Empty)(nil),            // 1: api.Empty
	(*ImageAuthRequest)(nil), // 2: api.ImageAuthRequest
	(*Images)(nil),           // 3: api.Images
	(*Page)(nil),             // 4: api.Page
	(*GalleryRequest)(nil),   // 5: api.GalleryRequest
	(*GalleryImage)(nil),     // 6: api.GalleryImage
}
var file_api_api_proto_depIdxs = []int32{
	0,  // 0: api.ImageAuthRequest.Image:type_name -> api.Image
	0,  // 1: api.Images.Image:type_name -> api.Image
	0,  // 2: api.GalleryImage.Image:type_name -> api.Image
	5,  // 3: api.ImgShare.GetGalleryImage:input_type -> api.GalleryRequest
	2,  // 4: api.ImgShare.CreateImage:input_type -> api.ImageAuthRequest
	0,  // 5: api.ImgShare.ReadImage:input_type -> api.Image
	1,  // 6: api.ImgShare.GetRandomImage:input_type -> api.Empty
	2,  // 7: api.ImgShare.SetDescriptionImage:input_type -> api.ImageAuthRequest
	2,  // 8: api.ImgShare.UpvoteImage:input_type -> api.ImageAuthRequest
	2,  // 9: api.ImgShare.DownvoteImage:input_type -> api.ImageAuthRequest
	2,  // 10: api.ImgShare.DeleteImage:input_type -> api.ImageAuthRequest
	4,  // 11: api.ImgShare.GetAllImages:input_type -> api.Page
	6,  // 12: api.ImgShare.GetGalleryImage:output_type -> api.GalleryImage
	1,  // 13: api.ImgShare.CreateImage:output_type -> api.Empty
	0,  // 14: api.ImgShare.ReadImage:output_type -> api.Image
	0,  // 15: api.ImgShare.GetRandomImage:output_type -> api.Image
	0,  // 16: api.ImgShare.SetDescriptionImage:output_type -> api.Image
	0,  // 17: api.ImgShare.UpvoteImage:output_type -> api.Image
	0,  // 18: api.ImgShare.DownvoteImage:output_type -> api.Image
	1,  // 19: api.ImgShare.DeleteImage:output_type -> api.Empty
	3,  // 20: api.ImgShare.GetAllImages:output_type -> api.Images
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageAuthRequest); i {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryRequest); i {
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
		file_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryImage); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
