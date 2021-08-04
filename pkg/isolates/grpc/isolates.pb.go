// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/isolates/grpc/isolates.proto

package grpc

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

type BaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Namespace *string `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Workflow  *string `protobuf:"bytes,3,opt,name=workflow,proto3,oneof" json:"workflow,omitempty"`
	Image     *string `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Cmd       *string `protobuf:"bytes,5,opt,name=cmd,proto3,oneof" json:"cmd,omitempty"`
	Size      *int32  `protobuf:"varint,6,opt,name=size,proto3,oneof" json:"size,omitempty"`
	MinScale  *int32  `protobuf:"varint,7,opt,name=minScale,proto3,oneof" json:"minScale,omitempty"`
}

func (x *BaseInfo) Reset() {
	*x = BaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseInfo) ProtoMessage() {}

func (x *BaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseInfo.ProtoReflect.Descriptor instead.
func (*BaseInfo) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{0}
}

func (x *BaseInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BaseInfo) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *BaseInfo) GetWorkflow() string {
	if x != nil && x.Workflow != nil {
		return *x.Workflow
	}
	return ""
}

func (x *BaseInfo) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *BaseInfo) GetCmd() string {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return ""
}

func (x *BaseInfo) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *BaseInfo) GetMinScale() int32 {
	if x != nil && x.MinScale != nil {
		return *x.MinScale
	}
	return 0
}

type CreateIsolateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *BaseInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
}

func (x *CreateIsolateRequest) Reset() {
	*x = CreateIsolateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIsolateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIsolateRequest) ProtoMessage() {}

func (x *CreateIsolateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIsolateRequest.ProtoReflect.Descriptor instead.
func (*CreateIsolateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIsolateRequest) GetInfo() *BaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ListIsolatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Annotations map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListIsolatesRequest) Reset() {
	*x = ListIsolatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIsolatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIsolatesRequest) ProtoMessage() {}

func (x *ListIsolatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIsolatesRequest.ProtoReflect.Descriptor instead.
func (*ListIsolatesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{2}
}

func (x *ListIsolatesRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type IsolateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info          *BaseInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
	Status        *string   `protobuf:"bytes,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	StatusMessage *string   `protobuf:"bytes,3,opt,name=statusMessage,proto3,oneof" json:"statusMessage,omitempty"`
	ServiceName   *string   `protobuf:"bytes,4,opt,name=serviceName,proto3,oneof" json:"serviceName,omitempty"`
}

func (x *IsolateInfo) Reset() {
	*x = IsolateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsolateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsolateInfo) ProtoMessage() {}

func (x *IsolateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsolateInfo.ProtoReflect.Descriptor instead.
func (*IsolateInfo) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{3}
}

func (x *IsolateInfo) GetInfo() *BaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *IsolateInfo) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *IsolateInfo) GetStatusMessage() string {
	if x != nil && x.StatusMessage != nil {
		return *x.StatusMessage
	}
	return ""
}

func (x *IsolateInfo) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

type ListIsolatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isolates []*IsolateInfo `protobuf:"bytes,1,rep,name=isolates,proto3" json:"isolates,omitempty"`
}

func (x *ListIsolatesResponse) Reset() {
	*x = ListIsolatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIsolatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIsolatesResponse) ProtoMessage() {}

func (x *ListIsolatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIsolatesResponse.ProtoReflect.Descriptor instead.
func (*ListIsolatesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{4}
}

func (x *ListIsolatesResponse) GetIsolates() []*IsolateInfo {
	if x != nil {
		return x.Isolates
	}
	return nil
}

type UpdateIsolateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info        *BaseInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
	ServiceName *string   `protobuf:"bytes,2,opt,name=ServiceName,proto3,oneof" json:"ServiceName,omitempty"`
}

func (x *UpdateIsolateRequest) Reset() {
	*x = UpdateIsolateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIsolateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIsolateRequest) ProtoMessage() {}

func (x *UpdateIsolateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIsolateRequest.ProtoReflect.Descriptor instead.
func (*UpdateIsolateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateIsolateRequest) GetInfo() *BaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *UpdateIsolateRequest) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

type GetIsolateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName *string `protobuf:"bytes,1,opt,name=ServiceName,proto3,oneof" json:"ServiceName,omitempty"`
}

func (x *GetIsolateRequest) Reset() {
	*x = GetIsolateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIsolateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIsolateRequest) ProtoMessage() {}

func (x *GetIsolateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIsolateRequest.ProtoReflect.Descriptor instead.
func (*GetIsolateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{6}
}

func (x *GetIsolateRequest) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Image         *string `protobuf:"bytes,2,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Cmd           *string `protobuf:"bytes,3,opt,name=cmd,proto3,oneof" json:"cmd,omitempty"`
	Size          *int32  `protobuf:"varint,4,opt,name=size,proto3,oneof" json:"size,omitempty"`
	MinScale      *int32  `protobuf:"varint,5,opt,name=minScale,proto3,oneof" json:"minScale,omitempty"`
	Generation    *int64  `protobuf:"varint,6,opt,name=generation,proto3,oneof" json:"generation,omitempty"`
	Created       *int64  `protobuf:"varint,7,opt,name=created,proto3,oneof" json:"created,omitempty"`
	Status        *string `protobuf:"bytes,8,opt,name=status,proto3,oneof" json:"status,omitempty"`
	StatusMessage *string `protobuf:"bytes,9,opt,name=statusMessage,proto3,oneof" json:"statusMessage,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{7}
}

func (x *Revision) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Revision) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *Revision) GetCmd() string {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return ""
}

func (x *Revision) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *Revision) GetMinScale() int32 {
	if x != nil && x.MinScale != nil {
		return *x.MinScale
	}
	return 0
}

func (x *Revision) GetGeneration() int64 {
	if x != nil && x.Generation != nil {
		return *x.Generation
	}
	return 0
}

func (x *Revision) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

func (x *Revision) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Revision) GetStatusMessage() string {
	if x != nil && x.StatusMessage != nil {
		return *x.StatusMessage
	}
	return ""
}

type GetIsolateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string     `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Namespace *string     `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Workflow  *string     `protobuf:"bytes,3,opt,name=workflow,proto3,oneof" json:"workflow,omitempty"`
	Revisions []*Revision `protobuf:"bytes,4,rep,name=revisions,proto3" json:"revisions,omitempty"`
}

func (x *GetIsolateResponse) Reset() {
	*x = GetIsolateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIsolateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIsolateResponse) ProtoMessage() {}

func (x *GetIsolateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_isolates_grpc_isolates_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIsolateResponse.ProtoReflect.Descriptor instead.
func (*GetIsolateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_isolates_grpc_isolates_proto_rawDescGZIP(), []int{8}
}

func (x *GetIsolateResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetIsolateResponse) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *GetIsolateResponse) GetWorkflow() string {
	if x != nil && x.Workflow != nil {
		return *x.Workflow
	}
	return ""
}

func (x *GetIsolateResponse) GetRevisions() []*Revision {
	if x != nil {
		return x.Revisions
	}
	return nil
}

var File_pkg_isolates_grpc_isolates_proto protoreflect.FileDescriptor

var file_pkg_isolates_grpc_isolates_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x9f, 0x02, 0x0a, 0x08, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x03, 0x63, 0x6d,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x06, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x63, 0x6d, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x48, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x6f,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x6f,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x49,
	0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x7f, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x84, 0x03, 0x0a,
	0x08, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x63, 0x6d,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x04, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x07, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x63, 0x6d, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x6f,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_isolates_grpc_isolates_proto_rawDescOnce sync.Once
	file_pkg_isolates_grpc_isolates_proto_rawDescData = file_pkg_isolates_grpc_isolates_proto_rawDesc
)

func file_pkg_isolates_grpc_isolates_proto_rawDescGZIP() []byte {
	file_pkg_isolates_grpc_isolates_proto_rawDescOnce.Do(func() {
		file_pkg_isolates_grpc_isolates_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_isolates_grpc_isolates_proto_rawDescData)
	})
	return file_pkg_isolates_grpc_isolates_proto_rawDescData
}

var file_pkg_isolates_grpc_isolates_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_isolates_grpc_isolates_proto_goTypes = []interface{}{
	(*BaseInfo)(nil),             // 0: grpc.BaseInfo
	(*CreateIsolateRequest)(nil), // 1: grpc.CreateIsolateRequest
	(*ListIsolatesRequest)(nil),  // 2: grpc.ListIsolatesRequest
	(*IsolateInfo)(nil),          // 3: grpc.IsolateInfo
	(*ListIsolatesResponse)(nil), // 4: grpc.ListIsolatesResponse
	(*UpdateIsolateRequest)(nil), // 5: grpc.UpdateIsolateRequest
	(*GetIsolateRequest)(nil),    // 6: grpc.GetIsolateRequest
	(*Revision)(nil),             // 7: grpc.Revision
	(*GetIsolateResponse)(nil),   // 8: grpc.GetIsolateResponse
	nil,                          // 9: grpc.ListIsolatesRequest.AnnotationsEntry
}
var file_pkg_isolates_grpc_isolates_proto_depIdxs = []int32{
	0, // 0: grpc.CreateIsolateRequest.info:type_name -> grpc.BaseInfo
	9, // 1: grpc.ListIsolatesRequest.annotations:type_name -> grpc.ListIsolatesRequest.AnnotationsEntry
	0, // 2: grpc.IsolateInfo.info:type_name -> grpc.BaseInfo
	3, // 3: grpc.ListIsolatesResponse.isolates:type_name -> grpc.IsolateInfo
	0, // 4: grpc.UpdateIsolateRequest.info:type_name -> grpc.BaseInfo
	7, // 5: grpc.GetIsolateResponse.revisions:type_name -> grpc.Revision
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_isolates_grpc_isolates_proto_init() }
func file_pkg_isolates_grpc_isolates_proto_init() {
	if File_pkg_isolates_grpc_isolates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_isolates_grpc_isolates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseInfo); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIsolateRequest); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIsolatesRequest); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsolateInfo); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIsolatesResponse); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIsolateRequest); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIsolateRequest); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
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
		file_pkg_isolates_grpc_isolates_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIsolateResponse); i {
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
	file_pkg_isolates_grpc_isolates_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_isolates_grpc_isolates_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_isolates_grpc_isolates_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_pkg_isolates_grpc_isolates_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_pkg_isolates_grpc_isolates_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_pkg_isolates_grpc_isolates_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_pkg_isolates_grpc_isolates_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_isolates_grpc_isolates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_isolates_grpc_isolates_proto_goTypes,
		DependencyIndexes: file_pkg_isolates_grpc_isolates_proto_depIdxs,
		MessageInfos:      file_pkg_isolates_grpc_isolates_proto_msgTypes,
	}.Build()
	File_pkg_isolates_grpc_isolates_proto = out.File
	file_pkg_isolates_grpc_isolates_proto_rawDesc = nil
	file_pkg_isolates_grpc_isolates_proto_goTypes = nil
	file_pkg_isolates_grpc_isolates_proto_depIdxs = nil
}
