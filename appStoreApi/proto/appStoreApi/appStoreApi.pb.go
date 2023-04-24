// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/appStoreApi/appStoreApi.proto

package appStoreApi

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

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appStoreApi_appStoreApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appStoreApi_appStoreApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_proto_appStoreApi_appStoreApi_proto_rawDescGZIP(), []int{0}
}

func (x *Pair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pair) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body   string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url    string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appStoreApi_appStoreApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appStoreApi_appStoreApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_appStoreApi_appStoreApi_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Request) GetHeader() map[string]*Pair {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetGet() map[string]*Pair {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *Request) GetPost() map[string]*Pair {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *Request) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32            `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body       string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appStoreApi_appStoreApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appStoreApi_appStoreApi_proto_msgTypes[2]
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
	return file_proto_appStoreApi_appStoreApi_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetHeader() map[string]*Pair {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_proto_appStoreApi_appStoreApi_proto protoreflect.FileDescriptor

var file_proto_appStoreApi_appStoreApi_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x41, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41,
	0x70, 0x69, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0xdf, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61,
	0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x1a, 0x4c, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61,
	0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x09, 0x50, 0x6f,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc7, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x1a, 0x4c, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69,
	0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x32, 0xc1, 0x05, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69,
	0x12, 0x37, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70,
	0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c,
	0x41, 0x70, 0x70, 0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x56, 0x69, 0x65, 0x77,
	0x4e, 0x75, 0x6d, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x69, 0x65, 0x77, 0x4e,
	0x75, 0x6d, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x70, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_appStoreApi_appStoreApi_proto_rawDescOnce sync.Once
	file_proto_appStoreApi_appStoreApi_proto_rawDescData = file_proto_appStoreApi_appStoreApi_proto_rawDesc
)

func file_proto_appStoreApi_appStoreApi_proto_rawDescGZIP() []byte {
	file_proto_appStoreApi_appStoreApi_proto_rawDescOnce.Do(func() {
		file_proto_appStoreApi_appStoreApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_appStoreApi_appStoreApi_proto_rawDescData)
	})
	return file_proto_appStoreApi_appStoreApi_proto_rawDescData
}

var file_proto_appStoreApi_appStoreApi_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_appStoreApi_appStoreApi_proto_goTypes = []interface{}{
	(*Pair)(nil),     // 0: appStoreApi.Pair
	(*Request)(nil),  // 1: appStoreApi.Request
	(*Response)(nil), // 2: appStoreApi.Response
	nil,              // 3: appStoreApi.Request.HeaderEntry
	nil,              // 4: appStoreApi.Request.GetEntry
	nil,              // 5: appStoreApi.Request.PostEntry
	nil,              // 6: appStoreApi.Response.HeaderEntry
}
var file_proto_appStoreApi_appStoreApi_proto_depIdxs = []int32{
	3,  // 0: appStoreApi.Request.header:type_name -> appStoreApi.Request.HeaderEntry
	4,  // 1: appStoreApi.Request.get:type_name -> appStoreApi.Request.GetEntry
	5,  // 2: appStoreApi.Request.post:type_name -> appStoreApi.Request.PostEntry
	6,  // 3: appStoreApi.Response.header:type_name -> appStoreApi.Response.HeaderEntry
	0,  // 4: appStoreApi.Request.HeaderEntry.value:type_name -> appStoreApi.Pair
	0,  // 5: appStoreApi.Request.GetEntry.value:type_name -> appStoreApi.Pair
	0,  // 6: appStoreApi.Request.PostEntry.value:type_name -> appStoreApi.Pair
	0,  // 7: appStoreApi.Response.HeaderEntry.value:type_name -> appStoreApi.Pair
	1,  // 8: appStoreApi.AppStoreApi.AddApp:input_type -> appStoreApi.Request
	1,  // 9: appStoreApi.AppStoreApi.DeleteAppByID:input_type -> appStoreApi.Request
	1,  // 10: appStoreApi.AppStoreApi.UpdateApp:input_type -> appStoreApi.Request
	1,  // 11: appStoreApi.AppStoreApi.QueryAppByID:input_type -> appStoreApi.Request
	1,  // 12: appStoreApi.AppStoreApi.QueryAllApps:input_type -> appStoreApi.Request
	1,  // 13: appStoreApi.AppStoreApi.AddInstallNum:input_type -> appStoreApi.Request
	1,  // 14: appStoreApi.AppStoreApi.QueryInstallNum:input_type -> appStoreApi.Request
	1,  // 15: appStoreApi.AppStoreApi.AddViewNum:input_type -> appStoreApi.Request
	1,  // 16: appStoreApi.AppStoreApi.QueryViewNum:input_type -> appStoreApi.Request
	1,  // 17: appStoreApi.AppStoreApi.AddComment:input_type -> appStoreApi.Request
	1,  // 18: appStoreApi.AppStoreApi.QueryAllCommentByID:input_type -> appStoreApi.Request
	2,  // 19: appStoreApi.AppStoreApi.AddApp:output_type -> appStoreApi.Response
	2,  // 20: appStoreApi.AppStoreApi.DeleteAppByID:output_type -> appStoreApi.Response
	2,  // 21: appStoreApi.AppStoreApi.UpdateApp:output_type -> appStoreApi.Response
	2,  // 22: appStoreApi.AppStoreApi.QueryAppByID:output_type -> appStoreApi.Response
	2,  // 23: appStoreApi.AppStoreApi.QueryAllApps:output_type -> appStoreApi.Response
	2,  // 24: appStoreApi.AppStoreApi.AddInstallNum:output_type -> appStoreApi.Response
	2,  // 25: appStoreApi.AppStoreApi.QueryInstallNum:output_type -> appStoreApi.Response
	2,  // 26: appStoreApi.AppStoreApi.AddViewNum:output_type -> appStoreApi.Response
	2,  // 27: appStoreApi.AppStoreApi.QueryViewNum:output_type -> appStoreApi.Response
	2,  // 28: appStoreApi.AppStoreApi.AddComment:output_type -> appStoreApi.Response
	2,  // 29: appStoreApi.AppStoreApi.QueryAllCommentByID:output_type -> appStoreApi.Response
	19, // [19:30] is the sub-list for method output_type
	8,  // [8:19] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_appStoreApi_appStoreApi_proto_init() }
func file_proto_appStoreApi_appStoreApi_proto_init() {
	if File_proto_appStoreApi_appStoreApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_appStoreApi_appStoreApi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_proto_appStoreApi_appStoreApi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_appStoreApi_appStoreApi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_appStoreApi_appStoreApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_appStoreApi_appStoreApi_proto_goTypes,
		DependencyIndexes: file_proto_appStoreApi_appStoreApi_proto_depIdxs,
		MessageInfos:      file_proto_appStoreApi_appStoreApi_proto_msgTypes,
	}.Build()
	File_proto_appStoreApi_appStoreApi_proto = out.File
	file_proto_appStoreApi_appStoreApi_proto_rawDesc = nil
	file_proto_appStoreApi_appStoreApi_proto_goTypes = nil
	file_proto_appStoreApi_appStoreApi_proto_depIdxs = nil
}
