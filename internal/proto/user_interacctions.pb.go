// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: internal/proto/user_interacctions.proto

package proto

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

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId        string               `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	WatchingTime   float32              `protobuf:"fixed32,2,opt,name=watching_time,json=watchingTime,proto3" json:"watching_time,omitempty"`
	WatchingRepeat int32                `protobuf:"varint,3,opt,name=watching_repeat,json=watchingRepeat,proto3" json:"watching_repeat,omitempty"`
	Interactions   *InteractionsRequest `protobuf:"bytes,4,opt,name=interactions,proto3" json:"interactions,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_user_interacctions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_user_interacctions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_user_interacctions_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoRequest) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *UserInfoRequest) GetWatchingTime() float32 {
	if x != nil {
		return x.WatchingTime
	}
	return 0
}

func (x *UserInfoRequest) GetWatchingRepeat() int32 {
	if x != nil {
		return x.WatchingRepeat
	}
	return 0
}

func (x *UserInfoRequest) GetInteractions() *InteractionsRequest {
	if x != nil {
		return x.Interactions
	}
	return nil
}

type InteractionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre       []string `protobuf:"bytes,1,rep,name=genre,proto3" json:"genre,omitempty"`
	Protagonist string   `protobuf:"bytes,2,opt,name=protagonist,proto3" json:"protagonist,omitempty"`
	Director    string   `protobuf:"bytes,3,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *InteractionsRequest) Reset() {
	*x = InteractionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_user_interacctions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionsRequest) ProtoMessage() {}

func (x *InteractionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_user_interacctions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionsRequest.ProtoReflect.Descriptor instead.
func (*InteractionsRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_user_interacctions_proto_rawDescGZIP(), []int{1}
}

func (x *InteractionsRequest) GetGenre() []string {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *InteractionsRequest) GetProtagonist() string {
	if x != nil {
		return x.Protagonist
	}
	return ""
}

func (x *InteractionsRequest) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

type MoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url         string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Genre       []string `protobuf:"bytes,4,rep,name=genre,proto3" json:"genre,omitempty"`
	Protagonist string   `protobuf:"bytes,5,opt,name=protagonist,proto3" json:"protagonist,omitempty"`
	Director    string   `protobuf:"bytes,6,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *MoviesRequest) Reset() {
	*x = MoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_user_interacctions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviesRequest) ProtoMessage() {}

func (x *MoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_user_interacctions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviesRequest.ProtoReflect.Descriptor instead.
func (*MoviesRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_user_interacctions_proto_rawDescGZIP(), []int{2}
}

func (x *MoviesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MoviesRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MoviesRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MoviesRequest) GetGenre() []string {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *MoviesRequest) GetProtagonist() string {
	if x != nil {
		return x.Protagonist
	}
	return ""
}

func (x *MoviesRequest) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

type PreprocessedDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string             `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Data   []*UserInfoRequest `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Movies []*MoviesRequest   `protobuf:"bytes,3,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *PreprocessedDataRequest) Reset() {
	*x = PreprocessedDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_user_interacctions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreprocessedDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreprocessedDataRequest) ProtoMessage() {}

func (x *PreprocessedDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_user_interacctions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreprocessedDataRequest.ProtoReflect.Descriptor instead.
func (*PreprocessedDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_user_interacctions_proto_rawDescGZIP(), []int{3}
}

func (x *PreprocessedDataRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PreprocessedDataRequest) GetData() []*UserInfoRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PreprocessedDataRequest) GetMovies() []*MoviesRequest {
	if x != nil {
		return x.Movies
	}
	return nil
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_user_interacctions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_user_interacctions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_user_interacctions_proto_rawDescGZIP(), []int{4}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_internal_proto_user_interacctions_proto protoreflect.FileDescriptor

var file_internal_proto_user_interacctions_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x12, 0x45, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x69, 0x0a, 0x13, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74,
	0x61, 0x67, 0x6f, 0x6e, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x61, 0x67, 0x6f, 0x6e, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x61, 0x67,
	0x6f, 0x6e, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x74, 0x61, 0x67, 0x6f, 0x6e, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x0f,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x6a, 0x0a, 0x13, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x25, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6b, 0x73, 0x75, 0x70, 0x2f, 0x74, 0x69, 0x6b, 0x73, 0x75,
	0x70, 0x2d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_user_interacctions_proto_rawDescOnce sync.Once
	file_internal_proto_user_interacctions_proto_rawDescData = file_internal_proto_user_interacctions_proto_rawDesc
)

func file_internal_proto_user_interacctions_proto_rawDescGZIP() []byte {
	file_internal_proto_user_interacctions_proto_rawDescOnce.Do(func() {
		file_internal_proto_user_interacctions_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_user_interacctions_proto_rawDescData)
	})
	return file_internal_proto_user_interacctions_proto_rawDescData
}

var file_internal_proto_user_interacctions_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_proto_user_interacctions_proto_goTypes = []any{
	(*UserInfoRequest)(nil),         // 0: interactions.UserInfoRequest
	(*InteractionsRequest)(nil),     // 1: interactions.InteractionsRequest
	(*MoviesRequest)(nil),           // 2: interactions.MoviesRequest
	(*PreprocessedDataRequest)(nil), // 3: interactions.PreprocessedDataRequest
	(*SuccessResponse)(nil),         // 4: interactions.SuccessResponse
}
var file_internal_proto_user_interacctions_proto_depIdxs = []int32{
	1, // 0: interactions.UserInfoRequest.interactions:type_name -> interactions.InteractionsRequest
	0, // 1: interactions.PreprocessedDataRequest.data:type_name -> interactions.UserInfoRequest
	2, // 2: interactions.PreprocessedDataRequest.movies:type_name -> interactions.MoviesRequest
	3, // 3: interactions.InteractionsService.ProcessData:input_type -> interactions.PreprocessedDataRequest
	4, // 4: interactions.InteractionsService.ProcessData:output_type -> interactions.SuccessResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_proto_user_interacctions_proto_init() }
func file_internal_proto_user_interacctions_proto_init() {
	if File_internal_proto_user_interacctions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_user_interacctions_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfoRequest); i {
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
		file_internal_proto_user_interacctions_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InteractionsRequest); i {
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
		file_internal_proto_user_interacctions_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MoviesRequest); i {
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
		file_internal_proto_user_interacctions_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PreprocessedDataRequest); i {
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
		file_internal_proto_user_interacctions_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SuccessResponse); i {
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
			RawDescriptor: file_internal_proto_user_interacctions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_user_interacctions_proto_goTypes,
		DependencyIndexes: file_internal_proto_user_interacctions_proto_depIdxs,
		MessageInfos:      file_internal_proto_user_interacctions_proto_msgTypes,
	}.Build()
	File_internal_proto_user_interacctions_proto = out.File
	file_internal_proto_user_interacctions_proto_rawDesc = nil
	file_internal_proto_user_interacctions_proto_goTypes = nil
	file_internal_proto_user_interacctions_proto_depIdxs = nil
}
