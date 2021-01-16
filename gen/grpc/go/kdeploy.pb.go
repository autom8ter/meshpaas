// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: kdeploy.proto

package kdeploypb

import (
	context "context"
	_ "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the application
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// application namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// docker image of application
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// k/v map of environmental variables
	Env map[string]string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k/v map of ports to expose ex: http: 80 https: 443
	Ports map[string]uint32 `protobuf:"bytes,5,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// number of deployment replicas
	Replicas uint32 `protobuf:"varint,6,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// status tracks the state of the application during it's lifecycle
	Status *Status `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{0}
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *App) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *App) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *App) GetPorts() map[string]uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *App) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *App) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type AppConstructor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the application
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// application namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// docker image of application
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// k/v map of environmental variables
	Env map[string]string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k/v map of ports to expose ex: http: 80 https: 443
	Ports map[string]uint32 `protobuf:"bytes,5,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// number of deployment replicas
	Replicas uint32 `protobuf:"varint,6,opt,name=replicas,proto3" json:"replicas,omitempty"`
}

func (x *AppConstructor) Reset() {
	*x = AppConstructor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConstructor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConstructor) ProtoMessage() {}

func (x *AppConstructor) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConstructor.ProtoReflect.Descriptor instead.
func (*AppConstructor) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{1}
}

func (x *AppConstructor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppConstructor) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AppConstructor) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *AppConstructor) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *AppConstructor) GetPorts() map[string]uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *AppConstructor) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

type AppUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the application
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// application namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// docker image of application
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// k/v map of environmental variables
	Env map[string]string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k/v map of ports to expose ex: http: 80 https: 443
	Ports map[string]uint32 `protobuf:"bytes,5,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// number of deployment replicas
	Replicas uint32 `protobuf:"varint,6,opt,name=replicas,proto3" json:"replicas,omitempty"`
}

func (x *AppUpdate) Reset() {
	*x = AppUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUpdate) ProtoMessage() {}

func (x *AppUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUpdate.ProtoReflect.Descriptor instead.
func (*AppUpdate) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{2}
}

func (x *AppUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppUpdate) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AppUpdate) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *AppUpdate) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *AppUpdate) GetPorts() map[string]uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *AppUpdate) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

type AppRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the application
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// application namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *AppRef) Reset() {
	*x = AppRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppRef) ProtoMessage() {}

func (x *AppRef) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppRef.ProtoReflect.Descriptor instead.
func (*AppRef) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{3}
}

func (x *AppRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Replica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase     string `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	Condition string `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
	Reason    string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Replica) Reset() {
	*x = Replica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replica) ProtoMessage() {}

func (x *Replica) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replica.ProtoReflect.Descriptor instead.
func (*Replica) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{4}
}

func (x *Replica) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *Replica) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Replica) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replicas []*Replica `protobuf:"bytes,1,rep,name=replicas,proto3" json:"replicas,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetReplicas() []*Replica {
	if x != nil {
		return x.Replicas
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kdeploy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_kdeploy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_kdeploy_proto_rawDescGZIP(), []int{6}
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_kdeploy_proto protoreflect.FileDescriptor

var file_kdeploy_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41,
	0x70, 0x70, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76,
	0x12, 0x2d, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a, 0x03, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e,
	0x2e, 0x7b, 0x31, 0x2c, 0x32, 0x32, 0x35, 0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e, 0x2e, 0x7b, 0x31, 0x2c, 0x32,
	0x32, 0x35, 0x7d, 0x24, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10,
	0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e, 0x2e, 0x7b, 0x31, 0x2c, 0x32, 0x32, 0x35, 0x7d, 0x24,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e,
	0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x38, 0x0a, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6b, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xe9, 0x02, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e, 0x2e, 0x7b, 0x31, 0x2c, 0x32, 0x32, 0x35, 0x7d,
	0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c,
	0x0a, 0x0a, 0x5e, 0x2e, 0x7b, 0x31, 0x2c, 0x32, 0x32, 0x35, 0x7d, 0x24, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x45,
	0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x33, 0x0a, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x1a, 0x36, 0x0a,
	0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x5e, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x52, 0x65, 0x66, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e,
	0x2e, 0x7b, 0x31, 0x2c, 0x32, 0x32, 0x35, 0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x0a, 0x0a, 0x5e, 0x2e, 0x7b, 0x31, 0x2c, 0x32,
	0x32, 0x35, 0x7d, 0x24, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22,
	0x55, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x22, 0x1f,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x85, 0x02, 0x0a, 0x0e, 0x4b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12,
	0x17, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x0c, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x12, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x6b, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x0f, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x66, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x0f, 0x2e, 0x6b, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x66, 0x1a, 0x0c, 0x2e, 0x6b,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x0f, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x66, 0x1a, 0x0c, 0x2e, 0x6b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x6b, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kdeploy_proto_rawDescOnce sync.Once
	file_kdeploy_proto_rawDescData = file_kdeploy_proto_rawDesc
)

func file_kdeploy_proto_rawDescGZIP() []byte {
	file_kdeploy_proto_rawDescOnce.Do(func() {
		file_kdeploy_proto_rawDescData = protoimpl.X.CompressGZIP(file_kdeploy_proto_rawDescData)
	})
	return file_kdeploy_proto_rawDescData
}

var file_kdeploy_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_kdeploy_proto_goTypes = []interface{}{
	(*App)(nil),            // 0: kdeploy.App
	(*AppConstructor)(nil), // 1: kdeploy.AppConstructor
	(*AppUpdate)(nil),      // 2: kdeploy.AppUpdate
	(*AppRef)(nil),         // 3: kdeploy.AppRef
	(*Replica)(nil),        // 4: kdeploy.Replica
	(*Status)(nil),         // 5: kdeploy.Status
	(*Log)(nil),            // 6: kdeploy.Log
	nil,                    // 7: kdeploy.App.EnvEntry
	nil,                    // 8: kdeploy.App.PortsEntry
	nil,                    // 9: kdeploy.AppConstructor.EnvEntry
	nil,                    // 10: kdeploy.AppConstructor.PortsEntry
	nil,                    // 11: kdeploy.AppUpdate.EnvEntry
	nil,                    // 12: kdeploy.AppUpdate.PortsEntry
	(*empty.Empty)(nil),    // 13: google.protobuf.Empty
}
var file_kdeploy_proto_depIdxs = []int32{
	7,  // 0: kdeploy.App.env:type_name -> kdeploy.App.EnvEntry
	8,  // 1: kdeploy.App.ports:type_name -> kdeploy.App.PortsEntry
	5,  // 2: kdeploy.App.status:type_name -> kdeploy.Status
	9,  // 3: kdeploy.AppConstructor.env:type_name -> kdeploy.AppConstructor.EnvEntry
	10, // 4: kdeploy.AppConstructor.ports:type_name -> kdeploy.AppConstructor.PortsEntry
	11, // 5: kdeploy.AppUpdate.env:type_name -> kdeploy.AppUpdate.EnvEntry
	12, // 6: kdeploy.AppUpdate.ports:type_name -> kdeploy.AppUpdate.PortsEntry
	4,  // 7: kdeploy.Status.replicas:type_name -> kdeploy.Replica
	1,  // 8: kdeploy.KdeployService.CreateApp:input_type -> kdeploy.AppConstructor
	2,  // 9: kdeploy.KdeployService.UpdateApp:input_type -> kdeploy.AppUpdate
	3,  // 10: kdeploy.KdeployService.DeleteApp:input_type -> kdeploy.AppRef
	3,  // 11: kdeploy.KdeployService.GetApp:input_type -> kdeploy.AppRef
	3,  // 12: kdeploy.KdeployService.Logs:input_type -> kdeploy.AppRef
	0,  // 13: kdeploy.KdeployService.CreateApp:output_type -> kdeploy.App
	0,  // 14: kdeploy.KdeployService.UpdateApp:output_type -> kdeploy.App
	13, // 15: kdeploy.KdeployService.DeleteApp:output_type -> google.protobuf.Empty
	0,  // 16: kdeploy.KdeployService.GetApp:output_type -> kdeploy.App
	6,  // 17: kdeploy.KdeployService.Logs:output_type -> kdeploy.Log
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_kdeploy_proto_init() }
func file_kdeploy_proto_init() {
	if File_kdeploy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kdeploy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
		file_kdeploy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConstructor); i {
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
		file_kdeploy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUpdate); i {
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
		file_kdeploy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppRef); i {
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
		file_kdeploy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replica); i {
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
		file_kdeploy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_kdeploy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_kdeploy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kdeploy_proto_goTypes,
		DependencyIndexes: file_kdeploy_proto_depIdxs,
		MessageInfos:      file_kdeploy_proto_msgTypes,
	}.Build()
	File_kdeploy_proto = out.File
	file_kdeploy_proto_rawDesc = nil
	file_kdeploy_proto_goTypes = nil
	file_kdeploy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KdeployServiceClient is the client API for KdeployService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KdeployServiceClient interface {
	CreateApp(ctx context.Context, in *AppConstructor, opts ...grpc.CallOption) (*App, error)
	UpdateApp(ctx context.Context, in *AppUpdate, opts ...grpc.CallOption) (*App, error)
	DeleteApp(ctx context.Context, in *AppRef, opts ...grpc.CallOption) (*empty.Empty, error)
	GetApp(ctx context.Context, in *AppRef, opts ...grpc.CallOption) (*App, error)
	Logs(ctx context.Context, in *AppRef, opts ...grpc.CallOption) (KdeployService_LogsClient, error)
}

type kdeployServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKdeployServiceClient(cc grpc.ClientConnInterface) KdeployServiceClient {
	return &kdeployServiceClient{cc}
}

func (c *kdeployServiceClient) CreateApp(ctx context.Context, in *AppConstructor, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, "/kdeploy.KdeployService/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kdeployServiceClient) UpdateApp(ctx context.Context, in *AppUpdate, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, "/kdeploy.KdeployService/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kdeployServiceClient) DeleteApp(ctx context.Context, in *AppRef, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kdeploy.KdeployService/DeleteApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kdeployServiceClient) GetApp(ctx context.Context, in *AppRef, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, "/kdeploy.KdeployService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kdeployServiceClient) Logs(ctx context.Context, in *AppRef, opts ...grpc.CallOption) (KdeployService_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KdeployService_serviceDesc.Streams[0], "/kdeploy.KdeployService/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &kdeployServiceLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KdeployService_LogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type kdeployServiceLogsClient struct {
	grpc.ClientStream
}

func (x *kdeployServiceLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KdeployServiceServer is the server API for KdeployService service.
type KdeployServiceServer interface {
	CreateApp(context.Context, *AppConstructor) (*App, error)
	UpdateApp(context.Context, *AppUpdate) (*App, error)
	DeleteApp(context.Context, *AppRef) (*empty.Empty, error)
	GetApp(context.Context, *AppRef) (*App, error)
	Logs(*AppRef, KdeployService_LogsServer) error
}

// UnimplementedKdeployServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKdeployServiceServer struct {
}

func (*UnimplementedKdeployServiceServer) CreateApp(context.Context, *AppConstructor) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (*UnimplementedKdeployServiceServer) UpdateApp(context.Context, *AppUpdate) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApp not implemented")
}
func (*UnimplementedKdeployServiceServer) DeleteApp(context.Context, *AppRef) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApp not implemented")
}
func (*UnimplementedKdeployServiceServer) GetApp(context.Context, *AppRef) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (*UnimplementedKdeployServiceServer) Logs(*AppRef, KdeployService_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}

func RegisterKdeployServiceServer(s *grpc.Server, srv KdeployServiceServer) {
	s.RegisterService(&_KdeployService_serviceDesc, srv)
}

func _KdeployService_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppConstructor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KdeployServiceServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kdeploy.KdeployService/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KdeployServiceServer).CreateApp(ctx, req.(*AppConstructor))
	}
	return interceptor(ctx, in, info, handler)
}

func _KdeployService_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KdeployServiceServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kdeploy.KdeployService/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KdeployServiceServer).UpdateApp(ctx, req.(*AppUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _KdeployService_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KdeployServiceServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kdeploy.KdeployService/DeleteApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KdeployServiceServer).DeleteApp(ctx, req.(*AppRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _KdeployService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KdeployServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kdeploy.KdeployService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KdeployServiceServer).GetApp(ctx, req.(*AppRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _KdeployService_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AppRef)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KdeployServiceServer).Logs(m, &kdeployServiceLogsServer{stream})
}

type KdeployService_LogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type kdeployServiceLogsServer struct {
	grpc.ServerStream
}

func (x *kdeployServiceLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _KdeployService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kdeploy.KdeployService",
	HandlerType: (*KdeployServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _KdeployService_CreateApp_Handler,
		},
		{
			MethodName: "UpdateApp",
			Handler:    _KdeployService_UpdateApp_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _KdeployService_DeleteApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _KdeployService_GetApp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _KdeployService_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kdeploy.proto",
}
