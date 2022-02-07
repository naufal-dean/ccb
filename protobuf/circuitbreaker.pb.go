// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: circuitbreaker.proto

package protobuf

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

type ListString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListString) Reset() {
	*x = ListString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListString) ProtoMessage() {}

func (x *ListString) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListString.ProtoReflect.Descriptor instead.
func (*ListString) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{0}
}

func (x *ListString) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type RequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url    string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Header map[string]*ListString `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestInput) Reset() {
	*x = RequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInput) ProtoMessage() {}

func (x *RequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInput.ProtoReflect.Descriptor instead.
func (*RequestInput) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{1}
}

func (x *RequestInput) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RequestInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RequestInput) GetHeader() map[string]*ListString {
	if x != nil {
		return x.Header
	}
	return nil
}

type GetInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Header map[string]*ListString `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetInput) Reset() {
	*x = GetInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInput) ProtoMessage() {}

func (x *GetInput) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInput.ProtoReflect.Descriptor instead.
func (*GetInput) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{2}
}

func (x *GetInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetInput) GetHeader() map[string]*ListString {
	if x != nil {
		return x.Header
	}
	return nil
}

type PostInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Header map[string]*ListString `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body   []byte                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PostInput) Reset() {
	*x = PostInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostInput) ProtoMessage() {}

func (x *PostInput) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostInput.ProtoReflect.Descriptor instead.
func (*PostInput) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{3}
}

func (x *PostInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PostInput) GetHeader() map[string]*ListString {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PostInput) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Proto         string                 `protobuf:"bytes,3,opt,name=proto,proto3" json:"proto,omitempty"`
	ProtoMajor    int32                  `protobuf:"varint,4,opt,name=proto_major,json=protoMajor,proto3" json:"proto_major,omitempty"`
	ProtoMinor    int32                  `protobuf:"varint,5,opt,name=proto_minor,json=protoMinor,proto3" json:"proto_minor,omitempty"`
	Header        map[string]*ListString `protobuf:"bytes,6,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body          []byte                 `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	ContentLength int64                  `protobuf:"varint,8,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *Response) GetProtoMajor() int32 {
	if x != nil {
		return x.ProtoMajor
	}
	return 0
}

func (x *Response) GetProtoMinor() int32 {
	if x != nil {
		return x.ProtoMinor
	}
	return 0
}

func (x *Response) GetHeader() map[string]*ListString {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Response) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

var File_circuitbreaker_proto protoreflect.FileDescriptor

var file_circuitbreaker_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x22, 0x24, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x4f, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5,
	0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x36, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x4f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x1a, 0x4f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xdf, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x61, 0x6a, 0x6f,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x69, 0x6e, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x69, 0x6e,
	0x6f, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x1a, 0x4f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa3, 0x01, 0x0a, 0x04, 0x48, 0x74, 0x74, 0x70, 0x12,
	0x37, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x75, 0x66, 0x61,
	0x6c, 0x2d, 0x64, 0x65, 0x61, 0x6e, 0x2f, 0x63, 0x63, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_circuitbreaker_proto_rawDescOnce sync.Once
	file_circuitbreaker_proto_rawDescData = file_circuitbreaker_proto_rawDesc
)

func file_circuitbreaker_proto_rawDescGZIP() []byte {
	file_circuitbreaker_proto_rawDescOnce.Do(func() {
		file_circuitbreaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_circuitbreaker_proto_rawDescData)
	})
	return file_circuitbreaker_proto_rawDescData
}

var file_circuitbreaker_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_circuitbreaker_proto_goTypes = []interface{}{
	(*ListString)(nil),   // 0: protobuf.ListString
	(*RequestInput)(nil), // 1: protobuf.RequestInput
	(*GetInput)(nil),     // 2: protobuf.GetInput
	(*PostInput)(nil),    // 3: protobuf.PostInput
	(*Response)(nil),     // 4: protobuf.Response
	nil,                  // 5: protobuf.RequestInput.HeaderEntry
	nil,                  // 6: protobuf.GetInput.HeaderEntry
	nil,                  // 7: protobuf.PostInput.HeaderEntry
	nil,                  // 8: protobuf.Response.HeaderEntry
}
var file_circuitbreaker_proto_depIdxs = []int32{
	5,  // 0: protobuf.RequestInput.header:type_name -> protobuf.RequestInput.HeaderEntry
	6,  // 1: protobuf.GetInput.header:type_name -> protobuf.GetInput.HeaderEntry
	7,  // 2: protobuf.PostInput.header:type_name -> protobuf.PostInput.HeaderEntry
	8,  // 3: protobuf.Response.header:type_name -> protobuf.Response.HeaderEntry
	0,  // 4: protobuf.RequestInput.HeaderEntry.value:type_name -> protobuf.ListString
	0,  // 5: protobuf.GetInput.HeaderEntry.value:type_name -> protobuf.ListString
	0,  // 6: protobuf.PostInput.HeaderEntry.value:type_name -> protobuf.ListString
	0,  // 7: protobuf.Response.HeaderEntry.value:type_name -> protobuf.ListString
	1,  // 8: protobuf.Http.Request:input_type -> protobuf.RequestInput
	2,  // 9: protobuf.Http.Get:input_type -> protobuf.GetInput
	3,  // 10: protobuf.Http.Post:input_type -> protobuf.PostInput
	4,  // 11: protobuf.Http.Request:output_type -> protobuf.Response
	4,  // 12: protobuf.Http.Get:output_type -> protobuf.Response
	4,  // 13: protobuf.Http.Post:output_type -> protobuf.Response
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_circuitbreaker_proto_init() }
func file_circuitbreaker_proto_init() {
	if File_circuitbreaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_circuitbreaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListString); i {
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
		file_circuitbreaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInput); i {
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
		file_circuitbreaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInput); i {
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
		file_circuitbreaker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostInput); i {
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
		file_circuitbreaker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_circuitbreaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_circuitbreaker_proto_goTypes,
		DependencyIndexes: file_circuitbreaker_proto_depIdxs,
		MessageInfos:      file_circuitbreaker_proto_msgTypes,
	}.Build()
	File_circuitbreaker_proto = out.File
	file_circuitbreaker_proto_rawDesc = nil
	file_circuitbreaker_proto_goTypes = nil
	file_circuitbreaker_proto_depIdxs = nil
}
