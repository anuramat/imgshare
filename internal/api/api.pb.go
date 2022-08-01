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

type State int32

const (
	State_StartState             State = 0
	State_NoState                State = 1
	State_UploadImageInitState   State = 2
	State_UploadImageState       State = 3
	State_UploadDescriptionState State = 4
	State_EditDescriptionState   State = 5
	State_RandomImageState       State = 6
	State_GalleryState           State = 7
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "StartState",
		1: "NoState",
		2: "UploadImageInitState",
		3: "UploadImageState",
		4: "UploadDescriptionState",
		5: "EditDescriptionState",
		6: "RandomImageState",
		7: "GalleryState",
	}
	State_value = map[string]int32{
		"StartState":             0,
		"NoState":                1,
		"UploadImageInitState":   2,
		"UploadImageState":       3,
		"UploadDescriptionState": 4,
		"EditDescriptionState":   5,
		"RandomImageState":       6,
		"GalleryState":           7,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_api_api_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID           int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	State            State  `protobuf:"varint,2,opt,name=State,proto3,enum=api.State" json:"State,omitempty"`
	LastUpload       string `protobuf:"bytes,3,opt,name=LastUpload,proto3" json:"LastUpload,omitempty"`
	LastDownload     string `protobuf:"bytes,4,opt,name=LastDownload,proto3" json:"LastDownload,omitempty"`
	LastGalleryIndex int64  `protobuf:"varint,5,opt,name=LastGalleryIndex,proto3" json:"LastGalleryIndex,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *User) GetState() State {
	if x != nil {
		return x.State
	}
	return State_StartState
}

func (x *User) GetLastUpload() string {
	if x != nil {
		return x.LastUpload
	}
	return ""
}

func (x *User) GetLastDownload() string {
	if x != nil {
		return x.LastDownload
	}
	return ""
}

func (x *User) GetLastGalleryIndex() int64 {
	if x != nil {
		return x.LastGalleryIndex
	}
	return 0
}

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
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
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
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
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
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageAuthRequest) ProtoMessage() {}

func (x *ImageAuthRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ImageAuthRequest.ProtoReflect.Descriptor instead.
func (*ImageAuthRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
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

	Image []*Image `protobuf:"bytes,1,rep,name=Image,proto3" json:"Image,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *Images) GetImage() []*Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0xb0, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4c,
	0x61, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x4c,
	0x61, 0x73, 0x74, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x47, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x79, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x55, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4c, 0x0a, 0x10, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x06, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x2a, 0xb2, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x07, 0x32, 0x91, 0x04, 0x0a, 0x05, 0x42,
	0x6f, 0x74, 0x44, 0x42, 0x12, 0x22, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
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
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x6d, 0x61, 0x74, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2d, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_api_proto_goTypes = []interface{}{
	(State)(0),               // 0: api.State
	(*User)(nil),             // 1: api.User
	(*Image)(nil),            // 2: api.Image
	(*Empty)(nil),            // 3: api.Empty
	(*ImageAuthRequest)(nil), // 4: api.ImageAuthRequest
	(*Images)(nil),           // 5: api.Images
}
var file_api_api_proto_depIdxs = []int32{
	0,  // 0: api.User.State:type_name -> api.State
	2,  // 1: api.ImageAuthRequest.Image:type_name -> api.Image
	2,  // 2: api.Images.Image:type_name -> api.Image
	1,  // 3: api.BotDB.CreateUser:input_type -> api.User
	1,  // 4: api.BotDB.ReadUser:input_type -> api.User
	1,  // 5: api.BotDB.UpdateUser:input_type -> api.User
	1,  // 6: api.BotDB.DeleteUser:input_type -> api.User
	4,  // 7: api.BotDB.CreateImage:input_type -> api.ImageAuthRequest
	2,  // 8: api.BotDB.ReadImage:input_type -> api.Image
	3,  // 9: api.BotDB.GetRandomImage:input_type -> api.Empty
	4,  // 10: api.BotDB.SetDescriptionImage:input_type -> api.ImageAuthRequest
	4,  // 11: api.BotDB.UpvoteImage:input_type -> api.ImageAuthRequest
	4,  // 12: api.BotDB.DownvoteImage:input_type -> api.ImageAuthRequest
	4,  // 13: api.BotDB.DeleteImage:input_type -> api.ImageAuthRequest
	3,  // 14: api.BotDB.GetAllImages:input_type -> api.Empty
	1,  // 15: api.BotDB.CreateUser:output_type -> api.User
	1,  // 16: api.BotDB.ReadUser:output_type -> api.User
	1,  // 17: api.BotDB.UpdateUser:output_type -> api.User
	1,  // 18: api.BotDB.DeleteUser:output_type -> api.User
	2,  // 19: api.BotDB.CreateImage:output_type -> api.Image
	2,  // 20: api.BotDB.ReadImage:output_type -> api.Image
	2,  // 21: api.BotDB.GetRandomImage:output_type -> api.Image
	2,  // 22: api.BotDB.SetDescriptionImage:output_type -> api.Image
	2,  // 23: api.BotDB.UpvoteImage:output_type -> api.Image
	2,  // 24: api.BotDB.DownvoteImage:output_type -> api.Image
	3,  // 25: api.BotDB.DeleteImage:output_type -> api.Empty
	5,  // 26: api.BotDB.GetAllImages:output_type -> api.Images
	15, // [15:27] is the sub-list for method output_type
	3,  // [3:15] is the sub-list for method input_type
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		EnumInfos:         file_api_api_proto_enumTypes,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
