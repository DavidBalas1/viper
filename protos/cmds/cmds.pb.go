// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: cmds.proto

package cmds

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

type ShellCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Cmd  string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *ShellCommandRequest) Reset() {
	*x = ShellCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellCommandRequest) ProtoMessage() {}

func (x *ShellCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellCommandRequest.ProtoReflect.Descriptor instead.
func (*ShellCommandRequest) Descriptor() ([]byte, []int) {
	return file_cmds_proto_rawDescGZIP(), []int{0}
}

func (x *ShellCommandRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ShellCommandRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type ShellCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ShellCommandResponse) Reset() {
	*x = ShellCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellCommandResponse) ProtoMessage() {}

func (x *ShellCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellCommandResponse.ProtoReflect.Descriptor instead.
func (*ShellCommandResponse) Descriptor() ([]byte, []int) {
	return file_cmds_proto_rawDescGZIP(), []int{1}
}

func (x *ShellCommandResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_cmds_proto_rawDescGZIP(), []int{2}
}

func (x *AgentInfo) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type EchoCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *EchoCommandRequest) Reset() {
	*x = EchoCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoCommandRequest) ProtoMessage() {}

func (x *EchoCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoCommandRequest.ProtoReflect.Descriptor instead.
func (*EchoCommandRequest) Descriptor() ([]byte, []int) {
	return file_cmds_proto_rawDescGZIP(), []int{3}
}

func (x *EchoCommandRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *EchoCommandRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type EchoCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *EchoCommandResponse) Reset() {
	*x = EchoCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoCommandResponse) ProtoMessage() {}

func (x *EchoCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoCommandResponse.ProtoReflect.Descriptor instead.
func (*EchoCommandResponse) Descriptor() ([]byte, []int) {
	return file_cmds_proto_rawDescGZIP(), []int{4}
}

func (x *EchoCommandResponse) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *EchoCommandResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_cmds_proto_msgTypes[5]
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
	return file_cmds_proto_rawDescGZIP(), []int{5}
}

var File_cmds_proto protoreflect.FileDescriptor

var file_cmds_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6d, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x13,
	0x53, 0x68, 0x65, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x53, 0x68, 0x65,
	0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x22, 0x3c, 0x0a, 0x12, 0x45, 0x63,
	0x68, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3d, 0x0a, 0x13, 0x45, 0x63, 0x68, 0x6f,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0x90, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x52, 0x75,
	0x6e, 0x45, 0x63, 0x68, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x45,
	0x63, 0x68, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x1a, 0x13, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x44, 0x0a,
	0x0f, 0x52, 0x75, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x15, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x14, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x32, 0xb4, 0x01, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x45, 0x63, 0x68, 0x6f, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53,
	0x68, 0x65, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f,
	0x63, 0x6d, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmds_proto_rawDescOnce sync.Once
	file_cmds_proto_rawDescData = file_cmds_proto_rawDesc
)

func file_cmds_proto_rawDescGZIP() []byte {
	file_cmds_proto_rawDescOnce.Do(func() {
		file_cmds_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmds_proto_rawDescData)
	})
	return file_cmds_proto_rawDescData
}

var file_cmds_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmds_proto_goTypes = []interface{}{
	(*ShellCommandRequest)(nil),  // 0: ShellCommandRequest
	(*ShellCommandResponse)(nil), // 1: ShellCommandResponse
	(*AgentInfo)(nil),            // 2: AgentInfo
	(*EchoCommandRequest)(nil),   // 3: EchoCommandRequest
	(*EchoCommandResponse)(nil),  // 4: EchoCommandResponse
	(*Empty)(nil),                // 5: Empty
}
var file_cmds_proto_depIdxs = []int32{
	4, // 0: Agent.RunEchoCommand:input_type -> EchoCommandResponse
	1, // 1: Agent.RunShellCommand:input_type -> ShellCommandResponse
	3, // 2: AgentManager.RunEchoCommand:input_type -> EchoCommandRequest
	0, // 3: AgentManager.RunShellCommand:input_type -> ShellCommandRequest
	5, // 4: AgentManager.GetAgents:input_type -> Empty
	3, // 5: Agent.RunEchoCommand:output_type -> EchoCommandRequest
	0, // 6: Agent.RunShellCommand:output_type -> ShellCommandRequest
	4, // 7: AgentManager.RunEchoCommand:output_type -> EchoCommandResponse
	1, // 8: AgentManager.RunShellCommand:output_type -> ShellCommandResponse
	2, // 9: AgentManager.GetAgents:output_type -> AgentInfo
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmds_proto_init() }
func file_cmds_proto_init() {
	if File_cmds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellCommandRequest); i {
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
		file_cmds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellCommandResponse); i {
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
		file_cmds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfo); i {
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
		file_cmds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoCommandRequest); i {
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
		file_cmds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoCommandResponse); i {
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
		file_cmds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cmds_proto_goTypes,
		DependencyIndexes: file_cmds_proto_depIdxs,
		MessageInfos:      file_cmds_proto_msgTypes,
	}.Build()
	File_cmds_proto = out.File
	file_cmds_proto_rawDesc = nil
	file_cmds_proto_goTypes = nil
	file_cmds_proto_depIdxs = nil
}
