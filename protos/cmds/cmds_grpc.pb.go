// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: cmds.proto

package cmds

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	RunShellCommand(ctx context.Context, opts ...grpc.CallOption) (Agent_RunShellCommandClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) RunShellCommand(ctx context.Context, opts ...grpc.CallOption) (Agent_RunShellCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/Agent/RunShellCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRunShellCommandClient{stream}
	return x, nil
}

type Agent_RunShellCommandClient interface {
	Send(*ShellCommandResponse) error
	Recv() (*ShellCommandRequest, error)
	grpc.ClientStream
}

type agentRunShellCommandClient struct {
	grpc.ClientStream
}

func (x *agentRunShellCommandClient) Send(m *ShellCommandResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRunShellCommandClient) Recv() (*ShellCommandRequest, error) {
	m := new(ShellCommandRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	RunShellCommand(Agent_RunShellCommandServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) RunShellCommand(Agent_RunShellCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method RunShellCommand not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_RunShellCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).RunShellCommand(&agentRunShellCommandServer{stream})
}

type Agent_RunShellCommandServer interface {
	Send(*ShellCommandRequest) error
	Recv() (*ShellCommandResponse, error)
	grpc.ServerStream
}

type agentRunShellCommandServer struct {
	grpc.ServerStream
}

func (x *agentRunShellCommandServer) Send(m *ShellCommandRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRunShellCommandServer) Recv() (*ShellCommandResponse, error) {
	m := new(ShellCommandResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Agent",
	HandlerType: (*AgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunShellCommand",
			Handler:       _Agent_RunShellCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cmds.proto",
}
